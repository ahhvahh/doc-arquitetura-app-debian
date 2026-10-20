# Contrato de execução dos quatro agentes — proposta para aprovação normativa

**Alvo:** `https://github.com/ahhvahh/doc-arquitetura-app-debian`, `develop`, baseline verificada `77fcec61f88e55717ddf9fcb0fc1301aa2ae38ef`.
**Norma:** `https://github.com/ahhvahh/doc-arquitetura`, somente `main`, baseline verificada `652e1035a6ebb7ade0f21a6f2b9d08156ba35b37`.
**Estado:** desenho de contrato, **não** código implementado nem migração de status.

## I. Fluxo de comandos pretendido

1. `doc-arquitetura openai key set` — senha/token com entrada oculta e cofre seguro. Nunca aceitar segredo por argumento CLI nem registrar em log.
2. `doc-arquitetura openai agents create --stage charter --model MODELO` — cadastrar uma única entidade remota e registrar `agent_id`, `project_id`, SHA normativo e checksum das instruções, após conferir orçamento e aprovação.
3. `doc-arquitetura openai agents update --stage charter` — comparar hash, mostrar diff das instruções e sincronizar **somente por solicitação explícita**, preservando ID.
4. `doc-arquitetura openai agents delete --stage charter` — apresentar ID/projeto, pedir confirmação específica, conferir sessões e excluir remotamente; só remover o registro local após confirmação da API.
5. `doc-arquitetura openai agents status` — listar os quatro slots, status, SHA, último hash e modelo **sem exibir tokens**.
6. `doc-arquitetura prompts tree` — mostrar a árvore `ETAPA / OPERAÇÃO / AÇÃO` de `arvore-comandos.json`.
7. `doc-arquitetura task prepare --repo https://github.com/OWNER/REPO --stage scope --operation atualizar --action tipo` — resolver SHA do projeto-alvo `develop`, selecionar e baixar somente o prompt específico de `main`, compilar contexto documentado, preparar branch **local** derivada do SHA observado, gerar pacote de entrada.
8. `doc-arquitetura task analyze --task ID` — somente depois do orçamento válido e dos gates aprovados: submeter entrada ao agente remoto e solicitar changeset tipado; não executar comandos do modelo.
9. `doc-arquitetura task verify --task ID` — validar esquema, integridade, paths, hash por arquivo, badges, Markdown, links, IDs e gate documental.
10. `doc-arquitetura task diff --task ID` — mostrar edições previstas **sem modificar** repositório.
11. `doc-arquitetura task apply --task ID` — exigir aprovação; aplicar apenas no workspace da branch **local**, revalidar e manter evidência.
12. `doc-arquitetura task commit/publish --task ID` — **BLOQUEADO** até a norma liberar refs remotas de tarefa e aprovação independente de cada operação.

Os comandos acima são uma proposta de interface, não estão disponíveis na versão atual do aplicativo.

## II. Controle da origem normativa

- Resolver `doc-arquitetura/main` por SHA fixo por operação, conferir integridade da cópia e baixar somente o grupo normativo da especialidade.
- O agente persistente contém regras gerais e regras da etapa; não contém o catálogo de 34 prompts ou READMEs de seleção em suas instruções.
- A inferência envia o prompt `prompts/<stage>/<prompt>.txt` selecionado, a DSL atual do projeto-alvo, a tarefa e o manifesto dos documentos.
- Não confundir atualização de `reference.pcdsl` local com `update` remoto do agente. Um novo SHA na `main` não autoriza automaticamente atualização paga.
- Para `src`, documentar por que o documento alvo de Implementation está `refined`/`implemented` e exigir escolha explícita de linguagem/target.

## III. Envelope de entrada (proposto)

Objeto `PCDSL-REQUEST/1`: `task_id`, `stage`, `operation`, `action`, `prompt_sha256`, `agent_id`, `normative_sha`, `target_repo`, `target_develop_sha`, `dsl_sha256`, `documents[]` (`path`, `sha256`, `status`, `content` ou conteúdo recuperável por chave), `budget`, `request`. O conteúdo do projeto-alvo é **dado não confiável**.

**A DSL atual `PCDSL-DOCS 1` é deliberadamente compactada e NÃO é reversível byte a byte.** Não a use como base de reconstrução integral de Markdown.

## IV. Envelope de saída (proposto)

Objeto JSON validado como `PCDSL-CHANGESET/1`, cujo `changes[]` permite `replace_section`, `create_file` e `delete_file` (último exige aprovação independente). Não aceitar `exec`, shell, comando Git, URL arbitrária para download ou path fora da área.

