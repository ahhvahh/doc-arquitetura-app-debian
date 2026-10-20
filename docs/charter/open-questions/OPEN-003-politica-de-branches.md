# OPEN-003 — Política de branches para propostas

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Conflito de requisitos.
## Questão
O responsável definiu publicação em uma nova branch gerada para cada inferência e derivada de `develop`. Entretanto, a norma externa atual restringe gravações do agente de documentação exclusivamente a `develop`. Como a norma será atualizada para liberar branches de tarefa sem autorizar edição de `main`, `develop` ou refs arbitrárias?
## Contexto
O aplicativo já descreve branches de tarefa em `src/docs/TASK-BRANCH-LIFECYCLE.md`, mas como proposta ainda não executada. A norma `ProjectCharter/1.0/AGENT-PROMPT.md` e os agentes documentais ainda determinam escrita somente em `develop`. A decisão funcional está tomada; falta compatibilizar a política normativa e suas verificações.
## Impacto
**Bloqueadora:** sim para publicação em branch de trabalho; a preparação de diff em ambiente local não depende de publicação.
## Próxima ação
Propor alteração explícita no documento normativo de governança por seu processo próprio, incluindo prefixos permitidos, lease da base, revisão, aprovações e proibição de `main`. Até lá, não publicar branches de tarefa; testes locais podem simular todo o ciclo.
## Destino previsto
Revisão normativa e, depois, Scope do fluxo Git.
