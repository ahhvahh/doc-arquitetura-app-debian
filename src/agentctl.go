package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const normGitURL = "https://github.com/ahhvahh/doc-arquitetura.git"
const normTree = "ProjectCharter/1.0"

type modelSpec struct{ ID, Repo, File, Purpose string }

var modelCatalog = []modelSpec{
	{"qwen3.5-0.8b", "ggml-org/Qwen3.5-0.8B-GGUF", "Qwen3.5-0.8B-Q4_0.gguf", "triagem, classificacao e extracao preliminar"},
	{"qwen3-1.7b", "bartowski/Qwen_Qwen3-1.7B-GGUF", "Qwen_Qwen3-1.7B-Q4_K_M.gguf", "analise semantica, contratos e inconsistencias"},
}

type modelRecord struct {
	ID          string `json:"id"`
	Repository  string `json:"repository"`
	Revision    string `json:"revision"`
	File        string `json:"file"`
	SHA256      string `json:"sha256"`
	Size        int64  `json:"size"`
	InstalledAt string `json:"installed_at"`
}
type refRecord struct {
	SHA       string `json:"sha"`
	Source    string `json:"source"`
	UpdatedAt string `json:"updated_at"`
}
type dslRecord struct {
	SHA          string `json:"reference_sha"`
	BundleSHA256 string `json:"bundle_sha256"`
	Documents    int    `json:"documents"`
	CompiledAt   string `json:"compiled_at"`
}

type appPaths struct{ Data, Cache, Config, State, Temp string }