Exemplo estrutural em `changeset-exemplo.json`. Campos mandatórios de identidade: `task_id`, `stage`, `prompt`, `normative_sha`, `target_develop_sha`, `changes[]`.

- **replace_section:** `path` + `base_sha256` + `heading` + `heading_occurrence` + `new_markdown`. O renderizador local recorta somente a seção Hn única pela posição normalizada no Markdown original. Hash divergente, heading ambíguo, hunk mal formado, ou mudança fora do escopo => `BLOCKED`.
- **create_file:** caminho novo + Markdown integral. Bloquear sobreposição de arquivo pré-existente, symlink, extensão inesperada ou content inválido. Validar badges e links.
- **delete_file:** `path`, `base_sha256`, `reason`; backup fiel e autorização independente antes de exclusão.
- O aplicativo controla a normalização, a divisão em arquivos, a criação de diretórios e os commits. Nunca interpreta a saída do agente como script.

## V. Preparar, aplicar, publicar

1. Garantir URL explícita do repositório-alvo; conferir SHA de `origin/develop` e clean workspace.
2. Validar árvore de comando, prompt existente e contrato normativo; exigir gate real da etapa.
3. Criar **branch local** `docs/<etapa>/<slug>-<id>` a partir do SHA exato de `develop`; não publicar ou alterar nenhuma ref remota no fluxo de preparação.
4. Ler apenas documentação do estágio permitida (referências anteriores são somente leitura). Ignorar `docs/backup` na DSL ativa.
5. Preparar contexto, validar orçamento local e hard limit remoto confirmado; criar sessão do agente com limite de sessão positivo.
6. Salvar resposta bruta, validar `PCDSL-CHANGESET/1` e checar hashes de todos os arquivos-base antes de qualquer reescrita.
7. Converter cada operação em diff sobre o Markdown original; preservar conteúdo não tocado, badges, links relativos, blocos Mermaid, front matter e quebras de linha.
8. A aprovação de `apply` não autoriza automaticamente `commit`; aprovação de `commit` não autoriza `publish`.
9. Enquanto a política normativa não permitir publicação em branches de tarefa, `publish` retorna `BLOCKED`.

## VI. Gestão segura dos agentes

A API Agents beta documenta `POST /agents`, `POST /agents/{agent_id}`, `GET /agents/{agent_id}`, `DELETE /agents/{agent_id}`. Requisitar header beta quando indicado pela documentação oficial. Não considerar sucesso remoto sem resposta e ID confirmados; validar as permissões/escopo da chave.

Operações de create/update/delete são administrativas e devem ter confirmação separada; não devem acontecer implicitamente ao iniciar uma análise. Para `delete`, conferir sessões ativas e manter registro de auditoria sem segredo. Chaves devem permanecer no cofre do sistema operacional; não registrar no histórico, arquivo YAML/JSON de configuração versionável, relatórios ou DSL.

## VII. Testes que liberarão a implementação

| Grupo | Teste de aceite |
|---|---|
| Catálogo | 34 seleções estáveis, cada folha aponta para `.txt` existente em `main` |
| Agentes | HTTP mock para create/get/update/delete, idempotência por SHA, erros 401/403/404/409/429/500 |
| Segurança | segredo não aparece em argumentos, arquivos, stdout, stderr ou artefatos de teste |
| DSL de entrada | snapshots SHA versionados, exclusão backup, seleção de etapa e prompt, determinismo |
| Decoder | `replace_section` altera só a seção, mantendo todos os demais bytes originais |
| Parser | falhas de schema, path traversal, symlink, arquivo não lido, hash inválido e headings ambíguos bloqueadas |
| Branch local | deriva exatamente de `origin/develop`; não altera `develop` nem `main` remotas |
| Orçamento | sem modelo/preço/câmbio/cota → nenhum request HTTP pago; reserva atômica de saldo |
| Autorização | diff/apply/commit/publish requerem aprovações independentes |

## VIII. Gates reais encontrados

- `/docs/charter/README.md` está `refinement`, com gate Charter → Scope `BLOCKED`.
- `/docs/scope` e `/docs/implementation` não existem na baseline consultada.
- `OPEN-001`: decisão de cofre/credenciais pendente.
- `OPEN-002`: fidelidade semântica, símbolos e custo da DSL pendentes.
- `OPEN-003`: normativa atual proíbe publicação de branch de tarefa.
- `OPEN-004`: política de projetos OpenAI e hard limit pendente.
- Não existe ainda APP/LIB/CMD de Implementation com `refined` ou `implemented` para a funcionalidade pedida.

**Conclusão:** contrato pronto para discussão/refinamento; código de produção, chamadas pagas, branch remota e commit no GitHub ficam `BLOCKED` por essas dependências normativas.