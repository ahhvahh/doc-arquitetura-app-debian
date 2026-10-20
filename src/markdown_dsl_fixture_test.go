package main

import (
 "bytes"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "fmt"
 "os"
 "path/filepath"
 "sort"
 "strings"
 "testing"
)

type stageMarkdownFixture struct {
 File string `json:"file"`
 Type string `json:"type"`
 Status string `json:"status"`
 Title string `json:"title"`
 Detail string `json:"detail"`
}

func readStageMarkdownManifest(t *testing.T) map[string][]stageMarkdownFixture {
 t.Helper()
 data, err := os.ReadFile(filepath.Join("test", "markdown", "manifest.json"))
 if err != nil { t.Fatal(err) }
 var manifest map[string][]stageMarkdownFixture
 if err := json.Unmarshal(data, &manifest); err != nil { t.Fatal(err) }
 if len(manifest) != 4 { t.Fatalf("want 4 stages; got %d", len(manifest)) }
 for _, stage := range []string{"charter", "scope", "implementation", "src"} {
  if len(manifest[stage]) != 5 { t.Fatalf("%s: want 5 fixtures, got %d", stage, len(manifest[stage])) }
 }
 return manifest
}

func expectedFixtureDSL(t *testing.T, stage, origin string, manifest map[string][]stageMarkdownFixture) string {
 t.Helper()
 fixtures := append([]stageMarkdownFixture(nil), manifest[stage]...)
 sort.Slice(fixtures, func(i,j int) bool {return fixtures[i].File < fixtures[j].File})
 var expected strings.Builder
 fmt.Fprintf(&expected, "PCDSL-DOCS 1\nORIGIN %s\nDOCUMENTS 5\n", origin)
 directory := filepath.Join("test", "markdown", "docs", stage)
 entries,err := os.ReadDir(directory)
 if err != nil {t.Fatal(err)}
 if len(entries) != 5 { t.Fatalf("%s: directory must have exactly 5 Markdown documents; found %d", stage,len(entries)) }
 seen := map[string]bool{}
 for _,fixture := range fixtures {
  if filepath.Base(fixture.File)!=fixture.File || !strings.HasSuffix(fixture.File,".md") || seen[fixture.File] {t.Fatalf("invalid fixture path: %s",fixture.File)}
  seen[fixture.File]=true
  path:=filepath.Join(directory,fixture.File)
  original,err:=os.ReadFile(path)
  if err!=nil {t.Fatal(err)}
  badges:=fmt.Sprintf("![Tipo: %s](badges/badge-type-%s.svg) ![Status: %s](badges/badge-status-%s.svg) ![Versão: v1](badges/badge-version-v1.svg)",fixture.Type,strings.ToLower(fixture.Type),fixture.Status,fixture.Status)
  fixtureRaw:=fmt.Sprintf("# %s\n\n%s\n\n## Cenário\n%s\n",fixture.Title,badges,fixture.Detail)
  if string(original)!=fixtureRaw {t.Fatalf("%s differs from independent fixture manifest",path)}
  payload:=fmt.Sprintf("# %s\n\n## Cenário\n%s",fixture.Title,fixture.Detail)
  hash:=sha256.Sum256(original)
  fmt.Fprintf(&expected,"DOC %q %s %s %d\n%s\n",fixture.File,hex.EncodeToString(hash[:]),fixture.Status,len([]byte(payload)),payload)
 }
 return expected.String()
}

func TestStageMarkdownDSLGolden(t *testing.T) {
 manifest:=readStageMarkdownManifest(t)
 for _,stage:=range []string{"charter","scope","implementation","src"} {
  t.Run(stage,func(t *testing.T){
   origin:="fixture:"+stage
   output:=filepath.Join(t.TempDir(),stage+".pcdsl")
   count,err:=convertCharter(filepath.Join("test","markdown","docs",stage),output,origin)
   if err!=nil {t.Fatal(err)}
   if count!=5 {t.Fatalf("want 5 docs, got %d",count)}
   actual,err:=os.ReadFile(output)
   if err!=nil {t.Fatal(err)}
   want:=expectedFixtureDSL(t,stage,origin,manifest)
   if string(actual)!=want {t.Fatalf("%s generated wrong DSL\nACTUAL:\n%s\nEXPECTED:\n%s",stage,string(actual),want)}
   if strings.Contains(string(actual),"badge-status-") || strings.Contains(string(actual),"![Tipo:") {
    t.Fatalf("%s: SVG presentation leaked to DSL",stage)
   }
   second:=filepath.Join(t.TempDir(),"second.pcdsl")
   if _,err:=convertCharter(filepath.Join("test","markdown","docs",stage),second,origin);err!=nil {t.Fatal(err)}
   again,err:=os.ReadFile(second);if err!=nil{t.Fatal(err)}
   if !bytes.Equal(actual,again) {t.Fatalf("%s: nondeterministic DSL",stage)}
  })
 }
}

