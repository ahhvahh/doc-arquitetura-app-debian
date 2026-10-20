# OBJ-002 — Contexto compacto e rastreável

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
Enviar todo o material bruto aumenta a quantidade de texto e dificulta localizar a origem dos requisitos.
## Resultado esperado
A aplicação fornece ao mecanismo de análise uma representação documental reduzida, mantendo rastreabilidade à versão e ao arquivo original.
## Critérios verificáveis
- Cada unidade documental processada aponta seu arquivo de origem e a versão verificada.
- O processo permite identificar quais documentos foram incluídos, omitidos ou descartados como backup.
- O estado local da branch `develop` do projeto-alvo é reconciliado com o remoto antes da extração, sem descartar trabalho local sem consentimento.
- Quando o resumo não for suficiente para revisão segura, o conteúdo original poderá ser consultado; não se perderá informação normativa silenciosamente.
- O custo e o volume real de contexto deverão ser mensuráveis.
## Dependências
- [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).
