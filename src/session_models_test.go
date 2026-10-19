package main

import (
    "bufio"
    "bytes"
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func fakeInstalledModel(t *testing.T,p appPaths,id string){
    t.Helper()
    path:=modelPath(p,id)
    if err:=os.MkdirAll(filepath.Dir(path),0700);err!=nil{t.Fatal(err)}
    if err:=os.WriteFile(path,[]byte("GGUFtest-local-model"),0600);err!=nil{t.Fatal(err)}
    sum,size,err:=shaFile(path);if err!=nil{t.Fatal(err)}
    if err:=writeJSONAtomic(modelMeta(p,id),modelRecord{ID:id,SHA256:sum,Size:size});err!=nil{t.Fatal(err)}
}

func TestModelMenuEmptyCatalog(t *testing.T){
    p:=isolatedPaths(t)
    options:=localModelOptions(p)
    if len(options)!=2||options[0].kind!="download"||options[1].kind!="skip"{t.Fatalf("unexpected options: %+v",options)}
}
func TestModelMenuSelectAndPersist(t *testing.T){
    p:=isolatedPaths(t)
    fakeInstalledModel(t,p,"qwen3-1.7b")
    var output bytes.Buffer
    app:=terminalApp{in:bufio.NewReader(strings.NewReader("1\n")),out:&output,paths:p,model:"local-model",session:terminalSession{ID:"session-model-test",Area:"charter"}}
    if err:=app.selectSessionModel();err!=nil{t.Fatal(err)}
    if app.model!="qwen3-1.7b"||app.session.Model!="qwen3-1.7b"{t.Fatalf("not selected: %+v",app.session)}
    var saved terminalSession
    if err:=readJSON(filepath.Join(p.State,"sessions",app.session.ID+".json"),&saved);err!=nil{t.Fatal(err)}
    if saved.Model!="qwen3-1.7b"{t.Fatalf("selection not persisted: %+v",saved)}
    resumed:=terminalApp{paths:p,session:saved,model:"local-model"}
    if err:=resumed.activateSessionModel();err!=nil||resumed.model!="qwen3-1.7b"{t.Fatalf("resume failed %s %v",resumed.model,err)}
}
func TestModelMenuRejectMissingAndCorrupt(t *testing.T){
    p:=isolatedPaths(t)
    a:=terminalApp{paths:p,out:&bytes.Buffer{},session:terminalSession{ID:"session-model-invalid"}}
    if err:=a.useSessionModel("qwen3-1.7b");err==nil{t.Fatal("missing model accepted")}
    fakeInstalledModel(t,p,"qwen3-1.7b")
    if err:=os.WriteFile(modelPath(p,"qwen3-1.7b"),[]byte("BAD!"),0600);err!=nil{t.Fatal(err)}
    if err:=a.useSessionModel("qwen3-1.7b");err==nil{t.Fatal("corrupted model accepted")}
    if a.session.Model!=""{t.Fatal("invalid selection persisted")}
}
func TestModelMenuSkipWithoutInstall(t *testing.T){
    p:=isolatedPaths(t)
    a:=terminalApp{in:bufio.NewReader(strings.NewReader("2\n")),out:&bytes.Buffer{},paths:p,session:terminalSession{ID:"session-no-model"}}
    if err:=a.selectSessionModel();err!=nil{t.Fatal(err)}
    if a.session.Model!=""||a.model!="local-model"{t.Fatalf("wrong mode: %+v",a.session)}
}
