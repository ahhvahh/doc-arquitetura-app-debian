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

## Proposta de regularização — Decisão 05 (2026-10-21)

**Origem:** aprovação explícita do responsável — opção A, preservar branches de tarefa e propor atualização normativa separada.
**Classificação:** PROPOSED. Esta decisão registra o objetivo de governança do aplicativo, mas não modifica nem revoga as regras de escrita atuais dos agentes.

- Manter a possibilidade de criar, para cada inferência aprovada, uma branch de tarefa específica derivada da `develop` do repositório-alvo, vinculada a um SHA-base verificável e a um identificador de tarefa.
- A futura política deverá limitar os prefixos e os nomes de branches permitidos, impedir referências arbitrárias, registrar a origem de cada alteração e verificar concorrência antes de publicar.
- Separar as autorizações para **aplicar alterações**, **criar commit** e **publicar remotamente**. Aprovação de análise ou prévia não autoriza implicitamente qualquer uma dessas etapas.
- Não autorizar gravação direta em `main`, merge, force-push ou elevação de permissão. O processo de revisão e eventual integração permanecerá sujeito a autorização independente.
- Preparar uma proposta de atualização do documento normativo de governança em `ahhvahh/doc-arquitetura` pelo processo próprio desse repositório, sem alterações normativas por este Charter e sem presumir que a mudança será aprovada.
- **Até a publicação autorizada da mudança normativa, não criar nem publicar branches de tarefa** pela automação documental; permitir somente a preparação local, a simulação, o diff e a revisão sem gravação em refs proibidas.
- A aplicação dessa futura política dependerá de validações em Scope/Implementation após os respectivos gates.

**Pendência impeditiva:** a norma `ProjectCharter/1.0`, na `main` de `ahhvahh/doc-arquitetura`, mantém a restrição de escrita exclusiva em `develop`. A decisão de produto não é exceção à norma. Manter a `OPEN-003` em `refinement` e seu bloqueio para publicação de branches de tarefa até alinhamento normativo verificável.
