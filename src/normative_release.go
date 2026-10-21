package main

import (
 "archive/zip"
 "bytes"
 "context"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "net/http"
 "os"
 "path/filepath"
 "regexp"
 "sort"
 "strings"
 "time"
)

const normativeAPI = "https://api.github.com/repos/ahhvahh/doc-arquitetura"
const releasePrefix = "projectcharter-1.0-"
const packageFilename = "projectcharter-1.0.zip"

type releaseAsset struct { Name string `json:"name"`; URL string `json:"url"` }
type normativeRelease struct { Tag string `json:"tag_name"`; Draft bool `json:"draft"`; Prerelease bool `json:"prerelease"`; Assets []releaseAsset `json:"assets"` }
type downloadedDSL struct {
 Tag string `json:"tag"`
 SHA string `json:"sha"`
 ArchiveSHA256 string `json:"archive_sha256"`
 DownloadedAt string `json:"downloaded_at"`
}
type normativeManifest struct {
 Commit string `json:"commit"`
 Schema string `json:"schema"`
 Packages []struct { Name string `json:"name"`; SHA256 string `json:"sha256"`; Entries int `json:"entries"` } `json:"packages"`
}
var sha40 = regexp.MustCompile("^[a-f0-9]{40}$")
var hash64 = regexp.MustCompile("^[a-f0-9]{64}$")
var bundledNames = []string{"charter.dsl","scope.dsl","implementation.dsl","source.dsl","prompts.dsl","manifest.json"}

