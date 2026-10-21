# OPEN-006 — Retenção e proteção do histórico de inferências

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão pendente de governança e segurança.

## Questão
Qual é o prazo de retenção de prompts integrais, respostas, avaliações, custos e medições? Quem pode consultar, exportar ou excluir esses registros? Como tratar dados privados dos repositórios e operadores diferentes?

## Contexto e origem
O histórico local solicitado inclui conteúdo enviado e recebido que pode carregar informações confidenciais. A decisão de manter um banco local não resolve controle de acesso, descarte ou tratamento de segredos.

## Impacto
**Bloqueadora:** sim para armazenamento de conteúdo integral em produção, não para modelos de dados fictícios e estudos.

## Próxima ação verificável
Aprovar políticas de retenção, descarte, acesso por operador, proteção de conteúdo e eventual exportação, incluindo evidência de eliminação. Contratos técnicos ficam para Scope/Implementation.

## Relações
- [OBJ-008 — Histórico](../objectives/OBJ-008-historico-local-de-inferencias.md)
- [CON-005 — Persistência](../conditions/CON-005-persistencia-local-de-inferencias.md)
- [CON-002 — Credenciais](../conditions/CON-002-protecao-credenciais.md)
- [OPEN-001 — Segurança](OPEN-001-politica-credenciais.md)
