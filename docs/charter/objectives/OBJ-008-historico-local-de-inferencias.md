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

## Requisito de interoperabilidade — métricas JSON (2026-10-21)
As métricas de cada solicitação/tentativa deverão ser **persistidas ou disponibilizadas como objetos JSON estruturados e recuperáveis**, permitindo que aplicações consumidoras consultem os dados e reconstruam relatórios em outro ambiente, sem depender da interface de análise original.

**Critérios verificáveis:**
- Um registro recuperado em JSON identifica sua solicitação, sessão, tentativa, provedor/modelo, timestamps, resultado, tokens e custos conhecidos, com referências de origem e contexto suficientes para comparação.
- Métricas nativas de cada provedor e indicadores comparáveis preservam sua proveniência, sem perda silenciosa de categorias ou alteração de unidades.
- Ausência de medição, estimativa e cobrança confirmada são diferenciadas; não se inventa zero em campos sem informação.
- Uma aplicação independente consegue recuperar as métricas e reconstruir uma visão histórica por sessão, inferência ou provedor sem repetir as chamadas externas.

O Charter não fixa nomes de propriedades, modelo físico de banco ou transporte de consulta; essas definições pertencem a Scope/Implementation e deverão respeitar [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md) e [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md).

## Decisão aprovada — retenção até exclusão manual (2026-10-21)
- Manter prompts, respostas, avaliações e métricas JSON recuperáveis até exclusão manual do operador autorizado, sem prazo automático de vencimento.
- Consultas e reconstrução externa de relatórios são permitidas somente ao operador e a consumidores expressamente autorizados; consulta não autoriza remoção.
- Omissões ou mascaramentos necessários para impedir vazamento de segredos são explicitados, sem registrar os próprios segredos.
- O descarte manual deverá registrar de modo verificável quem solicitou, quando, o que foi afetado e o resultado, sem manter o conteúdo descartado no evento de auditoria.
- Informar o alcance real da exclusão no histórico, em cópias e em exportações; não afirmar remoção de cópias fora do controle da aplicação.

**Critérios de verificação propostos:** não ocorrer expiração automática; registros continuam recuperáveis até descarte manual; exclusões não autorizadas são negadas; consultas indevidas e vazamentos de segredos são impedidos; o descarte gera evidência suficiente de sua execução.

Decisão: [OPEN-006 — Retenção e acesso](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md). Formato, armazenamento, proteção e consultas técnicas serão definidos em Scope/Implementation.
