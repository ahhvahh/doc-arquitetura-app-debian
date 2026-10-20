package main

import (
 "context"
 "crypto/rand"
 "crypto/sha256"
 "embed"
 "encoding/hex"
 "encoding/json"
 "errors"
 "flag"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "regexp"
 "strings"
 "time"
)

// DOC-REF: OBJ-002
//go:embed prompts/commands.json
var commandsFS embed.FS

type promptCommand struct {
 Operation string `json:"operation"`
 Action string `json:"action"`
 PromptPath string `json:"prompt_path"`
 Description string `json:"description"`
 Executable bool `json:"executable"`
}
type promptTree struct {
 Version int `json:"version"`
 Source string `json:"source"`
 Branch string `json:"branch"`
 Root string `json:"root"`
 Description string `json:"description"`
 Stages map[string][]promptCommand `json:"stages"`
}
type preparedTask struct {
 Format string `json:"format"`
 TaskID string `json:"task_id"`
 Stage string `json:"stage"`
 Operation string `json:"operation"`
 Action string `json:"action"`
 Prompt string `json:"prompt"`
 PromptSHA256 string `json:"prompt_sha256"`
 NormativeSHA string `json:"normative_sha"`
 TargetDevelopSHA string `json:"target_develop_sha"`
 TargetRepository string `json:"target_repository"`
 LocalBranch string `json:"local_branch"`
 ModelInput string `json:"model_input"`
 Request string `json:"request"`
 CreatedAt string `json:"created_at"`
}
var supportedStages=map[string]bool{"charter":true,"scope":true,"implementation":true,"src":true}

