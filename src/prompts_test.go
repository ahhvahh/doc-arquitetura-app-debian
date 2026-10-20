package main

import (
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func TestStagePromptCatalogUsesVersionedMain(t *testing.T){
    cache:=t.TempDir()
    sha:=strings.Repeat("a",40)
    dir:=filepath.Join(cache,"reference",sha,"ProjectCharter","1.0","charter","prompts")
    if err:=os.MkdirAll(dir,0700);err!=nil{t.Fatal(err)}
    for name,body:=range map[string]string{
        "README.md":"# Indice",
        "legacy-charter-migration.md":"# Prompt de migração",
        "objective-refinement.md":"# Prompt de objetivos",
        "PROTOCOL.md":"# Protocolo",
    } {
        if err:=os.WriteFile(filepath.Join(dir,name),[]byte(body),0600);err!=nil{t.Fatal(err)}
    }
    a:=terminalApp{paths:appPaths{Cache:cache},session:terminalSession{Area:"charter",ReferenceSHA:sha,Prompt:"legacy-charter-migration.md"}}
    items,err:=a.promptOptions();if err!=nil{t.Fatal(err)}
    if len(items)!=2 {t.Fatalf("deveria listar somente 2 prompts: %v",items)}
    chosen,err:=a.promptRead();if err!=nil||!strings.Contains(chosen,"migração"){t.Fatalf("prompt incorreto: %v %s",err,chosen)}
}

func TestCentralizedPlainTextPromptsTakePrecedence(t *testing.T) {
    cache := t.TempDir()
    sha := strings.Repeat("b", 40)
    root := filepath.Join(cache, "reference", sha, "ProjectCharter", "1.0")
    centralized := filepath.Join(root, "prompts", "charter")
    legacy := filepath.Join(root, "charter", "prompts")
    for _, dir := range []string{centralized, legacy} {
        if err := os.MkdirAll(dir, 0700); err != nil {t.Fatal(err)}
    }
    for name, body := range map[string]string{
        "objective-refinement.txt": "New normative plain text Charter prompt.",
        "charter-review.txt": "# Review Charter.",
        "README.md": "# Index, not an instruction.",
        "agente.txt": "Router, not a specialized prompt.",
        "PROTOCOL.md": "Rules, not a specialized prompt.",
    } {
        if err := os.WriteFile(filepath.Join(centralized, name), []byte(body), 0600); err != nil {t.Fatal(err)}
    }
    if err := os.WriteFile(filepath.Join(legacy, "old-only.md"), []byte("Legacy only"), 0600); err != nil {t.Fatal(err)}
    a := terminalApp{paths: appPaths{Cache: cache}, session: terminalSession{
        Area: "charter", ReferenceSHA: sha, Prompt: "objective-refinement.md",
    }}
    if a.promptDir() != centralized {t.Fatalf("promptDir=%s; expected %s",a.promptDir(),centralized)}
    options, err := a.promptOptions()
    if err != nil {t.Fatal(err)}
    if len(options) != 2 {t.Fatalf("should list only 2 specialized .txt prompts, got %+v",options)}
    for _, option := range options {
        if filepath.Ext(option.ID) != ".txt" || option.ID == "agente.txt" {
            t.Fatalf("unexpected catalog entry %+v",option)
        }
    }
    content, err := a.promptRead()
    if err != nil {t.Fatal(err)}
    if !strings.Contains(content,"New normative plain text Charter prompt") {t.Fatalf("incorrect prompt: %q",content)}
}

func TestNewCatalogDoesNotSilentlyFallBackWhenEmpty(t *testing.T) {
    cache := t.TempDir()
    sha := strings.Repeat("c", 40)
    dir := filepath.Join(cache,"reference",sha,"ProjectCharter","1.0","prompts","scope")
    if err := os.MkdirAll(dir,0700); err != nil {t.Fatal(err)}
    a := terminalApp{paths: appPaths{Cache: cache},session: terminalSession{Area:"scope",ReferenceSHA:sha}}
    if _,err := a.promptOptions(); err==nil {t.Fatal("empty normative prompt catalog accepted")}
}
