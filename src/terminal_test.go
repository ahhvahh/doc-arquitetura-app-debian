package main

import (
    "bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNormalizeGitHubURL(t *testing.T) {
	good, err := normalizeGitHubURL("https://github.com/ahhvahh/projeto")
	if err != nil || good != "https://github.com/ahhvahh/projeto.git" {
		t.Fatal(good, err)
	}
	for _, v := range []string{"file:///tmp/r", "https://evil.com/x/y", "https://github.com/x/y?token=secret", "https://github.com/x/y#ref", "https://github.com/x/y/../../bad", "https://github.com/x"} {
		if _, err := normalizeGitHubURL(v); err == nil {
			t.Errorf("aceitou %q", v)
		}
	}
}
func TestValidBranch(t *testing.T) {
	for _, b := range []string{"develop", "feature/test"} {
		if !validGitBranch(b) {
			t.Error(b)
		}
	}
	for _, b := range []string{"../../main", "a..b", "@{x}", "bad space", "x.lock"} {
		if validGitBranch(b) {
			t.Error("invalid:", b)
		}
	}
}
func TestDocPathAndSymlink(t *testing.T) {
	root := t.TempDir()
	s := terminalSession{Workspace: root, Area: "scope"}
	base := "docs/scope/"
	for _, p := range []string{base + "README.md", base + "types/TYPE-001.md"} {
		if !isDocPath(s, p) {
			t.Errorf("rejected %s", p)
		}
	}
	for _, p := range []string{"src/main.go", "docs/charter/foo.md", base + "../../charter/x.md", base + "x.txt", "docs/ProjectCharter/1.0/scope/wrong.md", "/tmp/file.md"} {
		if isDocPath(s, p) {
			t.Errorf("accepted %s", p)
		}
	}
	if err := os.MkdirAll(filepath.Join(root, "docs"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(t.TempDir(), filepath.Join(root, "docs", "scope")); err != nil {
		t.Fatal(err)
	}
	if _, err := safeDocFile(s, base+"README.md"); err == nil {
		t.Fatal("symlink accepted")
	}
}
func TestDiffForNewFile(t *testing.T) {
	root := t.TempDir()
	session := terminalSession{Workspace: root, Area: "charter"}
	var output bytes.Buffer
	app := terminalApp{out: &output, paths: appPaths{Temp: root}, session: session, proposals: []documentEdit{{Path: "docs/charter/TEST.md", Content: "# Test\n", Reason: "new"}}}
	if err := app.diffProposals(); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "NOVO ARQUIVO") {
		t.Fatal(output.String())
	}
}
func TestValidationRejectsOtherSteps(t *testing.T) {
	root := t.TempDir()
	app := terminalApp{session: terminalSession{Workspace: root, Area: "scope", Branch: "develop"}}
	if _, err := app.changedPaths(); err == nil {
		t.Fatal("should fail for non-git root")
	}
}

func TestLegacyDocumentationLocations(t *testing.T) {
    root := t.TempDir()
    old := filepath.Join(root, "docs", "ProjectCharter", "1.0", "scope")
    if err := os.MkdirAll(old, 0700); err != nil { t.Fatal(err) }
    found := legacyDocumentationLocations(root, "scope")
    if len(found) != 1 || found[0] != "docs/ProjectCharter/1.0/scope" {
        t.Fatalf("legacy folders: %v", found)
    }
    if got := legacyDocumentationLocations(root, "charter"); len(got) != 0 {
        t.Fatalf("unrelated area: %v", got)
    }
}

func TestBackupPathNotAgentEditable(t *testing.T) {
    s:=terminalSession{Workspace:t.TempDir(),Area:"scope"}
    for _,path:=range []string{"docs/backup/a/manifest.json","docs/backup/a/original/docs/scope/README.md","docs/charter/README.md"}{
        if isDocPath(s,path) {t.Fatalf("agente conseguiu escrever em %s",path)}
    }
}

func TestCharterLegacyContextForOrdinaryReview(t *testing.T) {
    root:=t.TempDir()
    old:=filepath.Join(root,"docs","01-Project_Charter","charter","README.md")
    if err:=os.MkdirAll(filepath.Dir(old),0700);err!=nil{t.Fatal(err)}
    if err:=os.WriteFile(old,[]byte("# Documento legado de Charter\\n\\nObjetivo preexistente do sistema.\\n"),0600);err!=nil{t.Fatal(err)}
    app:=terminalApp{session:terminalSession{Workspace:root,Area:"charter",Prompt:"charter-refinement.md"}}
    legacy,err:=app.legacyContext()
    if err!=nil {t.Fatal(err)}
    if !strings.Contains(legacy,"docs/01-Project_Charter/charter/README.md") ||
       !strings.Contains(legacy,"Objetivo preexistente") {
        t.Fatalf("documentação antiga omitida do contexto: %q",legacy)
    }
}

func TestCloseSessionRetainsWorkspaceAndHistory(t *testing.T) {
    root:=t.TempDir()
    workspace:=filepath.Join(root,"workspaces","session-closed","repository")
    if err:=os.MkdirAll(filepath.Join(workspace,".git"),0700);err!=nil{t.Fatal(err)}
    var output bytes.Buffer
    app:=terminalApp{
        paths:appPaths{State:root},
        out:&output,
        session:terminalSession{
            ID:"session-closed",Repo:"https://github.com/ahhvahh/a-bioma-sauva.git",
            Branch:"ajusteDoc",Area:"charter",Workspace:workspace,
        },
        proposals:[]documentEdit{{Path:"docs/charter/test.md"}},
    }
    if err:=app.saveSession();err!=nil{t.Fatal(err)}
    if err:=app.closeSession();err!=nil{t.Fatal(err)}
    if app.session.ClosedAt=="" {t.Fatal("close timestamp missing")}
    if len(app.proposals)!=0{t.Fatal("pending proposals retained after closing")}
    if _,err:=os.Stat(filepath.Join(workspace,".git"));err!=nil{t.Fatal("workspace deleted",err)}
    var saved terminalSession
    if err:=readJSON(filepath.Join(root,"sessions","session-closed.json"),&saved);err!=nil{t.Fatal(err)}
    if saved.ClosedAt==""{t.Fatal("close timestamp not persisted")}
    if err:=app.closeSession();err==nil{t.Fatal("closed session accepted twice")}
    // A closed session must no longer be listed for resumption.
    app.in=bufio.NewReader(strings.NewReader("1\n"))
    if err:=app.selectSession();err==nil{t.Fatal("closed session could be resumed")}
}
func TestCloseSavedSessionConfirmation(t *testing.T) {
    root:=t.TempDir()
    workspace:=filepath.Join(root,"workspaces","session-open","repository")
    if err:=os.MkdirAll(filepath.Join(workspace,".git"),0700);err!=nil{t.Fatal(err)}
    var output bytes.Buffer
    app:=terminalApp{paths:appPaths{State:root},out:&output,
        session:terminalSession{ID:"session-open",Repo:"https://github.com/ahhvahh/a-bioma-sauva.git",
        Branch:"develop",Area:"charter",Workspace:workspace}}
    if err:=app.saveSession();err!=nil{t.Fatal(err)}
    app.in=bufio.NewReader(strings.NewReader("1\nNAO\n"))
    if err:=app.closeSavedSession();err!=nil{t.Fatal(err)}
    var saved terminalSession
    if err:=readJSON(filepath.Join(root,"sessions","session-open.json"),&saved);err!=nil{t.Fatal(err)}
    if saved.ClosedAt!=""{t.Fatal("session closed without confirmation")}
    app.in=bufio.NewReader(strings.NewReader("1\nSIM\n"))
    if err:=app.closeSavedSession();err!=nil{t.Fatal(err)}
    if err:=readJSON(filepath.Join(root,"sessions","session-open.json"),&saved);err!=nil{t.Fatal(err)}
    if saved.ClosedAt==""{t.Fatal("confirmed closure not persisted")}
}

 
func TestTerminalChatReportsInferenceUnavailable(t *testing.T) {
    var output bytes.Buffer
    app := terminalApp{
        in: bufio.NewReader(strings.NewReader("analisar documentação\n/exit\n")),
        out: &output,
        session: terminalSession{Repo: "https://github.com/owner/repo.git", Branch: "develop", Area: "charter"},
    }
    if err := app.chat(); err != nil { t.Fatal(err) }
    if !strings.Contains(output.String(), "Análise automática local foi removida") ||
       !strings.Contains(output.String(), "processamento OpenAI ainda não está disponível") {
        t.Fatalf("chat must clearly report that no model inference occurs: %s", output.String())
    }
}
