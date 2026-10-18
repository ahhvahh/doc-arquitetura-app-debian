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
