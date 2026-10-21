# Doc Arquitetura — Integração OpenAI e DSL

![Tipo: GUIDE](../assets/documentation-style/badges/badge-type-guide.svg) ![Status: refinement](../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../assets/documentation-style/badges/badge-version-v1.svg)

## Resumo executivo
A aplicação Doc Arquitetura permitirá organizar a revisão e o refinamento documental de projetos GitHub por quatro especializações: Charter, Scope, Implementation e Source. Compilará localmente a referência normativa externa `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0`, manterá as configurações de cada agente reutilizável até uma atualização explícita, analisará o `/docs` atualizado do projeto-alvo e produzirá alterações rastreáveis para revisão antes da publicação em branch de tarefa, condicionada à autorização normativa.

## Propósito do projeto
| Campo | Definição |
|---|---|
| Problema | A revisão de um repositório documental extenso por modelos locais pode ser lenta; reenviar todo o Markdown eleva o custo e dificulta a rastreabilidade de alterações. |
| Propósito | Permitir análise documental assistida por IA com uso eficiente de contexto, resultados rastreáveis e revisão humana antes da publicação no GitHub. |
| Beneficiários | Desenvolvedor responsável, revisor documental e operador da aplicação Debian. |
| Valor esperado | Menos trabalho manual de preparação de contexto, atualizações consistentes com a norma e alterações verificáveis antes de commit. |

## Contexto e participantes
- [PCH-001 — Projeto e participantes](project-context/PCH-001-contexto-integracao-openai-dsl.md).

## Visão geral de alto nível
```mermaid
flowchart TD
    A[Referencia normativa main] --> B[Compilacao DSL normativa por etapa]
    B --> C[Publicacao manual no Vector Store]
    C -->|Consulta quando integrada| F[Agente especializado]
    D[Projeto-alvo develop atualizado] --> E[DSL de docs por execucao]
    E --> F
    F --> G[Changeset DSL com branch e commit]
    G --> H[Validacao local e revisao humana]
    H -->|Autorizado e politica alinhada| I[Branch da tarefa e publicacao]
    H -->|Rejeitado ou bloqueado| J[Nenhuma alteracao remota]
```
A DSL normativa é atualizada no Vector Store somente por comando explícito do operador, nunca por preparação de tarefa ou inferência. A cópia é derivada do SHA da main normativa e não substitui o Markdown original; consulta pelo agente depende de integração de pesquisa ainda pendente. Cada execução usa workspace isolado e projeto-alvo independente. O diagrama não autoriza gravar fora de develop.

## Objetivos e entregas esperadas
- [OBJ-001 — Quatro especializações](objectives/OBJ-001-agentes-documentais-especializados.md)
- [OBJ-002 — Contexto compacto e rastreável](objectives/OBJ-002-contexto-compacto-rastreavel.md)
- [OBJ-003 — Propostas verificáveis para revisão](objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md)
- [OBJ-004 — Evidência de testes automatizados](objectives/OBJ-004-evidencias-de-testes-automatizados.md)
- [OBJ-005 — Inferência com orçamento controlado](objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OBJ-006 — Publicação manual da base normativa](objectives/OBJ-006-publicacao-manual-base-normativa.md)

## Fora do escopo
- [NOBJ-001 — Sem publicação autônoma](non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md)

## Condições e premissas
- [CON-001 — Referência normativa versionada](conditions/CON-001-referencia-normativa-main.md)
- [CON-002 — Proteção de credenciais](conditions/CON-002-protecao-credenciais.md)
- [CON-003 — Separação entre norma externa e contratos do aplicativo](conditions/CON-003-contratos-do-aplicativo.md)
- [CON-004 — Orçamentos e aprovação prévia](conditions/CON-004-orcamento-obrigatorio.md)

## Questões abertas
- [OPEN-001 — Credenciais e armazenamento](open-questions/OPEN-001-politica-credenciais.md)
- [OPEN-002 — Fidelidade e orçamento de contexto](open-questions/OPEN-002-fidelidade-e-custo.md)
- [OPEN-003 — Política de branches](open-questions/OPEN-003-politica-de-branches.md)
- [OPEN-004 — Limites remotos por projeto OpenAI](open-questions/OPEN-004-projeto-openai-e-hard-limit.md)

## Base normativa pesquisável — decisão de Charter

O operador poderá solicitar a compilação e sincronização manual da norma, separadamente para Charter, Scope, Implementation e Development (src). O comando apresenta prévia sem publicação por padrão; aplicação explícita habilita o envio em formato textual pesquisável. A aplicação compara hashes, registra versões e confirma indexação, evitando uploads redundantes. A publicação não executa agentes nem altera repositórios. Regras de consulta e contratos da CLI serão definidos em Scope e Implementation após os gates. Veja [OBJ-006](objectives/OBJ-006-publicacao-manual-base-normativa.md) e [OPEN-002](open-questions/OPEN-002-fidelidade-e-custo.md).

## Situação da baseline
**BLOCKED — Charter → Scope.** Esta é uma proposta inicial em `refinement`. A implementação não está autorizada por este rascunho. Requer revisão do responsável, validação das OPEN e auditoria completa dos critérios normativos.

## Histórico e referência
- [Referência normativa](https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0) — somente leitura; `main`.
- [Projeto-alvo](https://github.com/ahhvahh/doc-arquitetura-app-debian/tree/develop) — branch de desenvolvimento.