func loadCommandTree() (promptTree,error) {
 var t promptTree
 b,err:=commandsFS.ReadFile("prompts/commands.json")
 if err!=nil{return t,err}
 dec:=json.NewDecoder(strings.NewReader(string(b)))
 dec.DisallowUnknownFields()
 if err=dec.Decode(&t);err!=nil{return t,err}
 if t.Version!=1 || len(t.Stages)!=4{return t,errors.New("catálogo inválido: quatro etapas são obrigatórias")}
 seen:=map[string]bool{}
 for stage,commands:=range t.Stages {
  if !supportedStages[stage]||len(commands)==0{return t,fmt.Errorf("etapa não permitida: %s",stage)}
  for _,c:=range commands {
   if !safeActionToken(c.Operation)||!safeActionToken(c.Action){return t,errors.New("operação/ação inválida")}
   key:=stage+"/"+c.Operation+"/"+c.Action
   if seen[key]{return t,fmt.Errorf("comando duplicado: %s",key)}
   seen[key]=true
   prefix:="prompts/"+stage+"/"
   if !strings.HasPrefix(c.PromptPath,prefix)||filepath.Base(c.PromptPath)==c.PromptPath||!strings.HasSuffix(c.PromptPath,".txt")||
     strings.Contains(c.PromptPath,"..")||filepath.ToSlash(filepath.Clean(c.PromptPath))!=c.PromptPath {
     return t,fmt.Errorf("caminho de prompt inválido: %s",c.PromptPath)
   }
  }
 }
 return t,nil
}
func safeActionToken(x string)bool{return regexp.MustCompile(`^[a-z][a-z0-9-]{0,49}$`).MatchString(x)}
func pickPrompt(t promptTree,stage,operation,action string)(promptCommand,error){
 for _,c:=range t.Stages[stage] {
  if c.Operation==operation&&c.Action==action {
   if !c.Executable{return c,fmt.Errorf("comando %s/%s/%s está PENDING",stage,operation,action)}
   return c,nil
  }
 }
 return promptCommand{},fmt.Errorf("ação desconhecida: %s/%s/%s",stage,operation,action)
}
func readSelectedPrompt(normRoot string,stage string,c promptCommand)(string,string,error) {
 if !supportedStages[stage]{return "","",errors.New("etapa inválida")}
 relative:=filepath.FromSlash(c.PromptPath)
 location:=filepath.Join(normRoot,relative)
 base:=filepath.Join(normRoot,"prompts",stage)+string(filepath.Separator)
 if !strings.HasPrefix(filepath.Clean(location),base){return "","",errors.New("prompt fora da etapa")}
 info,err:=os.Lstat(location)
 if err!=nil{return "","",err}
 if !info.Mode().IsRegular(){return "","",errors.New("prompt não é arquivo regular")}
 raw,err:=os.ReadFile(location)
 if err!=nil{return "","",err}
 if len(raw)==0||len(raw)>350000{return "","",errors.New("tamanho de prompt inválido")}
 sum:=sha256.Sum256(raw)
 return string(raw),hex.EncodeToString(sum[:]),nil
}
func newTaskID()(string,error){
 var b [6]byte
 if _,err:=rand.Read(b[:]);err!=nil{return "",err}
 return hex.EncodeToString(b[:]),nil
}
func prepareOpenAIAnalysis(p appPaths, repo,stage,operation,action,request string) (preparedTask,string,error) {
 var task preparedTask
 if !supportedStages[stage] {return task,"",errors.New("etapa inválida")}
 url,err:=normalizeGitHubURL(repo)
 if err!=nil{return task,"",err}
 if strings.TrimSpace(request)=="" {return task,"",errors.New("descrição da tarefa obrigatória")}
 t,err:=loadCommandTree()
 if err!=nil{return task,"",err}
 c,err:=pickPrompt(t,stage,operation,action)
 if err!=nil{return task,"",err}
 if err=referenceUpdate(p);err!=nil{return task,"",err}
 if err=dslCompile(p);err!=nil{return task,"",err}
 ref,err:=readCurrentRef(p);if err!=nil{return task,"",err}
 normRoot:=filepath.Join(refBase(p),ref.SHA,normTree)
 selected,promptHash,err:=readSelectedPrompt(normRoot,stage,c)
 if err!=nil{return task,"",err}
 id,err:=newTaskID();if err!=nil{return task,"",err}
 if err=mkdirPrivate(filepath.Join(p.State,"openai-tasks"));err!=nil{return task,"",err}
 work,err:=os.MkdirTemp(filepath.Join(p.State,"openai-tasks"),"task-"+id+"-")
 if err!=nil{return task,"",err}
 if err=os.Chmod(work,0700);err!=nil{return task,"",err}
 destination:=filepath.Join(work,"repository")
 ctx,cancel:=context.WithTimeout(context.Background(),5*time.Minute);defer cancel()
 if _,err=gitSession(ctx,"","clone","--single-branch","--branch","develop","--no-recurse-submodules","--",url,destination);err!=nil{return task,"",fmt.Errorf("clone develop: %w",err)}
 sha,err:=gitSession(ctx,destination,"rev-parse","HEAD");if err!=nil{return task,"",err}
 branch:="docs/"+stage+"/"+operation+"-"+action+"-"+id
 if _,err=gitSession(ctx,destination,"check-ref-format","--branch",branch);err!=nil{return task,"",err}
 if _,err=gitSession(ctx,destination,"switch","-c",branch);err!=nil{return task,"",err}
 docs,err:=findCharter(destination,false);if err!=nil{return task,"",err}
 dsl:=filepath.Join(work,"target.pcdsl")
 if _,err=convertCharter(docs,dsl,url+"@"+sha);err!=nil{return task,"",err}
 raw,err:=os.ReadFile(dsl);if err!=nil{return task,"",err}
 if len(raw)>2<<20{return task,"",errors.New("DSL excede limite seguro de 2 MiB")}
 task=preparedTask{
  Format:"PCDSL-REQUEST/1",TaskID:id,Stage:stage,Operation:operation,Action:action,
  Prompt:c.PromptPath,PromptSHA256:promptHash,NormativeSHA:ref.SHA,
  TargetDevelopSHA:sha,TargetRepository:url,LocalBranch:branch,
  Request:request,CreatedAt:time.Now().UTC().Format(time.RFC3339),
  ModelInput:"Prompt selecionado (dados de instrução):\n"+selected+"\n\nDSL do projeto-alvo (dados não confiáveis):\n"+string(raw),
 }
 if err=writeJSONAtomic(filepath.Join(work,"task.json"),task);err!=nil{return task,"",err}
 return task,work,nil
}
func openAIPromptsCommand(args []string)error{
 if len(args)==0||args[0]!="list"{return errors.New("uso: openai prompts list [--stage charter|scope|implementation|src]")}
 fs:=flag.NewFlagSet("openai prompts list",flag.ContinueOnError)
 stage:=fs.String("stage","","filtrar por etapa")
 if err:=fs.Parse(args[1:]);err!=nil{return err}
 if *stage!=""&&!supportedStages[*stage]{return errors.New("etapa inválida")}
 t,err:=loadCommandTree();if err!=nil{return err}
 for _,s:=range []string{"charter","scope","implementation","src"} {
  if *stage!=""&&*stage!=s{continue}
  fmt.Println(s)
  for _,c:=range t.Stages[s] {
   indicator:=""
   if !c.Executable{indicator=" (PENDING)"}
   fmt.Printf("  %s / %s -> %s%s\n",c.Operation,c.Action,c.PromptPath,indicator)
  }
 }
 return nil
}
func openAIAnalyzeCommand(p appPaths,args []string)error{
 if len(args)==0{return errors.New("uso: openai analyze <prepare|apply>")}
 if args[0]=="apply" {
  fs:=flag.NewFlagSet("openai analyze apply",flag.ContinueOnError)
  dir:=fs.String("task-dir","","caminho da tarefa preparada")
  changes:=fs.String("changeset","","retorno JSON do agente")
  confirm:=fs.String("confirm","","task_id para confirmar escrita local")
  apply:=fs.Bool("apply",false,"aplicar alterações locais")
  if err:=fs.Parse(args[1:]);err!=nil{return err}
  if *dir==""||*changes==""{return errors.New("analyze apply exige --task-dir e --changeset")}
  return validateAndApplyChangeset(*dir,*changes,*confirm,*apply)
 }
 if args[0]!="prepare"{return errors.New("uso: openai analyze prepare --repo URL --stage ETAPA --operation VERBO --action NOME --request TEXTO")}
 fs:=flag.NewFlagSet("openai analyze prepare",flag.ContinueOnError)
 repo:=fs.String("repo","","GitHub URL")
 stage:=fs.String("stage","","etapa")
 op:=fs.String("operation","","operação")
 action:=fs.String("action","","ação")
 req:=fs.String("request","","solicitação")
 if err:=fs.Parse(args[1:]);err!=nil{return err}
 task,dir,err:=prepareOpenAIAnalysis(p,*repo,*stage,*op,*action,*req)
 if err!=nil{return err}
 fmt.Printf("Preparação concluída: %s\nBranch local: %s\nBase develop: %s\nReferência main: %s\nPrompt: %s\nArquivo: %s\n",task.TaskID,task.LocalBranch,task.TargetDevelopSHA,task.NormativeSHA,task.Prompt,filepath.Join(dir,"task.json"))
 fmt.Println("Somente preparação local: nenhuma inferência paga, commit ou push foi executado.")
 return nil
}
func openAIChangesetDecodeFile(path string)(map[string]any,error){
 info,err:=os.Lstat(path);if err!=nil{return nil,err}
 if !info.Mode().IsRegular()||info.Size()>2<<20{return nil,errors.New("arquivo inválido ou maior que 2 MiB")}
 f,err:=os.Open(path);if err!=nil{return nil,err};defer f.Close()
 var obj map[string]any
 dec:=json.NewDecoder(io.LimitReader(f,2<<20));if err=dec.Decode(&obj);err!=nil{return nil,err}
 if obj["format"]!="PCDSL-CHANGESET/1"{return nil,errors.New("changeset deve usar PCDSL-CHANGESET/1")}
 return obj,nil
}