func appHome() (appPaths, error) {
	home, e := os.UserHomeDir()
	if e != nil {
		return appPaths{}, e
	}
	choose := func(env, xdg, def string) string {
		if v := os.Getenv(env); v != "" {
			return v
		}
		if v := os.Getenv(xdg); v != "" {
			return filepath.Join(v, "doc-arquitetura")
		}
		return filepath.Join(home, def, "doc-arquitetura")
	}
	p := appPaths{
		Data:   choose("DOC_ARQUITETURA_DATA_DIR", "XDG_DATA_HOME", ".local/share"),
		Cache:  choose("DOC_ARQUITETURA_CACHE_DIR", "XDG_CACHE_HOME", ".cache"),
		Config: choose("DOC_ARQUITETURA_CONFIG_DIR", "XDG_CONFIG_HOME", ".config"),
		State:  choose("DOC_ARQUITETURA_STATE_DIR", "XDG_STATE_HOME", ".local/state"),
		Temp:   os.TempDir(),
	}
	if s := os.Getenv("DOC_ARQUITETURA_TEMP_DIR"); s != "" {
		p.Temp = s
	}
	for _, v := range []string{p.Data, p.Cache, p.Config, p.State, p.Temp} {
		if !filepath.IsAbs(v) {
			return appPaths{}, fmt.Errorf("diretorio precisa ser absoluto: %s", v)
		}
	}
	return p, nil
}
func mkdirPrivate(p string) error { return os.MkdirAll(p, 0700) }
func writeJSONAtomic(dest string, value interface{}) error {
	if e := mkdirPrivate(filepath.Dir(dest)); e != nil {
		return e
	}
	b, e := json.MarshalIndent(value, "", "  ")
	if e != nil {
		return e
	}
	f, e := os.CreateTemp(filepath.Dir(dest), ".tmp-*")
	if e != nil {
		return e
	}
	defer os.Remove(f.Name())
	if e = f.Chmod(0600); e != nil {
		f.Close()
		return e
	}
	if _, e = f.Write(append(b, '\n')); e != nil {
		f.Close()
		return e
	}
	if e = f.Sync(); e != nil {
		f.Close()
		return e
	}
	if e = f.Close(); e != nil {
		return e
	}
	return os.Rename(f.Name(), dest)
}
func readJSON(path string, v interface{}) error {
	b, e := os.ReadFile(path)
	if e != nil {
		return e
	}
	return json.Unmarshal(b, v)
}
func shaFile(path string) (string, int64, error) {
	f, e := os.Open(path)
	if e != nil {
		return "", 0, e
	}
	defer f.Close()
	h := sha256.New()
	n, e := io.Copy(h, f)
	if e != nil {
		return "", 0, e
	}
	return hex.EncodeToString(h.Sum(nil)), n, nil
}
func safeToken(s string) bool { return regexp.MustCompile(`^[a-z][a-z0-9._-]{0,49}$`).MatchString(s) }
func getModel(name string) (modelSpec, error) {
	for _, m := range modelCatalog {
		if m.ID == name {
			return m, nil
		}
	}
	return modelSpec{}, fmt.Errorf("modelo desconhecido: %s; consulte 'models catalog'", name)
}
func modelPath(p appPaths, id string) string { return filepath.Join(p.Data, "models", id+".gguf") }
func modelMeta(p appPaths, id string) string { return filepath.Join(p.Data, "models", id+".json") }
func loadModel(p appPaths, id string) (modelRecord, error) {
	var m modelRecord
	if !safeToken(id) {
		return m, errors.New("identificador invalido")
	}
	if e := readJSON(modelMeta(p, id), &m); e != nil {
		return m, e
	}
	if m.ID != id {
		return m, errors.New("metadados do modelo inconsistentes")
	}
	return m, nil
}
func verifyModel(p appPaths, id string) (modelRecord, error) {
	m, e := loadModel(p, id)
	if e != nil {
		return m, e
	}
	path := modelPath(p, id)
	sum, size, e := shaFile(path)
	if e != nil {
		return m, e
	}
	if sum != m.SHA256 || size != m.Size {
		return m, errors.New("modelo nao passou na verificacao SHA-256")
	}
	f, e := os.Open(path)
	if e != nil {
		return m, e
	}
	defer f.Close()
	buf := make([]byte, 4)
	if _, e = io.ReadFull(f, buf); e != nil {
		return m, e
	}
	if string(buf) != "GGUF" {
		return m, errors.New("arquivo nao e GGUF")
	}
	return m, nil
}
func hfRevision(client *http.Client, repo string) (string, error) {
	req, e := http.NewRequest("GET", "https://huggingface.co/api/models/"+repo, nil)
	if e != nil {
		return "", e
	}
	resp, e := client.Do(req)
	if e != nil {
		return "", e
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Hugging Face API HTTP %d", resp.StatusCode)
	}
	var v struct {
		SHA string `json:"sha"`
	}
	if e = json.NewDecoder(io.LimitReader(resp.Body, 4<<20)).Decode(&v); e != nil {
		return "", e
	}
	if !regexp.MustCompile(`^[a-f0-9]{40}$`).MatchString(v.SHA) {
		return "", errors.New("revisao Hugging Face invalida")
	}
	return v.SHA, nil
}
func downloadModel(p appPaths, id, expectedSHA string) error {
	spec, e := getModel(id)
	if e != nil {
		return e
	}
	if expectedSHA != "" && !regexp.MustCompile(`^[a-fA-F0-9]{64}$`).MatchString(expectedSHA) {
		return errors.New("--sha256 exige 64 caracteres hexadecimais")
	}
	if _, e := verifyModel(p, id); e == nil {
		fmt.Println("Modelo ja instalado e verificado:", modelPath(p, id))
		return nil
	}
	dir := filepath.Join(p.Data, "models")
	if e := mkdirPrivate(dir); e != nil {
		return e
	}
	client := &http.Client{Timeout: 2 * time.Hour}
	revision, e := hfRevision(client, spec.Repo)
	if e != nil {
		return e
	}
	u := "https://huggingface.co/" + spec.Repo + "/resolve/" + revision + "/" + url.PathEscape(spec.File)
	req, e := http.NewRequest("GET", u, nil)
	if e != nil {
		return e
	}
	resp, e := client.Do(req)
	if e != nil {
		return e
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("download HTTP %d", resp.StatusCode)
	}
	f, e := os.CreateTemp(dir, ".download-*.gguf")
	if e != nil {
		return e
	}
	defer os.Remove(f.Name())
	if e = f.Chmod(0600); e != nil {
		f.Close()
		return e
	}
	h := sha256.New()
	n, e := io.Copy(io.MultiWriter(f, h), io.LimitReader(resp.Body, 16<<30))
	if e != nil {
		f.Close()
		return e
	}
	if e = f.Sync(); e != nil {
		f.Close()
		return e
	}
	if e = f.Close(); e != nil {
		return e
	}
	if n < 8 {
		return errors.New("arquivo GGUF pequeno demais")
	}
	probe, e := os.Open(f.Name())
	if e != nil {
		return e
	}
	magic := make([]byte, 4)
	_, e = io.ReadFull(probe, magic)
	probe.Close()
	if e != nil {
		return e
	}
	if string(magic) != "GGUF" {
		return errors.New("arquivo baixado nao e GGUF")
	}
	sum := hex.EncodeToString(h.Sum(nil))
	if expectedSHA != "" && !strings.EqualFold(sum, expectedSHA) {
		return fmt.Errorf("checksum inesperado: %s", sum)
	}
	rec := modelRecord{ID: id, Repository: spec.Repo, Revision: revision, File: spec.File, SHA256: sum, Size: n, InstalledAt: time.Now().UTC().Format(time.RFC3339)}
	if e = os.Rename(f.Name(), modelPath(p, id)); e != nil {
		return e
	}
	if e = writeJSONAtomic(modelMeta(p, id), rec); e != nil {
		return e
	}
	fmt.Printf("Instalado %s (%d bytes, SHA-256 %s, revisao %s)\n", id, n, sum, revision)
	return nil
}
func modelAssignmentsPath(p appPaths) string { return filepath.Join(p.Config, "agents.json") }
func assignModel(p appPaths, agent, id string) error {
	if !safeToken(agent) {
		return errors.New("nome de agente invalido")
	}
	if _, e := verifyModel(p, id); e != nil {
		return fmt.Errorf("instale e verifique o modelo primeiro: %w", e)
	}
	v := map[string]string{}
	if e := readJSON(modelAssignmentsPath(p), &v); e != nil && !os.IsNotExist(e) {
		return e
	}
	v[agent] = id
	if e := writeJSONAtomic(modelAssignmentsPath(p), v); e != nil {
		return e
	}
	fmt.Printf("Agente %s => %s\n", agent, id)
	return nil
}
func modelsCommand(p appPaths, args []string) error {
	if len(args) == 0 {
		return errors.New("uso: models <catalog|download|list|assign|serve>")
	}
	sub, args := args[0], args[1:]
	switch sub {
	case "catalog":
		for _, m := range modelCatalog {
			fmt.Printf("%-16s %-40s %s\n", m.ID, m.Repo, m.Purpose)
		}
		return nil
	case "download":
		f := flag.NewFlagSet("models download", flag.ContinueOnError)
		sha := f.String("sha256", "", "checksum SHA-256 esperado (opcional)")
		if e := f.Parse(args); e != nil {
			return e
		}
		if f.NArg() != 1 {
			return errors.New("uso: models download [--sha256 HEX] <modelo>")
		}
		return downloadModel(p, f.Arg(0), *sha)
	case "list":
		for _, spec := range modelCatalog {
			m, e := verifyModel(p, spec.ID)
			if e == nil {
				fmt.Printf("INSTALLED %-17s %d bytes SHA-256 %s\n", m.ID, m.Size, m.SHA256)
			} else {
				fmt.Printf("MISSING   %-17s (%v)\n", spec.ID, e)
			}
		}
		return nil
	case "assign":
		if len(args) != 2 {
			return errors.New("uso: models assign <agente> <modelo>")
		}
		return assignModel(p, args[0], args[1])
	case "serve":
		f := flag.NewFlagSet("models serve", flag.ContinueOnError)
		port := f.Int("port", 8080, "porta local")
		contextSize := f.Int("context", 2048, "contexto em tokens")
		if e := f.Parse(args); e != nil {
			return e
		}
		if f.NArg() != 1 {
			return errors.New("uso: models serve [--port N] [--context N] <modelo>")
		}
		if *port < 1 || *port > 65535 || *contextSize < 256 || *contextSize > 32768 {
			return errors.New("porta/contexto invalidos")
		}
		id := f.Arg(0)
		if _, e := verifyModel(p, id); e != nil {
			return e
		}
		server, e := exec.LookPath("llama-server")
		if e != nil {
			return errors.New("llama-server ausente: instale llama.cpp separadamente")
		}
		cmd := exec.Command(server, "-m", modelPath(p, id), "--host", "127.0.0.1", "--port", strconv.Itoa(*port), "-c", strconv.Itoa(*contextSize))
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("comando models desconhecido: %s", sub)
	}
}
func gitOutput(ctx context.Context, dir string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")
	b, e := cmd.CombinedOutput()
	if e != nil {
		return "", fmt.Errorf("git %s: %v: %s", args[0], e, strings.TrimSpace(string(b)))
	}
	return strings.TrimSpace(string(b)), nil
}
func remoteNormSHA() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	out, e := gitOutput(ctx, "", "ls-remote", "--heads", normGitURL, "main")
	if e != nil {
		return "", e
	}
	fields := strings.Fields(out)
	if len(fields) != 2 || fields[1] != "refs/heads/main" || !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(fields[0]) {
		return "", errors.New("HEAD remoto main invalido")
	}
	return fields[0], nil
}
func refBase(p appPaths) string        { return filepath.Join(p.Cache, "reference") }
func refCurrentPath(p appPaths) string { return filepath.Join(refBase(p), "current.json") }
func readCurrentRef(p appPaths) (refRecord, error) {
	var v refRecord
	e := readJSON(refCurrentPath(p), &v)
	if e != nil {
		return v, e
	}
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(v.SHA) {
		return v, errors.New("cache da referencia invalido")
	}
	return v, nil
}
func referenceUpdate(p appPaths) error {
	remote, e := remoteNormSHA()
	if e != nil {
		return e
	}
	base := refBase(p)
	if e = mkdirPrivate(base); e != nil {
		return e
	}
	path := filepath.Join(base, remote)
	if _, e = os.Stat(filepath.Join(path, normTree, "README.md")); e == nil {
		if e = writeJSONAtomic(refCurrentPath(p), refRecord{SHA: remote, Source: normGitURL, UpdatedAt: time.Now().UTC().Format(time.RFC3339)}); e != nil {
			return e
		}
		fmt.Println("Referencia atualizada (cache):", remote)
		return nil
	}
	staging, e := os.MkdirTemp(base, ".reference-*")
	if e != nil {
		return e
	}
	defer os.RemoveAll(staging)
	cloneDir := filepath.Join(staging, "repository")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if _, e = gitOutput(ctx, "", "clone", "--depth", "1", "--single-branch", "--branch", "main", "--no-recurse-submodules", "--", normGitURL, cloneDir); e != nil {
		return e
	}
	actual, e := gitOutput(ctx, cloneDir, "rev-parse", "HEAD")
	if e != nil {
		return e
	}
	if _, e = os.Stat(filepath.Join(cloneDir, normTree, "README.md")); e != nil {
		return fmt.Errorf("referencia normativa ausente: %w", e)
	}
	target := filepath.Join(base, actual)
	if _, e = os.Stat(target); os.IsNotExist(e) {
		if e = os.Rename(cloneDir, target); e != nil {
			return e
		}
	} else if e != nil {
		return e
	}
	if e = writeJSONAtomic(refCurrentPath(p), refRecord{SHA: actual, Source: normGitURL, UpdatedAt: time.Now().UTC().Format(time.RFC3339)}); e != nil {
		return e
	}
	if actual != remote {
		fmt.Printf("Aviso: main mudou durante o download (%s -> %s); armazenada revisao real\n", remote, actual)
	}
	fmt.Println("Referencia baixada:", actual)
	return nil
}
func referenceStatus(p appPaths) error {
	r, e := readCurrentRef(p)
	if e != nil {
		if !os.IsNotExist(e) {
			return e
		}
		fmt.Println("Local: nao baixada")
	} else {
		fmt.Println("Local:", r.SHA)
	}
	remote, e := remoteNormSHA()
	if e != nil {
		fmt.Println("Remoto: indisponivel:", e)
		return nil
	}
	fmt.Println("Remoto main:", remote)
	if r.SHA == remote {
		fmt.Println("Estado: atualizado")
	} else {
		fmt.Println("Estado: desatualizado ou nao instalado")
	}
	return nil
}
func dslBase(p appPaths) string        { return filepath.Join(p.Cache, "dsl") }
func dslCurrentPath(p appPaths) string { return filepath.Join(dslBase(p), "current.json") }
func dslCompile(p appPaths) error {
	ref, e := readCurrentRef(p)
	if e != nil {
		return fmt.Errorf("execute reference update antes: %w", e)
	}
	src := filepath.Join(refBase(p), ref.SHA, normTree)
	if _, e = os.Stat(filepath.Join(src, "README.md")); e != nil {
		return fmt.Errorf("documentos nao disponiveis para %s", ref.SHA)
	}
	base := dslBase(p)
	if e = mkdirPrivate(base); e != nil {
		return e
	}
	dest := filepath.Join(base, ref.SHA)
	if _, e = os.Stat(filepath.Join(dest, "manifest.json")); e == nil {
		var old dslRecord
		if e = readJSON(filepath.Join(dest, "manifest.json"), &old); e == nil && old.SHA == ref.SHA {
			sum, _, er := shaFile(filepath.Join(dest, "reference.pcdsl"))
			if er == nil && sum == old.BundleSHA256 {
				if e = writeJSONAtomic(dslCurrentPath(p), old); e != nil {
					return e
				}
				fmt.Println("DSL valida no cache:", ref.SHA)
				return nil
			}
		}
	}
	staging, e := os.MkdirTemp(base, ".dsl-*")
	if e != nil {
		return e
	}
	defer os.RemoveAll(staging)
	bundle := filepath.Join(staging, "reference.pcdsl")
	count, e := convertCharter(src, bundle, "doc-arquitetura/main@"+ref.SHA)
	if e != nil {
		return e
	}
	sum, _, e := shaFile(bundle)
	if e != nil {
		return e
	}
	rec := dslRecord{SHA: ref.SHA, BundleSHA256: sum, Documents: count, CompiledAt: time.Now().UTC().Format(time.RFC3339)}
	if e = writeJSONAtomic(filepath.Join(staging, "manifest.json"), rec); e != nil {
		return e
	}
	if _, e = os.Stat(dest); e == nil {
		return fmt.Errorf("cache DSL para %s corrompido; remova-o manualmente", ref.SHA)
	} else if !os.IsNotExist(e) {
		return e
	}
	if e = os.Rename(staging, dest); e != nil {
		return e
	}
	if e = writeJSONAtomic(dslCurrentPath(p), rec); e != nil {
		return e
	}
	fmt.Printf("DSL compilada: %d documentos, commit %s, SHA-256 %s\n", count, ref.SHA, sum)
	return nil
}
func dslStatus(p appPaths) error {
	var d dslRecord
	e := readJSON(dslCurrentPath(p), &d)
	if e != nil {
		if os.IsNotExist(e) {
			fmt.Println("DSL nao compilada")
			return nil
		}
		return e
	}
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(d.SHA) {
		return errors.New("manifest DSL invalido")
	}
	sum, _, e := shaFile(filepath.Join(dslBase(p), d.SHA, "reference.pcdsl"))
	if e != nil {
		return e
	}
	if sum != d.BundleSHA256 {
		return errors.New("DSL alterada: checksum divergente")
	}
	ref, e := readCurrentRef(p)
	if e != nil {
		return e
	}
	fmt.Printf("DSL: %s, %d documentos, SHA-256 %s\n", d.SHA, d.Documents, d.BundleSHA256)
	if d.SHA == ref.SHA {
		fmt.Println("Estado: compativel com a referencia local")
	} else {
		fmt.Println("Estado: desatualizada (referencia local mudou)")
	}
	return nil
}
func managedCommand(args []string) error {
	if len(args) < 2 {
		return errors.New("uso: <models|reference|dsl> <subcomando> [opcoes]")
	}
	p, e := appHome()
	if e != nil {
		return e
	}
	switch args[0] {
	case "models":
		return modelsCommand(p, args[1:])
	case "reference":
		if len(args) != 2 {
			return errors.New("uso: reference <update|status>")
		}
		switch args[1] {
		case "update":
			return referenceUpdate(p)
		case "status":
			return referenceStatus(p)
		}
	case "dsl":
		if len(args) != 2 {
			return errors.New("uso: dsl <update|compile|status>")
		}
		switch args[1] {
		case "update":
			if e := referenceUpdate(p); e != nil {
				return e
			}
			return dslCompile(p)
		case "compile":
			return dslCompile(p)
		case "status":
			return dslStatus(p)
		}
	}
	return fmt.Errorf("comando desconhecido: %s %s", args[0], args[1])
}
