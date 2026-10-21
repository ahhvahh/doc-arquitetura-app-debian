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

## Evidências obrigatórias para dupla avaliação — decisão 10 (2026-10-21)
O objetivo de qualidade aprovado prevê resultados de **testes estruturais**, **parecer de segundo modelo distinto do gerador** e **decisão de revisão humana**, todos rastreáveis ao mesmo conjunto de alterações e às baselines utilizadas.

**Critérios verificáveis de produto:** um erro de integridade reprova o resultado mesmo que ambos os modelos o considerem adequado; a ausência de uma revisão secundária exigida é evidenciada e impede declarar a proposta inteiramente validada; pareceres divergentes e intervenção humana devem permanecer auditáveis.

A homologação da semântica do conversor DSL, da cobertura normativa e dos testes de regressão ainda depende de [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md); esses testes não são considerados já existentes pela escolha desta política.

## Evidências de correção e reavaliação — decisão 12
Para cada divergência relevante, registrar: problema apontado, versão original, versão corrigida, **nova execução dos testes estruturais**, **nova avaliação por LLM distinto do gerador** e resultado de liberação ou bloqueio. Não considerar a correção validada com testes ou parecer de versão anterior.

Verificações negativas devem confirmar que falha estrutural, violação normativa ou divergência relevante não resolvida **impedem aceite e publicação de resultado validado**. Falta de orçamento ou de condições para novo parecer é motivo de bloqueio, não uma exceção. Critérios de relevância e limite de tentativas seguem pendentes em [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).
