# NOBJ-002 — Não executar LLMs locais na primeira entrega

![Tipo: NOBJ](../../assets/documentation-style/badges/badge-type-nobj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Exclusão
A primeira entrega de comparação entre modelos não terá execução de modelos de linguagem na máquina do operador como modalidade funcional. Serão utilizados diferentes provedores remotos.

## Justificativa
A opção B foi escolhida explicitamente pelo responsável em 2026-10-21, priorizando a comparação entre serviços remotos, modelos de preços e capacidades distintos, sem ampliar inicialmente os requisitos operacionais de inferência local.

## Efeito da exclusão
A aplicação não precisará operar ou distribuir motores de inferência locais para cumprir a primeira entrega. O fato de existirem evidências históricas de integração local no repositório não altera esta decisão. A exclusão não impede que a aplicação processe documentos e mantenha dados locais.

## Responsabilidade remanescente
A execução remota deve observar [OBJ-007 — Comparação](../objectives/OBJ-007-comparacao-controlada-de-modelos.md), [CON-004 — Orçamento](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-005 — Provedores](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).

## Condição de revisão
Reabrir apenas com nova decisão documentada do responsável sobre adoção futura de modelos locais, sem reinterpretar a exclusão como proibição permanente.
