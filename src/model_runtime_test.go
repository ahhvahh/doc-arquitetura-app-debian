package main

import (
 "testing"
 "bytes"
 "fmt"
 "net/http"
 "net/http/httptest"
 "strings"
)

func TestActiveModelMatching(t *testing.T){
    yes:=[]string{"qwen3-1.7b","/home/u/models/qwen3-1.7b.gguf","qwen3-1.7b.gguf"}
    for _,id:=range yes{if !activeModelIs(id,"qwen3-1.7b"){t.Errorf("model should match %s",id)}}
    for _,id:=range []string{"qwen3.5-0.8b","other.gguf","qwen3-1.7b-malicious.gguf"}{
        if activeModelIs(id,"qwen3-1.7b"){t.Errorf("model mismatched: %s",id)}
    }
}


func TestModelRuntimeHealthAndBenchmark(t *testing.T) {
    p:=isolatedPaths(t)
    fakeInstalledModel(t,p,"qwen3-1.7b")
    mux:=http.NewServeMux()
    mux.HandleFunc("/health",func(w http.ResponseWriter,r *http.Request){fmt.Fprint(w,`{"status":"ok"}`)})
    mux.HandleFunc("/v1/models",func(w http.ResponseWriter,r *http.Request){fmt.Fprint(w,`{"data":[{"id":"qwen3-1.7b"}]}`)})
    mux.HandleFunc("/completion",func(w http.ResponseWriter,r *http.Request){
        if r.Method!=http.MethodPost{t.Errorf("wrong method: %s",r.Method)}
        fmt.Fprint(w,`{"timings":{"prompt_n":20,"prompt_per_second":90,"predicted_n":50,"predicted_per_second":25}}`)
    })
    srv:=httptest.NewServer(mux);defer srv.Close()
    var output bytes.Buffer
    a:=terminalApp{paths:p,out:&output,api:srv.URL,session:terminalSession{Model:"qwen3-1.7b"}}
    if err:=a.modelReady();err!=nil{t.Fatal(err)}
    if err:=a.modelBenchmark();err!=nil{t.Fatal(err)}
    if !strings.Contains(output.String(),"25.00 tokens/s"){t.Fatal(output.String())}
}
func TestModelRuntimeRejectsWrongModel(t *testing.T) {
    p:=isolatedPaths(t)
    fakeInstalledModel(t,p,"qwen3-1.7b")
    server:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
        if r.URL.Path=="/health"{fmt.Fprint(w,`{"status":"ok"}`);return}
        fmt.Fprint(w,`{"data":[{"id":"qwen3.5-0.8b"}]}`)
    }));defer server.Close()
    a:=terminalApp{paths:p,api:server.URL,session:terminalSession{Model:"qwen3-1.7b"}}
    if err:=a.modelReady();err==nil||!strings.Contains(err.Error(),"incompatível"){
        t.Fatalf("deveria rejeitar modelo diferente: %v",err)
    }
}
