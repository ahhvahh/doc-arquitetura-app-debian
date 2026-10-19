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
	p := appPaths{Data: filepath.Join(root, "data"), Cache: filepath.Join(root, "cache"), Config: filepath.Join(root, "config"), State: filepath.Join(root, "state"), Temp: filepath.Join(root, "temp")}
	for _, d := range []string{p.Data, p.Cache, p.Config, p.State, p.Temp} {
		if e := os.MkdirAll(d, 0700); e != nil {
			t.Fatal(e)
		}
	}
	return p
}
func TestPathsRespectOverrides(t *testing.T) {
	root := t.TempDir()
	t.Setenv("DOC_ARQUITETURA_DATA_DIR", filepath.Join(root, "models"))
	t.Setenv("DOC_ARQUITETURA_CACHE_DIR", filepath.Join(root, "cache"))
	t.Setenv("DOC_ARQUITETURA_CONFIG_DIR", filepath.Join(root, "config"))
	t.Setenv("DOC_ARQUITETURA_STATE_DIR", filepath.Join(root, "state"))
	t.Setenv("DOC_ARQUITETURA_TEMP_DIR", filepath.Join(root, "tmp"))
	p, e := appHome()
	if e != nil {
		t.Fatal(e)
	}
	if p.Data != filepath.Join(root, "models") || p.Temp != filepath.Join(root, "tmp") {
		t.Fatalf("paths %+v", p)
	}
}
func TestModelValidationAndAssignment(t *testing.T) {
	p := isolatedPaths(t)
	spec, _ := getModel("qwen3.5-0.8b")
	if _, e := getModel("arbitrary-name"); e == nil {
		t.Fatal("unknown model accepted")
	}
	payload := []byte("GGUFtest-metadata")
	path := modelPath(p, spec.ID)
	if e := os.MkdirAll(filepath.Dir(path), 0700); e != nil {
		t.Fatal(e)
	}
	if e := os.WriteFile(path, payload, 0600); e != nil {
		t.Fatal(e)
	}
	sha, size, e := shaFile(path)
	if e != nil {
		t.Fatal(e)
	}
	rec := modelRecord{ID: spec.ID, SHA256: sha, Size: size}
	if e = writeJSONAtomic(modelMeta(p, spec.ID), rec); e != nil {
		t.Fatal(e)
	}
	if _, e = verifyModel(p, spec.ID); e != nil {
		t.Fatal(e)
	}
	if e = assignModel(p, "triagem", spec.ID); e != nil {
		t.Fatal(e)
	}
	var saved map[string]string
	if e = readJSON(modelAssignmentsPath(p), &saved); e != nil {
		t.Fatal(e)
	}
	if saved["triagem"] != spec.ID {
		t.Fatalf("assignment %+v", saved)
	}
	if e = os.WriteFile(path, []byte("BAD!"), 0600); e != nil {
		t.Fatal(e)
	}
	if _, e = verifyModel(p, spec.ID); e == nil {
		t.Fatal("corrupted model verified")
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
func TestBadToken(t *testing.T) {
	for _, id := range []string{"../other", "", "A/Model", "../../home"} {
		if safeToken(id) {
			t.Errorf("accepted %q", id)
		}
	}
}

func TestModelCatalogGGUFFilenames(t *testing.T) {
    tests := []struct{
        id   string
        repo string
        file string
    }{
        {"qwen3.5-0.8b", "ggml-org/Qwen3.5-0.8B-GGUF", "Qwen3.5-0.8B-Q4_0.gguf"},
        {"qwen3-1.7b", "bartowski/Qwen_Qwen3-1.7B-GGUF", "Qwen_Qwen3-1.7B-Q4_K_M.gguf"},
    }
    for _, tc := range tests {
        t.Run(tc.id, func(t *testing.T) {
            model, err := getModel(tc.id)
            if err != nil {
                t.Fatal(err)
            }
            if model.Repo != tc.repo || model.File != tc.file {
                t.Fatalf("GGUF catalog mismatched: repo=%q file=%q; want repo=%q file=%q", model.Repo, model.File, tc.repo, tc.file)
            }
        })
    }
}

func TestAdditionalReviewModelsInCatalog(t *testing.T) {
    for _,tc:=range []struct{ID,Repo,File string}{
        {"gemma3-1b","ggml-org/gemma-3-1b-it-GGUF","gemma-3-1b-it-Q4_K_M.gguf"},
        {"qwen3.5-4b","lmstudio-community/Qwen3.5-4B-GGUF","Qwen3.5-4B-Q4_K_M.gguf"},
    }{
        model,err:=getModel(tc.ID)
        if err!=nil||model.Repo!=tc.Repo||model.File!=tc.File{
            t.Errorf("catalog mismatch for %s: %+v (%v)",tc.ID,model,err)
        }
    }
}
