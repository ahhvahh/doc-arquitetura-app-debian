package main

import (
  "os"
  "path/filepath"
  "strings"
  "testing"
)

func configTestHome(t *testing.T) string {
    t.Helper()
    dir:=t.TempDir()
    t.Setenv("DOC_ARQUITETURA_CONFIG_DIR",dir)
    t.Setenv("DOC_ARQUITETURA_LLAMA_CONTEXT","")
    os.Unsetenv("DOC_ARQUITETURA_LLAMA_CONTEXT")
    return dir
}
func TestConfiguredLlamaContextPreference(t *testing.T) {
    dir:=configTestHome(t)
    n,source,err:=configuredLlamaContext()
    if err!=nil||n!=4096||source!="padrão"{t.Fatalf("%d %s %v",n,source,err)}
    if err=localConfigCommand([]string{"set-context","16384"});err!=nil{t.Fatal(err)}
    n,source,err=configuredLlamaContext()
    if err!=nil||n!=16384||source!="config.json"{t.Fatalf("%d %s %v",n,source,err)}
    st,e:=os.Stat(filepath.Join(dir,"config.json"))
    if e!=nil||st.Mode().Perm()!=0600{t.Fatalf("config permissions: %v %v",st,e)}
    t.Setenv("DOC_ARQUITETURA_LLAMA_CONTEXT","12288")
    n,source,err=configuredLlamaContext()
    if err!=nil||n!=12288||!strings.Contains(source,"ambiente"){t.Fatalf("%d %s %v",n,source,err)}
}
func TestConfiguredLlamaContextRejectsBadFile(t *testing.T) {
    dir:=configTestHome(t)
    for _,bad:=range []string{`{"llama":{"context_tokens":32769}}`,`{"llama":{"context_tokens":-1}}`,`{"llama":{"context_tokens":0}}`,`{"llama":{"context_tokens":8192},"typo":true}`,`{"llama":{"context_tokens":8192}}{}`} {
        if err:=os.WriteFile(filepath.Join(dir,"config.json"),[]byte(bad),0600);err!=nil{t.Fatal(err)}
        if _,_,err:=configuredLlamaContext();err==nil{t.Errorf("accepted invalid config: %s",bad)}
    }
}
func TestConfiguredLlamaContextRejectsInsecureFiles(t *testing.T){
    dir:=configTestHome(t)
    path:=filepath.Join(dir,"config.json")
    if err:=os.WriteFile(path,[]byte(`{"llama":{"context_tokens":16384}}`),0644);err!=nil{t.Fatal(err)}
    if _,_,err:=configuredLlamaContext();err==nil{t.Fatal("accepted world-readable config")}
    if err:=os.Remove(path);err!=nil{t.Fatal(err)}
    other:=filepath.Join(dir,"elsewhere")
    if err:=os.WriteFile(other,[]byte("{}"),0600);err!=nil{t.Fatal(err)}
    if err:=os.Symlink(other,path);err!=nil{t.Fatal(err)}
    if _,_,err:=configuredLlamaContext();err==nil{t.Fatal("accepted symlink")}
}
func TestConfiguredLlamaContextRejectsInvalidEnvironment(t *testing.T){
    configTestHome(t)
    for _,v:=range []string{"16385abc","0","32769"} {
        t.Setenv("DOC_ARQUITETURA_LLAMA_CONTEXT",v)
        if _,_,err:=configuredLlamaContext();err==nil{t.Errorf("accepted env %q",v)}
    }
}
