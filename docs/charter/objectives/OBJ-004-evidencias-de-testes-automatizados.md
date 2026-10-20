# OBJ-004 — Evidências de testes automatizados

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
A integração possui operações com impacto em credenciais, arquivos, branches e custos externos. Resultados de testes precisam ser reproduzíveis e auditáveis.

## Resultado esperado
Cada ação funcional possui testes unitários identificáveis, executados de modo automatizado com evidências consultáveis no GitHub Actions antes da publicação de nova versão.

## Critérios verificáveis
- Cada operação funcional definida em Scope possui casos positivos e negativos com identificadores estáveis.
- Execuções automatizadas disponibilizam resultado individual dos testes e relatório de cobertura.
- Falhas de teste impedem que o fluxo seja declarado concluído.
- Testes unitários usam simulações para OpenAI e Git; não consomem saldo da OpenAI nem alteram repositórios remotos.

## Relações
- [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).
- [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).
