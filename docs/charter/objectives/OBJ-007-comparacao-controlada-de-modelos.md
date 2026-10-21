# OBJ-007 — Comparação controlada de modelos de linguagem

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
A eficiência de LLMs simples e avançados para refinar documentação deve ser comparada com evidências de qualidade, consumo e custos, e não apenas pelo tamanho ou reputação do modelo.

## Resultado esperado
Permitir que uma mesma tarefa de Charter, Scope, Implementation ou Source seja avaliada por modelos diferentes, incluindo opções econômicas e avançadas, mantendo a mesma baseline documental para comparação de qualidade, tempo, tokens e valores.

## Critérios verificáveis
- Cada comparação associa tarefa, projeto, especialização, prompt efetivamente enviado, hashes dos documentos e SHAs normativo e do alvo.
- Cada tentativa identifica provedor, modelo/versão e condições relevantes; variações de entrada são declaradas.
- O operador pode comparar pelo menos dois modelos usando a mesma tarefa e contexto, sem publicação automática do resultado.
- Comparativos mostram validade documental, preservação do conteúdo, aderência à norma, necessidade de correção humana, duração, tokens previstos/realizados e custos previstos/realizados quando disponíveis.
- Ausência de metadados ou falha de uma tentativa é explicitada, jamais substituída por dados inventados.
- A aceitação final das alterações cabe ao operador, independentemente do modelo mais barato.

## Limites
A seleção de **OpenAI, Anthropic e Google Gemini** como provedores previstos foi registrada em [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md). Sua viabilidade e os modelos específicos permanecem pendentes de verificação. A escala de qualidade depende de [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md). Modelos executados localmente não fazem parte da primeira entrega.

## Relações
- [OBJ-001 — Agentes](OBJ-001-agentes-documentais-especializados.md)
- [OBJ-005 — Orçamentos](OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OBJ-008 — Histórico](OBJ-008-historico-local-de-inferencias.md)

## Aplicação da decisão 06 — comparação entre provedores
A primeira entrega pretendida permitirá executar ao menos uma mesma tarefa em dois provedores remotos distintos, com condições rastreáveis e resultados apresentados sem pressupor equivalência técnica de modelos, tokenizadores, modalidades de cobrança ou limites de gasto.

A aceitação de um resultado documental e a autorização de seu merge são decisões humanas distintas da inferência e dos commits intermediários condicionados à [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md). A lista nominal pretendida está em [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e ainda depende de homologação independente de cada serviço.

## Comparabilidade prevista dos três provedores (decisão 08)
Para a primeira entrega, o objetivo é permitir a comparação controlada dos modelos selecionados de **OpenAI, Anthropic e Google Gemini**, desde que cada serviço seja previamente considerado apto. O operador deverá poder identificar o provedor de cada resultado e compará-lo com a mesma tarefa, referência normativa e baseline documental.

**Verificações esperadas:** registrar pelo menos uma tentativa rastreável por provedor habilitado, identificar indisponibilidades e diferenças de contagem de tokens, custo e duração; não tratar métricas ausentes como zero. Não declarar os três provedores disponíveis antes da validação individual.
