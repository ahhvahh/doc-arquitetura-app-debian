# Doc Arquitetura — Revisão documental e comparação de modelos

![Tipo: GUIDE](../assets/documentation-style/badges/badge-type-guide.svg) ![Status: refinement](../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../assets/documentation-style/badges/badge-version-v1.svg)

## Resumo executivo
A aplicação Doc Arquitetura permitirá organizar a revisão e o refinamento documental de projetos GitHub, utilizando diferentes provedores remotos de LLM, por quatro especializações: Charter, Scope, Implementation e Source. Compilará localmente a referência normativa externa `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0`, manterá as configurações de cada agente reutilizável até uma atualização explícita, analisará o `/docs` atualizado do projeto-alvo e produzirá alterações rastreáveis para revisão antes da publicação em branch de tarefa, condicionada à autorização normativa.

## Propósito do projeto
| Campo | Definição |
|---|---|
| Problema | Analisar projetos extensos exige contexto elevado, gera custos diferentes por modelo e ainda não oferece histórico comparável de qualidade, tokens e despesas. |
| Propósito | Permitir refinamento documental por diferentes provedores remotos de LLM, com sessões contínuas de inferência, histórico local em métricas JSON recuperáveis, comparação de qualidade e custos e aprovação humana para integrar alterações à `develop`. |
| Beneficiários | Desenvolvedor responsável, revisor documental e operador da aplicação Debian. |
| Valor esperado | Menor trabalho manual de contexto, consistência normativa, alterações verificáveis e escolha de modelos apoiada por indicadores históricos de qualidade, consumo, tempo e custo. |

## Contexto e participantes
- [PCH-001 — Projeto e participantes](project-context/PCH-001-contexto-integracao-openai-dsl.md).

## Contexto multimodelo proposto
O diagrama representa as relações pretendidas, não funcionalidades já implementadas.

```mermaid
flowchart LR
    O[Operador] --> A[Aplicação documental]
    R[Repositório GitHub] --> A
    N[Norma main] --> A
    A -. Avaliação proposta .-> M[Modelos distintos]
    M -. Respostas e medições .-> A
    A -. Histórico previsto .-> B[Banco de dados local]
    A --> H[Revisão humana]
```

## Visão geral de alto nível
```mermaid
flowchart TD
    A[Referencia normativa main] --> B[Compilacao DSL normativa por etapa]
    B --> C[Publicacao manual no Vector Store]
    C -->|Consulta quando integrada| F[Agente especializado]
    D[Projeto-alvo develop atualizado] --> E[DSL de docs por execucao]
    E --> F
    F --> L[Historico local previsto]
    F --> G[Changeset DSL com branch e commit]
    G --> V[Validacao local]
    V -. Somente apos alinhamento normativo .-> S[Branch da sessao]
    S --> K[Commit automatico validado]
    K --> P[Nova inferencia na mesma sessao]
    P --> F
    K --> H[Revisao para integracao]
    H -->|Aprovado e permitido| I[Integracao pretendida a develop]
    H -->|Nao aprovado ou bloqueado| J[Sem integracao]
```
A DSL normativa é atualizada no Vector Store somente por comando explícito do operador, nunca por preparação de tarefa ou inferência. A cópia é derivada do SHA da main normativa e não substitui o Markdown original; consulta pelo agente depende de integração de pesquisa ainda pendente. Cada execução usa workspace isolado e projeto-alvo independente. O diagrama descreve intenção de produto, **não concede permissão** de escrever em branches de sessão nem de executar merge: a [OPEN-003](open-questions/OPEN-003-politica-de-branches.md) continua bloqueadora até a atualização autorizada da política normativa. A autorização de gastos de inferência é independente da revisão para integração Git.

## Objetivos e entregas esperadas
- [OBJ-001 — Quatro especializações](objectives/OBJ-001-agentes-documentais-especializados.md)
- [OBJ-002 — Contexto compacto e rastreável](objectives/OBJ-002-contexto-compacto-rastreavel.md)
- [OBJ-003 — Propostas verificáveis para revisão](objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md)
- [OBJ-004 — Evidência de testes automatizados](objectives/OBJ-004-evidencias-de-testes-automatizados.md)
- [OBJ-005 — Inferência com orçamento controlado](objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OBJ-006 — Publicação manual da base normativa](objectives/OBJ-006-publicacao-manual-base-normativa.md)
- [OBJ-007 — Comparação entre modelos](objectives/OBJ-007-comparacao-controlada-de-modelos.md)
- [OBJ-008 — Histórico de inferências](objectives/OBJ-008-historico-local-de-inferencias.md)

