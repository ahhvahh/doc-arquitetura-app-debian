package main

import (
 "context"
 "crypto/sha256"
 "encoding/hex"
 "encoding/json"
 "errors"
 "fmt"
 "io"
 "net/http"
 "net/http/httptest"
 "os"
 "path/filepath"
 "strings"
 "testing"
)

func TestCommandTreeHasAllNormativePrompts(t *testing.T){
 tree,err:=loadCommandTree();if err!=nil{t.Fatal(err)}
 count:=0;pending:=0
 for _,stage:=range []string{"charter","scope","implementation","src"}{
  for _,command:=range tree.Stages[stage]{
   count++
   if !command.Executable{
    pending++
    if _,err:=pickPrompt(tree,stage,command.Operation,command.Action);err==nil{t.Fatalf("PENDING prompt accepted: %+v",command)}
   } else {
    got,err:=pickPrompt(tree,stage,command.Operation,command.Action)
    if err!=nil||got.PromptPath!=command.PromptPath{t.Fatalf("missing route %s/%s/%s: %v",stage,command.Operation,command.Action,err)}
   }
  }
 }
 if count!=34||pending<1{t.Fatalf("expected 34 prompts including at least one pending; got %d pending=%d",count,pending)}
 if _,err:=pickPrompt(tree,"charter","mapear","../scope");err==nil{t.Fatal("path traversal allowed")}
}
func TestSelectedPromptStrictlyScoped(t *testing.T){
 root:=t.TempDir()
 area:=filepath.Join(root,"prompts","scope")
 if err:=os.MkdirAll(area,0700);err!=nil{t.Fatal(err)}
 path:=filepath.Join(area,"type-refinement.txt")
 content:="Prompt selected, others not loaded"
 if err:=os.WriteFile(path,[]byte(content),0600);err!=nil{t.Fatal(err)}
 got,hash,err:=readSelectedPrompt(root,"scope",promptCommand{PromptPath:"prompts/scope/type-refinement.txt"})
 if err!=nil||got!=content{t.Fatalf("prompt: %q error=%v",got,err)}
 digest:=sha256.Sum256([]byte(content))
 if hash!=hex.EncodeToString(digest[:]){t.Fatal("SHA-256 mismatch")}
 _,_,err=readSelectedPrompt(root,"scope",promptCommand{PromptPath:"prompts/charter/objective-refinement.txt"})
 if err==nil{t.Fatal("cross-area prompt accepted")}
}
func TestAgentNormativeInstructionsExcludeSpecializedPrompts(t *testing.T){
 root:=t.TempDir()
 for name,body:=range map[string]string{
  "AGENT-PROMPT.md":"Norma global",
  "charter/AGENT-PROMPT.md":"Norma charter",
  "charter/README.md":"Índice charter",
  "prompts/charter/objective-refinement.txt":"PROMPT SEPARADO NÃO INSERIR",
 }{
  path:=filepath.Join(root,name)
  if err:=os.MkdirAll(filepath.Dir(path),0700);err!=nil{t.Fatal(err)}
  if err:=os.WriteFile(path,[]byte(body),0600);err!=nil{t.Fatal(err)}
 }
 input,hash,err:=openAIInstructions(root,"charter")
 if err!=nil{t.Fatal(err)}
 if !strings.Contains(input,"Norma global")||!strings.Contains(input,"Norma charter")||strings.Contains(input,"PROMPT SEPARADO NÃO INSERIR"){t.Fatalf("leakage or incomplete instructions: %q",input)}
 if len(hash)!=64{t.Fatal("invalid SHA")}
}
func TestRemoteAgentLifecycleMock(t *testing.T){
 operations:=[]string{}
 var remote openAIAgentData
 srv:=httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
  if r.Header.Get("Authorization")!="Bearer test-key"||r.Header.Get("OpenAI-Beta")!="agents=v1"{t.Errorf("invalid API headers")}
  operations=append(operations,r.Method+" "+r.URL.Path)
  w.Header().Set("Content-Type","application/json")
  switch{
  case r.Method=="POST"&&r.URL.Path=="/v1/agents":
   if err:=json.NewDecoder(r.Body).Decode(&remote);err!=nil{t.Error(err)}
   remote.ID="agent_mock_1"
   _=json.NewEncoder(w).Encode(remote)
  case r.Method=="GET"&&r.URL.Path=="/v1/agents/agent_mock_1":
   _=json.NewEncoder(w).Encode(remote)
  case r.Method=="POST"&&r.URL.Path=="/v1/agents/agent_mock_1":
   if err:=json.NewDecoder(r.Body).Decode(&remote);err!=nil{t.Error(err)}
   remote.ID="agent_mock_1"
   _=json.NewEncoder(w).Encode(remote)
  case r.Method=="DELETE"&&r.URL.Path=="/v1/agents/agent_mock_1":
   _=json.NewEncoder(w).Encode(map[string]any{"id":"agent_mock_1","deleted":true})
  default:
   w.WriteHeader(http.StatusNotFound)
  }
 }))
 defer srv.Close()
 c:=agentHTTPClient{HTTP:srv.Client(),URL:srv.URL+"/v1",Key:"test-key"}
 record,err:=remoteSyncAgent(context.Background(),c,remoteAgentRecord{},"charter","model-test","a-sha","first instructions","first-hash")
 if err!=nil||record.ID!="agent_mock_1"{t.Fatalf("create failed: %+v %v",record,err)}
 repeated,err:=remoteSyncAgent(context.Background(),c,record,"charter","model-test","a-sha","first instructions","first-hash")
 if err!=nil||repeated!=record{t.Fatalf("idempotency error: %+v %v",repeated,err)}
 changed,err:=remoteSyncAgent(context.Background(),c,record,"charter","model-test","b-sha","new instructions","second-hash")
 if err!=nil||changed.NormativeSHA!="b-sha"{t.Fatalf("update failed: %+v %v",changed,err)}
 if err=remoteDeleteAgent(context.Background(),c,record.ID);err!=nil{t.Fatal(err)}
 want:=[]string{"POST /v1/agents","GET /v1/agents/agent_mock_1","GET /v1/agents/agent_mock_1","POST /v1/agents/agent_mock_1","DELETE /v1/agents/agent_mock_1"}
 if fmt.Sprint(operations)!=fmt.Sprint(want){t.Fatalf("unexpected requests: %v",operations)}
}
func TestRemoteAgentMissingKeyBlocksNetwork(t *testing.T){
 c:=agentHTTPClient{HTTP:http.DefaultClient,URL:"http://127.0.0.1:1",Key:""}
 if err:=c.call(context.Background(),"GET","/agents",nil,nil);err==nil{t.Fatal("missing key was accepted")}
}
func TestChangeSetChecksMarkdownPreservation(t *testing.T){
 workspace:=t.TempDir()
 root:=filepath.Join(workspace,"docs","charter","objectives")
 if err:=os.MkdirAll(root,0700);err!=nil{t.Fatal(err)}
 path:=filepath.Join(root,"OBJ-001.md")
 original:="# OBJ-001\n\n![Tipo: OBJ](badge-type-obj.svg)\n\n## Motivo\nOriginal\n\n## Critérios\nAntes\n\n## Rastreio\nPreservado\n"
 if err:=os.WriteFile(path,[]byte(original),0600);err!=nil{t.Fatal(err)}
 task:=preparedTask{TaskID:"safe-123",Stage:"charter",Prompt:"prompts/charter/objective-refinement.txt",NormativeSHA:"normative",TargetDevelopSHA:"head"}
 cs:=markdownChangeSet{
  Format:"PCDSL-CHANGESET/1",TaskID:task.TaskID,Stage:task.Stage,Prompt:task.Prompt,NormativeSHA:task.NormativeSHA,TargetDevelopSHA:task.TargetDevelopSHA,
  Changes:[]markdownChange{{Operation:"replace_section",Path:"docs/charter/objectives/OBJ-001.md",BaseSHA256:mdDigest([]byte(original)),Heading:"## Critérios",HeadingOccurrence:1,NewMarkdown:"## Critérios\nDepois\n",Reason:"teste"}},
 }
 edits,err:=planChangeset(workspace,task,cs)
 if err!=nil{t.Fatal(err)}
 if len(edits)!=1||!strings.Contains(string(edits[0].Data),"Depois")||!strings.Contains(string(edits[0].Data),"Preservado")||!strings.Contains(string(edits[0].Data),"Original") {t.Fatalf("incorrect reconstitution: %+v",edits)}
 // A stale SHA must fail, without modifying the on-disk document.
 cs.Changes[0].BaseSHA256=strings.Repeat("0",64)
 if _,err:=planChangeset(workspace,task,cs);err==nil{t.Fatal("stale source allowed")}
 raw,err:=os.ReadFile(path)
 if err!=nil||string(raw)!=original{t.Fatal("preview modified original")}
}
func TestChangesetRejectsCrossStageAndTraversal(t *testing.T){
 workspace:=t.TempDir()
 for _,testPath:=range []string{"docs/scope/README.md","docs/charter/../scope/README.md","docs/charter/../../.git/config","docs/charter/invalid.txt"}{
  if _,err:=safeTargetMarkdownPath(workspace,"charter",testPath);err==nil{t.Fatalf("unsafe path allowed: %s",testPath)}
 }
 if _,err:=safeTargetMarkdownPath(workspace,"src","docs/src/test.md");err==nil{t.Fatal("source markdown applied without contract")}
 for _,jsonInput:=range []string{`{"format":"PCDSL-CHANGESET/1","task_id":"x","stage":"charter","changes":[{"operation":"exec","path":"docs/charter/x.md","reason":"x"}]}`, `{"format":"PCDSL-CHANGESET/1","task_id":"x","stage":"charter","changes":[]}`, `{"format":"PCDSL-CHANGESET/1","task_id":"x","stage":"charter","changes":[{"operation":"create_file","path":"docs/charter/x.md","reason":"x","new_markdown":"# A"}],"unexpected":"bad"}`,
 }{
  if _,err:=decodeChangeset(strings.NewReader(jsonInput));err==nil{t.Fatalf("invalid changeset accepted: %s",jsonInput)}
 }
}
func TestCreateFileRequirements(t *testing.T){
 work:=t.TempDir()
 task:=preparedTask{TaskID:"x",Stage:"scope",Prompt:"prompts/scope/type-refinement.txt",NormativeSHA:"n",TargetDevelopSHA:"s"}
 c:=markdownChangeSet{Format:"PCDSL-CHANGESET/1",TaskID:"x",Stage:"scope",Prompt:task.Prompt,NormativeSHA:"n",TargetDevelopSHA:"s",Changes:[]markdownChange{{
  Operation:"create_file",Path:"docs/scope/types/TYPE-001-example.md",Reason:"teste",
  NewMarkdown:"# TYPE-001\n\n![Tipo](badge-type-type.svg) ![Status](badge-status-refinement.svg) ![Versão](badge-version-v1.svg)\n",
 }}}
 writes,err:=planChangeset(work,task,c)
 if err!=nil||len(writes)!=1||!writes[0].IsNew{t.Fatalf("create file failed: %+v %v",writes,err)}
 c.Changes[0].NewMarkdown="# missing badges"
 if _,err:=planChangeset(work,task,c);err==nil{t.Fatal("accepted missing metadata")}
}
func TestDecodeChangeSetReader(t *testing.T){
 payload:=`{"format":"PCDSL-CHANGESET/1","task_id":"x","stage":"charter","changes":[{"operation":"create_file","path":"docs/charter/test.md","new_markdown":"# T","reason":"R"}]}`
 obj,err:=decodeChangeset(strings.NewReader(payload))
 if err!=nil||obj.TaskID!="x"{t.Fatalf("failed decoding: %+v %v",obj,err)}
 if _,err:=decodeChangeset(io.MultiReader(strings.NewReader(payload),strings.NewReader("{}")));err==nil{t.Fatal("accepted two JSON roots")}
 if _,err:=decodeChangeset(strings.NewReader(strings.Replace(payload,"create_file","delete_file",1)));err==nil{t.Fatal("accepted destructive delete")}
 if _,err:=decodeChangeset(strings.NewReader(strings.Replace(payload,"create_file","create_file",1)));err!=nil&& !errors.Is(err,io.EOF) {t.Fatal(err)}
}
