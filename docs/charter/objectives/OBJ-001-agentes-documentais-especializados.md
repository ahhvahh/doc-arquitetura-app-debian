# OBJ-001 — Agentes documentais especializados

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
O uso de regras indistintas pode misturar responsabilidades das etapas e gerar propostas inconsistentes.
## Resultado esperado
A aplicação oferece quatro especializações independentes — Charter, Scope, Implementation e Source — aderentes à norma versionada, com identificação da especialização e baseline utilizada.
## Critérios verificáveis
- A interface apresenta exatamente as quatro especializações solicitadas.
- Cada análise indica a especialização, a referência normativa persistida no agente e o commit-base.
- A atualização das instruções normativas de cada agente ocorre apenas em comando explícito; a inferência comum não as reenvia.
- Uma especialização não executa silenciosamente a responsabilidade de outra.
## Dependências
- [CON-001](../conditions/CON-001-referencia-normativa-main.md).
- [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).
