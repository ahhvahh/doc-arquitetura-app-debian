# Controle do Charter — Doc Arquitetura App Debian

![Tipo: GUIDE](../assets/documentation-style/badges/badge-type-guide.svg) ![Status: refinement](../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../assets/documentation-style/badges/badge-version-v1.svg)

> Painel operacional de continuidade do Charter. Os documentos canônicos vinculados são a fonte das decisões de produto; este registro não altera seus contratos, não autoriza operações pagas e não substitui o gate normativo.

## Ideia geral do projeto

O **Doc Arquitetura App Debian** pretende auxiliar operadores autorizados a revisar e refinar documentos de projetos GitHub usando quatro especializações (Charter, Scope, Implementation e Source), DSL rastreável, inferência remota, avaliação independente, comparação de modelos e histórico local. Busca reduzir custos de contexto, preservar a fidelidade documental e impedir publicação ou despesas não autorizadas. Beneficia responsáveis por projetos, revisores e operadores. A primeira entrega tem escopo reduzido: referência normativa fixa na `main`, análise qualitativa sem pontuação numérica não calibrada, modelos remotos (OpenAI e AWS AgentCore planejados) e automação de branches/PR adiada. **Nada disso comprova homologação ou implementação.** Fonte: [apresentação e propósito](README.md#propósito-do-projeto), [PCH-001](project-context/PCH-001-contexto-integracao-openai-dsl.md) e [OPEN-003](open-questions/OPEN-003-politica-de-branches.md).

## O que o projeto fará

| Documento | O que fará | Por quê / valor esperado | Situação de revisão | Demanda |
|---|---|---|---|---|
| [OBJ-001](objectives/OBJ-001-agentes-documentais-especializados.md) | Selecionar agentes documentais por Charter, Scope, Implementation e Source | Separar responsabilidades e manter a norma associada a cada agente | ● Revisão pendente (`refinement`) | D-002 |
| [OBJ-002](objectives/OBJ-002-contexto-compacto-rastreavel.md) | Compactar e rastrear o contexto documental | Reduzir consumo sem perder conteúdo obrigatório | ● Revisão pendente (`refinement`) | D-003 |
| [OBJ-003](objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md) | Gerar propostas de mudança identificáveis e revisáveis | Evitar alterações sem governança ou autorização | ● Revisão pendente (`refinement`) | D-004 |
| [OBJ-004](objectives/OBJ-004-evidencias-de-testes-automatizados.md) | Registrar evidências automatizadas por ação | Permitir avaliação reprodutível dos critérios de aceite | ● Revisão pendente (`refinement`) | D-003 |
| [OBJ-005](objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) | Impedir inferências pagas sem proteção financeira comprovada | Evitar despesas não autorizadas e não bloqueáveis | ● Revisão pendente (`refinement`) | D-001 |
| [OBJ-006](objectives/OBJ-006-publicacao-manual-base-normativa.md) | Preparar e publicar manualmente base normativa consultável | Evitar reenvio integral da norma a cada solicitação | ● Revisão pendente (`refinement`) | D-003 |
| [OBJ-007](objectives/OBJ-007-comparacao-controlada-de-modelos.md) | Comparar modelos e revisores com a mesma baseline | Avaliar aderência documental, tempo, consumo e custo | ● Revisão pendente (`refinement`) | D-002 / D-003 |
| [OBJ-008](objectives/OBJ-008-historico-local-de-inferencias.md) | Manter histórico local auditável de solicitações e resultados | Permitir retomada e análise comparativa posterior | ● Revisão pendente (`refinement`) | D-001 / D-003 |
| [OBJ-009](objectives/OBJ-009-execucao-automatica-de-tarefas-elegiveis.md) | Executar tarefas elegíveis automaticamente sem aprovação geral do plano | Reduzir interrupções, mantendo autorização financeira individual e validação documental | ● Revisão pendente (`refinement`) | D-007 |

## O que o projeto não fará

| Documento | Exclusão | Justificativa e consequência | Responsabilidade remanescente | Revisão | Demanda |
|---|---|---|---|---|---|
| [NOBJ-001](non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md) | Não integrar alterações em develop autonomamente | Preservar aprovação humana final; integração permanece responsabilidade do operador | Operador/repositório | ● Revisão pendente | D-004 |
| [NOBJ-002](non-objectives/NOBJ-002-sem-inferencia-local-na-primeira-entrega.md) | Não executar LLM local na primeira entrega | Reduzir complexidade operacional do MVP | Evolução futura, por nova decisão | ● Revisão pendente | D-002 |
| [NOBJ-003](non-objectives/NOBJ-003-nao-gerenciar-concorrencia-na-branch-da-sessao.md) | Não conciliar gravações concorrentes de terceiros na branch de sessão | A sessão planejada é exclusiva; conflitos externos exigem intervenção | Operador/responsável pelo repositório | ● Revisão pendente | D-004 |
| [NOBJ-004](non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md) | Não integrar Anthropic, Gemini ou outros provedores diretamente no MVP | Concentrar a homologação em OpenAI e AWS AgentCore | Nova evolução autorizada | ● Revisão pendente | D-002 |

## Mapa de documentos e cobertura

Baseline do inventário: `develop` em `a512f36fdbf9f9c6776bf214e8c25a5c3d50d93a` (anterior à correção editorial). Hashes abaixo são dos *blobs* canônicos nesta baseline; não representam confirmação de implementação. Os **27 documentos de negócio** identificados (1 PCH, 9 OBJ, 4 NOBJ, 5 CON, 6 OPEN e 2 DSC) têm badge `refinement`. A apresentação [README](README.md) também está em `refinement`.

| Documento canônico | Assunto | Status de origem | Hash do blob | Revisão | Demanda |
|---|---|---|---|---|---|
| [PCH-001](project-context/PCH-001-contexto-integracao-openai-dsl.md) | Contexto, atores e limites | `refinement` | `0d5e46420bd4` | ● Revisão pendente | D-005 |
| [OBJ-001](objectives/OBJ-001-agentes-documentais-especializados.md) | Agentes especializados | `refinement` | `5bb837fa6b3d` | ● Revisão pendente | D-002 |
| [OBJ-002](objectives/OBJ-002-contexto-compacto-rastreavel.md) | Contexto compacto | `refinement` | `57f8f423373c` | ● Revisão pendente | D-003 |
| [OBJ-003](objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md) | Propostas auditáveis | `refinement` | `78ae50aaf49f` | ● Revisão pendente | D-004 |
| [OBJ-004](objectives/OBJ-004-evidencias-de-testes-automatizados.md) | Evidências de testes | `refinement` | `17c787efa8d7` | ● Revisão pendente | D-003 |
| [OBJ-005](objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) | Orçamento e bloqueios | `refinement` | `7b99c182f5d9` | ● Revisão pendente | D-001 |
| [OBJ-006](objectives/OBJ-006-publicacao-manual-base-normativa.md) | Acervo normativo pesquisável | `refinement` | `4e804af770f1` | ● Revisão pendente | D-003 |
| [OBJ-007](objectives/OBJ-007-comparacao-controlada-de-modelos.md) | Comparação de modelos | `refinement` | `d4b274f7bad8` | ● Revisão pendente | D-002 / D-003 |
| [OBJ-008](objectives/OBJ-008-historico-local-de-inferencias.md) | Histórico local | `refinement` | `e5021ef01338` | ● Revisão pendente | D-003 |
| [OBJ-009](objectives/OBJ-009-execucao-automatica-de-tarefas-elegiveis.md) | Orquestração automática de tarefas elegíveis | `refinement` | `820d56002501` | ● Revisão pendente | D-007 |
| [NOBJ-001](non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md) | Publicação sem autorização | `refinement` | `6d715d19c6e6` | ● Revisão pendente | D-004 |
| [NOBJ-002](non-objectives/NOBJ-002-sem-inferencia-local-na-primeira-entrega.md) | LLM local no MVP | `refinement` | `9539c9d1814c` | ● Revisão pendente | D-002 |
| [NOBJ-003](non-objectives/NOBJ-003-nao-gerenciar-concorrencia-na-branch-da-sessao.md) | Concorrência externa na sessão | `refinement` | `efd118109d26` | ● Revisão pendente | D-004 |
| [NOBJ-004](non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md) | Provedores diretos extras no MVP | `refinement` | `7a70b36162c2` | ● Revisão pendente | D-002 |
| [CON-001](conditions/CON-001-referencia-normativa-main.md) | Fonte normativa fixa | `refinement` | `4fc8609e55f0` | ● Revisão pendente | D-004 |
| [CON-002](conditions/CON-002-protecao-credenciais.md) | Proteção de credenciais | `refinement` | `b86298e18b07` | ● Revisão pendente | D-001 / D-005 |
| [CON-003](conditions/CON-003-contratos-do-aplicativo.md) | Separação dos contratos | `refinement` | `bb155e6b40f4` | ● Revisão pendente | D-003 |
| [CON-004](conditions/CON-004-orcamento-obrigatorio.md) | Proteção financeira obrigatória | `refinement` | `30dacdd233a1` | ● Revisão pendente | D-001 |
| [CON-005](conditions/CON-005-persistencia-local-de-inferencias.md) | Persistência local | `refinement` | `a969d0b75ab0` | ● Revisão pendente | D-003 |
| [OPEN-001](open-questions/OPEN-001-politica-credenciais.md) | Credenciais e acesso | `refinement` | `bd857b2214cc` | ● Revisão pendente | D-005 / D-001 |
| [OPEN-002](open-questions/OPEN-002-fidelidade-e-custo.md) | Fidelidade DSL e custo | `refinement` | `f9ea1d49ae99` | ● Revisão pendente | D-003 |
| [OPEN-003](open-questions/OPEN-003-politica-de-branches.md) | Branches e PRs | `refinement` | `b6eb02114e08` | ● Revisão pendente | D-004 |
| [OPEN-004](open-questions/OPEN-004-projeto-openai-e-hard-limit.md) | Proteções financeiras remotas | `refinement` | `6e721316b99f` | ● Revisão pendente | D-001 |
| [OPEN-005](open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) | Homologação de provedores | `refinement` | `2bb637d3e1a1` | ● Revisão pendente | D-002 |
| [OPEN-006](open-questions/OPEN-006-retencao-e-protecao-do-historico.md) | Retenção e descarte | `refinement` | `6f4036d7af8b` | ● Revisão pendente | D-003 |
| [DSC-001](discovery/openai-api/DSC-001-viabilidade-openai-api.md) | Viabilidade e proteção financeira OpenAI | `refinement` | `c8753877f7ab` | ● Revisão pendente | D-001 / D-002 |
| [DSC-002](discovery/aws-agentcore/DSC-002-viabilidade-aws-agentcore.md) | Viabilidade AWS AgentCore e cobertura financeira | `refinement` | `4bba8f963983` | ● Revisão pendente | D-001 / D-002 |

**Cobertura:** README principal, PCH-001, OBJ-001 a OBJ-009, NOBJ-001 a NOBJ-004, CON-001 a CON-005, OPEN-001 a OPEN-006, DSC-001/002 e seus índices verificados na baseline. Não se identificaram documentos `ASM-*`; sua necessidade exige avaliação, não criação automática. As duas DSC estão publicadas, mas sua homologação administrativa permanece pendente. `docs/README.md` não existe e está fora do escopo de escrita deste agente; `badge-status-refined.svg` ainda não é referenciado por documentos atuais. Os 37 Markdown do Charter foram verificados quanto a links relativos de arquivos e presença do trio de badges; a verificação integral de âncoras, URLs externas e implementação em `/src` não foi realizada.

## Pontos a revisar e questões abertas

| Demanda | Questão / documento | Estado de trabalho | Situação de revisão | Criticidade / motivo | Depende de | Bloqueia | Critério de conclusão | Próxima ação |
|---|---|---|---|---|---|---|---|---|
| D-001 | OPEN-004 / CON-004 / OBJ-005 — comprovar bloqueio financeiro remoto | em_analise | ● Revisão pendente | P0 — limite local de US$ 10 por sessão não substitui proteção remota comprovada; chamadas pagas continuam bloqueadas | — | D-002; gate de inferência paga; execução remunerada de D-007 | Evidência administrativa datada por projeto OpenAI e, separadamente, por provedor; eficácia do bloqueio remoto compatível com a política financeira | Priorizar DSC-001 e os quatro projetos OpenAI, sem chamadas pagas; manter AWS bloqueada e DSC-002 inconclusiva até homologação independente |
| D-002 | OPEN-005 / NOBJ-004 — viabilidade de OpenAI e AWS AgentCore | a_analisar | ● Revisão pendente | P0 — DSC documentadas, porém elegibilidade, preços, contas e proteção financeira ainda não homologados | D-001 para chamadas pagas | OBJ-001; OBJ-007; gate do MVP | Discovery conclusivo, elegibilidade, modelos e restrições validados | Homologar evidências de DSC-001/002 por provedor, sem ativar cobrança |
| D-003 | OPEN-002 / OBJ-002 / OBJ-004 — equivalência e validação DSL | a_analisar | ● Revisão pendente | P1 — compromete confiança nas propostas e seus dependentes | — | OBJ-003; OBJ-006; OBJ-007 | Critérios de fidelidade e testes de preservação rastreáveis | Confirmar evidências de equivalência e critérios aceitos |
| D-004 | OPEN-003 / OBJ-003 / NOBJ-001 — automação de branches e PR | bloqueado | ● Revisão pendente | P1 — conflito normativo para função futura; automação adiada no MVP | Processo separado de alteração da norma | Automação Git de sessões; não o MVP reduzido | Regra compatível aprovada e controles Git verificados | Manter função desabilitada e preservar decisão de adiamento |
| D-005 | PCH-001 — consolidar ACT-005, ACT-006 e ACT-007 | concluido | ● Revisão pendente no PCH-001 | P1 — cadastro canônico necessário para governança dos atores | Decisões já documentadas em OPEN-001 e PCH-001 | Rastreabilidade dos papéis | Seções únicas por ID com papel, interesse, interações e limites | Regularização documental comprovada; revisão de PCH-001 permanece independente |
| D-006 | Controle de continuidade, README e índice OPEN-001 | concluido | ● Revisão pendente nos documentos de Charter | P2 — lacuna de navegação e retomada | Inventário e decisões registradas em develop | Retomadas incrementais | Painel único e índices alinhados à baseline | Entrega inicial comprovada; atualizar índice quando surgirem novos documentos |
| D-007 | OBJ-009 / OBJ-005 — execução automática de tarefas elegíveis (opção B) | concluido | ● Revisão pendente nos documentos canônicos (`refinement`) | P1 — orquestração depende de controles ainda não homologados | D-001; D-002; D-003 para execução paga | Fluxo remunerado continua condicionado; sem liberação Git | Decisão B e limites documentados, IDs e links verificados em develop | Entrega documental concluída; manter dependências de homologação e implementação |

Legenda: **● Revisão pendente**, **○ Não verificado — pendente**, **✓ Validado**, **— Cancelado/histórico**. Criticidades P0 crítica, P1 alta, P2 média e P3 baixa. D-005 e D-006 referem-se à **regularização documental desta rodada**, não à promoção de PCH/OPEN para `refined`.

## Decisões confirmadas

| Fonte / demanda | Decisão de negócio e justificativa | Origem verificável | Documento | Autorização do diff | Materialização |
|---|---|---|---|---|---|
| D-002 | Primeira entrega sem LLM local; evitar complexidade | [NOBJ-002](non-objectives/NOBJ-002-sem-inferencia-local-na-primeira-entrega.md) | [NOBJ-002](non-objectives/NOBJ-002-sem-inferencia-local-na-primeira-entrega.md) | Nenhum diff canônico novo | DOCUMENTED |
| D-002 | Inicialmente OpenAI direta e agentes Amazon Bedrock AgentCore; não habilitar sem homologação | [OPEN-005](open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) | [NOBJ-004](non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md) | Nenhum diff canônico novo | DOCUMENTED; homologação UNKNOWN |
| D-001 | Orçamento local, autorização individual e bloqueio remoto comprovado; negar chamadas sem evidência | [OPEN-004](open-questions/OPEN-004-projeto-openai-e-hard-limit.md) | [CON-004](conditions/CON-004-orcamento-obrigatorio.md) | Nenhum diff canônico novo | DOCUMENTED; efetividade UNKNOWN |
| D-003 | Avaliação qualitativa por critério e revisor distinto; sem pontuação numérica não calibrada no MVP | [OPEN-002](open-questions/OPEN-002-fidelidade-e-custo.md) | [OBJ-007](objectives/OBJ-007-comparacao-controlada-de-modelos.md) | Nenhum diff canônico novo | DOCUMENTED |
| D-004 | Automação Git de branches/PR adiada no MVP; alteração normativa externa separada | [OPEN-003](open-questions/OPEN-003-politica-de-branches.md) | [NOBJ-001](non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md) | Nenhum diff canônico novo | DOCUMENTED; função BLOCKED |
| D-005 | Administrador Linux gerencia credenciais; serviço com identidade exclusiva e credenciais restritas via systemd | [OPEN-001](open-questions/OPEN-001-politica-credenciais.md) | [PCH-001](project-context/PCH-001-contexto-integracao-openai-dsl.md) | **Aprovado pelo usuário em 2026-10-22**: regularizar atores | Seções ACT-005/006/007 registradas em `8e187821780fca45ed4af05d78f442447665f8c8` |
| D-003 | Histórico local retido até exclusão manual autorizada, acesso compartilhado somente via aplicação | [OPEN-006](open-questions/OPEN-006-retencao-e-protecao-do-historico.md) | [OBJ-008](objectives/OBJ-008-historico-local-de-inferencias.md) | Nenhum diff canônico novo | DOCUMENTED; comprovação pendente |
| D-006 | Controle de continuidade e índices vinculados para evitar rediscutir decisões | Aprovação específica do usuário em 2026-10-22 (conversa corrente) | [README](README.md), [índice OPEN](open-questions/README.md) | **Aprovado** para quatro arquivos, somente develop | Painel e links registrados em `8e187821780fca45ed4af05d78f442447665f8c8` |
| D-007 | Executar automaticamente tarefas elegíveis sem aprovação geral do plano; chamadas pagas continuam com pré-autorizações individuais por operação | Escolha B expressa pelo responsável em 2026-10-22 (conversa corrente) | [OBJ-009](objectives/OBJ-009-execucao-automatica-de-tarefas-elegiveis.md) e [OBJ-005](objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) | **Aprovado especificamente** em 2026-10-22 para quatro arquivos Charter, somente develop | DOCUMENTED; regularização aplicada no commit `75d0c1d0c9bfe0ff4c53d912a21ef6665d244c2b` |
| D-001 / D-007 | Parâmetros padrão configuráveis: 1.000.000 de tokens por inferência e orçamento acumulado de US$ 10,00 por sessão; contabilização por chamada e por sessão, reserva preventiva, suspensão antes de exceder limites locais e preservação das autorizações individuais e da proteção remota obrigatória | Decisão expressa do responsável em 2026-10-22 (conversa corrente) | [OBJ-005](objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [OBJ-009](objectives/OBJ-009-execucao-automatica-de-tarefas-elegiveis.md), [CON-004](conditions/CON-004-orcamento-obrigatorio.md) | **Aprovado especificamente** em 2026-10-22 para estes três documentos e CONTINUIDADE.md, somente develop | DOCUMENTED; execução e homologação ainda pendentes |
| D-001 — alternativa B | Homologar prioritariamente a proteção financeira dos quatro projetos OpenAI; manter AWS AgentCore prevista, mas bloqueada para operações pagas até homologação independente, sem afrouxar a política de falha fechada | Escolha expressa do responsável em 2026-10-23 (conversa corrente) | [OPEN-004](open-questions/OPEN-004-projeto-openai-e-hard-limit.md), [DSC-001](discovery/openai-api/DSC-001-viabilidade-openai-api.md), [DSC-002](discovery/aws-agentcore/DSC-002-viabilidade-aws-agentcore.md) | Somente o diff específico aprovado poderá ser persistido; não autoriza testes pagos | Ordem de homologação DOCUMENTED; eficácia remota UNKNOWN; D-001 segue aberta |

**Distinção:** aprovação de decisão funcional não aprova automaticamente diffs futuros. Históricos superados (como a exclusão da branch na decisão 33) permanecem preservados nas fontes, sem reativação.

## Progresso e próxima ação

- **Rodada:** 2026-10-22 — definição dos limites configuráveis por inferência e sessão, sem habilitação de despesas.
- **Universo coberto:** 27 documentos canônicos (1 PCH, 9 OBJ, 4 NOBJ, 5 CON, 6 OPEN e 2 DSC), mais apresentação e índices.
- **Demanda atual:** D-001 (P0), comprovação ainda pendente de bloqueio financeiro remoto; limites locais de inferência/sessão definidos e documentados. D-007 regularizada documentalmente, sem execução implementada.
- **Concluídos por entrega documental já verificada:** D-005, D-006 e D-007. A regularização de D-007 está comprovada no commit `75d0c1d0c9bfe0ff4c53d912a21ef6665d244c2b`; status documentais continuam `refinement`.
- **Revisões de negócio pendentes:** D-001, D-002, D-003 e D-004 (quatro IDs únicos); a regularização documental D-007 foi concluída, mas os 27 documentos canônicos permanecem em `refinement` e seus requisitos exigem revisão/homologação posteriores.
- **Aguardando decisão nova:** nenhuma decisão previamente tomada deve ser perguntada novamente sem evidência de reabertura.
- **Aguardando gravação:** nenhuma demanda documental de negócio previamente aprovada permanece sem publicação; os limites de consumo e os estudos DSC foram registrados em `develop`. Manutenções editoriais de índices exigem aprovação específica do respectivo diff. Nenhuma funcionalidade foi implementada.
- **Bloqueados:** D-004 para automação Git; D-001 e D-002 para inferência remunerada/homologação.
- **Reabertos na rodada:** nenhum.
- **Próxima demanda:** D-001 (P0); comprovar primeiro a proteção financeira dos quatro projetos OpenAI, com D-002 (P0) na sequência. A AWS AgentCore segue planejada, porém não habilitada para chamadas pagas sem homologação própria. D-007 continua dependente de D-001, D-002 e D-003 para execução remunerada. Nenhum uso pago decorre desta decisão.

## Histórico de rodadas e reaberturas

| Rodada | Itens | Mudança / resultado | Evidência nova / impacto | Commit conhecido ou publicação |
|---|---|---|---|---|
| 2026-10-22 — inventário | D-001 a D-006 | Consolidação sem criar decisões; atores e índices incompletos identificados | `develop` 84c58183848f; norma 20b49b182ea5 | Sem commit da auditoria |
| 2026-10-22 — regularização | D-005, D-006 | Painel inicial, cadastro ACT-005/006/007 e índice OPEN-001 atualizados | Aprovação expressa limitada a quatro arquivos; estados de negócio preservados | `8e187821780fca45ed4af05d78f442447665f8c8` em develop, verificado |
| 2026-10-22 — execução automática | D-007 | Opção B aceita: tarefas elegíveis são executadas sem aprovação geral do plano; chamadas pagas continuam individualmente autorizadas | Objetivo novo OBJ-009, índice e OBJ-005; sem homologação financeira ou serviço implantado | Commit `75d0c1d0c9bfe0ff4c53d912a21ef6665d244c2b` confirmado em develop |
| 2026-10-22 — limites de consumo da sessão | D-001, D-007 | Padrões de 1.000.000 tokens por inferência e US$ 10,00 acumulados por sessão, configuráveis; preservadas autorizações financeiras por chamada e proteção remota | OBJ-005, OBJ-009, CON-004 e painel; requisito `DOCUMENTED`, homologação `UNKNOWN` | Limites documentados em `develop`; execução e homologação ainda pendentes |
| 2026-10-22 — pesquisa financeira | D-001 | Identificados hard spend limits OpenAI e spend limits AWS Settings | Evidência documental oficial, sem acesso administrativo ou teste de bloqueio | DSC-001/002 publicadas em `a512f36fdbf9f9c6776bf214e8c25a5c3d50d93a`; homologação pendente |
| 2026-10-23 — priorização financeira | D-001 / alternativa B | OpenAI primeiro, quatro projetos por especialização; AWS AgentCore retida no planejamento sem liberação paga | Diretriz de trabalho escolhida pelo responsável; prova de controles externos não apresentada | Proposta documental preparada para aprovação específica sobre baseline `380690539d9afeac81f3a172f73c4683edc7a046`; sem autorização de execução remunerada |

## Identificação, baseline e autorização

- **Repositório-alvo:** https://github.com/ahhvahh/doc-arquitetura-app-debian
- **Branch autorizada:** `develop`; `main` é somente leitura.
- **Referência normativa:** https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0
- **SHA do alvo antes da rodada:** `75d0c1d0c9bfe0ff4c53d912a21ef6665d244c2b`.
- **SHA de main do alvo antes da rodada:** `865ec1efc79c9a7dd2de19f2943922b66fabc785` (somente leitura).
- **SHA normativo:** `20b49b182ea529f8b809e8e009132ab0fd5fe267` (somente leitura).
- **Skill operacional:** Agente Charter 1.2.0; prompts `objective-refinement.txt` e `constraint-assumption-refinement.txt` da normativa main.
- **Autorização de escrita da rodada:** usuário aprovou especificamente em 2026-10-22 a edição de `docs/charter/objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md`, `docs/charter/objectives/OBJ-009-execucao-automatica-de-tarefas-elegiveis.md`, `docs/charter/conditions/CON-004-orcamento-obrigatorio.md` e `docs/charter/CONTINUIDADE.md` para os limites padrão configuráveis; somente `develop`. A regularização anterior de D-007 está comprovada no commit `75d0c1d0c9bfe0ff4c53d912a21ef6665d244c2b`. Não há autorização contínua para novos diffs ou outras branches/etapas.
- **Cobertura/limitações:** inventário de 27 documentos com DSC-001/002 publicadas; homologação remota e validação total da cadeia implementada permanecem pendentes.
- **Persistência:** esta linha descreve a baseline **anterior** ao commit da rodada de limites; conferir commit e branch após a gravação, sem inserir o SHA do próprio commit no arquivo que o contém.

## Gate Charter → Scope

**BLOCKED** para a baseline `84c58183848fcdb3dce0ebdd4963bfa09d9c027c`, conforme auditoria histórica do Charter (sem nova auditoria integral nesta rodada D-007): OPEN-004/005 sem homologação financeira e de provedor; OPEN-002 sem comprovação de fidelidade DSL; OPEN-003 mantém automação Git planejada incompatível com a norma (fora do MVP reduzido); ausência de Discovery conclusiva pertinente; documentos em `refinement`. O gate não é liberado por esta regularização editorial. A avaliação do escopo **reduzido** exige validação própria das dependências necessárias, sem presumir que decisões `DOCUMENTED` comprovem implementação.

## Complemento de continuidade — Discoveries D-001/D-002 (publicadas em 2026-10-22)

- **Baseline de leitura:** `develop` `5ac41c79111f5c2273650d86faa1cc85e35e0335`; norma `ahhvahh/doc-arquitetura/main` `20b49b182ea529f8b809e8e009132ab0fd5fe267`.
- **Documentos presentes no mapa:** [DSC-001](discovery/openai-api/DSC-001-viabilidade-openai-api.md) e [DSC-002](discovery/aws-agentcore/DSC-002-viabilidade-aws-agentcore.md), ambos `refinement`, vinculados a D-001 e D-002, sem homologação de contas.
- **D-001 (P0):** capacidade de bloqueio mensal OpenAI documentada; AWS Settings com spend limit condicionado; comprovação de ativação, abrangência e eficácia pendente. Não liberar operação remunerada.
- **D-002 (P0):** estudo de capacidades dos dois fornecedores preparado; modelos, regiões, conta/projetos, preços efetivos, identidade do revisor e custos remotos ainda dependem de evidência administrativa.
- **Badge de tipo DSC:** já publicado em `docs/charter/assets/documentation-style/badges/badge-type-dsc.svg` (blob `387f0ffd90422c127dd1bf85aa38ec8ed424b57f`), com referências relativas nos cinco Markdown Discovery; a pendência antiga está superada pelo commit `a512f36fdbf9f9c6776bf214e8c25a5c3d50d93a`.
- **Próxima ação:** homologar conta, projeto, custo e proteção remota de cada provedor; manter ambos DSC em `refinement` até evidência suficiente. A manutenção editorial do mapa e do README não amplia autorizações operacionais.
- **Gate:** `GATE_NOT_EVALUATED` nesta rodada pontual; auditoria histórica continua `BLOCKED`.

## Revisão de 2026-10-23 — regularização pontual de D-002

- Pedido: revisão do projeto solicitada explicitamente pelo usuário.
- Baseline: 9367a81153e02115456178ec0311d6a4ee703d64.
- Norma main: 20b49b182ea529f8b809e8e009132ab0fd5fe267.
- Evidência: PCH-001 apresenta a escolha AWS como pendente; OPEN-005, decisão 23, já selecionou AgentCore.
- Escolha A: alinhar a referência de PCH-001 à decisão existente.
- Autorização: seleção expressa da alternativa A apresentada nesta revisão.
- Resultado: correção documental; D-002 permanece aberta para homologação.
- Próxima ação: retomar D-001 para verificar evidências financeiras; nenhuma operação paga foi autorizada.
