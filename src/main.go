package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Check struct{ ID, Severity, Kind, Arg, Message string }
type DSL struct {
	Base     string
	Policies map[string]string
	Routes   map[string]string
	Checks   []Check
}
type Finding struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Status   string `json:"status"`
	Path     string `json:"path,omitempty"`
	Detail   string `json:"detail"`
}
type Report struct {
	Root     string         `json:"root"`
	Base     string         `json:"base"`
	Verdict  string         `json:"verdict"`
	Findings []Finding      `json:"findings"`
	Stats    map[string]int `json:"stats"`
}

func parseDSL(path string) (DSL, error) {
	b, e := os.ReadFile(path)
	if e != nil {
		return DSL{}, e
	}
	d := DSL{Policies: map[string]string{}, Routes: map[string]string{}}
	lines := strings.Split(string(b), "\n")
	seen := false
	kinds := map[string]bool{"exists": true, "min_docs": true, "statuses": true, "no_pending": true, "links": true, "badges": true, "unique_ids": true, "semantic": true}
	for i, line := range lines {
		s := strings.TrimSpace(line)
		if s == "" || strings.HasPrefix(s, "#") {
			continue
		}
		if !seen {
			if s != "PCDSL 1" {
				return d, fmt.Errorf("linha %d: versao DSL invalida", i+1)
			}
			seen = true
			continue
		}
		switch {
		case strings.HasPrefix(s, "BASE "):
			d.Base = strings.TrimSpace(strings.TrimPrefix(s, "BASE "))
		case strings.HasPrefix(s, "POL "):
			z := strings.SplitN(strings.TrimPrefix(s, "POL "), "=", 2)
			if len(z) != 2 {
				return d, fmt.Errorf("linha %d: POL invalida", i+1)
			}
			d.Policies[z[0]] = z[1]
		case strings.HasPrefix(s, "ROUTE "):
			z := strings.Split(strings.TrimPrefix(s, "ROUTE "), "|")
			if len(z) != 2 {
				return d, fmt.Errorf("linha %d: ROUTE invalida", i+1)
			}
			d.Routes[z[0]] = z[1]
		case strings.HasPrefix(s, "CHECK "):
			z := strings.Split(strings.TrimPrefix(s, "CHECK "), "|")
			if len(z) != 5 {
				return d, fmt.Errorf("linha %d: CHECK requer 5 campos", i+1)
			}
			c := Check{z[0], z[1], z[2], z[3], z[4]}
			if c.Severity != "BLOCKER" && c.Severity != "REVIEW" || !kinds[c.Kind] {
				return d, fmt.Errorf("linha %d: severidade/validador desconhecido", i+1)
			}
			d.Checks = append(d.Checks, c)
		default:
			return d, fmt.Errorf("linha %d: operacao desconhecida", i+1)
		}
	}
	if !seen || d.Base == "" || len(d.Checks) == 0 {
		return d, errors.New("DSL incompleta")
	}
	if filepath.IsAbs(d.Base) || strings.HasPrefix(filepath.Clean(d.Base), "..") {
		return d, errors.New("BASE deve ser relativa")
	}
	return d, nil
}

var badgeType = regexp.MustCompile(`badge-type-[a-z0-9-]+\.svg`)
var badgeStatus = regexp.MustCompile(`badge-status-([a-z]+)\.svg`)
var badgeVersion = regexp.MustCompile(`badge-version-[a-z0-9-]+\.svg`)
var linkRe = regexp.MustCompile(`!?\[[^\]]*\]\(([^\s\)]+)(?:\s+[^)]*)?\)`)
var idRe = regexp.MustCompile(`^(OBJ|NOBJ|CON|ASM|OPEN|DSC|PCH)-\d{3,}`)

