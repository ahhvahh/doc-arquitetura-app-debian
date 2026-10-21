# OBJ-005 — Inferências com custo previsível e limitado

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
O uso de modelos remotos gera cobrança variável por tokens e ferramentas, em dólares. Sem prévia de consumo e barreiras independentes, uma análise extensa pode exceder o orçamento do operador.

## Resultado esperado
Cada inferência paga apresenta estimativas de entrada, saída e custo em reais, solicita autorização explícita e aplica as políticas de orçamento por agente e globais antes de transmitir o pedido, com controle complementar no serviço OpenAI.

## Critérios verificáveis
- Nenhuma inferência inicia quando faltam modelo, tarifa atual, cotação, orçamento autorizado ou identificação do agente/projeto OpenAI.
- A aplicação apresenta valor estimado e reserva conservadora por inferência, incluindo saída e possível raciocínio, e exige confirmação do operador.
- Limites por operação, agente/dia, agente/mês e globais não podem ser excedidos por reservas simultâneas aprovadas pela aplicação.
- Uma sessão nova e independente recebe limite monetário quando a API disponibiliza esse recurso; ausência de configuração bloqueia uso desse modo.
- A configuração administrativa do hard limit por projeto OpenAI é distinta do limite local e deve ser documentada e verificada antes da primeira operação de produção.
- O registro de consumo posterior identifica diferença entre orçamento reservado e cobrança medida, sem alegar limite absoluto na fatura da OpenAI.
- Nenhuma chave ou segredo de pagamento aparece em logs, DSL, Git ou relatórios de CI.

## Relações
- [CON-004](../conditions/CON-004-orcamento-obrigatorio.md)
- [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md)
- [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md)

## Extensão para múltiplos modelos e provedores (2026-10-21)
O controle financeiro esperado deverá abranger provedores pagos adicionais quando aprovados, preservando os controles específicos já documentados para OpenAI. Cada inferência relaciona previsão e realização de tokens de entrada/saída e de valores, com moeda, tarifa, câmbio e origem dos dados.

Cobrança efetivamente informada pelo provedor e cálculo local são informações diferentes. Dados não disponibilizados permanecem desconhecidos, sem assumir custo zero. As proteções remotas variam entre serviços; não se presume que todo provedor suporte os mesmos limites da OpenAI.

Relações: [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md), [OBJ-008](OBJ-008-historico-local-de-inferencias.md) e [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).
