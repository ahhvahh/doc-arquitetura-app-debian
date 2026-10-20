# Plano de integração OpenAI — doc-arquitetura-app-debian

**Projeto-alvo:** `ahhvahh/doc-arquitetura-app-debian`, branch `develop`.
**Referência exclusiva:** `ahhvahh/doc-arquitetura`, branch `main`, diretório `ProjectCharter/1.0`, SHA consultado `652e1035a6ebb7ade0f21a6f2b9d08156ba35b37`.
**Estado:** proposta para detalhamento em `/docs/charter`, `/docs/scope`, `/docs/implementation`. Não é código executável nem aprovação de gate.

## Regras que devem permanecer invariantes

- Manter a referência normativa da OpenAI **somente leitura** em `doc-arquitetura/main`, com SHA registrado. Atualizar instruções remotas somente quando o operador solicitar; inferências não atualizam a norma.
- Criar/reconciliar quatro agentes independentes: `charter`, `scope`, `implementation`, `src`. Cada um recebe apenas os normativos da própria etapa e as regras comuns exigidas pelo orquestrador, compactados com referência rastreável a documentos e SHA. Não misturar prompts especializados com base persistente do agente.
- Para cada operação, selecionar exatamente **um** `ProjectCharter/1.0/prompts/<etapa>/<prompt>.txt`, com possibilidade de múltiplas tarefas deliberadamente separadas. Não enviar os índices README e os outros 33 prompts como contexto.
- Analisar a documentação real `/docs` do projeto-alvo na `develop`, após reconciliar o HEAD, compilando DSL documental fresca para **cada nova operação**. Não manter cache compartilhado de conteúdo entre sessões.
- A aplicação recebe changesets estruturados e **não executa comandos vindos da LLM**. Mudanças só são aplicadas depois da validação e das aprovações explicitamente previstas em `TASK-BRANCH-LIFECYCLE.md`. Impedir escrita na `main` normativa.
- **Fail closed para faturamento:** sem preço/modelo atual, limites positivos por operação/agente/global e câmbio vigente, não enviar inferência paga. Limite de sessão `spend_control.limit` em centavos USD e hard limit de projeto configurado separadamente no painel OpenAI. Hard limit não é instantâneo.

## CLI de configuração proposta (ainda não implementada)

```text
doc-arquitetura openai key set                   # credencial padrão, entrada oculta
doc-arquitetura openai key set --agent charter
doc-arquitetura openai key set --agent scope
doc-arquitetura openai key set --agent implementation
doc-arquitetura openai key set --agent src

doc-arquitetura openai key status            # apenas presente/ausente, sem revelar valor
doc-arquitetura openai agents sync --agent charter --model <model-id>
doc-arquitetura openai agents sync --agent scope --model <model-id>
doc-arquitetura openai agents sync --agent implementation --model <model-id>
doc-arquitetura openai agents sync --agent src --model <model-id>
doc-arquitetura openai agents status

doc-arquitetura openai prompts list --agent charter
doc-arquitetura openai analyze --repo https://github.com/OWNER/REPO --agent charter --prompt project-intake-review
```

### Entrada segura da chave

- `key set` solicita segredo diretamente do terminal com **echo desativado**, sem aceitá-lo como argumento de linha de comando. Evitar exposição em histórico, `/proc` e logs. Não persistir em `config.json`, Git nem no DSL. Preferir um armazenamento de segredos do sistema (ex.: Secret Service/libsecret no Linux), com identificadores do agente/projeto. Em ambiente sem cofre seguro, exigir procedimento explícito; nunca usar fallback silencioso.
- `OPENAI_API_KEY` pode ser admitida como substituição transitória, com indicação de origem e sem imprimir valor. Evitar passar chaves por mensagens neste chat.
- Uma chave padrão pode atender às quatro etapas **se todas estiverem no mesmo projeto OpenAI**. A configuração por agente sobrescreve a padrão apenas para aquela etapa.
- Se quatro projetos OpenAI forem usados, cada agente deve ter `project_id`, identidade de credencial e hard limit distintos. Um projeto compartilhado tem teto remoto **agregado**, não quatro tetos independentes.

## Composição normativa por agente

| Agente | Norma persistida (mesmo SHA main) | Prompts disponíveis sob demanda |
|---|---|---|
| Charter | `ProjectCharter/1.0/AGENT-PROMPT.md` + `charter/AGENT-PROMPT.md`, `charter/GUIDE.md` e DSL necessária do catálogo de exemplo Charter | `prompts/charter/*.txt`, excluindo `agente.txt` quando for apenas roteador |
| Scope | Regras comuns + `scope/AGENT-PROMPT.md`, `scope/GUIDE.md` e DSL de exemplo Scope | `prompts/scope/*.txt` |
| Implementation | Regras comuns + `implementation/AGENT-PROMPT.md`, documentação normativa de LIB/SVC/APP | `prompts/implementation/*.txt` |
| src | Regras comuns + `src/AGENT-PROMPT.md`, `src/README.md`, modelos de arquitetura aplicáveis | `prompts/src/*.txt` |

