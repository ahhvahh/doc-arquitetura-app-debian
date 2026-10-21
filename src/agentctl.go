package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const normGitURL = "https://github.com/ahhvahh/doc-arquitetura.git"
const normTree = "ProjectCharter/1.0"

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

type appPaths struct{ Cache, Config, State, Temp string }

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
		Cache:  choose("DOC_ARQUITETURA_CACHE_DIR", "XDG_CACHE_HOME", ".cache"),
		Config: choose("DOC_ARQUITETURA_CONFIG_DIR", "XDG_CONFIG_HOME", ".config"),
		State:  choose("DOC_ARQUITETURA_STATE_DIR", "XDG_STATE_HOME", ".local/state"),
		Temp:   os.TempDir(),
	}
	if s := os.Getenv("DOC_ARQUITETURA_TEMP_DIR"); s != "" {
		p.Temp = s
	}
	for _, v := range []string{p.Cache, p.Config, p.State, p.Temp} {
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
func gitOutput(ctx context.Context, dir string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	gitEnv, err := githubGitEnv()
	if err != nil {return "", err}
	cmd.Env = gitEnv
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
		return errors.New("uso: <reference|dsl> <subcomando> [opcoes]")
	}
	p, e := appHome()
	if e != nil {
		return e
	}
	switch args[0] {
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
			return dslReleaseUpdate(p)
		case "compile":
			return dslCompile(p)
		case "status":
			return dslReleaseStatus(p)
		}
	}
	return fmt.Errorf("comando desconhecido: %s %s", args[0], args[1])
}
