# Arquitetura proposta — Serviço documental isolado

> Estado: **refinement / proposta técnica**. Esta especificação não afirma que o daemon ou a segurança multiusuário já estejam implementados.
> Referência normativa de leitura: `ahhvahh/doc-arquitetura`, branch `main`, diretório `ProjectCharter/1.0`.
> Projeto aplicativo: `ahhvahh/doc-arquitetura-app-debian`, desenvolvimento somente em `develop`. **Nenhuma escrita em `main`.**

## 1. Mudança de topologia

```text
Usuário autorizado (membro de docarq-clients)
  /usr/bin/doc-arquitetura     CLI sem Git/GGUF/runtime local
          │ comando + argumentos (sem shell)
          │ /run/doc-arquitetura-app/control.sock (AF_UNIX, 0660)
          ▼
systemd doc-arquitetura-app.service
  conta técnica docarq-svc (sem login interativo)
  ├─ SO_PEERCRED e controle de titularidade da tarefa
  ├─ fila de tarefas e eventos JSONL persistentes
  ├─ clona GitHub em workspace privado
  ├─ lê normativa main como referência (somente leitura)
  ├─ modelo GGUF e prompts sob controle do daemon
  ├─ llama-server filho
  │     └─ /run/doc-arquitetura-app/private/<id>/llama.sock (0700)
  └─ commit/push SOMENTE na branch da tarefa, após aprovação
```

O CLI será apenas cliente do protocolo. O serviço será o **único** responsável por clonagem, modelos, inferência, escrita, Git e publicação. Deve ser removido/impedido o caminho legado em que `doc-arquitetura terminal` edita repositórios com o UID do usuário. O serviço **não** poderá receber comandos arbitrários de shell/exec do cliente ou do modelo.

### Permissões Linux propostas

| Objeto | Proprietário | Grupo | Modo | Observação |
|---|---|---|---|---|
| Conta do daemon | `docarq-svc` | `docarq-clients` (primário) | sem login | Identidade exclusiva do serviço |
| Grupo de operadores | administrado por root | `docarq-clients` | n/a | Apenas membros podem conectar ao socket |
| Grupo técnico interno | administrado por root | `docarq-runtime` | n/a | Apenas conta do serviço pertence a esse grupo |
| `/run/doc-arquitetura-app/` | `docarq-svc` | `docarq-clients` | `0750` | Travessia do diretório do socket público |
| `control.sock` | `docarq-svc` | `docarq-clients` | `0660` | Socket público de comandos estruturados |
| `private/` | `docarq-svc` | `docarq-runtime` | `0700` | Não acessível ao grupo de clientes |
| `private/<task>/llama.sock` | `docarq-svc` | `docarq-runtime` | `0600` | Apenas serviço pode interagir com o Llama |
| `/usr/lib/doc-arquitetura-app/llama-server` | `root` | `docarq-runtime` | `0750` | Usuários clientes não podem executar diretamente |
| `/var/lib/doc-arquitetura-app/` | `docarq-svc` | `docarq-clients` | `0700` | Workspace, modelos, logs e dados privados |

O serviço roda com `User=docarq-svc`, `Group=docarq-clients` e `SupplementaryGroups=docarq-runtime`. O diretório `private/` precisa continuar modo 0700, independente do grupo herdado. Embora pertençam a `docarq-clients`, os operadores não recebem acesso a modelos, tokens, workspaces ou ao socket do Llama. Root e processos sob o UID técnico têm as prerrogativas habituais do Linux; o serviço não promete proteção contra root.

Autorização obrigatória por Linux DAC no socket **e** `SO_PEERCRED` para registrar UID/PID do cliente; o serviço não aceita UID informado pelo próprio cliente como prova de identidade. Aprovação de uma tarefa: somente o UID do solicitante, ou administrador explicitamente configurado; outro membro do grupo não pode aprová-la implicitamente.

### Unidade systemd proposta

