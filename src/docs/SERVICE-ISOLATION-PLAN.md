# Arquitetura proposta — Serviço documental isolado

> **Atualização 2026-10-20:** proposta histórica, não implementada e superada para inferência local. llama.cpp, GGUF e sockets de modelo foram removidos do cliente e do pacote Debian. Não construir esse runtime com base neste plano. A inferência futura será especificada para OpenAI, incluindo controles de gastos.

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
  │     └─ /run/doc-arquitetura-app-private/<id>/llama.sock (0700)
  └─ commit/push SOMENTE na branch da tarefa, após aprovação
```

O CLI será apenas cliente do protocolo. O serviço será o **único** responsável por clonagem, modelos, inferência, escrita, Git e publicação. Deve ser removido/impedido o caminho legado em que `doc-arquitetura terminal` edita repositórios com o UID do usuário. O serviço **não** poderá receber comandos arbitrários de shell/exec do cliente ou do modelo.

### Configuração administrativa root e permissões

A configuração administrativa é definida em [ADMIN-SETUP.md](ADMIN-SETUP.md). Todo cadastro de contas/grupos Linux, credenciais GitHub, usuários autorizados e parâmetros do daemon será realizado **somente por root** via terminal (`sudo doc-arquitetura admin`).

O esquema anterior, em que `docarq-svc` era membro primário de `docarq-clients`, foi substituído para impedir que a conta técnica integre o grupo dos usuários finais. O serviço usará exclusivamente `User=docarq-svc` e `Group=docarq-runtime`; o grupo `docarq-clients` será usado apenas no socket público, de propriedade do systemd (socket activation).

| Objeto | Proprietário/grupo | Modo | Finalidade |
|---|---|---|---|
| `/etc/doc-arquitetura-app/config.yaml` | `root:docarq-runtime` | `0640` | Configuração administrativa, sem segredo |
| `/etc/doc-arquitetura-app/secrets/` | `root:root` | `0700` | Credenciais sob custódia root |
| `/run/doc-arquitetura-app/control.sock` | `root:docarq-clients` | `0660` | Comandos autorizados, validação por `SO_PEERCRED` e allowlist |
| `/run/doc-arquitetura-app-private/` | `docarq-svc:docarq-runtime` | `0700` | Socket do Llama inacessível a clientes |
| `/usr/lib/doc-arquitetura-app/llama-server` | `root:docarq-runtime` | `0750` | Executável privado do serviço |
| `/var/lib/doc-arquitetura-app/` | `docarq-svc:docarq-runtime` | `0700` | Clones, modelos, tarefas e registros privados |

Os grupos/clientes podem usar IPC apenas se o usuário estiver simultaneamente no grupo e em `authorized_users`; aprovar operações de outra pessoa não decorre da participação no grupo. Root gerencia mudanças por menu, mas **não executa o daemon como root**. O socket externo é separado do socket interno e dos diretórios de modelos.

A unidade `doc-arquitetura-app.socket` define `SocketGroup=docarq-clients` e `SocketMode=0660` para o socket de controle; `doc-arquitetura-app.service` usa `User=docarq-svc`, `Group=docarq-runtime`, `RuntimeDirectory=doc-arquitetura-app-private` e `RuntimeDirectoryMode=0700`. O arquivo de configuração é root-owned, somente leitura pelo daemon; credenciais são entregues pelo `LoadCredential=` do systemd. Configuração inválida bloqueia o serviço. Esta definição substitui os exemplos antigos de permissões e lifecycle.

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
- Credenciais GitHub são cadastradas somente pelo root (GitHub App de instalação ou chave SSH de serviço), protegidas em secrets root-only e entregues ao daemon por systemd LoadCredential; não transitam no cliente nem são incluídas em logs/URLs. Consulte ADMIN-SETUP.md. Acesso a repositórios privados ou push fica `BLOCKED` até configuração segura dessas credenciais. O daemon não tem acesso automático ao SSH agent/credenciais do usuário.
- Criar branch de trabalho baseada no SHA exato de `origin/develop`, sem modificar o ponteiro de `develop`.
- Commit local e publicação na branch da tarefa exigem aprovação separada; push sem force e bloqueio de troca inesperada de HEAD.
- Quando houver conflito ou mudança remota, suspender publicação e solicitar nova revisão.
- Não criar PR/merge automático em `develop` sem autorização específica; a entrega termina com branch de tarefa publicada e evidências da execução.
- Na referência normativa, continuar usando exclusivamente `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0`, somente leitura. No alvo, documentação ativa em `/docs/charter`, `/docs/scope`, `/docs/implementation`; histórico em **`/docs/backup`**.

## 4. Modelos, fila e isolamento

Devido ao ambiente CPU de 4 núcleos/8 GiB, limitar a **uma inferência/modelo ativo** por vez; enfileirar as demais tarefas. Pode manter o Llama carregado por um período de inatividade configurável, mas nunca permitir que o cliente se conecte diretamente ao socket privado. O modelo não recebe permissões Git/shell, apenas texto e ferramentas estritamente catalogadas; a aplicação valida cada proposta antes de aplicar.

Requisitos de segurança a verificar nos testes: ACL/DAC de sockets e binário, `SO_PEERCRED`, credenciais Git, symlinks, diretórios temporários, reinício do daemon, queda do modelo, cancelamento, ausência de shell injection, isolamento de tarefas e aprovação vinculada ao UID/revisão.

## 5. Estratégia de entrega

1. **Documentar contratos** em Charter, Scope e Implementation, incluindo o novo [cadastro administrativo root](ADMIN-SETUP.md), e revisar gates normativos. Estado inicial `refinement`/`BLOCKED` para dependências ainda não aprovadas.
2. Implementar daemon de controle, IPC com credenciais e persistência de tarefas; testes locais Unix/grupos.
3. Mover clonagem, repositório normativo, seleção de modelos, prompts e inferência para o processo do serviço. **Desabilitar os caminhos legados de escrita no CLI**; não ligar o serviço com lógica incompleta.
4. Implementar branches por operação e fluxo separado de aprovação, commit e push.
5. Adicionar streaming de eventos/tempos e benchmark de execução.
6. Atualizar `package.yaml`, scripts Debian e unidade systemd, incluindo contas/grupos, binário privado e teste de instalação com usuário autorizado/não autorizado. Alterar versão do pacote em release próprio.
7. Teste de ponta a ponta em Debian: instalação/atualização, daemon ativo, cliente com e sem grupo, inferência, clone, task branch, diff, commit, push, reinício/cancelamento e logs.

**Não ativar automaticamente um daemon de testes sem implementar todas as fronteiras de segurança acima.**
