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
A lista inicial de provedores, modelos locais/remotos e a escala de qualidade dependem de [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md). Este objetivo não declara executores locais como implementados.

## Relações
- [OBJ-001 — Agentes](OBJ-001-agentes-documentais-especializados.md)
- [OBJ-005 — Orçamentos](OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OBJ-008 — Histórico](OBJ-008-historico-local-de-inferencias.md)
