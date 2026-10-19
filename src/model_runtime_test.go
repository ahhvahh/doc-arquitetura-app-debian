package main

import (
 "context"
 "encoding/json"
 "fmt"
 "net"
 "net/http"
 "os/exec"
 "os"
 "path/filepath"
 "testing"
)

func TestActiveModelMatching(t *testing.T) {
    for _,v:=range []string{"qwen3-1.7b","/models/qwen3-1.7b.gguf"} {
        if !activeModelIs(v,"qwen3-1.7b"){t.Fatal(v)}
    }
    if activeModelIs("another.gguf","qwen3-1.7b"){t.Fatal("wrong model accepted")}
}
func TestManagedUnixSocketEndpoints(t *testing.T) {
    dir:=t.TempDir()
    sock:=filepath.Join(dir,"test.sock")
    ln,err:=net.Listen("unix",sock);if err!=nil{t.Fatal(err)}
    mux:=http.NewServeMux()
    mux.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){fmt.Fprint(w,`{"status":"ok"}`)})
    mux.HandleFunc("/v1/models",func(w http.ResponseWriter,r *http.Request){json.NewEncoder(w).Encode(map[string]any{"data":[]any{map[string]string{"id":"qwen3-1.7b"}}})})
    mux.HandleFunc("/completion",func(w http.ResponseWriter,r *http.Request){
        fmt.Fprint(w,`{"timings":{"prompt_n":64,"prompt_per_second":100,"predicted_n":32,"predicted_per_second":15.5}}`)
    })
    server:=&http.Server{Handler:mux}
    go server.Serve(ln)
    defer server.Shutdown(context.Background())
    p:=isolatedPaths(t)
    a:=terminalApp{paths:p,session:terminalSession{Model:"qwen3-1.7b"},serverSocket:sock,serverProcess:&exec.Cmd{},out:os.Stdout}
    if err:=a.modelReady();err!=nil{t.Fatal(err)}
    if err:=a.modelBenchmark();err!=nil{t.Fatal(err)}
}
func TestPrivateRuntimeExecutableOverride(t *testing.T) {
    t.Setenv("DOC_ARQUITETURA_LLAMA_BIN","relative/llama-server")
    if _,err:=managedLlamaExecutable();err==nil{t.Fatal("relative path accepted")}
}
func TestPrivateRuntimeContextBounds(t *testing.T) {
    t.Setenv("DOC_ARQUITETURA_LLAMA_CONTEXT","99999")
    if managedLlamaContext()!=4096{t.Fatal("out-of-range context was accepted")}
    t.Setenv("DOC_ARQUITETURA_LLAMA_CONTEXT","2048")
    if managedLlamaContext()!=2048{t.Fatal("valid context rejected")}
}
