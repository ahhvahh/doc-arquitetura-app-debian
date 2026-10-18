package main

import (
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
	base := "docs/ProjectCharter/1.0/scope/"
	for _, p := range []string{base + "README.md", base + "types/TYPE-001.md"} {
		if !isDocPath(s, p) {
			t.Errorf("rejected %s", p)
		}
	}
	for _, p := range []string{"src/main.go", "docs/ProjectCharter/1.0/charter/foo.md", base + "../../charter/x.md", base + "x.txt", "/tmp/file.md"} {
		if isDocPath(s, p) {
			t.Errorf("accepted %s", p)
		}
	}
	if err := os.MkdirAll(filepath.Join(root, "docs", "ProjectCharter", "1.0"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(t.TempDir(), filepath.Join(root, "docs", "ProjectCharter", "1.0", "scope")); err != nil {
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
	app := terminalApp{out: &output, paths: appPaths{Temp: root}, session: session, proposals: []documentEdit{{Path: "docs/ProjectCharter/1.0/charter/TEST.md", Content: "# Test\n", Reason: "new"}}}
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