func TestWholeProjectMarkdownDSLCoversFourStages(t *testing.T) {
 manifest:=readStageMarkdownManifest(t)
 path:=filepath.Join(t.TempDir(),"project.pcdsl")
 docs:=filepath.Join("test","markdown","docs")
 n,err:=convertCharter(docs,path,"fixture:project")
 if err!=nil {t.Fatal(err)}
 if n!=20 {t.Fatalf("want 20 documents in full project, got %d",n)}
 raw,err:=os.ReadFile(path);if err!=nil{t.Fatal(err)}
 if !strings.Contains(string(raw),"DOCUMENTS 20\n"){t.Fatalf("wrong header: %s",raw)}
 for area,items:=range manifest {
  for _,item:=range items {
   id:=filepath.ToSlash(filepath.Join(area,item.File))
   if !strings.Contains(string(raw),fmt.Sprintf("DOC %q ",id)){t.Fatalf("missing file %s in full DSL",id)}
  }
 }
}

func TestStageMarkdownDSLDetectsModifiedContent(t *testing.T) {
 manifest:=readStageMarkdownManifest(t)
 stage:="charter"
 work:=t.TempDir()
 for _,f:=range manifest[stage] {
  raw,err:=os.ReadFile(filepath.Join("test","markdown","docs",stage,f.File))
  if err!=nil{t.Fatal(err)}
  if f.File=="PCH-001-contexto.md" {
   raw=[]byte(strings.Replace(string(raw),"badge-status-refined.svg","badge-status-cancelled.svg",1))
  }
  if err:=os.WriteFile(filepath.Join(work,f.File),raw,0600);err!=nil{t.Fatal(err)}
 }
 output:=filepath.Join(t.TempDir(),"mutated.pcdsl")
 n,err:=convertCharter(work,output,"fixture:mutated")
 if err!=nil||n!=5 {t.Fatalf("modified data failed: %d %v",n,err)}
 raw,err:=os.ReadFile(output);if err!=nil{t.Fatal(err)}
 if !strings.Contains(string(raw)," cancelled ") {t.Fatal("modified status not recognized")}
 if strings.Contains(string(raw),"badge-status-cancelled") {t.Fatal("SVG marker leaked to payload")}
}

func TestFixtureMarkdownDSLExcludesBackup(t *testing.T) {
 docs:=t.TempDir()
 for _,f:=range readStageMarkdownManifest(t)["scope"] {
  raw,err:=os.ReadFile(filepath.Join("test","markdown","docs","scope",f.File));if err!=nil{t.Fatal(err)}
  if err:=os.WriteFile(filepath.Join(docs,f.File),raw,0600);err!=nil{t.Fatal(err)}
 }
 legacy:=filepath.Join(docs,"backup","snapshot","TYPE-999-legacy.md")
 if err:=os.MkdirAll(filepath.Dir(legacy),0700);err!=nil{t.Fatal(err)}
 if err:=os.WriteFile(legacy,[]byte("# Historical document\n"),0600);err!=nil{t.Fatal(err)}
 dst:=filepath.Join(t.TempDir(),"without-backup.pcdsl")
 n,err:=convertCharter(docs,dst,"fixture:no-backup")
 if err!=nil||n!=5 {t.Fatalf("backup incorrectly included: %d %v",n,err)}
 raw,err:=os.ReadFile(dst);if err!=nil{t.Fatal(err)}
 if strings.Contains(string(raw),"TYPE-999") {t.Fatal("backup leaked into DSL")}
}
