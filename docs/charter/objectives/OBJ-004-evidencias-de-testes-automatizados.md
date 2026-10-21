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

Verificações negativas devem confirmar que falha estrutural, violação normativa ou divergência relevante não resolvida **impedem aceite e publicação de resultado validado**. Falta de orçamento ou de condições para novo parecer é motivo de bloqueio, não uma exceção. Critérios de relevância e de fidelidade continuam pendentes em [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md); a decisão 13 substitui o limite numérico de tentativas por ciclo condicionado aos bloqueios de segurança e custo.

## Verificabilidade do ciclo com contexto integral — decisão 13
As evidências previstas devem permitir constatar que **todas as tentativas de correção** conservam o conjunto original completo e a mesma baseline normativa/documental, modificando apenas o complemento da **última proposta e última observação de divergência**, sem omissão ou troca silenciosa de material de origem.

Cada rodada deve fornecer evidências próprias dos testes estruturais, da nova avaliação secundária, da divergência identificada, da correção apresentada e das condições financeiras autorizadas. Testar bloqueios quando houver erro normativo, ausência de proteção remota comprovada, orçamento insuficiente ou **capacidade de contexto inferior ao material integral**.

Não exigir completar todas as quatro correções quando uma resposta válida for alcançada antes: cada nova versão deve passar nos testes estruturais e na segunda avaliação, seguida da revisão humana. Esgotadas as quatro correções sem validade, a execução termina **bloqueada e não validada**. O procedimento ainda requer homologação técnica posterior. Relação: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Teste do limite de quatro iterações — decisão 14
Os cenários de evidência deverão distinguir **geração inicial** dos **quatro ciclos máximos de correção**, com nova validação estrutural e revisão por LLM em cada ciclo.

Verificar aceitação para encaminhamento humano quando a divergência for sanada na primeira, segunda, terceira ou quarta correção; e verificar bloqueio, sem quinta tentativa automática, quando a quarta não resolver a falha. Preservar também o bloqueio antecipado por orçamento, proteção remota, identidade, segurança, norma ou contexto integral incompatível, com registro auditável.

O critério é intenção de produto vinculada a [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), não comprovação de testes implementados.
