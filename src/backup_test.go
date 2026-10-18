package main

import (
    "bufio"
    "bytes"
    "context"
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func TestBackupSourceAndDestination(t *testing.T){
    if got:=migrationTarget("scope","docs/ProjectCharter/1.0/scope/types/TYPE-001.md");got!="docs/scope/types/TYPE-001.md"{t.Fatal(got)}
    if got:=migrationTarget("scope","docs/scope/README.md");got!="docs/scope/README.md"{t.Fatal(got)}
    if got:=migrationTarget("scope","docs/legacy/README.md");got!=""{t.Fatal(got)}
    for _,v:=range []string{"docs/backup/test/original/a.md","src/secret.md","../other.md",".git/config","docs/../../x.md"}{
        if _,err:=withinRepository(t.TempDir(),v);err==nil{t.Errorf("aceitou origem: %s",v)}
    }
}
func TestBackupSnapshotForDevelopment(t *testing.T) {
    root:=t.TempDir()
    current:=filepath.Join(root,"docs","scope","README.md")
    legacy:=filepath.Join(root,"docs","ProjectCharter","1.0","scope","types","TYPE-001.md")
    for _,p:=range []string{current,legacy}{
        if err:=os.MkdirAll(filepath.Dir(p),0700);err!=nil{t.Fatal(err)}
        if err:=os.WriteFile(p,[]byte("# Contrato\n\nTexto da fonte original.\n"),0600);err!=nil{t.Fatal(err)}
    }
    ctx:=context.Background()
    run:=func(args ...string)string{
        out,e:=gitSession(ctx,root,args...)
        if e!=nil{t.Fatal(args,e)}
        return out
    }
    run("init","-q")
    run("config","user.name","Documentation Test")
    run("config","user.email","test@example.invalid")
    run("add","--","docs")
    run("commit","-m","baseline")
    sha:=run("rev-parse","HEAD")
    a:=terminalApp{
        in:bufio.NewReader(strings.NewReader("SIM\n")),
        out:&bytes.Buffer{},
        paths:appPaths{State:t.TempDir()},
        session:terminalSession{Repo:"https://github.com/owner/repo.git",Branch:"develop",Area:"scope",BaseSHA:sha,ReferenceSHA:strings.Repeat("a",40),Workspace:root},
    }
    if err:=a.backupCreate();err!=nil{t.Fatal(err)}
    if !strings.HasPrefix(a.session.BackupID,"20"){t.Fatalf("ID: %s",a.session.BackupID)}
    manifest,err:=a.verifiedBackup();if err!=nil{t.Fatal(err)}
    if len(manifest.Files)!=2{t.Fatalf("esperava 2 arquivos: %d",len(manifest.Files))}
    for _,entry:=range manifest.Files{
        if !strings.HasPrefix(entry.BackupPath,"docs/backup/"){t.Fatal(entry.BackupPath)}
        if !a.backupAllowedPath(entry.BackupPath){t.Fatal("backup bloqueado:",entry.BackupPath)}
    }
    if a.backupAllowedPath("docs/backup/"+a.session.BackupID+"/original/extra.md"){t.Fatal("arquivo não manifesto aceito")}
    changed,err:=a.changedPaths();if err!=nil{t.Fatal(err)}
    if len(changed)!=3{t.Fatalf("esperava 2 originais e manifesto, mudou %d: %v",len(changed),changed)}
    if err=a.validate();err!=nil{t.Fatal(err)}
}
