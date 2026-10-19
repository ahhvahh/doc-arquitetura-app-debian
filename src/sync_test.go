package main
import ("os";"path/filepath";"strings";"testing")
func TestCompactDoc(t *testing.T) {
 s:=compactDoc("# X\n\n\n![Status](x/badge-status-refined.svg)\n\nTexto")
 if strings.Contains(s,"badge-status") || !strings.Contains(s,"Texto"){t.Fatal(s)}
}
func TestConvertCharter(t *testing.T) {
 source:=t.TempDir();out:=filepath.Join(t.TempDir(),"doc.pcdsl")
 if e:=os.WriteFile(filepath.Join(source,"OBJ-001.md"),[]byte("# OBJ\n\nTexto"),0644);e!=nil{t.Fatal(e)}
 n,e:=convertCharter(source,out,"teste");if e!=nil||n!=1{t.Fatalf("%d %v",n,e)}
 b,e:=os.ReadFile(out);if e!=nil||!strings.Contains(string(b),"PCDSL-DOCS 1") {t.Fatal(e)}
}

func TestFindCharterDirectDocs(t *testing.T) {
 root:=t.TempDir()
 direct:=filepath.Join(root,"docs","scope")
 legacy:=filepath.Join(root,"docs","ProjectCharter","1.0","scope")
 if err:=os.MkdirAll(direct,0700);err!=nil{t.Fatal(err)}
 if err:=os.MkdirAll(legacy,0700);err!=nil{t.Fatal(err)}
 found,err:=findCharter(root,false)
 if err!=nil || found!=filepath.Join(root,"docs"){t.Fatalf("should prefer direct docs root: %q, %v",found,err)}
 // Normative source remains versioned; do not confuse it with target /docs.
 norm:=filepath.Join(root,"ProjectCharter","1.0")
 if err=os.MkdirAll(norm,0700);err!=nil{t.Fatal(err)}
 found,err=findCharter(root,true)
 if err!=nil || found!=norm{t.Fatalf("normative path changed: %q, %v",found,err)}
}
func TestReferenceRejectsDevelop(t *testing.T) {
 if err:=cloneReference("develop",filepath.Join(t.TempDir(),"reference"));err==nil {
  t.Fatal("normative develop must never be used")
 }
}

func TestConvertExcludesBackupDirectory(t *testing.T) {
    root:=t.TempDir()
    docs:=filepath.Join(root,"docs")
    active:=filepath.Join(docs,"charter","README.md")
    legacy:=filepath.Join(docs,"backup","snapshot","original","docs","charter","OBJ-001.md")
    for _,p:=range []string{active,legacy}{
        if err:=os.MkdirAll(filepath.Dir(p),0700);err!=nil{t.Fatal(err)}
        if err:=os.WriteFile(p,[]byte("# Teste\\n"),0600);err!=nil{t.Fatal(err)}
    }
    output:=filepath.Join(t.TempDir(),"bundle.pcdsl")
    n,err:=convertCharter(docs,output,"target")
    if err!=nil || n!=1 {t.Fatalf("deve excluir backup: %d %v",n,err)}
    b,err:=os.ReadFile(output)
    if err!=nil || strings.Contains(string(b),"OBJ-001.md"){t.Fatalf("backup incluido no bundle: %v",err)}
}
