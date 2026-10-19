package main

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

func activeModelIs(id, wanted string) bool {
    if id == wanted { return true }
    return strings.HasSuffix(id, "/"+wanted+".gguf") ||
        strings.HasSuffix(id, "\\"+wanted+".gguf") ||
        id == wanted+".gguf"
}

func (a *terminalApp) modelReady() error {
    if a.session.Model == "" { return errors.New("nenhum modelo configurado; use /model select") }
    if st,err:=os.Stat(modelPath(a.paths,a.session.Model));err!=nil||!st.Mode().IsRegular(){return errors.New("GGUF selecionado indisponível; use /model select")}
    endpoint,err:=localhostAPI(a.api);if err!=nil{return err}
    cli:=&http.Client{Timeout:3*time.Second}
    res,err:=cli.Get(endpoint+"/health")
    if err!=nil {
        return fmt.Errorf("servidor local inacessível em %s: inicie com /model start (exige llama-server instalado), ou em outro terminal: doc-arquitetura models serve %s: %w", endpoint,a.session.Model,err)
    }
    io.Copy(io.Discard,io.LimitReader(res.Body,2048));res.Body.Close()
    if res.StatusCode!=http.StatusOK {
        return fmt.Errorf("llama-server não pronto em %s (HTTP %d); aguarde carregar o modelo e use /model check",endpoint,res.StatusCode)
    }
    models,err:=cli.Get(endpoint+"/v1/models")
    if err!=nil{return fmt.Errorf("não foi possível verificar o modelo do servidor: %w",err)}
    defer models.Body.Close()
    if models.StatusCode!=http.StatusOK{return fmt.Errorf("modelo do servidor não verificável: HTTP %d",models.StatusCode)}
    var response struct{Data []struct{ID string `json:"id"`} `json:"data"`}
    if err=json.NewDecoder(io.LimitReader(models.Body,1<<20)).Decode(&response);err!=nil{return err}
    if len(response.Data)==0{return errors.New("llama-server não anunciou modelo ativo")}
    for _,m:=range response.Data {if activeModelIs(m.ID,a.session.Model){return nil}}
    ids:=make([]string,0,len(response.Data));for _,m:=range response.Data{ids=append(ids,m.ID)}
    return fmt.Errorf("modelo incompatível: a sessão usa %s mas o servidor oferece %s. Encerre o servidor anterior e execute: doc-arquitetura models serve %s",a.session.Model,strings.Join(ids,", "),a.session.Model)
}

func (a *terminalApp) modelStart() error {
    if a.session.Model==""{return errors.New("selecione um modelo com /model select")}
    if err:=a.modelReady();err==nil{fmt.Fprintln(a.out,"llama-server já está pronto com o modelo da sessão");return nil}
    endpoint,err:=localhostAPI(a.api);if err!=nil{return err}
    // If a server already holds the port, do not launch another one over it.
    probe:=&http.Client{Timeout:2*time.Second}
    if res,e:=probe.Get(endpoint+"/health");e==nil{
        res.Body.Close()
        return fmt.Errorf("porta %s já responde, mas não com o modelo esperado; encerre o servidor existente antes de iniciar %s",endpoint,a.session.Model)
    }
    if a.serverProcess!=nil{return errors.New("servidor da sessão já iniciado; use /model check ou /model stop")}
    u,err:=url.Parse(endpoint);if err!=nil{return err}
    port:=u.Port()
    if port==""{port="80"}
    n,err:=strconv.Atoi(port);if err!=nil||n<1||n>65535{return errors.New("porta de API inválida")}
    exe,err:=exec.LookPath("llama-server")
    if err!=nil{return errors.New("llama-server não está instalado no PATH; instale o runtime llama.cpp antes de usar /model start")}
    if _,err=verifyModel(a.paths,a.session.Model);err!=nil{return err}
    cmd:=exec.Command(exe,"-m",modelPath(a.paths,a.session.Model),"--alias",a.session.Model,
        "--host","127.0.0.1","--port",port,"-c","8192")
    cmd.Stdout=io.Discard
    cmd.Stderr=os.Stderr
    if err=cmd.Start();err!=nil{return fmt.Errorf("falha ao iniciar llama-server: %w",err)}
    a.serverProcess=cmd
    done:=make(chan error,1)
    a.serverDone=done
    go func(){done<-cmd.Wait()}()
    fmt.Fprintln(a.out,"Carregando modelo no llama-server local (até 60 segundos)...")
    deadline:=time.NewTimer(60*time.Second);defer deadline.Stop()
    tick:=time.NewTicker(600*time.Millisecond);defer tick.Stop()
    for{
        select{
        case err=<-done:
            a.serverProcess=nil;a.serverDone=nil
            return fmt.Errorf("llama-server terminou durante a inicialização: %v",err)
        case <-deadline.C:
            a.modelStop()
            return errors.New("timeout ao carregar modelo (60s); verifique logs, memória disponível e versão do llama.cpp")
        case <-tick.C:
            if err=a.modelReady();err==nil{
                fmt.Fprintf(a.out,"Modelo %s pronto em %s. Use /benchmark para medir tokens/s.\n",a.session.Model,endpoint)
                return nil
            }
        }
    }
}

