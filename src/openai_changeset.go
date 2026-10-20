package main

import (
 "context"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "os"
 "path/filepath"
 "strings"
)

// DOC-REF: OBJ-003
type markdownChangeSet struct {
 Format string `json:"format"`
 TaskID string `json:"task_id"`
 Stage string `json:"stage"`
 Prompt string `json:"prompt"`
 NormativeSHA string `json:"normative_sha"`
 TargetDevelopSHA string `json:"target_develop_sha"`
 Changes []markdownChange `json:"changes"`
}
type markdownChange struct {
 Operation string `json:"operation"`
 Path string `json:"path"`
 BaseSHA256 string `json:"base_sha256,omitempty"`
 Heading string `json:"heading,omitempty"`
 HeadingOccurrence int `json:"heading_occurrence,omitempty"`
 NewMarkdown string `json:"new_markdown"`
 Reason string `json:"reason"`
}
func decodeChangeset(reader io.Reader)(markdownChangeSet,error){
 var s markdownChangeSet
 d:=json.NewDecoder(io.LimitReader(reader,2<<20))
 d.DisallowUnknownFields()
 if err:=d.Decode(&s);err!=nil{return s,err}
 var extra any
 if err:=d.Decode(&extra);err!=io.EOF{return s,errors.New("changeset JSON contém dados extras")}
 if s.Format!="PCDSL-CHANGESET/1"||s.TaskID==""||!supportedStages[s.Stage]||len(s.Changes)==0||len(s.Changes)>8 {
  return s,errors.New("formato, etapa ou quantidade de alterações inválidos")
 }
 seen:=map[string]bool{}
 for _,c:=range s.Changes {
  if c.Path==""||seen[c.Path]||c.Reason==""||len(c.NewMarkdown)>250000{return s,errors.New("changeset contém caminho duplicado, justificativa vazia ou conteúdo excessivo")}
  seen[c.Path]=true
  if c.Operation!="create_file"&&c.Operation!="replace_section" {return s,errors.New("operação não permitida: "+c.Operation)}
 }
 return s,nil
}
func loadChangesetFile(path string)(markdownChangeSet,error){
 var empty markdownChangeSet
 i,err:=os.Lstat(path);if err!=nil{return empty,err}
 if !i.Mode().IsRegular()||i.Size()>2<<20{return empty,errors.New("changeset deve ser arquivo regular menor que 2 MiB")}
 f,err:=os.Open(path);if err!=nil{return empty,err};defer f.Close()
 return decodeChangeset(f)
}
func safeTargetMarkdownPath(workspace,stage,path string)(string,error){
 if !supportedStages[stage]||stage=="src"{return "",errors.New("alterações de fonte exigem contrato de Development próprio; Markdown permitido apenas em Charter/Scope/Implementation")}
 wanted:="docs/"+stage+"/"
 if !strings.HasPrefix(path,wanted)||!strings.HasSuffix(path,".md")||filepath.IsAbs(path)||strings.Contains(path,"\\")||
  filepath.ToSlash(filepath.Clean(path))!=path||strings.Contains(path,"..") {return "",errors.New("destino documental não autorizado")}
 abs:=filepath.Join(workspace,filepath.FromSlash(path))
 base:=filepath.Join(workspace,"docs",stage)
 if !strings.HasPrefix(abs,base+string(filepath.Separator)){return "",errors.New("destino fora da etapa")}
 for dir:=filepath.Dir(abs);;dir=filepath.Dir(dir){
  if i,e:=os.Lstat(dir);e==nil&&i.Mode()&os.ModeSymlink!=0{return "",errors.New("diretório ancestral é symlink")}
  if dir==workspace||dir==filepath.Dir(dir){break}
 }
 return abs,nil
}
func mdDigest(b []byte)string {s:=sha256.Sum256(b);return hex.EncodeToString(s[:])}
func replaceMarkdownSection(original []byte, heading string, occurrence int, replacement string)([]byte,error){
 if !strings.HasPrefix(heading,"## ")&&!strings.HasPrefix(heading,"### "){return nil,errors.New("heading deve ser ## ou ###")}
 if occurrence<1||occurrence>20{return nil,errors.New("ocorrência de heading inválida")}
 if !strings.HasPrefix(replacement,heading+"\n")&&!strings.HasPrefix(replacement,heading+"\r\n"){
  return nil,errors.New("nova seção deve manter o mesmo heading")}
 source:=string(original)
 lines:=strings.SplitAfter(source,"\n")
 counter:=0;start:=-1;end:=len(source);offset:=0
 level:=len(heading)-len(strings.TrimLeft(heading,"#"))
 for _,line:=range lines {
  clean:=strings.TrimSuffix(strings.TrimSuffix(line,"\n"),"\r")
  if start<0{
   if clean==heading{
    counter++
    if counter==occurrence{start=offset}
   }
  }else if strings.HasPrefix(clean,"#"){
    hashes:=len(clean)-len(strings.TrimLeft(clean,"#"))
    if hashes>0&&hashes<=level&&strings.HasPrefix(clean,strings.Repeat("#",hashes)+" "){end=offset;break}
  }
  offset+=len(line)
 }
 if start<0{return nil,fmt.Errorf("heading não encontrado (%d): %s",occurrence,heading)}
 result:=source[:start]+strings.TrimRight(replacement,"\r\n")+"\n\n"+source[end:]
 return []byte(result),nil
}
type pendingMarkdownWrite struct{Path string;Data []byte;IsNew bool}
func planChangeset(workspace string,task preparedTask,s markdownChangeSet)([]pendingMarkdownWrite,error){
 if s.TaskID!=task.TaskID||s.Stage!=task.Stage||s.Prompt!=task.Prompt||
   s.NormativeSHA!=task.NormativeSHA||s.TargetDevelopSHA!=task.TargetDevelopSHA{
   return nil,errors.New("changeset não corresponde ao task_id/prompt/SHA da sessão")
 }
 var writes []pendingMarkdownWrite
 for _,op:=range s.Changes{
  dst,err:=safeTargetMarkdownPath(workspace,task.Stage,op.Path);if err!=nil{return nil,err}
  info,e:=os.Lstat(dst)
  switch op.Operation {
  case "create_file":
   if e==nil{return nil,errors.New("arquivo já existe: "+op.Path)}
   if !os.IsNotExist(e){return nil,e}
   if op.BaseSHA256!=""||!strings.HasPrefix(strings.TrimSpace(op.NewMarkdown),"# "){
    return nil,errors.New("create_file exige Markdown com H1 e sem base_sha256")
   }
   if !strings.Contains(op.NewMarkdown,"badge-type-")||!strings.Contains(op.NewMarkdown,"badge-status-")||!strings.Contains(op.NewMarkdown,"badge-version-"){
    return nil,errors.New("novo documento deve incluir badges normativos")
   }
   writes=append(writes,pendingMarkdownWrite{Path:dst,Data:[]byte(op.NewMarkdown),IsNew:true})
  case "replace_section":
   if e!=nil{return nil,e}
   if !info.Mode().IsRegular(){return nil,errors.New("arquivo original não regular")}
   raw,err:=os.ReadFile(dst);if err!=nil{return nil,err}
   if op.BaseSHA256==""||mdDigest(raw)!=strings.ToLower(op.BaseSHA256){
    return nil,errors.New("SHA-256 divergente: "+op.Path)
   }
   modified,err:=replaceMarkdownSection(raw,op.Heading,op.HeadingOccurrence,op.NewMarkdown)
   if err!=nil{return nil,err}
   writes=append(writes,pendingMarkdownWrite{Path:dst,Data:modified})
  }
 }
 return writes,nil
}
func validateAndApplyChangeset(taskDir,changesPath,confirm string,apply bool)error{
 taskPath:=filepath.Join(taskDir,"task.json")
 ti,err:=os.Lstat(taskPath);if err!=nil{return err}
 if !ti.Mode().IsRegular(){return errors.New("task.json inválido")}
 var task preparedTask
 if err:=readJSON(taskPath,&task);err!=nil{return err}
 if task.Format!="PCDSL-REQUEST/1"||task.TaskID==""{return errors.New("tarefa inválida")}
 cs,err:=loadChangesetFile(changesPath);if err!=nil{return err}
 workspace:=filepath.Join(taskDir,"repository")
 if i,e:=os.Lstat(workspace);e!=nil||!i.IsDir(){return errors.New("workspace de tarefa ausente")}
 ctx:=context.Background()
 branch,e:=gitSession(ctx,workspace,"branch","--show-current");if e!=nil{return e}
 if strings.TrimSpace(branch)!=task.LocalBranch{return errors.New("branch local não corresponde à tarefa")}
 head,e:=gitSession(ctx,workspace,"rev-parse","HEAD");if e!=nil{return e}
 if strings.TrimSpace(head)!=task.TargetDevelopSHA{return errors.New("HEAD local diverge da develop original")}
 status,e:=gitSession(ctx,workspace,"status","--porcelain");if e!=nil{return e}
 if strings.TrimSpace(status)!=""{return errors.New("workspace já possui modificações: revisão obrigatória")}
 writes,err:=planChangeset(workspace,task,cs);if err!=nil{return err}
 if !apply{fmt.Printf("PREVIEW: %d alteração(ões) válidas na branch local %s; nenhuma escrita.\n",len(writes),task.LocalBranch);return nil}
 if confirm!=task.TaskID{return errors.New("para aplicar, confirme --confirm com o task_id")}
 // Stage files and validate the complete change set before replacing originals.
 staged:=make([]string,len(writes))
 for i,w:=range writes {
  if err=os.MkdirAll(filepath.Dir(w.Path),0700);err!=nil{return err}
  f,e:=os.CreateTemp(filepath.Dir(w.Path),".dsl-staged-*")
  if e!=nil{return e}
  staged[i]=f.Name()
  if e=f.Chmod(0600);e!=nil{f.Close();return e}
  if _,e=f.Write(w.Data);e!=nil{f.Close();return e}
  if e=f.Close();e!=nil{return e}
 }
 defer func(){for _,f:=range staged{if f!=""{_ = os.Remove(f)}}}()
 // Only locally, and only after the explicit task ID confirmation.
 for i,w:=range writes{
  if e:=os.Rename(staged[i],w.Path);e!=nil{return e}
  staged[i]=""
 }
 target:=filepath.Join(taskDir,"updated.pcdsl")
 n,err:=convertCharter(filepath.Join(workspace,"docs"),target,task.TargetRepository+"@local-task")
 if err!=nil{return err}
 fmt.Printf("APLICADO LOCALMENTE: %d mudanças; DSL pós-alteração com %d documentos em %s. Nenhum commit/push foi executado.\n",len(writes),n,target)
 return nil
}
