package main

import (
    "context"
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "regexp"
    "strings"
    "syscall"
    "time"
)

// Per-user GitHub configuration is intentionally separate from source repositories.
// Never commit github.json, github.token or any token to Git.
type githubFileConfig struct {
    GitHub struct {
        Authentication string `json:"authentication"`
        Username string `json:"username"`
        TokenFile string `json:"token_file"`
        Token string `json:"token,omitempty"`
    } `json:"github"`
}
var githubUsernameRE=regexp.MustCompile(`^[A-Za-z0-9](?:[A-Za-z0-9-]{0,37})$`)

func githubConfigPath() (string,error) {
    p,err:=appHome();if err!=nil{return "",err}
    return filepath.Join(p.Config,"github.json"),nil
}
func githubPrivateRead(path string) ([]byte,error) {
    info,err:=os.Lstat(path);if err!=nil{return nil,err}
    if !info.Mode().IsRegular(){return nil,fmt.Errorf("%s: arquivo deve ser regular, sem symlink",path)}
    if info.Mode().Perm()&0077!=0 {return nil,fmt.Errorf("%s: permissões inseguras; execute chmod 600",path)}
    if stat,ok:=info.Sys().(*syscall.Stat_t);ok&&int(stat.Uid)!=os.Geteuid(){
        return nil,fmt.Errorf("%s: proprietário não corresponde ao usuário em execução",path)
    }
    return os.ReadFile(path)
}
func githubLoad() (*githubFileConfig,error) {
    file,err:=githubConfigPath();if err!=nil{return nil,err}
    content,err:=githubPrivateRead(file)
    if os.IsNotExist(err){return nil,nil}
    if err!=nil{return nil,err}
    var cfg githubFileConfig
    if err=json.Unmarshal(content,&cfg);err!=nil{return nil,fmt.Errorf("github.json inválido: %w",err)}
    if cfg.GitHub.Authentication!="token"{
        return nil,errors.New("github.json: authentication deve ser token")
    }
    if !githubUsernameRE.MatchString(cfg.GitHub.Username){
        return nil,errors.New("github.json: username GitHub inválido")
    }
    if cfg.GitHub.TokenFile!=""&&cfg.GitHub.Token!=""{
        return nil,errors.New("github.json: escolha token_file ou token, não ambos")
    }
    if cfg.GitHub.TokenFile==""&&cfg.GitHub.Token==""{
        return nil,errors.New("github.json: informe token_file ou token")
    }
    if cfg.GitHub.TokenFile!=""&&!filepath.IsAbs(cfg.GitHub.TokenFile){
        return nil,errors.New("github.json: token_file deve ser caminho absoluto")
    }
    return &cfg,nil
}
func (c *githubFileConfig) secret() (string,error) {
    if c==nil {return "",errors.New("GitHub não configurado")}
    secret:=c.GitHub.Token
    if c.GitHub.TokenFile!=""{
        b,err:=githubPrivateRead(c.GitHub.TokenFile);if err!=nil{return "",fmt.Errorf("token_file: %w",err)}
        secret=strings.TrimSpace(string(b))
    }
    if len(secret)<8||len(secret)>4096||strings.ContainsAny(secret," \t\r\n"){
        return "",errors.New("token GitHub ausente ou com caracteres inválidos")
    }
    return secret,nil
}
func githubGitEnv() ([]string,error) {
    env:=append([]string(nil),os.Environ()...)
    env=append(env,"GIT_TERMINAL_PROMPT=0")
    cfg,err:=githubLoad();if err!=nil{return nil,err}
    if cfg==nil{return env,nil} // public repositories work without authentication
    if _,err=cfg.secret();err!=nil{return nil,err}
    executable,err:=os.Executable();if err!=nil{return nil,err}
    // Git invokes the same Go binary as an askpass helper. Credentials are never
    // passed through command arguments, environment values or Git remote URLs.
    // Reset other credential helpers to prevent credential mix-ups.
    env=append(env,
       "GIT_ASKPASS="+executable,
       "DOC_ARQUITETURA_GITHUB_ASKPASS=1",
       "GIT_CONFIG_COUNT=1",
       "GIT_CONFIG_KEY_0=credential.helper",
       "GIT_CONFIG_VALUE_0=",
    )
    return env,nil
}
func githubAskPass() error {
    if len(os.Args)!=2{return errors.New("helper GitHub: prompt ausente")}
    prompt:=os.Args[1]
    // Do not return credentials for another host or protocol, even when a cloned
    // repository supplies a malicious remote.
    isGithub:=regexp.MustCompile(`^((Username|Password) for 'https://([^'@/]+@)?github\.com(:443)?(/[^']*)?': ?)$`).MatchString(prompt)
    if !isGithub{return errors.New("helper GitHub: domínio não autorizado")}
    cfg,err:=githubLoad();if err!=nil{return err}
    if cfg==nil{return errors.New("GitHub não configurado")}
    if strings.HasPrefix(prompt,"Username"){
        fmt.Println(cfg.GitHub.Username)
        return nil
    }
    secret,err:=cfg.secret();if err!=nil{return err}
    fmt.Println(secret)
    return nil
}
func githubCommand(args []string) error {
    if len(args)==0{return errors.New("uso: github <config|check> [URL HTTPS GitHub]")}
    cfgPath,err:=githubConfigPath();if err!=nil{return err}
    switch args[0]{
    case "config":
        if len(args)!=1{return errors.New("uso: github config")}
        fmt.Println("Arquivo GitHub:",cfgPath)
        cfg,err:=githubLoad();if err!=nil{return err}
        if cfg==nil{
            fmt.Println("Não configurado. Crie github.json e proteja com chmod 600.")
        }else{
            fmt.Printf("Autenticação: %s | conta: %s | segredo: [protegido]\n",cfg.GitHub.Authentication,cfg.GitHub.Username)
        }
        return nil
    case "check":
        if len(args)!=2{return errors.New("uso: github check https://github.com/OWNER/REPO")}
        remote,err:=normalizeGitHubURL(args[1]);if err!=nil{return err}
        ctx,cancel:=context.WithTimeout(context.Background(),30*time.Second);defer cancel()
        // Read-only check; no clone, commit or push is performed.
        out,err:=gitSession(ctx,"","ls-remote","--heads",remote,"develop")
        if err!=nil{return fmt.Errorf("GitHub indisponível: %w",err)}
        if strings.TrimSpace(out)==""{return errors.New("conectado, mas branch develop não encontrada")}
        fmt.Printf("GitHub acessível; develop encontrada em %s\n",remote)
        return nil
    default:return fmt.Errorf("comando github desconhecido: %s",args[0])
    }
}
func githubGitCommand(ctx context.Context,args ...string)(*exec.Cmd,error){
    env,err:=githubGitEnv();if err!=nil{return nil,err}
    cmd:=exec.CommandContext(ctx,"git",args...)
    cmd.Env=env
    return cmd,nil
}
