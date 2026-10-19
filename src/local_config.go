package main

import (
 "bytes"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "strconv"
 "strings"
)

// Per-user preferences are separate from github.json, so no GitHub token is
// exposed when reading or displaying the runtime settings.
type localConfig struct {
    Llama struct {
        ContextTokens *int `json:"context_tokens"`
    } `json:"llama"`
}
const defaultContextTokens = 4096
const minContextTokens = 512
const maxContextTokens = 32768

func localConfigPath() (string,error) {
    paths,err:=appHome()
    if err!=nil{return "",err}
    return filepath.Join(paths.Config,"config.json"),nil
}
func validateContextTokens(n int) error {
    if n<minContextTokens||n>maxContextTokens {
        return fmt.Errorf("context_tokens deve estar entre %d e %d; valor recebido: %d",minContextTokens,maxContextTokens,n)
    }
    return nil
}
func readLocalConfig() (localConfig,error) {
    var cfg localConfig
    path,err:=localConfigPath()
    if err!=nil{return cfg,err}
    // Only the current Linux user can read or modify the configuration.
    data,err:=githubPrivateRead(path)
    if os.IsNotExist(err){return cfg,nil}
    if err!=nil{return cfg,err}
    dec:=json.NewDecoder(bytes.NewReader(data))
    dec.DisallowUnknownFields()
    if err:=dec.Decode(&cfg);err!=nil{return cfg,fmt.Errorf("config.json inválido: %w",err)}
    var extra any
    if e:=dec.Decode(&extra);e!=io.EOF {
        return cfg,errors.New("config.json deve conter somente um objeto JSON")
    }
    if cfg.Llama.ContextTokens!=nil {
        if err:=validateContextTokens(*cfg.Llama.ContextTokens);err!=nil{return cfg,err}
    }
    return cfg,nil
}
// Environment setting overrides the persisted preference; invalid values
// are errors, never silently replaced by 4096.
func configuredLlamaContext() (int,string,error) {
    cfg,err:=readLocalConfig();if err!=nil{return 0,"",err}
    n:=defaultContextTokens
    source:="padrão"
    if cfg.Llama.ContextTokens!=nil{n=*cfg.Llama.ContextTokens;source="config.json"}
    if raw,ok:=os.LookupEnv("DOC_ARQUITETURA_LLAMA_CONTEXT");ok {
        parsed,e:=strconv.Atoi(strings.TrimSpace(raw))
        if e!=nil{return 0,"",fmt.Errorf("DOC_ARQUITETURA_LLAMA_CONTEXT inválido: %q",raw)}
        if err:=validateContextTokens(parsed);err!=nil{return 0,"",err}
        n=parsed;source="ambiente DOC_ARQUITETURA_LLAMA_CONTEXT"
    }
    return n,source,nil
}
func localConfigCommand(args []string) error {
    path,err:=localConfigPath();if err!=nil{return err}
    if len(args)==0 || args[0]=="show" {
        if len(args)>1{return errors.New("uso: doc-arquitetura config show")}
        n,source,err:=configuredLlamaContext()
        if err!=nil{return err}
        fmt.Printf("Configuração: %s\nJanela de contexto: %d tokens (%s)\n",path,n,source)
        return nil
    }
    if len(args)!=2 || args[0]!="set-context" {
        return errors.New("uso: doc-arquitetura config show | config set-context <tokens>")
    }
    n,err:=strconv.Atoi(args[1]);if err!=nil{return fmt.Errorf("número de tokens inválido: %s",args[1])}
    if err=validateContextTokens(n);err!=nil{return err}
    cfg,err:=readLocalConfig();if err!=nil{return err}
    cfg.Llama.ContextTokens=&n
    if err=writeJSONAtomic(path,cfg);err!=nil{return err}
    fmt.Printf("Janela padrão salva: %d tokens em %s\n",n,path)
    if raw,present:=os.LookupEnv("DOC_ARQUITETURA_LLAMA_CONTEXT");present {
        fmt.Printf("ATENÇÃO: variável DOC_ARQUITETURA_LLAMA_CONTEXT=%s substitui o valor do arquivo nesta execução. No chat, /context set remove a substituição; fora dele, reabra sem a variável.\n",raw)
    }
    fmt.Println("A configuração entra em vigor quando uma nova instância do modelo for iniciada. Use /model stop e /model start no chat, ou reabra o aplicativo.")
    return nil
}
