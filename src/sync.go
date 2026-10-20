package main

import (
  "context"
 "crypto/sha256"
 "encoding/hex"
  "errors"
 "flag"
 "fmt"
  "io/fs"
   "os"
  "path/filepath"
 "regexp"
 "sort"
 "strings"
 "time"
)

var statusBadge = regexp.MustCompile("badge-status-([a-z]+)\\.svg")
var badgeExpr = regexp.MustCompile("!\\[[^\\]]*\\]\\([^)]*badge-(?:type|status|version)-[^)]*\\)")

func compactDoc(input string) string {
 lines := strings.Split(strings.ReplaceAll(input,"\r\n","\n"),"\n")
 out:=make([]string,0,len(lines)); prevBlank:=false
 for _,s:=range lines {
  s=strings.TrimSpace(badgeExpr.ReplaceAllString(strings.TrimRight(s," \t"),""))
  if s=="" { if prevBlank {continue};prevBlank=true } else {prevBlank=false}
  out=append(out,s)
 }
 return strings.TrimSpace(strings.Join(out,"\n"))
}

func convertCharter(source,destination,origin string)(int,error){
 var docs []string
 err:=filepath.WalkDir(source,func(p string,d fs.DirEntry,e error)error{
  if e!=nil{return e};if d.Type()&os.ModeSymlink!=0{return nil};if d.IsDir(){
       relative,er:=filepath.Rel(source,p);if er!=nil{return er}
       normalized:=filepath.ToSlash(relative)
       if normalized=="backup"||strings.HasPrefix(normalized,"backup/"){return filepath.SkipDir}
       return nil
     }
  if strings.HasSuffix(strings.ToLower(p),".md"){docs=append(docs,p)};return nil
 });if err!=nil{return 0,err};if len(docs)==0{return 0,errors.New("Charter nao contem Markdown")}
 sort.Strings(docs)
 if err=os.MkdirAll(filepath.Dir(destination),0755);err!=nil{return 0,err}
 f,err:=os.Create(destination+".tmp");if err!=nil{return 0,err}
 defer func(){f.Close();os.Remove(destination+".tmp")}()
 fmt.Fprintf(f,"PCDSL-DOCS 1\nORIGIN %s\nDOCUMENTS %d\n",origin,len(docs))
 for _,file:=range docs {
  raw,e:=os.ReadFile(file);if e!=nil{return 0,e};relative,e:=filepath.Rel(source,file);if e!=nil{return 0,e}
  sha:=sha256.Sum256(raw);status:="unknown";if m:=statusBadge.FindStringSubmatch(string(raw));len(m)>1{status=m[1]}
  content:=compactDoc(string(raw))
  fmt.Fprintf(f,"DOC %q %s %s %d\n%s\n",filepath.ToSlash(relative),hex.EncodeToString(sha[:]),status,len([]byte(content)),content)
 }
 if err=f.Sync();err!=nil{return 0,err};if err=f.Close();err!=nil{return 0,err};if err=os.Rename(destination+".tmp",destination);err!=nil{return 0,err}
 return len(docs),nil
}

func findCharter(repo string,reference bool)(string,error){
 choices:=[]string{"docs/ProjectCharter/1.0","docs/01-Project_Charter","ProjectCharter/1.0"}
  if !reference {
    docs:=filepath.Join(repo,"docs")
    for _, area:=range []string{"charter","scope","implementation"} {
      st,e:=os.Lstat(filepath.Join(docs,area))
      if e==nil && st.IsDir() {return docs,nil}
    }
  }
 if reference {choices=[]string{"ProjectCharter/1.0"}}
 for _,s:=range choices{p:=filepath.Join(repo,filepath.FromSlash(s));if x,e:=os.Stat(p);e==nil&&x.IsDir(){return p,nil}}
 return "",errors.New("documentação não encontrada: esperado /docs/{charter,scope,implementation} no projeto-alvo ou ProjectCharter/1.0 na referência normativa")
}

func cloneReference(ref,dest string)error{
 if ref!="main"{return errors.New("referencia normativa permitida somente na branch main")}
 if _,e:=os.Stat(dest);e==nil{return errors.New("destino ja existe: "+dest)}
 if e:=os.MkdirAll(filepath.Dir(dest),0755);e!=nil{return e}
 ctx,cancel:=context.WithTimeout(context.Background(),4*time.Minute);defer cancel()
 cmd,e:=githubGitCommand(ctx,"-c","protocol.file.allow=never","clone","--depth","1","--single-branch","--branch",ref,"--no-recurse-submodules","--","https://github.com/ahhvahh/doc-arquitetura.git",dest)
 if e!=nil{return e}
 cmd.Stderr=os.Stderr;return cmd.Run()
}

func runExtra(args []string)error{
 f:=flag.NewFlagSet(args[0],flag.ContinueOnError)
 target:=f.String("target","","URL GitHub do cliente")
 work:=f.String("workspace","./workspace","workspace local")
 ref:=f.String("norm-ref","main","referencia normativa: somente main")
 src:=f.String("source","","pasta Charter a converter")
 out:=f.String("output","","arquivo .pcdsl")
 origin:=f.String("origin","local","origem dos arquivos")
 if e:=f.Parse(args[1:]);e!=nil{return e}
 switch args[0]{
 case "sync":
  if *target==""{return errors.New("sync exige -target")}
  if *ref!="main"{return errors.New("-norm-ref deve ser main")}
  root,e:=filepath.Abs(*work);if e!=nil{return e}
  client:=filepath.Join(root,"client");reference:=filepath.Join(root,"reference")
  if e=cloneRepo(*target,"develop",client);e!=nil{return e}
  if e=cloneReference(*ref,reference);e!=nil{return e}
  p,e:=findCharter(reference,true);if e!=nil{return e}
  n,e:=convertCharter(p,filepath.Join(root,"reference.pcdsl"),"doc-arquitetura@"+*ref);if e!=nil{return e}
  fmt.Printf("Referencia ProjectCharter/1.0: %d documentos convertidos\n",n)
  p,e=findCharter(client,false);if e!=nil{fmt.Fprintln(os.Stderr,"Aviso: cliente sem ProjectCharter; etapas dependentes bloqueadas");return nil}
  n,e=convertCharter(p,filepath.Join(root,"client.pcdsl"),*target+"@develop");if e!=nil{return e}
  fmt.Printf("Cliente ProjectCharter: %d documentos convertidos\n",n);return nil
 case "convert":
  if *src==""||*out==""{return errors.New("convert exige -source e -output")}
  n,e:=convertCharter(*src,*out,*origin);if e==nil{fmt.Printf("%d documentos convertidos\n",n)};return e
 default:return errors.New("comando desconhecido")
 }
}