```ini
[Unit]
Description=Doc Arquitetura — orquestrador documental privado
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=docarq-svc
Group=docarq-clients
SupplementaryGroups=docarq-runtime
ExecStart=/usr/bin/doc-arquitetura service
WorkingDirectory=/var/lib/doc-arquitetura-app
Environment=HOME=/var/lib/doc-arquitetura-app
StateDirectory=doc-arquitetura-app
StateDirectoryMode=0700
CacheDirectory=doc-arquitetura-app
CacheDirectoryMode=0700
RuntimeDirectory=doc-arquitetura-app
RuntimeDirectoryMode=0750
UMask=0077
NoNewPrivileges=yes
ProtectSystem=strict
ProtectHome=yes
PrivateTmp=yes
ProtectKernelTunables=yes
ProtectControlGroups=yes
RestrictAddressFamilies=AF_UNIX AF_INET AF_INET6
KillMode=control-group
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

O daemon cria o socket de controle no RuntimeDirectory e aplica `chmod 0660` depois do bind (o `UMask=0077` impediria isso automaticamente). O `llama.cpp` será **iniciado pelo serviço** quando necessário, sempre pelo socket privado, nunca TCP. O socket privado é criado sob a árvore `0700`, não dentro do diretório acessível ao grupo de clientes. `KillMode=control-group` evita processos órfãos do motor após a parada do serviço.

O `llama-server` atual já é compilado e distribuído privadamente no pacote. A mudança é no **usuário executor**, **permissões do binário**, **ciclo de vida** e **protocolo**; não implica recompilar por solicitação.

## 2. Protocolo cliente → daemon

Mensagens JSON estruturadas e versionadas, limitadas em tamanho (não transportar conteúdo arbitrário ilimitado); `request_id`, `task_id`, `operation` pertencem ao contrato, não são strings shell.

```json
{
  "version": 1,
  "command": "task.submit",
  "request_id": "b9c17d",
  "payload": {
    "repository": "https://github.com/OWNER/REPO.git",
    "source_branch": "develop",
    "area": "charter",
    "prompt": "charter-review.md",
    "model": "qwen3-1.7b",
    "instruction": "Revisar objetivos e apontar inconsistências"
  }
}
```

Comandos de controle: `task.submit`, `task.status`, `task.events`, `task.diff`, `task.approve`, `task.commit`, `task.publish`, `task.cancel`, `models.list`, `models.select`. Aprovações exigem `task_id`, UID verificado, revisão/sha do plano e token de operação antirreplay. Solicitações repetidas com o mesmo `request_id` devem ser idempotentes.

O serviço registra eventos de auditoria com parâmetros sensíveis censurados (nunca registrar token GitHub, resposta com segredo, conteúdo completo de prompts ou URLs com credenciais).

## 3. GitHub e credenciais

- A **única base de tarefas editáveis** é `origin/develop`. Nenhum commit ou push direto em `main` ou `develop`.
- O daemon clona como `docarq-svc` em workspace privado, com `--no-recurse-submodules`; não executa código, hooks ou filtros fornecidos pelo projeto-alvo.
- Credenciais GitHub ficam sob controle do serviço (GitHub App de instalação ou credencial de menor privilégio), protegidas em arquivo/secret do systemd, não transmitidas por stdin do cliente nem incluídas em logs/URLs. Acesso a repositórios privados ou push fica `BLOCKED` até configuração segura dessas credenciais. O daemon não tem acesso automático ao SSH agent/credenciais do usuário.
- Criar branch de trabalho baseada no SHA exato de `origin/develop`, sem modificar o ponteiro de `develop`.
- Commit local e publicação na branch da tarefa exigem aprovação separada; push sem force e bloqueio de troca inesperada de HEAD.
- Quando houver conflito ou mudança remota, suspender publicação e solicitar nova revisão.
- Não criar PR/merge automático em `develop` sem autorização específica; a entrega termina com branch de tarefa publicada e evidências da execução.
- Na referência normativa, continuar usando exclusivamente `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0`, somente leitura. No alvo, documentação ativa em `/docs/charter`, `/docs/scope`, `/docs/implementation`; histórico em **`/docs/backup`**.

## 4. Modelos, fila e isolamento

Devido ao ambiente CPU de 4 núcleos/8 GiB, limitar a **uma inferência/modelo ativo** por vez; enfileirar as demais tarefas. Pode manter o Llama carregado por um período de inatividade configurável, mas nunca permitir que o cliente se conecte diretamente ao socket privado. O modelo não recebe permissões Git/shell, apenas texto e ferramentas estritamente catalogadas; a aplicação valida cada proposta antes de aplicar.

Requisitos de segurança a verificar nos testes: ACL/DAC de sockets e binário, `SO_PEERCRED`, credenciais Git, symlinks, diretórios temporários, reinício do daemon, queda do modelo, cancelamento, ausência de shell injection, isolamento de tarefas e aprovação vinculada ao UID/revisão.

## 5. Estratégia de entrega

1. **Documentar contratos** em Charter, Scope e Implementation, e revisar gates normativos. Estado inicial `refinement`/`BLOCKED` para dependências ainda não aprovadas.
2. Implementar daemon de controle, IPC com credenciais e persistência de tarefas; testes locais Unix/grupos.
3. Mover clonagem, repositório normativo, seleção de modelos, prompts e inferência para o processo do serviço. **Desabilitar os caminhos legados de escrita no CLI**; não ligar o serviço com lógica incompleta.
4. Implementar branches por operação e fluxo separado de aprovação, commit e push.
5. Adicionar streaming de eventos/tempos e benchmark de execução.
6. Atualizar `package.yaml`, scripts Debian e unidade systemd, incluindo contas/grupos, binário privado e teste de instalação com usuário autorizado/não autorizado. Alterar versão do pacote em release próprio.
7. Teste de ponta a ponta em Debian: instalação/atualização, daemon ativo, cliente com e sem grupo, inferência, clone, task branch, diff, commit, push, reinício/cancelamento e logs.

**Não ativar automaticamente um daemon de testes sem implementar todas as fronteiras de segurança acima.**