func apiRequest(ctx context.Context, url, accept string)([]byte,error){
 cfg,err:=githubLoad();if err!=nil{return nil,err}
 req,err:=http.NewRequestWithContext(ctx,http.MethodGet,url,nil);if err!=nil{return nil,err}
 req.Header.Set("Accept",accept)
 req.Header.Set("X-GitHub-Api-Version","2022-11-28")
 req.Header.Set("User-Agent","doc-arquitetura-app-debian")
 if cfg!=nil{token,e:=cfg.secret();if e!=nil{return nil,e};req.Header.Set("Authorization","Bearer "+token)}
 client:= &http.Client{Timeout:90*time.Second,CheckRedirect:func(req *http.Request,via []*http.Request)error{
   if len(via)>=5{return errors.New("excesso de redirecionamentos")}
   host:=strings.ToLower(req.URL.Hostname())
   if req.URL.Scheme!="https" || (host!="github.com" && host!="api.github.com" && !strings.HasSuffix(host,".githubusercontent.com")){return errors.New("redirecionamento não autorizado")}
   req.Header.Del("Authorization")
   return nil
 }}
 resp,err:=client.Do(req);if err!=nil{return nil,err};defer resp.Body.Close()
 if resp.StatusCode!=200{return nil,fmt.Errorf("GitHub HTTP %d (%s)",resp.StatusCode,url)}
 return io.ReadAll(io.LimitReader(resp.Body,32<<20))
}
func selectRelease(releases []normativeRelease)(normativeRelease,error){
 for _,r:=range releases{
  sha:=strings.TrimPrefix(r.Tag,releasePrefix)
  if r.Draft||r.Prerelease||!strings.HasPrefix(r.Tag,releasePrefix)||!sha40.MatchString(sha){continue}
  names:=map[string]bool{};for _,a:=range r.Assets{names[a.Name]=true}
  if names[packageFilename] && names[packageFilename+".sha256"] {return r,nil}
 }
 return normativeRelease{},errors.New("nenhum Release ProjectCharter DSL estável encontrado")
}
func releaseAssetURL(r normativeRelease, name string)(string,error){
 for _,a:=range r.Assets{if a.Name==name && strings.HasPrefix(a.URL,normativeAPI+"/releases/assets/"){return a.URL,nil}}
 return "",fmt.Errorf("asset inválido ou ausente: %s",name)
}
func verifyDSLArchive(data, checksum []byte, tag string)(map[string][]byte,normativeManifest,error){
 expectedSHA:=strings.TrimPrefix(tag,releasePrefix)
 sum:=sha256.Sum256(data)
 actual:=hex.EncodeToString(sum[:])
 fields:=strings.Fields(string(checksum))
 if len(fields)!=2 || fields[0]!=actual || (fields[1]!=packageFilename && fields[1]!="*"+packageFilename){return nil,normativeManifest{},errors.New("checksum ZIP divergente")}
 zr,err:=zip.NewReader(bytes.NewReader(data),int64(len(data)));if err!=nil{return nil,normativeManifest{},err}
 content:=map[string][]byte{}
 allowed:=map[string]bool{};for _,s:=range bundledNames{allowed[s]=true}
 for _,file:=range zr.File{
  if !allowed[file.Name]||file.FileInfo().Mode()&os.ModeSymlink!=0||file.UncompressedSize64>16<<20||len(content)>=len(bundledNames){return nil,normativeManifest{},errors.New("arquivo ZIP inesperado ou inválido")}
  if _,ok:=content[file.Name];ok{return nil,normativeManifest{},errors.New("arquivo ZIP duplicado")}
  reader,e:=file.Open();if e!=nil{return nil,normativeManifest{},e}
  blob,e:=io.ReadAll(io.LimitReader(reader,16<<20+1));reader.Close();if e!=nil{return nil,normativeManifest{},e}
  if len(blob)>16<<20{return nil,normativeManifest{},errors.New("arquivo ZIP muito grande")}
  content[file.Name]=blob
 }
 if len(content)!=len(bundledNames){return nil,normativeManifest{},errors.New("arquivos DSL obrigatórios ausentes")}
 var manifest normativeManifest
 if err=json.Unmarshal(content["manifest.json"],&manifest);err!=nil{return nil,manifest,err}
 if manifest.Schema!="projectcharter-package/1" || manifest.Commit!=expectedSHA {return nil,manifest,errors.New("manifest não corresponde ao Release")}
 if len(manifest.Packages)!=5{return nil,manifest,errors.New("lista de pacotes incompleta")}
 seen:=map[string]bool{}
 for _,p:=range manifest.Packages{
  if p.Name=="manifest.json"||seen[p.Name]||p.Entries<=0||!hash64.MatchString(p.SHA256){return nil,manifest,errors.New("pacote inválido")}
  blob,ok:=content[p.Name];if !ok{return nil,manifest,errors.New("DSL desconhecida")}
  h:=sha256.Sum256(blob);if hex.EncodeToString(h[:])!=p.SHA256{return nil,manifest,errors.New("checksum DSL divergente: "+p.Name)}
  if !bytes.HasPrefix(blob,[]byte("PROJECTCHARTER-DSL/1
")){return nil,manifest,errors.New("cabeçalho DSL inválido: "+p.Name)}
  seen[p.Name]=true
 }
 return content,manifest,nil
}
func releaseCurrentFile(p appPaths)string{return filepath.Join(dslBase(p),"release-current.json")}
func dslReleaseUpdate(p appPaths)error{
 ctx,cancel:=context.WithTimeout(context.Background(),4*time.Minute);defer cancel()
 body,err:=apiRequest(ctx,normativeAPI+"/releases?per_page=30","application/vnd.github+json");if err!=nil{return err}
 var releases []normativeRelease;if err=json.Unmarshal(body,&releases);err!=nil{return err}
 latest,err:=selectRelease(releases);if err!=nil{return err}
 sha:=strings.TrimPrefix(latest.Tag,releasePrefix)
 archiveURL,err:=releaseAssetURL(latest,packageFilename);if err!=nil{return err}
 checksumURL,err:=releaseAssetURL(latest,packageFilename+".sha256");if err!=nil{return err}
 archive,err:=apiRequest(ctx,archiveURL,"application/octet-stream");if err!=nil{return err}
 checksum,err:=apiRequest(ctx,checksumURL,"application/octet-stream");if err!=nil{return err}
 content,_,err:=verifyDSLArchive(archive,checksum,latest.Tag);if err!=nil{return err}
 base:=dslBase(p);if err=mkdirPrivate(base);err!=nil{return err}
 target:=filepath.Join(base,sha)
 if _,err=os.Stat(target);os.IsNotExist(err){
  stage,e:=os.MkdirTemp(base,".release-*");if e!=nil{return e};defer os.RemoveAll(stage)
  for name,data:=range content{if e=os.WriteFile(filepath.Join(stage,name),data,0600);e!=nil{return e}}
  if e=os.Rename(stage,target);e!=nil{return e}
 }else if err!=nil{return err}
 record:=downloadedDSL{Tag:latest.Tag,SHA:sha,ArchiveSHA256:strings.Fields(string(checksum))[0],DownloadedAt:time.Now().UTC().Format(time.RFC3339)}
 if err=writeJSONAtomic(releaseCurrentFile(p),record);err!=nil{return err}
 names:=append([]string{},bundledNames...);sort.Strings(names)
 fmt.Printf("ProjectCharter DSL recebido: %s (%s)
Arquivos: %s
",latest.Tag,sha,strings.Join(names,", "))
 return nil
}
func dslReleaseStatus(p appPaths)error{
 var record downloadedDSL
 err:=readJSON(releaseCurrentFile(p),&record)
 if os.IsNotExist(err){return dslStatus(p)}
 if err!=nil{return err}
 if !sha40.MatchString(record.SHA)||record.Tag!=releasePrefix+record.SHA{return errors.New("registro de Release inválido")}
 for _,name:=range bundledNames{if _,err=os.Stat(filepath.Join(dslBase(p),record.SHA,name));err!=nil{return fmt.Errorf("cache DSL incompleto: %w",err)}}
 fmt.Printf("DSL Release: %s
Diretório: %s
",record.Tag,filepath.Join(dslBase(p),record.SHA))
 return nil
}
