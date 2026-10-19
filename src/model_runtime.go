package main

import (
    "bytes"
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strconv"
    "strings"
    "time"
)

// llama.cpp is a private subprocess, reached only through a UNIX domain socket.
// No TCP listener, server daemon, global installation or web UI is required.
const privateLlamaBinary = "/usr/lib/doc-arquitetura-app/llama-server"
const privateLlamaHost = "http://doc-arquitetura.internal"

func managedLlamaExecutable() (string,error) {
    paths:=[]string{}
    if override:=os.Getenv("DOC_ARQUITETURA_LLAMA_BIN");override!="" {
        if !filepath.IsAbs(override) {return "",errors.New("DOC_ARQUITETURA_LLAMA_BIN precisa ser absoluto")}
        paths=append(paths,override)
    }else{
        paths=append(paths,privateLlamaBinary)
        if exe,err:=os.Executable();err==nil {
            paths=append(paths,filepath.Join(filepath.Dir(exe),"llama-server"))
        }
    }
    for _,p:=range paths {
        info,e:=os.Stat(p)
        if e==nil && info.Mode().IsRegular() && info.Mode()&0111!=0{return p,nil}
    }
    return "",fmt.Errorf("llama-server privado ausente; recompile e instale o pacote Debian com src/scripts/build-private-llama.sh (esperado: %s)",paths[0])
}
func activeModelIs(id,wanted string) bool {
    if id==wanted{return true}
    return strings.HasSuffix(id,"/"+wanted+".gguf") || id==wanted+".gguf"
}
func (a *terminalApp) managedHTTP() (*http.Client,string,error) {
    if a.serverSocket==""{return nil,"",errors.New("runtime ainda não iniciado; use /model start")}
    socket:=a.serverSocket
    transport:=&http.Transport{
        Proxy:nil,
        DialContext:func(ctx context.Context,_,_ string)(net.Conn,error){
            var d net.Dialer
            return d.DialContext(ctx,"unix",socket)
        },
        DisableKeepAlives:true,
    }
    return &http.Client{Transport:transport,Timeout:15*time.Minute},privateLlamaHost,nil
}
func (a *terminalApp) modelReady() error {
    if a.serverProcess==nil||a.serverSocket==""{return errors.New("modelo não iniciado; a aplicação o carregará na primeira análise ou use /model start")}
    if a.session.Model==""{return errors.New("nenhum modelo selecionado")}
    cli,host,err:=a.managedHTTP();if err!=nil{return err}
    ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second);defer cancel()
    req,err:=http.NewRequestWithContext(ctx,http.MethodGet,host+"/health",nil);if err!=nil{return err}
    res,err:=cli.Do(req);if err!=nil{return fmt.Errorf("runtime privado indisponível: %w",err)}
    io.Copy(io.Discard,io.LimitReader(res.Body,1024));res.Body.Close()
    if res.StatusCode!=http.StatusOK{return fmt.Errorf("modelo ainda carregando (HTTP %d)",res.StatusCode)}
    request,err:=http.NewRequestWithContext(ctx,http.MethodGet,host+"/v1/models",nil);if err!=nil{return err}
    response,err:=cli.Do(request);if err!=nil{return err}
    defer response.Body.Close()
    if response.StatusCode!=http.StatusOK{return fmt.Errorf("catálogo privado HTTP %d",response.StatusCode)}
    var result struct { Data []struct{ID string `json:"id"`} `json:"data"` }
    if err=json.NewDecoder(io.LimitReader(response.Body,1<<20)).Decode(&result);err!=nil{return err}
    for _,entry:=range result.Data{if activeModelIs(entry.ID,a.session.Model){return nil}}
    return fmt.Errorf("modelo do runtime não corresponde ao selecionado: %s",a.session.Model)
}
func (a *terminalApp) cleanupRuntime() {
    if a.serverSocket!="" { _=os.Remove(a.serverSocket) }
    if a.serverDir!="" {_=os.Remove(a.serverDir)}
    a.serverSocket=""
    a.serverDir=""
}
func (a *terminalApp) modelStart() error {
    if a.session.Model=="" {return errors.New("selecione o modelo com /model select")}
    if a.serverProcess!=nil {
        return a.modelReady()
    }
    spec,err:=managedLlamaExecutable();if err!=nil{return err}
    contextTokens,contextSource,err:=configuredLlamaContext();if err!=nil{return err}
    if _,err=verifyModel(a.paths,a.session.Model);err!=nil{return err}
    // mktemp mode 0700 prevents cross-user access; name fits UNIX sun_path.
    parent:=os.TempDir()
    if runtimeDir:=os.Getenv("XDG_RUNTIME_DIR");runtimeDir!=""{
        if info,e:=os.Stat(runtimeDir);e==nil && info.IsDir() {parent=runtimeDir}
    }
    dir,err:=os.MkdirTemp(parent,"docarq-llm-");if err!=nil{return err}
    if err=os.Chmod(dir,0700);err!=nil{os.Remove(dir);return err}
    socket:=filepath.Join(dir,"llama.sock")
    if len(socket)>100{os.Remove(dir);return errors.New("caminho do socket Unix muito longo; defina TMPDIR menor")}
    a.serverDir,a.serverSocket=dir,socket
    cores:=runtime.NumCPU()
    if cores>4{cores=4};if cores<1{cores=1}
    cmd:=exec.Command(spec,"-m",modelPath(a.paths,a.session.Model),
       "--alias",a.session.Model,"--host",socket,"--no-webui",
       "-c",strconv.Itoa(contextTokens),"-np","1", // single session: one full-size KV cache slot
       "-t",strconv.Itoa(cores),"-tb",strconv.Itoa(cores),
       "-ngl","0")
    cmd.Stdout=io.Discard
    cmd.Stderr=os.Stderr
    if err=cmd.Start();err!=nil{a.cleanupRuntime();return fmt.Errorf("não foi possível iniciar runtime privado: %w",err)}
    a.serverProcess=cmd
    done:=make(chan error,1)
    a.serverDone=done
    go func(){done<-cmd.Wait()}()
    fmt.Fprintf(a.out,"Iniciando llama.cpp privado: %s, contexto %d (%s), %d threads, 1 slot. Socket: %s\n",
        a.session.Model,contextTokens,contextSource,cores,socket)
    deadline:=time.NewTimer(90*time.Second);defer deadline.Stop()
    ticker:=time.NewTicker(800*time.Millisecond);defer ticker.Stop()
    for {
        select {
        case exit:=<-done:
            a.serverProcess=nil;a.serverDone=nil;a.cleanupRuntime()
            return fmt.Errorf("runtime privado terminou antes de ficar pronto: %v",exit)
        case <-deadline.C:
            a.modelStop()
            return errors.New("modelo não ficou pronto em 90 segundos; verifique memória, GGUF e logs do processo")
        case <-ticker.C:
            if err=a.modelReady();err==nil{
                fmt.Fprintln(a.out,"Runtime privado pronto para análise.")
                return nil
            }
        }
    }
}
func (a *terminalApp) ensureRuntime() error {
    if a.serverProcess==nil{return a.modelStart()}
    return a.modelReady()
}
func (a *terminalApp) modelStop() {
    cmd,done:=a.serverProcess,a.serverDone
    a.serverProcess=nil
    a.serverDone=nil
    if cmd!=nil {
        _=cmd.Process.Signal(os.Interrupt)
        select {
        case <-done:
        case <-time.After(3*time.Second):
            _=cmd.Process.Kill()
            <-done
        }
    }
    a.cleanupRuntime()
}
func (a *terminalApp) modelBenchmark() error {
    if err:=a.ensureRuntime();err!=nil{return err}
    client,host,err:=a.managedHTTP();if err!=nil{return err}
    data,err:=json.Marshal(map[string]any{
        "prompt":"Descreva quatro critérios objetivos para revisar uma documentação Charter em português.",
        "n_predict":128,"temperature":0.0,"stream":false,
    });if err!=nil{return err}
    request,err:=http.NewRequest(http.MethodPost,host+"/completion",bytes.NewReader(data));if err!=nil{return err}
    request.Header.Set("Content-Type","application/json")
    start:=time.Now()
    response,err:=client.Do(request);if err!=nil{return err}
    defer response.Body.Close()
    if response.StatusCode!=200{return fmt.Errorf("benchmark do runtime HTTP %d",response.StatusCode)}
    var result struct {
        Timings struct{
            PromptN int `json:"prompt_n"`
            PromptPerSecond float64 `json:"prompt_per_second"`
            PredictedN int `json:"predicted_n"`
            PredictedPerSecond float64 `json:"predicted_per_second"`
        } `json:"timings"`
    }
    if err=json.NewDecoder(io.LimitReader(response.Body,2<<20)).Decode(&result);err!=nil{return err}
    t:=result.Timings
    if t.PredictedN<=0||t.PredictedPerSecond<=0{return errors.New("runtime não forneceu métricas de desempenho válidas")}
    fmt.Fprintf(a.out,"BENCHMARK %s — Geração: %.2f tokens/s (%d tokens); prompt: %.2f tokens/s (%d tokens); duração: %.2f s\n",
        a.session.Model,t.PredictedPerSecond,t.PredictedN,t.PromptPerSecond,t.PromptN,time.Since(start).Seconds())
    return nil
}
