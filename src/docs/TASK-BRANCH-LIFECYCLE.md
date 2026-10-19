# Contrato — Branch por operação, aprovação e telemetria

> Estado: **refinement**. Especificação proposta, ainda não implantada.
> Regra inegociável: `main` é somente leitura; `develop` é exclusivamente a origem das branches de tarefa.

## 1. Nomear a operação antes de processar

Ao receber a descrição completa da tarefa, produzir **título em português até 100 caracteres** e identificar a etapa (`charter`, `scope`, `implementation`). Primeiro sugerir um título pelo agente, quando disponível, ou gerar slug determinístico como fallback. A geração e sanitização do nome são **feitas pelo serviço**, não por execução de código do modelo.

Formato proposto do nome da branch:

```text
docs/<etapa>/<resumo-ascii-lowercase>-<id-unico>
docs/charter/revisar-objetivos-de-recepcao-7d21c6
docs/scope/atualizar-fluxo-de-integracao-598bc4
docs/implementation/revisar-contratos-de-servico-f8ad03
```

- Limite máximo: **100 bytes** para o nome completo de `refs/heads/<branch>` (incluindo o prefixo), sem caracteres inválidos, barras duplas, `..`, `@{`, `.lock` ou referência vazia.
- Sanitizar espaços/acentos/pontuação, manter apenas `[a-z0-9/-]`, limitar o trecho descritivo, acrescentar ID único derivado de aleatoriedade segura, verificar formato com `git check-ref-format --branch`.
- Capturar SHA de `origin/develop` da operação e validar a branch remota; em caso de colisão gerar outra. Nunca reaproveitar uma branch de tarefa para operação não relacionada.
- A branch pode existir apenas no clone até a publicação **aprovada**. Não publicar branch vazia no GitHub por padrão.
- Cada nova operação cria **novo workspace e nova branch**, não continua commits de outra operação. Retomada usa o `task_id` da mesma operação e não cria outra branch.

## 2. Etapas e relógios

Estados persistidos: `QUEUED → NAMING → FETCHING → BRANCH_CREATED → READING_NORMS → INVENTORY → LOADING_MODEL → ANALYZING → PROPOSAL_READY → AWAITING_APPROVAL → APPLYING → VALIDATING → AWAITING_COMMIT_APPROVAL → COMMITTING → AWAITING_PUSH_APPROVAL → PUBLISHING → COMPLETED`. Transições alternativas: `BLOCKED`, `FAILED`, `CANCELLED`.

Cada evento contém:

```json
{
  "task_id": "t-47c8d5",
  "event_seq": 17,
  "timestamp": "2026-10-19T12:47:00-03:00",
  "stage": "ANALYZING",
  "state": "running",
  "message": "Analisando documentos Charter",
  "stage_elapsed_ms": 12700,
  "task_elapsed_ms": 45900,
  "current": null,
  "total": null,
  "percent": null
}
```

O percentual só aparece quando existe denominador confiável (ex.: `12/40 documentos` => `30%`, bytes de clone conhecidos). Para carregar modelo, inferência ou operações Git sem total conhecido, exibir `em execução` e duração crescente, **não inventar porcentagem**. O tempo por estágio começa na sua transição; tempo total começa no momento do aceite da operação e inclui fila/espera por aprovação, com `active_elapsed` distinguindo processamento de `waiting_elapsed`.

Consolidação no terminal (ilustrativa):

```text
[12:47:03] T-47C8D5 | BRANCH_CREATED | 100% | etapa 0,8s | total 4,3s
[12:47:11] T-47C8D5 | INVENTORY      | 30% (12/40) | etapa 7,4s | total 12,1s
[12:48:42] T-47C8D5 | ANALYZING      | em execução | etapa 1m31s | total 1m43s
[12:49:17] T-47C8D5 | AWAITING_APPROVAL | 0s | total 2m18s
[12:51:12] T-47C8D5 | PUBLISHING    | em execução | etapa 3s | total 4m13s
[12:51:16] T-47C8D5 | COMPLETED     | total 4m17s | ativo 2m22s | espera 1m55s
```

Eventos numéricos são persistidos com sequência monotônica e consumidos pelo CLI em streaming, com reconnect/replay por `event_seq`. CLI mostra progresso e tempo continuamente mesmo durante ações sem eventos novos. Eventos com falha incluem operação, código seguro, tempo decorrido e possibilidade de retry idempotente. `SIGINT` cancela apenas a assinatura do console; cancelamento do trabalho exige comando explícito `task.cancel`.

## 3. Aprovação e publicação

O modelo propõe documentos somente em `/docs/<etapa>/`; backups ficam em `/docs/backup/` com manifesto e aprovação próprios. O serviço mostra diff, limita escopo, valida gates, hashes e autoria. Se detectar outro arquivo alterado, symlink ou falha de gate, bloqueia.

As aprovações são operações distintas, vinculadas a **UID do dono da tarefa, task_id, revisão do plano, SHA base e SHA do diff**:

1. Autorizar snapshot em `/docs/backup/`, quando necessário.
2. Autorizar aplicação das alterações documentais.
3. Autorizar criação de commit local na branch da tarefa.
4. Autorizar push remoto exclusivamente para `refs/heads/docs/<etapa>/<nome-unico>`.

Para publicação, confirmar novamente branch alvo e SHA base, usar `git push origin HEAD:refs/heads/<task-branch>` sem `--force`. Validar que `refs/heads/develop` e `refs/heads/main` são destinos proibidos, mesmo se o cliente tentar forjar nomes ou enviar outro comando. O serviço nunca faz merge automático em `develop`.

## 4. Critérios de aceite e pendências

- Comandos da mesma tarefa recebidos mais de uma vez não criam branches duplicadas.
- Dois usuários diferentes não conseguem ler histórico privado/aprovar/publicar tarefas um do outro sem política administrativa explícita.
- A branch existe e deriva exatamente do SHA de `develop` lido para a tarefa.
- Falha de rede, reinício systemd ou perda do console não altera `main` ou `develop`.
- Todos os eventos apresentam `stage_elapsed_ms` e `task_elapsed_ms`; `percent` pode ser nulo.
- Sem token de GitHub de serviço com permissão adequada, o push fica `BLOCKED`, sem tentativa de tomar credenciais do usuário.
- **Pendente para refinamento:** mecanismo oficial de credenciais GitHub no daemon; política de administração/approvação entre operadores; retenção de eventos, diffs e históricos; migração dos GGUF existentes do diretório do usuário para o diretório privado do serviço; limites e política de concorrência/retry. Esses pontos precisam de decisão antes de implementar operações remotas efetivas.


## 5. Controle administrativo de acesso

O cadastro de usuários autorizados, a identidade técnica e a credencial GitHub são definidos apenas em [ADMIN-SETUP.md](ADMIN-SETUP.md), por root. O daemon valida `SO_PEERCRED`, grupo Linux e allowlist no arquivo root-owned antes de aceitar comandos. O `task_id` e o UID do criador são obrigatórios para aprovações; não transmitir credenciais do GitHub pela CLI.
