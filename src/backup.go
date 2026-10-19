package main

import (
    "context"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
    "sort"
    "strings"
    "time"
)

const backupRoot = "docs/backup"
const maxBackupFileSize int64 = 4 << 20
const maxBackupTotalSize int64 = 32 << 20

type backupItem struct { Source string; Size int64 }
type backupEntry struct {
    SourcePath string `json:"source_path"`
    BackupPath string `json:"backup_path"`
    SHA256 string `json:"sha256"`
    SizeBytes int64 `json:"size_bytes"`
    ProposedDestination string `json:"proposed_destination,omitempty"`
}
type backupManifest struct {
    Schema int `json:"schema"`
    SnapshotID string `json:"snapshot_id"`
    SourceRepository string `json:"source_repository"`
    SourceBranch string `json:"source_branch"`
    SourceHead string `json:"source_head"`
    NormativeHead string `json:"normative_head"`
    CreatedAt string `json:"created_at"`
    Files []backupEntry `json:"files"`
}

func documentSourceExtension(name string) bool {
    switch strings.ToLower(filepath.Ext(name)) {
    case ".md", ".markdown", ".svg", ".png", ".jpg", ".jpeg", ".webp", ".drawio", ".json", ".yaml", ".yml", ".pdf":
        return true
    }
    return false
}
func withinRepository(root, relative string) (string, error) {
    if relative == "" || filepath.IsAbs(relative) || strings.Contains(relative, "\\") || filepath.ToSlash(filepath.Clean(relative)) != relative ||
        relative == "." || strings.HasPrefix(relative, "../") || strings.ContainsAny(relative, "\x00\r\n") ||
        strings.HasPrefix(relative, ".git/") || relative == ".git" ||
        strings.HasPrefix(relative, backupRoot+"/") || relative == backupRoot ||
        (relative=="src" || strings.HasPrefix(relative,"src/")) {
        return "", errors.New("origem não autorizada para backup documental: "+relative)
    }
    current := root
    for _,part := range strings.Split(relative,"/") {
        if part == "." || part == ".." || part == "" { return "",errors.New("origem inválida") }
        current=filepath.Join(current,part)
        st,err:=os.Lstat(current)
        if os.IsNotExist(err) { continue }
        if err!=nil {return "",err}
        if st.Mode() & os.ModeSymlink != 0 {return "",errors.New("symlink rejeitado: "+relative)}
    }
    return filepath.Join(root,filepath.FromSlash(relative)),nil
}
func (a *terminalApp) backupSources() ([]backupItem,error) {
    area:=a.session.Area
    roots:=[]string{
        "docs/"+area,
        "docs/ProjectCharter/1.0/"+area,
        "docs/01-Project_Charter/"+area,
        "ProjectCharter/1.0/"+area,
    }
    roots=append(roots,a.session.BackupSources...)
    paths:=map[string]int64{}
    var total int64
    for _,relative:=range roots {
        absolute,err:=withinRepository(a.session.Workspace,relative)
        if err!=nil{return nil,err}
        st,err:=os.Lstat(absolute)
        if os.IsNotExist(err){continue}
        if err!=nil{return nil,err}
        if st.Mode()&os.ModeSymlink!=0 {return nil,errors.New("origem symlink: "+relative)}
        visit:=func(path string, de fs.DirEntry, err error) error {
            if err!=nil{return err}
            rp,err:=filepath.Rel(a.session.Workspace,path);if err!=nil{return err}
            rp=filepath.ToSlash(rp)
            if strings.HasPrefix(rp,backupRoot+"/") {return filepath.SkipDir}
            if de.Type()&os.ModeSymlink!=0{return fmt.Errorf("symlink em fonte documental: %s",rp)}
            if de.IsDir(){return nil}
            if !de.Type().IsRegular() || !documentSourceExtension(path){return nil}
            if _,err:=withinRepository(a.session.Workspace,rp);err!=nil{return err}
            info,err:=de.Info();if err!=nil{return err}
            if info.Size()>maxBackupFileSize {return fmt.Errorf("arquivo excessivamente grande para backup: %s",rp)}
            if _,ok:=paths[rp];!ok{
                paths[rp]=info.Size();total+=info.Size()
                if len(paths)>300||total>maxBackupTotalSize {return errors.New("backup excede limite de 300 arquivos ou 32 MiB; reduza as fontes")}
            }
            return nil
        }
        if st.IsDir(){if err=filepath.WalkDir(absolute,visit);err!=nil{return nil,err}}else{
            de,err:=os.ReadDir(filepath.Dir(absolute));if err!=nil{return nil,err}
            for _,item:=range de{if item.Name()==filepath.Base(absolute){if err=visit(absolute,item,nil);err!=nil{return nil,err};break}}
        }
    }
    result:=make([]backupItem,0,len(paths))
    for p,n:=range paths{result=append(result,backupItem{p,n})}
    sort.Slice(result,func(i,j int)bool{return result[i].Source<result[j].Source})
    return result,nil
}
func (a *terminalApp) backupPlan() ([]backupItem,error) {
    files,err:=a.backupSources();if err!=nil{return nil,err}
    fmt.Fprintf(a.out,"\nFontes documentais para cópia em /%s/: %d\n",backupRoot,len(files))
    var total int64
    for _,item:=range files{fmt.Fprintf(a.out,"  %-70s %d bytes\n",item.Source,item.Size);total+=item.Size}
    fmt.Fprintf(a.out,"Total: %d bytes. Revise possíveis segredos e dados sensíveis ANTES da cópia e do push.\n",total)
    return files,nil
}
func migrationTarget(area,path string) string {
    if strings.HasPrefix(path,"docs/"+area+"/"){return path}
    for _,root:=range []string{"docs/ProjectCharter/1.0/"+area+"/","docs/01-Project_Charter/"+area+"/","ProjectCharter/1.0/"+area+"/"}{
        if strings.HasPrefix(path,root){return "docs/"+area+"/"+strings.TrimPrefix(path,root)}
    }
    return "" // análise humana obrigatória para origens não canônicas
}
func (a *terminalApp) migrationPlan() error {
    files,err:=a.backupPlan();if err!=nil{return err}
    for _,item:=range files{
        dest:=migrationTarget(a.session.Area,item.Source)
        if dest==""{fmt.Fprintf(a.out,"  REVISAR destino: %s\n",item.Source);continue}
        if dest==item.Source {fmt.Fprintf(a.out,"  PRESERVAR: %s\n",item.Source);continue}
        path,err:=withinRepository(a.session.Workspace,dest);if err!=nil{return err}
        if _, err = os.Lstat(path); err == nil {
            fmt.Fprintf(a.out, "  CONFLITO: %s -> %s (destino já existe)\n", item.Source, dest)
        } else if !os.IsNotExist(err) {
            return err
        } else {
            fmt.Fprintf(a.out, "  PROPOSTO: %s -> %s\n", item.Source, dest)
        }
    }
    fmt.Fprintln(a.out,"Plano é preliminar; o agente deve atualizar documentos e links. Nenhum arquivo foi movido.")
    return nil
}
func (a *terminalApp) backupDir() string {
    return filepath.Join(a.session.Workspace,filepath.FromSlash(backupRoot),a.session.BackupID)
}
func (a *terminalApp) backupCreate() error {
    if a.session.Branch!="develop" {return errors.New("backup versionado permitido somente em develop")}
    if a.session.BackupID!="" {return errors.New("snapshot já criado nesta sessão: "+a.session.BackupID)}
    ctx, cancel := context.WithTimeout(context.Background(),20*time.Second);defer cancel()
    status,err:=gitSession(ctx,a.session.Workspace,"status","--porcelain=v1","--untracked-files=all")
    if err!=nil{return err}
    if status!="" {return errors.New("workspace precisa estar limpo antes do snapshot")}
    head,err:=gitSession(ctx,a.session.Workspace,"rev-parse","HEAD")
    if err!=nil{return err}
    if head!=a.session.BaseSHA{return errors.New("HEAD da sessão divergiu; reinicie o inventário")}
    files,err:=a.backupPlan();if err!=nil{return err}
    if len(files)==0{return errors.New("nenhum arquivo documental encontrado para backup; inclua uma fonte com /backup add")}
    ok,err:=a.confirmation("Autoriza copiar estas fontes para /docs/backup/ e registrar no commit?");if err!=nil{return err};if !ok{return nil}
    base:=filepath.Join(a.session.Workspace,filepath.FromSlash(backupRoot))
    if _,err=withinRepository(a.session.Workspace,"docs");err!=nil{return err}
    if st,e:=os.Lstat(base);e==nil&&st.Mode()&os.ModeSymlink!=0{return errors.New("/docs/backup é um symlink")}
    if err=os.MkdirAll(base,0700);err!=nil{return err}
    snapshot:=time.Now().UTC().Format("20060102T150405Z")+"-"+head[:12]
    dest:=filepath.Join(base,snapshot)
    if _,err=os.Lstat(dest);err==nil{return errors.New("snapshot já existe")}else if !os.IsNotExist(err){return err}
    stage,err:=os.MkdirTemp(base,".preparing-");if err!=nil{return err}
    defer os.RemoveAll(stage)
    manifest:=backupManifest{Schema:1,SnapshotID:snapshot,SourceRepository:a.session.Repo,SourceBranch:a.session.Branch,SourceHead:head,NormativeHead:a.session.ReferenceSHA,CreatedAt:time.Now().UTC().Format(time.RFC3339)}
    for _,item:=range files{
        src,err:=withinRepository(a.session.Workspace,item.Source);if err!=nil{return err}
        data,err:=os.ReadFile(src);if err!=nil{return err}
        if int64(len(data))!=item.Size {return fmt.Errorf("fonte mudou durante backup: %s",item.Source)}
        sum:=sha256.Sum256(data);hash:=hex.EncodeToString(sum[:])
        target:=filepath.Join(stage,"original",filepath.FromSlash(item.Source))
        if err=os.MkdirAll(filepath.Dir(target),0700);err!=nil{return err}
        if err=os.WriteFile(target,data,0600);err!=nil{return err}
        verified,_,err:=shaFile(target);if err!=nil||verified!=hash{return fmt.Errorf("hash do backup não confere: %s: %v",item.Source,err)}
        current,_,err:=shaFile(src);if err!=nil||current!=hash{return fmt.Errorf("fonte mudou após cópia: %s: %v",item.Source,err)}
        manifest.Files=append(manifest.Files,backupEntry{SourcePath:item.Source,BackupPath:backupRoot+"/"+snapshot+"/original/"+item.Source,SHA256:hash,SizeBytes:int64(len(data)),ProposedDestination:migrationTarget(a.session.Area,item.Source)})
    }
    content,err:=json.MarshalIndent(manifest,"","  ");if err!=nil{return err}
    content=append(content,'\n')
    if err=os.WriteFile(filepath.Join(stage,"manifest.json"),content,0600);err!=nil{return err}
    hash:=sha256.Sum256(content)
    if err=os.Rename(stage,dest);err!=nil{return err}
    a.session.BackupID=snapshot
    a.session.BackupManifestSHA=hex.EncodeToString(hash[:])
    if err=a.saveSession();err!=nil{return err}
    fmt.Fprintf(a.out,"Snapshot verificado: /%s/%s/ (%d arquivos). Use /migration plan e depois o agente para propor os novos documentos.\n",backupRoot,snapshot,len(manifest.Files))
    return nil
}
func (a *terminalApp) verifiedBackup() (backupManifest,error) {
    var result backupManifest
    if a.session.BackupID=="" {return result,errors.New("backup ausente; use /backup plan e /backup create")}
    if strings.Contains(a.session.BackupID,"/")||strings.Contains(a.session.BackupID,".."){return result,errors.New("ID de backup inválido")}
    location:=a.backupDir()
    b,err:=os.ReadFile(filepath.Join(location,"manifest.json"));if err!=nil{return result,err}
    sum:=sha256.Sum256(b)
    if hex.EncodeToString(sum[:])!=a.session.BackupManifestSHA{return result,errors.New("manifesto de backup modificado")}
    if err=json.Unmarshal(b,&result);err!=nil{return result,err}
    if result.Schema!=1||result.SnapshotID!=a.session.BackupID||result.SourceRepository!=a.session.Repo||result.SourceHead!=a.session.BaseSHA||result.NormativeHead!=a.session.ReferenceSHA{
        return result,errors.New("metadados de backup divergentes da sessão")
    }
    if len(result.Files)==0{return result,errors.New("backup vazio")}
    seen:=map[string]bool{}
    for _,file:=range result.Files {
        expected:=backupRoot+"/"+a.session.BackupID+"/original/"+file.SourcePath
        if file.BackupPath!=expected||seen[expected]{return result,errors.New("arquivo de backup inválido")}
        seen[expected]=true
        path:=filepath.Join(a.session.Workspace,filepath.FromSlash(expected))
        if _,err=withinRepository(a.session.Workspace,file.SourcePath);err!=nil{return result,err}
        // The backup subtree may only consist of real files/directories.
        current:=a.session.Workspace
        for _,part:=range strings.Split(expected,"/"){
            current=filepath.Join(current,part)
            st,e:=os.Lstat(current);if e!=nil{return result,e}
            if st.Mode()&os.ModeSymlink!=0{return result,errors.New("symlink em backup")}
        }
        verified,n,e:=shaFile(path);if e!=nil||verified!=file.SHA256||n!=file.SizeBytes{return result,fmt.Errorf("snapshot corrompido: %s",expected)}
    }
    return result,nil
}
func (a *terminalApp) backupAllowedPath(path string) bool {
    if a.session.BackupID==""||strings.Contains(a.session.BackupID,"/")||strings.Contains(a.session.BackupID,".."){return false}
    raw,err:=os.ReadFile(filepath.Join(a.backupDir(),"manifest.json"));if err!=nil{return false}
    sum:=sha256.Sum256(raw)
    if hex.EncodeToString(sum[:])!=a.session.BackupManifestSHA{return false}
    var manifest backupManifest
    if json.Unmarshal(raw,&manifest)!=nil||manifest.Schema!=1||manifest.SnapshotID!=a.session.BackupID||
       manifest.SourceRepository!=a.session.Repo||manifest.NormativeHead!=a.session.ReferenceSHA||
       manifest.SourceHead!=a.session.BaseSHA{return false}
    if path==backupRoot+"/"+a.session.BackupID+"/manifest.json"{return true}
    for _,f:=range manifest.Files{if f.BackupPath==path{return true}}
    return false
}

func (a *terminalApp) legacyContext() (string,error) {
    if !strings.HasPrefix(a.session.Prompt,"legacy-"){return "",nil}
    sources,err:=a.backupSources();if err!=nil{return "",err}
    var out strings.Builder
    limit:=12000
    for _,item:=range sources{
        if !strings.HasSuffix(strings.ToLower(item.Source),".md") {continue}
        if strings.HasPrefix(item.Source,"docs/"+a.session.Area+"/"){continue}
        if item.Size>8000 || int(item.Size)>limit {
            fmt.Fprintf(&out,"LEGADO OMITIDO POR TAMANHO: %s (%d bytes). Solicite revisão específica antes de editar.\n",item.Source,item.Size)
            continue
        }
        abs,err:=withinRepository(a.session.Workspace,item.Source);if err!=nil{return "",err}
        b,err:=os.ReadFile(abs);if err!=nil{return "",err}
        fmt.Fprintf(&out,"\n--- ORIGEM LEGADA (SOMENTE LEITURA) %s ---\n%s\n--- FIM ---\n",item.Source,b)
        limit-=len(b)
    }
    return out.String(),nil
}
