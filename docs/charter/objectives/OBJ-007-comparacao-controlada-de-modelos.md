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
A seleção anterior de **OpenAI, Anthropic e Google Gemini** foi **substituída pela decisão 22**, que prevê somente a **integração direta OpenAI** e a **integração com agentes Amazon Bedrock (AWS)** na primeira entrega, conforme [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md). Sua viabilidade e os modelos específicos permanecem pendentes de verificação; a oferta AWS foi selecionada como **Amazon Bedrock AgentCore** na decisão 23. A escala de qualidade depende de [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md). Modelos executados localmente não fazem parte da primeira entrega.

## Relações
- [OBJ-001 — Agentes](OBJ-001-agentes-documentais-especializados.md)
- [OBJ-005 — Orçamentos](OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OBJ-008 — Histórico](OBJ-008-historico-local-de-inferencias.md)

## Aplicação da decisão 06 — comparação entre provedores
A primeira entrega pretendida permitirá executar ao menos uma mesma tarefa em dois provedores remotos distintos, com condições rastreáveis e resultados apresentados sem pressupor equivalência técnica de modelos, tokenizadores, modalidades de cobrança ou limites de gasto.

A aceitação de um resultado documental e a autorização de seu merge são decisões humanas distintas da inferência e dos commits intermediários condicionados à [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md). A lista nominal pretendida está em [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e ainda depende de homologação independente de cada serviço.

## Comparabilidade prevista dos três provedores (decisão 08)
**Registro histórico substituído pela decisão 22:** a meta atual passou a abranger apenas as integrações OpenAI e agentes AWS.

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

O ciclo está limitado, pela decisão 14, a **quatro correções com reavaliação após a tentativa inicial** e depende também de custo, proteção, disponibilidade do contexto integral e validação estrutural. Se o quarto ciclo não resolver a divergência, ou qualquer requisito obrigatório falhar antes, o processo termina bloqueado, não aceito. Relação: [OPEN-002 — Decisão 13](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Limite de ciclos avaliados — decisão 14
Cada comparação deverá identificar a avaliação inicial e os pareceres das **até quatro correções subsequentes**, sem confundir quatro correções com quatro chamadas de API no total. Uma iteração pode envolver a geração e a revisão secundária, ambas sujeitas à autorização financeira apropriada.

Registrar a conclusão antecipada quando a proposta passar nas verificações ou o bloqueio sem quinta correção quando o máximo se esgotar. A validade semântica continua não homologada. Relação: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Limitação de contexto e separação de baselines — decisão 16
A insuficiência de janela de contexto é motivo de **interrupção documentada**, não de compactação oculta nem de troca automática de modelo/provedor. A escolha aprovada é solicitar ao operador que **redefina expressamente a tarefa e a baseline**, identificando o que deixa de estar coberto pela nova análise.

Resultados produzidos com baselines, conteúdo ou escopos diferentes **não devem ser comparados como equivalentes** sem explicitar a alteração. Registrar a tentativa interrompida, a decisão de redefinição, modelos autorizados, contexto cabível e custos das novas chamadas, se realizadas, respeitando [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md).

Permanecem válidas as quatro correções máximas por tarefa e as barreiras normativas, estruturais, financeiras e de aprovação humana de [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Classificação explicável do parecer secundário — decisão 20
O modelo revisor distinto do gerador deverá apresentar divergências classificadas e **fundamentadas por evidência documental**: «material» quando alterar requisitos, significado, consistência, integridade ou norma; «somente editorial» quando houver apenas melhoria de forma que não modifique conteúdo nem requisitos obrigatórios. Registrar ambas as categorias para comparação entre modelos, incluindo ocorrências omitidas e dúvidas de classificação.

O parecer editorial isolado não obriga outra inferência; permanece visível na revisão humana. Divergência material exige correção com nova avaliação e respeito ao teto de **quatro correções por tarefa** e às autorizações financeiras individuais. Uma classificação sem fundamento suficiente não comprova aceitação. Relação: [OPEN-002 — decisão 20](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Comparação e disponibilidade em entrega parcial — decisão 21
A **decisão 21** estabeleceu a possibilidade de **entrega parcial declarada** com somente as integrações e modelos **individualmente homologados**; sua lista nominal anterior de três provedores foi substituída. O universo atual da primeira entrega inclui **OpenAI e agentes Amazon Bedrock (AWS)**, conforme [OPEN-005 — decisão 22](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md). A falta de uma das duas integrações não deve ser mascarada como implementação ou atendimento completo do objetivo.

A comparação terá de mostrar **quais provedores estão habilitados, quais estão pendentes/desabilitados e quais evidências verificadas** justificam essa situação. Comparar qualidade, tempo, consumo e custo apenas para chamadas efetivamente autorizadas e com contexto suficiente para comparabilidade; dados ausentes **não são zero** nem prova de inferioridade. O não homologado não pode ser utilizado como revisor secundário obrigatório da [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

Não são relaxados os controles de identidade, credenciais e custo de [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md). A validação técnica continua pendente; trata-se de regra de produto, não de provedor já habilitado.

## Comparação das duas integrações planejadas — decisão 22
A primeira entrega prevê comparação controlada das capacidades efetivamente homologadas da **OpenAI direta** e dos **agentes Amazon Bedrock (AWS)**, considerando que **um agente orquestrador não representa necessariamente um único modelo de linguagem**. A aplicação deverá registrar o serviço/agent e o **modelo subjacente efetivamente utilizado** quando esse dado for disponibilizado e verificável.

Somente comparar resultados de tarefas e baselines compatíveis, declarando diferenças funcionais, de contexto, cobertura e métricas; nunca inferir equivalência de medição de tokens, custos ou semântica entre agente AWS e modelo OpenAI. A **segunda avaliação por modelo distinto**, exigida pela [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), depende de modelos efetivamente distintos e elegíveis; não basta alternar dois nomes de serviço. A exclusão de outras integrações diretas na primeira entrega é [NOBJ-004](../non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md).

## Modelo subjacente no Amazon Bedrock AgentCore — decisão 23
A integração AWS da primeira entrega será planejada para **Amazon Bedrock AgentCore**, conforme [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md). A comparação de resultados e a independência da segunda avaliação devem identificar o **modelo efetivamente utilizado**, quando verificável, separando esse modelo da modalidade de execução por agente. Não presumir que AgentCore, isoladamente, constitui um segundo LLM distinto ou que suas métricas são equivalentes às de OpenAI.

A escolha documenta **intenção de produto**, não disponibilidade AWS, qualidade medida, capacidade de auditoria ou implementação. A revisão humana e as restrições financeiras continuam obrigatórias.

## Modelos gerador e revisor predefinidos por operação — decisão 24
A configuração da aplicação antes da abertura de sessão deverá permitir estabelecer **um motor e modelo elegível para geração** e **outro modelo elegível para revisão independente**, associados à especialização documental e à operação. O padrão pode combinar **OpenAI direta** e **Amazon Bedrock AgentCore**, inclusive papéis dentro do mesmo motor, **desde que os modelos efetivamente utilizados sejam distintos e estejam homologados**.

**Critério positivo:** para cada inferência concluída como válida, os registros identificam **especialização, operação, papel, motor/integração, modelo/versão efetivos, origem e versão do prompt, configuração padrão carregada e eventuais alterações escolhidas pelo operador**. A revisão deve usar modelo distinto do gerador e preservar os controles financeiros por chamada, a integridade de contexto e a avaliação de divergências materiais da [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

**Critério negativo:** **dois agentes/serviços diferentes não são prova de dois modelos de linguagem diferentes**. Falta de segundo modelo distinto, homologado ou identificável impede declarar avaliação independente concluída e **bloqueia o aceite/commit de resultado validado**, sem converter avaliação humana em substituto automático para o segundo LLM. A seleção de um segundo motor, modelo ou prompt durante a sessão exige escolha explícita e deve ser auditável.

Os padrões são diretrizes de produto e não comprovam integração ou homologação; formatos de arquivo, interface e contratos serão definidos após o Charter. Relações: [OBJ-001](OBJ-001-agentes-documentais-especializados.md), [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md).

## Rastreabilidade de valores herdados e validados — decisão 25
Para cada papel da análise, o produto deverá identificar se **motor, modelo e prompt** vieram do **padrão da especialização**, de uma **exceção da operação** ou de uma **escolha explícita naquela inferência**. Valores específicos da operação têm prioridade sobre o padrão geral da **mesma especialização e mesmo papel**; escolhas por inferência continuam possíveis e auditáveis.

**Critério positivo:** o parecer de revisão de resultado declarado validado identifica **dois modelos efetivamente distintos e homologados**, com valores resolvidos, proveniência e versão da configuração, além da baseline, da autorização financeira por chamada e da comparação da qualidade da proposta.

**Critério negativo:** ausência de configuração completa, modelo ou motor não homologado, par gerador/revisor equivalente ou incompatibilidade entre valores herdados e específicos **não podem liberar aceite**, gerar commit de resultado validado nem iniciar chamada paga indevida. O sistema deve sinalizar o conflito em vez de recorrer silenciosamente a outro modelo. A alteração manual autorizada para uma inferência precisa preservar o registro de origem e os controles anteriores.

A regra de herança não pressupõe homologação da OpenAI ou do Amazon Bedrock AgentCore. Referências: [OBJ-001](OBJ-001-agentes-documentais-especializados.md), [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md), [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Comparações rastreáveis sob perfil inicial de sessão — decisão 26
A rastreabilidade de geração e revisão deverá distinguir a **versão dos padrões carregados ao abrir a sessão** dos **modelos/motores, prompts e contextos efetivamente usados em cada inferência**, inclusive escolhas explícitas que divergem dos padrões. Uma alteração posterior do perfil geral não muda a baseline de configuração de uma sessão existente; só novas sessões carregam os novos padrões.

A comparação deverá informar a origem dos valores (herança da especialização, exceção da operação e seleção pontual), as versões e o modelo realmente empregado. **Não pressupor** que um modelo configurado anteriormente continua elegível. Um resultado só é considerado validado quando houver **gerador e revisor distintos e homologados**, com evidências de revisão independente, integridade e autorizações aplicáveis.

Esta regra não comprova existência de modelos homologados, nem autoriza chamadas financeiras sem controles ou escrita Git fora de `develop`. Relações: [OBJ-001](OBJ-001-agentes-documentais-especializados.md), [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md), [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).