## Fora do escopo
- [NOBJ-001 — Sem integração autônoma à develop](non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md)
- [NOBJ-002 — Sem inferência local na primeira entrega](non-objectives/NOBJ-002-sem-inferencia-local-na-primeira-entrega.md)

## Condições e premissas
- [CON-001 — Referência normativa versionada](conditions/CON-001-referencia-normativa-main.md)
- [CON-002 — Proteção de credenciais](conditions/CON-002-protecao-credenciais.md)
- [CON-003 — Separação entre norma externa e contratos do aplicativo](conditions/CON-003-contratos-do-aplicativo.md)
- [CON-004 — Orçamentos e aprovação prévia](conditions/CON-004-orcamento-obrigatorio.md)
- [CON-005 — Persistência local](conditions/CON-005-persistencia-local-de-inferencias.md)

## Questões abertas
- [OPEN-001 — Credenciais e armazenamento](open-questions/OPEN-001-politica-credenciais.md)
- [OPEN-002 — Fidelidade e orçamento de contexto](open-questions/OPEN-002-fidelidade-e-custo.md)
- [OPEN-003 — Política de branches](open-questions/OPEN-003-politica-de-branches.md)
- [OPEN-004 — Limites remotos por projeto OpenAI](open-questions/OPEN-004-projeto-openai-e-hard-limit.md)
- [OPEN-005 — Provedores e modelos](open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md)
- [OPEN-006 — Retenção e sigilo do histórico](open-questions/OPEN-006-retencao-e-protecao-do-historico.md)

## Base normativa pesquisável — decisão de Charter

O operador poderá solicitar a compilação e sincronização manual da norma, separadamente para Charter, Scope, Implementation e Development (src). O comando apresenta prévia sem publicação por padrão; aplicação explícita habilita o envio em formato textual pesquisável. A aplicação compara hashes, registra versões e confirma indexação, evitando uploads redundantes. A publicação não executa agentes nem altera repositórios. Regras de consulta e contratos da CLI serão definidos em Scope e Implementation após os gates. Veja [OBJ-006](objectives/OBJ-006-publicacao-manual-base-normativa.md) e [OPEN-002](open-questions/OPEN-002-fidelidade-e-custo.md).

## Situação da baseline
**BLOCKED — Charter → Scope.** Esta é uma proposta inicial em `refinement`. A implementação não está autorizada por este rascunho. Requer revisão do responsável, validação das OPEN e auditoria completa dos critérios normativos.

## Histórico e referência
- [Referência normativa](https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0) — somente leitura; `main`.
- [Projeto-alvo](https://github.com/ahhvahh/doc-arquitetura-app-debian/tree/develop) — branch de desenvolvimento.

## Registro da ampliação — 2026-10-21
A inclusão de comparação multimodelo e histórico local foi autorizada para refinamento documental; não equivale à conclusão de implementação ou do gate. As questões [OPEN-001](open-questions/OPEN-001-politica-credenciais.md), [OPEN-002](open-questions/OPEN-002-fidelidade-e-custo.md), [OPEN-003](open-questions/OPEN-003-politica-de-branches.md), [OPEN-004](open-questions/OPEN-004-projeto-openai-e-hard-limit.md), [OPEN-005](open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e [OPEN-006](open-questions/OPEN-006-retencao-e-protecao-do-historico.md) permanecem abertas para análise individual.

## Decisões funcionais de 2026-10-21 — provedores, sessões e métricas
- A [OPEN-005](open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) registra a opção B, múltiplos provedores remotos, mantendo a escolha inicial dos serviços pendente de viabilidade.
- A [OPEN-003](open-questions/OPEN-003-politica-de-branches.md) registra o fluxo desejado de **uma branch por sessão**, commits automáticos de resultados validados, novos prompts na sessão ativa e aprovação humana **somente da integração Git a `develop`**. A execução permanece bloqueada pela norma atual.
- O [OBJ-008](objectives/OBJ-008-historico-local-de-inferencias.md) exige métricas JSON recuperáveis para consulta e reconstrução de análises por aplicações autorizadas, sem impor formato interno do banco nesta etapa.
