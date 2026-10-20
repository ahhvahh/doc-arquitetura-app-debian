package main

import (
    "errors"
    "fmt"
    "os"
    "path/filepath"
    "sort"
    "strings"
)
type promptOption struct{ ID, Label string }
var promptLabels=map[string]string{
 "project-intake-review.md":"Diagnóstico inicial",
 "existing-project-catalog.md":"Catalogar documentação existente",
 "legacy-charter-migration.md":"Reutilizar e migrar Charter",
 "legacy-scope-migration.md":"Reutilizar e migrar Scope",
 "legacy-implementation-migration.md":"Reutilizar e migrar Implementation",
 "charter-review.md":"Auditoria geral de Charter",
 "scope-review.md":"Auditoria geral de Scope",
 "implementation-review.md":"Auditoria geral de Implementation",
 "charter-refinement.md":"Refinar Charter",
 "scope-refinement.md":"Refinar Scope",
 "objective-refinement.md":"Alterar objetivos",
 "discovery-refinement.md":"Nova funcionalidade / Discovery",
 "flow-refinement.md":"Alterar fluxos",
 "type-refinement.md":"Alterar tipos",
 "integration-refinement.md":"Alterar integrações",
 "library-refinement.md":"Alterar biblioteca",
 "service-refinement.md":"Alterar serviço",
 "application-cli-refinement.md":"Alterar aplicação CLI",
}
// promptDir selects the centralized text prompt catalog when present,
// falling back to the legacy layout while normative main is unchanged.
func (a *terminalApp) promptDir() string {
    base := filepath.Join(refBase(a.paths), a.session.ReferenceSHA, normTree)
    centralized := filepath.Join(base, "prompts", a.session.Area)
    if info, err := os.Stat(centralized); err == nil && info.IsDir() {
        return centralized
    }
    return filepath.Join(base, a.session.Area, "prompts")
}
func (a *terminalApp) promptOptions() ([]promptOption,error) {
 entries,err:=os.ReadDir(a.promptDir());if err!=nil{return nil,err}
 result:=[]promptOption{}
 for _,entry:=range entries {
   name := entry.Name()
   extension := filepath.Ext(name)
   if entry.IsDir() || (extension != ".md" && extension != ".txt") ||
      name == "README.md" || name == "PROTOCOL.md" ||
      name == "agente.md" || name == "agente.txt" {continue}
   label:=promptLabels[entry.Name()]
   if label == "" {
      label = promptLabels[strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))+".md"]
   }
   if label==""{
     b,e:=os.ReadFile(filepath.Join(a.promptDir(),entry.Name()));if e!=nil{return nil,e}
     line:=strings.SplitN(string(b),"\n",2)[0];label=strings.TrimSpace(strings.TrimPrefix(line,"#"))
     if label==""{label=entry.Name()}
   }
   result=append(result,promptOption{ID:entry.Name(),Label:label})
 }
 sort.Slice(result,func(i,j int)bool{return result[i].ID<result[j].ID})
 if len(result)==0{return nil,errors.New("nenhum prompt normativo disponível na main versionada")}
 return result,nil
}
func (a *terminalApp) promptSelect() error{
 opts,err:=a.promptOptions();if err!=nil{return err}
 display:=make([]string,len(opts));for i,v:=range opts{display[i]=v.Label+" ("+v.ID+")"}
 choice,err:=a.choose("Escolha um prompt especializado da etapa "+a.session.Area+" (normativa main):",display);if err!=nil{return err}
 for i,v:=range display{if choice==v{a.session.Prompt=opts[i].ID;return a.saveSession()}}
 return errors.New("seleção inválida")
}
func (a *terminalApp) promptRead() (string,error) {
 if a.session.Prompt=="" {return "",nil}
 opts,err:=a.promptOptions();if err!=nil{return "",err}
 for _,opt:=range opts {
    if opt.ID==a.session.Prompt ||
       strings.TrimSuffix(opt.ID, filepath.Ext(opt.ID)) == strings.TrimSuffix(a.session.Prompt, filepath.Ext(a.session.Prompt)) {
      b,err:=os.ReadFile(filepath.Join(a.promptDir(),opt.ID))
      if err!=nil{return "",err}
      return string(b),nil
    }
 }
 return "",errors.New("prompt selecionado não existe no catálogo da referência versionada")
}
func (a *terminalApp) showPrompts() error{
 opts,err:=a.promptOptions();if err!=nil{return err}
 fmt.Fprintf(a.out,"Prompts de %s, referência main SHA %s:\n",a.session.Area,a.session.ReferenceSHA)
 for i,v:=range opts{
   marker:=" "
   if v.ID==a.session.Prompt ||
      strings.TrimSuffix(v.ID,filepath.Ext(v.ID))==strings.TrimSuffix(a.session.Prompt,filepath.Ext(a.session.Prompt)){marker="*"}
   fmt.Fprintf(a.out,"%s %d. %s (%s)\n",marker,i+1,v.Label,v.ID)
 }
 fmt.Fprintln(a.out,"Use /prompt select para escolher ou /prompt reset para retirar a seleção.")
 return nil
}
