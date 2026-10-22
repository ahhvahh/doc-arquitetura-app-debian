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

## Disponibilidade dos modelos de ambos os papéis como pré-condição — decisão 27
O fluxo documental **destinado a validação** só poderá começar depois de verificar os **dois LLMs distintos, homologados, identificáveis e disponíveis** que exercerão os papéis de gerador e revisor, inclusive a identidade efetiva do modelo associado a Amazon Bedrock AgentCore. Dois agentes não constituem necessariamente dois modelos.

A aplicação deve apresentar o **pareamento efetivo**, indicar a **versão dos padrões iniciais de sessão** e as escolhas explícitas naquela inferência e registrar bloqueio fundamentado quando faltar uma das condições. **Não** permitir geração apenas por indisponibilidade de revisor, fallback automático para outro modelo, aceite sem parecer independente ou commit de resultado validado. Disponibilidade confirmada previamente **não garante** execução futura: falha posterior suspende o ciclo e impede validação, conservando evidências e gastos autorizados.

Mantêm-se comparação por baseline íntegra, autorização financeira separada para chamadas de geração e revisão, homologação individual, avaliação das divergências materiais, revisão humana e limites de quatro correções por tarefa. Relação: [OPEN-002 — decisão 27](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Revisão independente retomada após indisponibilidade — decisão 28
A revisão interrompida **depois da geração** pode ser retomada **sobre a mesma proposta íntegra, mesma tarefa e mesma baseline**, desde que um **revisor efetivamente distinto do gerador, homologado e disponível** seja escolhido ou confirmado **expressamente pelo operador** e que nova chamada remunerada seja autorizada com todos os controles de custo, segurança e contexto. A escolha pode diferir do padrão original da sessão, sem alterar os padrões iniciais congelados nem apagar a geração e o histórico anterior.

A comparação deverá distinguir **modelo gerador original** e **modelo revisor efetivo da retomada**, registrar suas versões, baselines, validade de integridade, evidências da falha e novos custos/pareceres. **Não comparar como equivalentes** resultados de contexto/baseline modificados sem explicitar a diferença; se essa diferença impedir uma avaliação independente íntegra da proposta original, bloquear a retomada.

**Não repetir a geração** apenas porque o revisor falhou, nem considerar a revisão pendente como correção documental automática. O teto de quatro correções continua atrelado à mesma tarefa, incluindo correções eventualmente executadas **antes** da interrupção. Sem revisor elegível ou proteção financeira suficiente, nenhum aceite/commit de resultado validado pode ocorrer. Referências: [OPEN-002 — decisão 28](../open-questions/OPEN-002-fidelidade-e-custo.md) e [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).

## Avaliação de qualidade com barreiras de validade — política aprovada (2026-10-22)

**Resultado esperado adicional:** cada comparação deverá separar a **elegibilidade documental obrigatória** da **avaliação comparativa de qualidade**, conforme a [política da OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md). Os resultados devem permitir compreender a qualidade sem que uma pontuação compense falhas normativas, estruturais, de fidelidade, contexto, segurança ou da avaliação independente.

**Critérios verificáveis:**
- Apresentar, por tentativa e modelo efetivamente utilizado, o estado de cada verificação eliminatória, evidências e razões de bloqueio. Com divergência material pendente, falha de integridade, ausência de revisor distinto ou verificação inconclusiva, informar **não elegível**, sem nota capaz de liberar aceite.
- Quando elegível, apresentar avaliações individuais de **completude**, **clareza**, **consistência**, **rastreabilidade** e **necessidade/importância das correções**, com justificativas e proveniência; sugestões somente editoriais não geram correções obrigatórias por si.
- Exibir **pontuação consolidada apenas se houver método quantitativo calibrado e homologado**, identificando versão, critérios e condições de aplicação. Pesos, fórmula, escala e nota mínima ainda não estão definidos; antes da homologação, mostrar **pontuação indisponível/não calibrada**, sem inventar valores ou classificar resultados sem base comparável.
- Garantir que tarefa, prompt efetivo, contexto, SHA normativo e baseline documental sejam equivalentes na comparação; variações de modelo, escopo ou fonte devem ser sinalizadas e impedem atribuir equivalência indevida.
- Mostrar **tempo, tokens, custos e correções como indicadores separados da nota de qualidade**; ausência de medição significa desconhecido, não zero. Não tratar menor custo como maior qualidade nem a nota como homologação automática.
- Manter revisão humana e a autorização específica para publicação Git independentes de qualquer pontuação. O controle prévio de dois modelos distintos e o limite de quatro correções por tarefa permanecem exigíveis.

**Situação:** critérios de produto `DOCUMENTED`, não `IMPLEMENTED`; escalas, pesos, fórmula, testes de calibração e homologação ficam pendentes, sem escolhas técnicas neste Charter. Relações: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md) e [OBJ-008](OBJ-008-historico-local-de-inferencias.md).

## Resultados da primeira entrega — comparação qualitativa e resumo auditável (2026-10-22)

**Origem:** escolhas de qualidade **1A — avaliação por critério**, **2A — adiar pontuação numérica** e **3A — resumo de divergências e evidências para revisão humana**. Esta delimitação da primeira entrega **não revoga** a política de critérios eliminatórios nem elimina a pontuação quantitativa como possibilidade futura após homologação.

**Resultados esperados verificáveis:**
- Antes de avaliação positiva, exibir o estado e a razão de cada **verificação eliminatória**: integridade e fidelidade documental/DSL, aderência normativa, escopo, baseline e contexto, segurança e **parecer independente por segundo LLM efetivamente distinto, disponível e homologado**. Falha ou evidência inconclusiva impede declarar o resultado elegível e aprovado.
- Para resultados elegíveis, exibir individualmente **completude, clareza, consistência, rastreabilidade e necessidade/importância de correções**, com justificativa, origem, evidência e indicação de lacuna ou incerteza; sem síntese numérica de qualidade.
- No escopo inicial, indicar a nota quantitativa como **indisponível/não calibrada**; não calcular nem mostrar pesos, escala ou nota geral não homologados. A evolução futura exige método documentado, calibrado, versionado e validado, e não substitui julgamento humano.
- Na revisão humana, oferecer **resumo das falhas, divergências materiais, sugestões editoriais e evidências correspondentes**, com acesso pelo serviço ao **histórico integral de pareceres e verificações**. Divergência material não resolvida continua bloqueadora; sugestão editorial isolada permanece opcional e não consome ciclo por obrigação.
- Preservar até **quatro correções por tarefa**; custos, tokens, duração e quantidade de correções são indicadores separados, não notas substitutas de qualidade. Comparações somente entre tarefas, prompts, contextos e baselines efetivamente equivalentes ou com diferenças declaradas sem ranking indevido.
- A revisão humana, a pré-autorização financeira **por chamada paga**, a prova de proteção remota bloqueante, a autorização Git independente e a restrição normativa de escrita permanecem inalteradas, mesmo após avaliação qualitativa favorável.

**Critérios negativos:** nunca transformar falha eliminatória em nota compensável, avaliação inconclusiva em aceite, falta de métrica em zero, dois agentes em prova de dois modelos distintos, relatório resumido em perda do histórico ou parecer de revisão em publicação automática.

**Pendências de comprovação:** evidências de comparações reais, verificações negativas, confiabilidade das justificativas, acesso ao histórico e fidelidade documental; pontuação numérica permanece fora da primeira entrega e sem calibração. Intenção `DOCUMENTED`, não `IMPLEMENTED`; status `refinement` / `v1`. Relações: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md) e [OBJ-008](OBJ-008-historico-local-de-inferencias.md).
