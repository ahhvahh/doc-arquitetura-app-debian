# CON-001 — Referência normativa versionada

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
O sistema deve utilizar exclusivamente `ahhvahh/doc-arquitetura`, branch `main`, caminho `ProjectCharter/1.0`, como origem normativa; essa referência será somente leitura.
## Origem verificável
Decisão explícita do responsável e regra descrita no [AGENT-PROMPT normativo](https://github.com/ahhvahh/doc-arquitetura/blob/main/ProjectCharter/1.0/AGENT-PROMPT.md).
## Motivo e impacto
A referência precisa ser reproduzível; sem identificação da versão, análises diferentes não seriam comparáveis.
## Verificação
Uma execução deve registrar o SHA da referência normativa usada, manter uma cópia local verificável sem alterar o repositório de referência e vincular o SHA normativo ao cadastro de cada agente. A norma só será reenviada quando o operador solicitar atualização do agente.
## Aplicação à base consultável
O Vector Store é cópia derivada e pesquisável, não uma nova fonte normativa. Cada publicação identifica SHA da main e integridade do compilado. Sincronização somente por comando explícito, nunca silenciosamente durante inferência.

## Relações
- [OBJ-001](../objectives/OBJ-001-agentes-documentais-especializados.md).
- [OBJ-006](../objectives/OBJ-006-publicacao-manual-base-normativa.md).
## Condição de revisão
Alteração explícita da referência normativa pelo responsável.
