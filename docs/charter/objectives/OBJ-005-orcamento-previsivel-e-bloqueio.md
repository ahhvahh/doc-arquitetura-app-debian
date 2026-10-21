# OBJ-005 — Inferências com custo previsível e limitado

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
O uso de modelos remotos gera cobrança variável por tokens e ferramentas, em dólares. Sem prévia de consumo e barreiras independentes, uma análise extensa pode exceder o orçamento do operador.

## Resultado esperado
Cada inferência paga apresenta estimativas de entrada, saída e custo em reais, solicita autorização financeira explícita e respeita os limites locais por especialização e globais **e a proteção remota comprovadamente bloqueante do provedor/projeto aplicável**. Se qualquer camada obrigatória não estiver verificada, a inferência paga não é iniciada.

## Critérios verificáveis
- Nenhuma inferência paga inicia quando faltam modelo, tarifa atual, cotação, orçamento autorizado, identificação do provedor/projeto aplicável ou evidência válida de proteção remota suficiente e efetiva.
- A aplicação apresenta valor estimado e reserva conservadora por inferência, incluindo saída e possível raciocínio, e exige confirmação do operador.
- Limites por operação, agente/dia, agente/mês e globais não podem ser excedidos por reservas simultâneas aprovadas pela aplicação.
- Uma sessão nova e independente recebe limite monetário quando a API disponibiliza esse recurso; ausência de configuração bloqueia uso desse modo.
- Para a integração OpenAI, manter quatro projetos separados — Charter, Scope, Implementation e Source — e exigir verificação individual da eficácia e abrangência do mecanismo remoto de bloqueio financeiro de cada projeto habilitado; a configuração administrativa de orçamento ou alerta não demonstra, por si só, um hard limit.
- O registro de consumo posterior identifica diferenças entre orçamento reservado e cobrança medida, sem alegar garantia absoluta sobre faturas externas; a ausência ou perda da proteção remota comprovada torna novas inferências pagas indisponíveis.
- Nenhuma chave ou segredo de pagamento aparece em logs, DSL, Git ou relatórios de CI.

## Relações
- [CON-004](../conditions/CON-004-orcamento-obrigatorio.md)
- [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md)
- [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md)

## Extensão para múltiplos modelos e provedores (2026-10-21)
O controle financeiro esperado deverá abranger provedores pagos adicionais quando aprovados, preservando os controles específicos já documentados para OpenAI. Cada inferência relaciona previsão e realização de tokens de entrada/saída e de valores, com moeda, tarifa, câmbio e origem dos dados.

Cobrança efetivamente informada pelo provedor e cálculo local são informações diferentes. Dados não disponibilizados permanecem desconhecidos, sem assumir custo zero. As proteções remotas variam entre serviços; não se presume que todo provedor suporte os mesmos limites da OpenAI.

Relações: [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md), [OBJ-008](OBJ-008-historico-local-de-inferencias.md) e [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).

## Regra transversal — provedores remotos (decisão 06)
A política de orçamentos locais aplica-se a todos os provedores remotos admitidos, sem confundir a arquitetura própria de projetos OpenAI com recursos inexistentes em outros serviços. Antes de cada inferência paga, conferir tarifa e estimativas e exigir autorização financeira conforme a [CON-004](../conditions/CON-004-orcamento-obrigatorio.md). A aprovação futura de merge não substitui a autorização de despesa.

A comparação deverá conservar indicadores nativos do provedor e métricas reconciliáveis no histórico local; valores indisponíveis continuam desconhecidos.

## Política aprovada — bloquear chamadas pagas sem proteção remota (decisão 09)
A decisão do responsável é **não utilizar provedores para inferência paga sem proteção remota suficiente e verificada**. Aplica-se a **OpenAI, Anthropic e Google Gemini**, além de quaisquer provedores futuramente aprovados. Estimativa, teto local, alerta remoto ou aceite do usuário, isoladamente, não liberam operação remunerada.

**Critérios esperados de verificação:**
- Demonstrar para cada provedor e identidade/projeto de faturamento utilizado que existe mecanismo remoto efetivo e bloqueante, com alcance compatível com o orçamento aprovado; **desconhecido = bloqueado**.
- Comprovar também preço, câmbio, saldo/reserva local, autorização por inferência e identidade de cobrança segregada quando exigida.
- Suspender novas chamadas pagas quando a evidência do controle expirar, se tornar incompatível ou não puder ser verificada.
- Diferenciar testes comprovadamente gratuitos de inferências sujeitas à cobrança; não presumir gratuidade.

**A organização de quatro projetos OpenAI independentes da decisão 04 permanece vigente**. Somente as garantias que forem efetivamente verificadas poderão ser afirmadas. Relações: [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).

## Previsão de custo da segunda avaliação (decisão 10)
O orçamento autorizado de uma revisão documental deverá considerar separadamente a inferência geradora e a **segunda avaliação por modelo distinto**, incluindo estimativa e medição de custos/tokens para ambas as chamadas quando aplicáveis.

A revisão secundária **não é isenta** da política de bloqueio da [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md): sem proteção remota suficiente comprovada, limite local ou autorização exigida, negar a inferência paga. A indisponibilidade financeira do segundo avaliador impede declarar concluído o fluxo de validação dupla; não justifica publicar sem revisão.

## Atribuição financeira ao usuário autorizado — decisão 11
Quando houver tarefas organizadas por serviço Linux, a exigência de autorização financeira permanece vinculada ao **usuário responsável pela inferência** e à identidade de provedor/projeto e especialização aplicáveis. O organizador não pode usar permissão de outro usuário ou credencial de serviço para contornar bloqueios.

O uso de quatro projetos OpenAI segregados por especialização permanece previsto; a definição de identidade do usuário **não substitui** a verificação de credenciais, contexto de consumo, limite local e proteção remota por projeto de [CON-004](../conditions/CON-004-orcamento-obrigatorio.md). Sem comprovação de identidade ou autorização, a chamada paga deve permanecer bloqueada.

Relações: [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md) e [CON-002](../conditions/CON-002-protecao-credenciais.md). Nenhum mecanismo de execução ou faturamento é homologado por esta decisão.
