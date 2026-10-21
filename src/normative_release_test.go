package main

import (
 "archive/zip"
 "bytes"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "fmt"
 "strings"
 "testing"
)
func testZipBundle(t *testing.T,corrupt bool)([]byte,[]byte,string){
 t.Helper()
 tag:=releasePrefix+strings.Repeat("a",40)
 b:=&bytes.Buffer{}
 zipw:=zip.NewWriter(b)
 manifest:=normativeManifest{Commit:strings.Repeat("a",40),Schema:"projectcharter-package/1"}
 for _,name:=range bundledNames{
  if name=="manifest.json"{continue}
  data:=[]byte("PROJECTCHARTER-DSL/1\nFILE\t"+strings.TrimSuffix(name,".dsl")+"/README.md\nTEXT\t\"# heading\\n\"\nEND\n")
  sum:=sha256.Sum256(data)
  manifest.Packages=append(manifest.Packages,struct{Name string `json:"name"`;SHA256 string `json:"sha256"`;Entries int `json:"entries"`}{name,hex.EncodeToString(sum[:]),1})
  if corrupt && name=="scope.dsl"{data=append(data,'x')}
  file,err:=zipw.Create(name);if err!=nil{t.Fatal(err)};if _,err=file.Write(data);err!=nil{t.Fatal(err)}
 }
 meta,err:=json.Marshal(manifest);if err!=nil{t.Fatal(err)}
 f,err:=zipw.Create("manifest.json");if err!=nil{t.Fatal(err)};if _,err=f.Write(meta);err!=nil{t.Fatal(err)}
 if err=zipw.Close();err!=nil{t.Fatal(err)}
 sum:=sha256.Sum256(b.Bytes())
 checksum:=[]byte(fmt.Sprintf("%x  projectcharter-1.0.zip
",sum))
 return b.Bytes(),checksum,tag
}
func TestReleaseSelection(t *testing.T){
 valid:=releasePrefix+strings.Repeat("a",40)
 releases:=[]normativeRelease{{Tag:releasePrefix+strings.Repeat("b",40),Draft:true},{Tag:valid,Assets:[]releaseAsset{{Name:packageFilename},{Name:packageFilename+".sha256"}}}}
 r,err:=selectRelease(releases);if err!=nil||r.Tag!=valid{t.Fatalf("unexpected release: %v %v",r,err)}
}
func TestVerifyNormativeRelease(t *testing.T){
 data,sum,tag:=testZipBundle(t,false)
 result,manifest,err:=verifyDSLArchive(data,sum,tag)
 if err!=nil||len(result)!=6||len(manifest.Packages)!=5{t.Fatalf("bundle not validated: %v",err)}
 sum[0]='0'
 if _,_,err=verifyDSLArchive(data,sum,tag);err==nil{t.Fatal("tampered checksum accepted")}
}
func TestRejectModifiedDSL(t *testing.T){
 data,sum,tag:=testZipBundle(t,true)
 if _,_,err:=verifyDSLArchive(data,sum,tag);err==nil{t.Fatal("tampered DSL accepted")}
}