O conteúdo persistente não incorpora os 34 prompts. Armazenar localmente apenas ID remoto, SHA normativo, hash SHA-256 das instruções, projeto associado e estado da sincronização. Para atualizar um agente, resolver SHA da `main`, baixar os documentos necessários, montar conteúdo determinístico, estimar custo aplicável, mostrar diff/hash, solicitar aprovação e chamar API de criação/atualização. Não atualizar os quatro agentes implicitamente a cada inicialização.

## OpenAI API (verificar capacidades atuais na implementação)

- Cadastro de agente persistente: `POST /agents` com `model`, `instructions`, `name` e `metadata` (o objeto Agent pertence a um projeto OpenAI).
- Atualização explícita: `POST /agents/{agent_id}` com instruções e metadata, confirmando ID/hashes antes e depois.
- Execução: uma nova sessão `POST /agents/sessions` com `agent_id`, entrada documental filtrada e **limite de sessão positivo** `spend_control.limit` em centavos de USD. Não reutilizar sessão entre operações independentes. Validar autenticação, resposta, cobrança e status antes de permitir efeitos no workspace.
- Projetos OpenAI e `hard spend limit` administrados no painel / APIs administrativas com permissões específicas; **nunca assumir** que `spend_control` é um limite mensal do agente.

## Contratos/documentação a refinar antes de alterar /src

1. `/docs/charter`: validar status, resolver OPEN-001 (cofre/segredos), OPEN-002 (compactação e custos), OPEN-003 (task branches versus norma que exige develop) e OPEN-004 (projetos OpenAI/hard limits). Auditar a política de autorização e os objetivos.
2. `/docs/scope`: descrever atores e FLWs `configurar credencial`, `sincronizar agente`, `listar/selecionar prompt`, `preflight de custo`, `analisar projeto`, `propor changeset`, `aprovar e validar` e contratos TYPE/INT/ALG para payloads, quotas e estados.
3. `/docs/implementation`: definir APP/CMD para CLI, LIB para cliente OpenAI e montagem DSL, DEP para armazenamento de segredos e persistência de metadata, testes de segurança e mocks HTTP, Services só se realmente necessários.
4. `/src`: implementar a partir de unidades refinadas, preservando runtime llama.cpp existente e incluindo DOC-REFs de Implementation nos módulos pertinentes. Testes no GitHub Actions sem chaves e sem requisições pagas.

## Matriz inicial de testes automatizados

| ID | Cenário | Critério |
|---|---|---|
| KEY-001 | Prompt do token sem eco e sem argumento | Nenhum segredo exibido em stdout/stderr ou CLI args |
| KEY-002 | `status` após `key set` | Exibe apenas presença e projeto associado |
| KEY-003 | Falha de cofre / permissões | Bloqueia sem salvar em texto puro |
| NORM-001 | `main` tem quatro pastas | Resolver um SHA único e caminhos corretos |
| NORM-002 | Mudança de SHA sem `sync` | Não atualiza agentes remotos automaticamente |
| NORM-003 | Repetição do `sync` sem mudança | Idempotente: não duplica agente nem reenvia instruções |
| PROMPT-001 | Selecionar um prompt | Apenas um `.txt` no payload da inferência |
| PROMPT-002 | Prompt inexistente ou fora do diretório | Bloqueio sem rede |
| DSL-001 | Alteração na documentação-alvo | Nova DSL por operação, sem cache entre inferências |
| COST-001 | Preço/câmbio/limite ausente | Nenhuma chamada à API |
| COST-002 | Limite de sessão | `spend_control.limit` positivo e em centavos USD |
| API-001 | Criação/atualização do agente | Mock HTTP verifica método, URL, ID e hash |
| API-002 | Resposta inválida/429/timeout | Nenhuma edição/publicação remota |
| GIT-001 | Proposal/approval/commit/push | Etapas separadas, scope verificado, nenhuma alteração na main |

**Gate atual observado:** o README Charter do projeto-alvo declara `BLOCKED — Charter → Scope`, e não foram encontrados diretórios `/docs/scope` e `/docs/implementation`. Portanto, este arquivo é uma proposta de atualização, não autorização para gerar o código de produção.

## Referências verificadas em 2026-10-20

- https://developers.openai.com/api/reference/typescript/resources/beta/subresources/agents
- https://developers.openai.com/api/reference/ruby/resources/beta/subresources/agents/subresources/sessions/methods/create
- https://developers.openai.com/api/docs/guides/spend-limits