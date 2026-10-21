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

## Revisor LLM independente da geração — decisão 10 (2026-10-21)
O fluxo de avaliação de uma alteração documental incluirá o parecer de **modelo distinto daquele responsável pela geração**, seguido de revisão humana. O avaliador deve identificar omissões, possíveis contradições, riscos de perda de conteúdo e aderência à referência normativa, conservando o mesmo contexto rastreável.

Identificar tanto o modelo gerador quanto o avaliador, suas versões, consumo/custo e eventuais divergências. A concordância entre eles não comprova automaticamente fidelidade semântica nem supera falhas detectadas por validação estrutural. A chamada do avaliador também está sujeita a [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md).

A política para divergências relevantes foi escolhida em [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md): corrigir e reavaliar antes do aceite humano. Os parâmetros quantitativos e os critérios para identificar relevância ainda exigem refinamento.

## Reavaliação da proposta corrigida — decisão 12
Pareceres discordantes relevantes obrigam nova revisão da **versão corrigida**, incluindo testes estruturais e parecer independente de LLM diferente do gerador, com origem, versões e custos rastreáveis. Discordância relevante não resolvida impede considerar o resultado apto a aceite humano ou integração Git. A concordância dos modelos não substitui validação normativa ou estrutural.

Relação: [OPEN-002 — Tratamento de divergências](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Avaliação repetida com baseline original preservada — decisão 13
Uma proposta corrigida após divergência relevante deve passar por **nova avaliação crítica de modelo distinto do gerador**. Em cada iteração, a geração recebe o **material de origem integral e invariável** e, como complemento, o resultado e a observação da avaliação **mais recentes**.

Manter rastreáveis todas as versões de propostas, comparações e pareceres ao longo das iterações, incluindo modelo/provedor/versão, hashes, SHAs, tokens, duração e custos. Não confundir reenvio apenas do último feedback com exclusão do histórico persistido nem tratar concordância de modelos como homologação semântica.

O ciclo não tem número máximo fixo, porém depende de condições de custo, proteção, disponibilidade de contexto e validação estrutural; se qualquer requisito for violado, o processo fica bloqueado, não aceito. Relação: [OPEN-002 — Decisão 13](../open-questions/OPEN-002-fidelidade-e-custo.md).
