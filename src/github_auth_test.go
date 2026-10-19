package main

import (
 "os"
 "path/filepath"
 "strings"
 "testing"
)

func writeGithubFixture(t *testing.T,token string) string {
 t.Helper()
 p:=isolatedPaths(t)
 t.Setenv("DOC_ARQUITETURA_CONFIG_DIR",p.Config)
 tokenPath:=filepath.Join(p.Config,"github.token")
 if err:=os.WriteFile(tokenPath,[]byte(token+"\n"),0600);err!=nil{t.Fatal(err)}
 cfg:=filepath.Join(p.Config,"github.json")
 data:=[]byte(`{"github":{"authentication":"token","username":"example-user","token_file":"`+tokenPath+`"}}`)
 if err:=os.WriteFile(cfg,data,0600);err!=nil{t.Fatal(err)}
 return cfg
}
func TestGitHubConfiguration(t *testing.T){
 writeGithubFixture(t,"github_pat_testingexample")
 cfg,err:=githubLoad()
 if err!=nil{t.Fatal(err)}
 token,err:=cfg.secret()
 if err!=nil||token!="github_pat_testingexample"{t.Fatalf("token mismatch: %v",err)}
 env,err:=githubGitEnv()
 if err!=nil{t.Fatal(err)}
 all:=strings.Join(env,"\n")
 if !strings.Contains(all,"GIT_ASKPASS="){t.Fatal("askpass missing")}
 if strings.Contains(all,"github_pat_testingexample"){t.Fatal("token leaked to environment")}
}
func TestGitHubRejectsLoosePermissions(t *testing.T){
 cfg:=writeGithubFixture(t,"github_pat_testingexample")
 if err:=os.Chmod(cfg,0644);err!=nil{t.Fatal(err)}
 if _,err:=githubLoad();err==nil{t.Fatal("insecure file accepted")}
}
func TestGitHubRejectsTokenSymlink(t *testing.T){
 cfg:=writeGithubFixture(t,"github_pat_testingexample")
 p:=filepath.Dir(cfg)
 link:=filepath.Join(p,"linkedtoken")
 if err:=os.Symlink(filepath.Join(p,"github.token"),link);err!=nil{t.Fatal(err)}
 data:=[]byte(`{"github":{"authentication":"token","username":"example-user","token_file":"`+link+`"}}`)
 if err:=os.WriteFile(cfg,data,0600);err!=nil{t.Fatal(err)}
 x,err:=githubLoad();if err!=nil{t.Fatal(err)}
 if _,err=x.secret();err==nil{t.Fatal("symlink accepted")}
}
func TestGitHubConfigAbsent(t *testing.T){
 p:=isolatedPaths(t);t.Setenv("DOC_ARQUITETURA_CONFIG_DIR",p.Config)
 c,err:=githubLoad();if err!=nil||c!=nil{t.Fatalf("expected optional config: %+v %v",c,err)}
}
