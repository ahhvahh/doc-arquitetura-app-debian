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
