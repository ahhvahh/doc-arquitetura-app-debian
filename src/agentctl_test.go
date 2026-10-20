package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func isolatedPaths(t *testing.T) appPaths {
	t.Helper()
	root := t.TempDir()
	p := appPaths{Cache: filepath.Join(root, "cache"), Config: filepath.Join(root, "config"), State: filepath.Join(root, "state"), Temp: filepath.Join(root, "temp")}
	for _, d := range []string{p.Cache, p.Config, p.State, p.Temp} {
		if e := os.MkdirAll(d, 0700); e != nil {
			t.Fatal(e)
		}
	}
	return p
}
func TestPathsRespectOverrides(t *testing.T) {
	root := t.TempDir()
	t.Setenv("DOC_ARQUITETURA_CACHE_DIR", filepath.Join(root, "cache"))
	t.Setenv("DOC_ARQUITETURA_CONFIG_DIR", filepath.Join(root, "config"))
	t.Setenv("DOC_ARQUITETURA_STATE_DIR", filepath.Join(root, "state"))
	t.Setenv("DOC_ARQUITETURA_TEMP_DIR", filepath.Join(root, "tmp"))
	p, e := appHome()
	if e != nil {
		t.Fatal(e)
	}
	if p.Config != filepath.Join(root, "config") || p.Temp != filepath.Join(root, "tmp") {
		t.Fatalf("paths %+v", p)
	}
}
func TestDSLCompileCacheAndIntegrity(t *testing.T) {
	p := isolatedPaths(t)
	sha := strings.Repeat("a", 40)
	src := filepath.Join(refBase(p), sha, normTree)
	if e := os.MkdirAll(src, 0700); e != nil {
		t.Fatal(e)
	}
	if e := os.WriteFile(filepath.Join(src, "README.md"), []byte("# Normative reference\n"), 0600); e != nil {
		t.Fatal(e)
	}
	if e := writeJSONAtomic(refCurrentPath(p), refRecord{SHA: sha, Source: normGitURL}); e != nil {
		t.Fatal(e)
	}
	if e := dslCompile(p); e != nil {
		t.Fatal(e)
	}
	if e := dslStatus(p); e != nil {
		t.Fatal(e)
	}
	var rec dslRecord
	if e := readJSON(dslCurrentPath(p), &rec); e != nil {
		t.Fatal(e)
	}
	if rec.SHA != sha || rec.Documents != 1 || len(rec.BundleSHA256) != 64 {
		t.Fatalf("manifest %+v", rec)
	}
	if e := dslCompile(p); e != nil {
		t.Fatal("cache not reusable:", e)
	}
	bundle := filepath.Join(dslBase(p), sha, "reference.pcdsl")
	if e := os.WriteFile(bundle, []byte("tampered"), 0600); e != nil {
		t.Fatal(e)
	}
	if e := dslStatus(p); e == nil {
		t.Fatal("tampered DSL verified")
	}
}
