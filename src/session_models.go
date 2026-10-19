package main

import (
    "errors"
    "fmt"
    "os"
    "strings"
)

// Session model selection is independent from global agent assignments.
// Only installed and verified GGUF files can be activated.
type modelMenuItem struct {
    label string
    kind string
    id string
}

func localModelOptions(p appPaths) []modelMenuItem {
    options := []modelMenuItem{}
    for _, spec := range modelCatalog {
        gguf, errFile := os.Stat(modelPath(p, spec.ID))
        meta, errMeta := os.Stat(modelMeta(p, spec.ID))
        if errFile == nil && errMeta == nil && gguf.Mode().IsRegular() && meta.Mode().IsRegular() {
            options = append(options, modelMenuItem{
                label: spec.ID + " — instalado — " + spec.Purpose,
                kind: "installed", id: spec.ID,
            })
        }
    }
    return append(options,
        modelMenuItem{label:"Baixar um modelo do catálogo",kind:"download"},
        modelMenuItem{label:"Continuar sem modelo (configurar depois)",kind:"skip"},
    )
}

func (a *terminalApp) useSessionModel(id string) error {
    if a.session.ID == "" { return errors.New("nenhuma sessão GitHub iniciada") }
    if _, err := getModel(id); err != nil { return err }
    if _, err := verifyModel(a.paths, id); err != nil {
        return fmt.Errorf("modelo %s não está instalado ou íntegro: %w",id,err)
    }
    oldSessionModel, oldRuntimeModel := a.session.Model, a.model
    if oldSessionModel != id {a.modelStop()}
    a.session.Model = id
    a.model = id
    if err := a.saveSession(); err != nil {
        a.session.Model, a.model = oldSessionModel, oldRuntimeModel
        return err
    }
    fmt.Fprintf(a.out,"Modelo da sessão: %s. O runtime privado será iniciado sob demanda pela aplicação; nenhum servidor externo é necessário.\n",id)
    return nil
}

func (a *terminalApp) selectSessionModel() error {
    for {
        options := localModelOptions(a.paths)
        names := make([]string,len(options))
        for i,item := range options { names[i] = item.label }
        chosen, err := a.choose("Modelo para esta sessão GitHub:",names)
        if err != nil { return err }
        for _,item := range options {
            if chosen != item.label { continue }
            switch item.kind {
            case "installed":
                return a.useSessionModel(item.id)
            case "skip":
                a.modelStop()
                a.session.Model = ""
                a.model = "local-model"
                if err = a.saveSession();err != nil {return err}
                fmt.Fprintln(a.out,"Sessão sem modelo: selecione posteriormente com /model select. Comandos documentais continuam disponíveis.")
                return nil
            case "download":
                catalog := make([]string,len(modelCatalog))
                for j,spec := range modelCatalog { catalog[j] = spec.ID + " — " + spec.Purpose }
                choice, er := a.choose("Selecione o modelo para baixar:",catalog)
                if er != nil {return er}
                for j,spec := range modelCatalog {
                    if choice != catalog[j] {continue}
                    yes,er := a.confirmation("Autoriza baixar "+spec.ID+" do Hugging Face?")
                    if er != nil {return er}
                    if !yes {break}
                    if er = downloadModel(a.paths,spec.ID,"");er != nil {return er}
                    return a.useSessionModel(spec.ID)
                }
            }
            break
        }
    }
}

func (a *terminalApp) activateSessionModel() error {
    if a.session.Model == "" {a.model="local-model";return nil}
    if _,err:=verifyModel(a.paths,a.session.Model);err!=nil{
        return fmt.Errorf("modelo salvo na sessão (%s) está indisponível: %w; use /model select",a.session.Model,err)
    }
    a.model=a.session.Model
    return nil
}

func (a *terminalApp) clearSessionModel() error {
    if a.session.ID==""{return errors.New("nenhuma sessão GitHub iniciada")}
    previous,previousRuntime:=a.session.Model,a.model
    a.modelStop()
    a.session.Model=""
    a.model="local-model"
    if err:=a.saveSession();err!=nil{
        a.session.Model,a.model=previous,previousRuntime
        return err
    }
    return nil
}

func (a *terminalApp) modelStatus(){
    id:=strings.TrimSpace(a.session.Model)
    if id==""{
        fmt.Fprintln(a.out,"Nenhum modelo selecionado. Use /model select. Disponíveis: doc-arquitetura models catalog.")
        return
    }
    fmt.Fprintf(a.out,"Modelo selecionado: %s; runtime privado gerenciado pela aplicação\n",id)
    if a.serverSocket==""{fmt.Fprintln(a.out,"Estado: descarregado; será iniciado na próxima análise")}else{fmt.Fprintln(a.out,"Socket privado:",a.serverSocket)}
}
