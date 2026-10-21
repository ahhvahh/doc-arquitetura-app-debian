# OPEN-006 — Retenção e proteção do histórico de inferências

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão funcional de retenção e proteção resolvida; execução dos controles ainda sujeita a validação.

## Questão
**Decidida:** manter prompts, respostas, avaliações, custos, medições e métricas JSON até exclusão manual. Somente o operador e consumidores expressamente autorizados podem consultar o histórico, observadas as medidas de proteção de segredos.

## Contexto e origem
O histórico local solicitado inclui conteúdo enviado e recebido que pode carregar informações confidenciais. A decisão de manter um banco local não resolve controle de acesso, descarte ou tratamento de segredos.

## Impacto
**Bloqueadora para decisão de Charter:** não; o responsável definiu retenção, acesso e descarte. Para produção, a implementação e validação efetivas de proteção, identidade e autorização continuam pendentes, inclusive [OPEN-001](OPEN-001-politica-credenciais.md).

## Próxima ação verificável
Verificar em Scope/Implementation os controles de acesso, a proteção de conteúdo, a extensão do descarte a backups/cópias, as autorizações de exportação e a auditoria da exclusão. Não declarar eliminação de cópias que a aplicação não controla.

## Relações
- [OBJ-008 — Histórico](../objectives/OBJ-008-historico-local-de-inferencias.md)
- [CON-005 — Persistência](../conditions/CON-005-persistencia-local-de-inferencias.md)
- [CON-002 — Credenciais](../conditions/CON-002-protecao-credenciais.md)
- [OPEN-001 — Segurança](OPEN-001-politica-credenciais.md)

## Resultado da análise — decisão C (2026-10-21)
**Triagem: RESOLVIDA.** O responsável selecionou a opção C e autorizou o registro documental nesta `develop`.

- **Retenção:** conservar prompts, respostas, avaliações e métricas JSON locais **sem expiração automática**, até solicitação manual de exclusão pelo operador autorizado.
- **Consulta:** permitir acesso ao operador e às aplicações consumidoras expressamente autorizadas; a autorização para consultar não implica permissão para excluir.
- **Proteção:** credenciais e segredos não devem integrar registros, exportações, logs nem trilhas de descarte. Trechos omitidos ou mascarados por segurança devem ser identificados sem expor seus valores.
- **Exclusão e rastreabilidade:** exigir comando ou pedido manual autorizado; conservar somente evidência mínima de quem solicitou, data, identificação dos dados afetados, alcance e resultado, sem manter o conteúdo descartado na auditoria.
- **Cópias:** indicar o alcance real da exclusão em backups e exportações; não presumir descarte de cópias sob controle externo.
- **Interoperabilidade:** as métricas JSON obedecem à mesma regra de retenção, acesso e exclusão manual dos registros de inferência.

**Destinação canônica:** [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md) e [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md). Mecanismos técnicos para Scope e Implementation.

**Situação do documento:** decisão de produto concluída, mas o badge permanece em `refinement`, pois falta o asset `badge-status-refined.svg` em `docs/assets/` e sua criação está fora do escopo desta alteração de Charter. O encerramento da análise não é comprovação de implementação; pendências de segurança continuam em [OPEN-001](OPEN-001-politica-credenciais.md).