func isContract(relative string) bool {
	name := filepath.Base(relative)
	if !strings.HasSuffix(name, ".md") || strings.Contains(strings.ToLower(name), "exemplo") || strings.Contains(strings.ToLower(name), "example") {
		return false
	}
	return idRe.MatchString(name)
}
func markdowns(root string) ([]string, error) {
	var paths []string
	e := filepath.WalkDir(root, func(p string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if d.IsDir() {
			if d.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(strings.ToLower(p), ".md") {
			paths = append(paths, p)
		}
		return nil
	})
	sort.Strings(paths)
	return paths, e
}
func rel(base, p string) string {
	x, e := filepath.Rel(base, p)
	if e != nil {
		return p
	}
	return filepath.ToSlash(x)
}
func countDocs(base, glob string) int {
	matches, _ := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
	n := 0
	for _, f := range matches {
		if isContract(rel(base, f)) {
			n++
		}
	}
	return n
}
func checkLinks(base string) ([]string, error) {
	files, e := markdowns(base)
	if e != nil {
		return nil, e
	}
	var problems []string
	for _, file := range files {
		b, e := os.ReadFile(file)
		if e != nil {
			return nil, e
		}
		for _, m := range linkRe.FindAllStringSubmatch(string(b), -1) {
			link := m[1]
			if strings.HasPrefix(link, "#") || strings.HasPrefix(link, "http:") || strings.HasPrefix(link, "https:") || strings.HasPrefix(link, "mailto:") || strings.HasPrefix(link, "data:") {
				continue
			}
			decoded, e := url.PathUnescape(strings.Split(link, "#")[0])
			if e != nil {
				problems = append(problems, rel(base, file)+": URL invalida")
				continue
			}
			if decoded == "" {
				continue
			}
			target := filepath.Clean(filepath.Join(filepath.Dir(file), decoded))
			// Permit links a arquivos de outras pastas do mesmo repositorio; nao executamos conteudo.
			if _, e := os.Stat(target); e != nil {
				problems = append(problems, rel(base, file)+" -> "+link)
			}
		}
	}
	return problems, nil
}
func runChecks(repo string, d DSL) Report {
	base := filepath.Join(repo, filepath.FromSlash(d.Base))
	r := Report{Root: repo, Base: d.Base, Stats: map[string]int{}}
	for _, c := range d.Checks {
		f := Finding{ID: c.ID, Severity: c.Severity, Status: "PASS", Detail: c.Message}
		checkPath := filepath.Clean(filepath.Join(base, filepath.FromSlash(c.Arg)))
		if c.Kind != "min_docs" && c.Kind != "semantic" && checkPath != base && !strings.HasPrefix(checkPath, base+string(filepath.Separator)) {
			f.Status = "BLOCKED"
			f.Detail = "Caminho DSL fora da BASE"
			r.Findings = append(r.Findings, f)
			continue
		}
		switch c.Kind {
		case "exists":
			if _, e := os.Stat(checkPath); e != nil {
				f.Status = "FAIL"
				f.Path = c.Arg
				f.Detail = "Arquivo obrigatorio ausente"
			}
		case "min_docs":
			z := strings.Split(c.Arg, ":")
			if len(z) != 2 {
				f.Status = "FAIL"
				f.Detail = "Sintaxe min_docs invalida"
				break
			}
			min, e := strconv.Atoi(z[1])
			if e != nil {
				f.Status = "FAIL"
				break
			}
			n := countDocs(base, z[0])
			if n < min {
				f.Status = "FAIL"
				f.Path = z[0]
				f.Detail = fmt.Sprintf("Encontrados %d; minimo %d", n, min)
			}
		case "statuses":
			files, e := markdowns(base)
			if e != nil {
				f.Status = "FAIL"
				f.Detail = e.Error()
				break
			}
			var issues []string
			for _, p := range files {
				if !isContract(rel(base, p)) {
					continue
				}
				b, _ := os.ReadFile(p)
				m := badgeStatus.FindStringSubmatch(string(b))
				if len(m) < 2 || (m[1] != "refined" && m[1] != "implemented") {
					issues = append(issues, rel(base, p))
				}
			}
			if len(issues) > 0 {
				f.Status = "FAIL"
				f.Detail = "Status pendente/invalido: " + strings.Join(issues, ", ")
			}
		case "no_pending":
			b, e := os.ReadFile(checkPath)
			if e != nil {
				f.Status = "FAIL"
				f.Detail = "Arquivo ausente"
				break
			}
			if strings.Contains(strings.ToUpper(string(b)), "PENDING") {
				f.Status = "FAIL"
				f.Path = c.Arg
				f.Detail = "README possui PENDING"
			}
		case "links":
			bad, e := checkLinks(base)
			if e != nil {
				f.Status = "FAIL"
				f.Detail = e.Error()
			} else if len(bad) > 0 {
				f.Status = "FAIL"
				f.Detail = fmt.Sprintf("%d links quebrados: %s", len(bad), strings.Join(bad, "; "))
			}
		case "badges":
			files, e := markdowns(base)
			if e != nil {
				f.Status = "FAIL"
				f.Detail = e.Error()
				break
			}
			var issues []string
			for _, p := range files {
				if !isContract(rel(base, p)) && rel(base, p) != "README.md" {
					continue
				}
				b, _ := os.ReadFile(p)
				s := string(b)
				if len(badgeType.FindAllString(s, -1)) != 1 || len(badgeStatus.FindAllString(s, -1)) != 1 || len(badgeVersion.FindAllString(s, -1)) != 1 {
					issues = append(issues, rel(base, p))
				}
			}
			if len(issues) > 0 {
				f.Status = "FAIL"
				f.Detail = "Badges ausentes/duplicados: " + strings.Join(issues, ", ")
			}
		case "unique_ids":
			files, e := markdowns(base)
			if e != nil {
				f.Status = "FAIL"
				f.Detail = e.Error()
				break
			}
			seen := map[string]string{}
			var dup []string
			for _, p := range files {
				if !isContract(rel(base, p)) {
					continue
				}
				id := idRe.FindString(filepath.Base(p))
				if prior, ok := seen[id]; ok {
					dup = append(dup, id+" ("+prior+", "+rel(base, p)+")")
				} else {
					seen[id] = rel(base, p)
				}
			}
			if len(dup) > 0 {
				f.Status = "FAIL"
				f.Detail = "IDs repetidos: " + strings.Join(dup, "; ")
			}
		case "semantic":
			f.Status = "REVIEW"
			f.Path = c.Arg
			f.Detail = "Necessita revisao semantica e evidencia: " + c.Message
		}
		r.Findings = append(r.Findings, f)
		r.Stats[f.Status]++
	}
	if r.Stats["FAIL"] > 0 || r.Stats["BLOCKED"] > 0 {
		r.Verdict = "BLOCKED"
	} else {
		r.Verdict = "NEEDS_SEMANTIC_REVIEW"
	}
	return r
}
func printReport(r Report, asJSON bool) {
	if asJSON {
		b, _ := json.MarshalIndent(r, "", "  ")
		fmt.Println(string(b))
		return
	}
	fmt.Printf("Projeto: %s\nDSL base: %s\nGate preliminar: %s\n\n", r.Root, r.Base, r.Verdict)
	for _, f := range r.Findings {
		fmt.Printf("%-9s %-9s %-8s %s\n", f.ID, f.Severity, f.Status, f.Detail)
	}
	fmt.Println("\nO gate definitivo depende de charter-review.md, fontes normativas e decisao do usuario.")
}
func cloneRepo(source, branch, dest string) error {
	if branch != "develop" {
		return errors.New("o prototipo clona somente develop")
	}
	parsed, e := url.Parse(source)
	if e != nil {
		return e
	}
	if parsed.Scheme != "https" || parsed.Host != "github.com" || parsed.User != nil || parsed.RawQuery != "" || parsed.Fragment != "" {
		return errors.New("aceita apenas URL https://github.com/owner/repo")
	}
	path := strings.TrimSuffix(strings.Trim(parsed.Path, "/"), ".git")
	parts := strings.Split(path, "/")
	valid := regexp.MustCompile(`^[A-Za-z0-9_.-]+$`)
	if len(parts) != 2 || !valid.MatchString(parts[0]) || !valid.MatchString(parts[1]) || parts[0] == ".." || parts[1] == ".." {
		return errors.New("proprietario/repositorio invalido")
	}
	abs, e := filepath.Abs(dest)
	if e != nil {
		return e
	}
	if _, e := os.Stat(abs); e == nil {
		return errors.New("destino ja existe; nao sera sobrescrito")
	}
	if e:=os.MkdirAll(filepath.Dir(abs),0755);e!=nil{return e}
	cmd, e := githubGitCommand(context.Background(),"-c", "protocol.file.allow=never", "clone", "--depth", "1", "--single-branch", "--branch", "develop", "--no-recurse-submodules", "--", "https://github.com/"+path+".git", abs)
	if e!=nil{return e}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func defaultDSLPath() string {
 if override := os.Getenv("DOC_ARQUITETURA_DSL"); override != "" {
  return override
 }
 const installed = "/usr/share/doc-arquitetura-app/dsl/charter.pcdsl"
 if _, err := os.Stat(installed); err == nil {
  return installed
 }
 return "dsl/charter.pcdsl"
}
func main() {
 if os.Getenv("DOC_ARQUITETURA_GITHUB_ASKPASS")=="1" {
     if err:=githubAskPass();err!=nil{os.Exit(1)}
     return
 }
 if len(os.Args)>1 && os.Args[1]=="openai" {
     if err:=openAICommand(os.Args[2:]);err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(1)}
     return
 }
 if len(os.Args)>1 && os.Args[1]=="github" {
     if err:=githubCommand(os.Args[2:]);err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(1)}
     return
 }
 if len(os.Args)==1 || os.Args[1]=="terminal" {
  var args []string
  if len(os.Args)>1 { args=os.Args[2:] }
  if err:=runTerminal(args);err!=nil {fmt.Fprintln(os.Stderr,err);os.Exit(1)}
  return
 }
 if len(os.Args)>1 && (os.Args[1]=="reference" || os.Args[1]=="dsl") {
  if err:=managedCommand(os.Args[1:]);err!=nil {fmt.Fprintln(os.Stderr,err);os.Exit(1)}
  return
 }

 if len(os.Args)>1 && (os.Args[1]=="sync" || os.Args[1]=="convert") {
  if err:=runExtra(os.Args[1:]);err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(1)};return
 }

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Uso: doc-arquitetura [terminal] | <github|check|analyze|clone|routes|reference|dsl> [opcoes]")
		os.Exit(2)
	}
	command := os.Args[1]
	fs := flag.NewFlagSet(command, flag.ExitOnError)
	repo := fs.String("repo", ".", "diretorio de projeto cliente")
	dsl := fs.String("dsl", defaultDSLPath(), "regras DSL")
	jsonFlag := fs.Bool("json", false, "relatorio JSON")
	gitURL := fs.String("url", "", "URL GitHub do cliente")
	destination := fs.String("dest", "", "destino de clonagem")
	_ = fs.Parse(os.Args[2:])
	if command == "clone" {
		if *destination == "" || *gitURL == "" {
			fmt.Fprintln(os.Stderr, "clone exige -url e -dest")
			os.Exit(2)
		}
		if e := cloneRepo(*gitURL, "develop", *destination); e != nil {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(1)
		}
		return
	}
	d, e := parseDSL(*dsl)
	if e != nil {
		fmt.Fprintln(os.Stderr, "DSL:", e)
		os.Exit(2)
	}
	switch command {
	case "routes":
		keys := make([]string, 0, len(d.Routes))
		for k := range d.Routes {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Printf("%-15s %s\n", k, d.Routes[k])
		}
	case "check", "analyze":
		root, e := filepath.Abs(*repo)
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			os.Exit(2)
		}
		report := runChecks(root, d)
		printReport(report, *jsonFlag)

	default:
		fmt.Fprintln(os.Stderr, "Comando desconhecido:", command)
		os.Exit(2)
	}
}