# OBJ-008 — Histórico local e auditável de inferências

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
Sem persistir solicitações, respostas e métricas torna-se difícil revisitar experimentos, comparar modelos e compreender diferenças entre previsão e cobrança.

## Resultado esperado
Manter um histórico em banco de dados local para análise futura de solicitações e tentativas de inferência, com prompts e resultados completos e medições de consumo e valores previstos e realizados, quando conhecidos.

## Critérios verificáveis
- Cada solicitação e tentativa tem ID, horário, duração, estado, projeto-alvo, etapa, prompt selecionado, provedor, modelo, SHA normativo e baseline documental identificados.
- O registro contém o prompt efetivamente transmitido e a resposta integral recebida, preservando o vínculo ao contexto de origem e a sinalização de trechos omitidos por segurança.
- Tokens previstos de entrada e saída e tokens realizados de entrada e saída são independentes; categorias extras de uso podem ser registradas com sua origem.
- Custos previstos e realizados são campos distintos; tarifa, moeda, cotação, data e origem diferenciam valores confirmados por provedor dos calculados localmente.
- Medições não fornecidas pelo provedor são desconhecidas, não zero; falhas, cancelamentos e respostas incompletas são rastreáveis.
- O histórico permite consultas e comparações por tarefa, modelo, provedor, especialização, qualidade, período, duração, tokens e custos.

## Limites
Este objetivo não escolhe banco, esquema ou política de exclusão. A retenção e o acesso dependem de [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md).

## Relações
- [OBJ-007 — Comparação](OBJ-007-comparacao-controlada-de-modelos.md)
- [OBJ-005 — Orçamentos](OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [CON-005 — Persistência local](../conditions/CON-005-persistencia-local-de-inferencias.md)
- [CON-002 — Credenciais](../conditions/CON-002-protecao-credenciais.md)