func (a *terminalApp) modelStop() {
    if a.serverProcess==nil{return}
    cmd,done:=a.serverProcess,a.serverDone
    a.serverProcess=nil;a.serverDone=nil
    _=cmd.Process.Signal(os.Interrupt)
    select{
    case <-done:
    case <-time.After(3*time.Second):
        _=cmd.Process.Kill()
        <-done
    }
}

func (a *terminalApp) modelBenchmark() error {
    if err:=a.modelReady();err!=nil{return err}
    endpoint,err:=localhostAPI(a.api);if err!=nil{return err}
    // Native llama.cpp timings report prompt processing and generation separately.
    request:=map[string]any{
        "prompt":"Leia esta solicitação de documentação técnica em português. Descreva resumidamente quatro critérios de qualidade de um Charter, um por frase.",
        "n_predict":128,"temperature":0.0,"stream":false,"cache_prompt":false,
    }
    body,err:=json.Marshal(request);if err!=nil{return err}
    cli:=&http.Client{Timeout:3*time.Minute}
    req,err:=http.NewRequest(http.MethodPost,endpoint+"/completion",bytes.NewReader(body));if err!=nil{return err}
    req.Header.Set("Content-Type","application/json")
    start:=time.Now()
    response,err:=cli.Do(req);if err!=nil{return err}
    defer response.Body.Close()
    if response.StatusCode!=http.StatusOK{return fmt.Errorf("benchmark HTTP %d",response.StatusCode)}
    var result struct{
        Timings struct{
            PromptN int `json:"prompt_n"`
            PromptPerSecond float64 `json:"prompt_per_second"`
            PredictedN int `json:"predicted_n"`
            PredictedPerSecond float64 `json:"predicted_per_second"`
        } `json:"timings"`
    }
    if err=json.NewDecoder(io.LimitReader(response.Body,3<<20)).Decode(&result);err!=nil{return err}
    t:=result.Timings
    if t.PredictedN<1||t.PredictedPerSecond<=0{return errors.New("servidor não forneceu métricas válidas de geração")}
    fmt.Fprintf(a.out,"BENCHMARK REAL — %s\n",a.session.Model)
    fmt.Fprintf(a.out,"Tokens gerados: %d | Geração: %.2f tokens/s\n",t.PredictedN,t.PredictedPerSecond)
    if t.PromptPerSecond>0{fmt.Fprintf(a.out,"Processamento do prompt: %d tokens, %.2f tokens/s\n",t.PromptN,t.PromptPerSecond)}
    fmt.Fprintf(a.out,"Tempo total da solicitação: %.2f s. Resultado local, depende da CPU, memória, contexto e carga.\n",time.Since(start).Seconds())
    return nil
}
