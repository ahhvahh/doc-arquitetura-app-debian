# OPEN-003 — Política de branches para propostas

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Conflito de requisitos.
## Questão
O responsável substituiu a proposta de **branch por inferência** por **uma branch por sessão**, com commits automáticos após validação e aprovação humana apenas para a integração posterior à `develop`. A norma externa atual ainda restringe gravações de agentes exclusivamente a `develop`. Como compatibilizar o fluxo de sessão e a revisão humana com essa norma sem liberar refs arbitrárias?
## Contexto
O aplicativo já descreve branches de tarefa em `src/docs/TASK-BRANCH-LIFECYCLE.md`, mas como proposta ainda não executada. A norma `ProjectCharter/1.0/AGENT-PROMPT.md` e os agentes documentais ainda determinam escrita somente em `develop`. A decisão funcional está tomada; falta compatibilizar a política normativa e suas verificações.
## Impacto
**Bloqueadora:** sim para publicação em branch de trabalho; a preparação de diff em ambiente local não depende de publicação.
## Próxima ação
Propor alteração explícita no documento normativo de governança por seu processo próprio, incluindo prefixos permitidos, lease da base, revisão, aprovações e proibição de `main`. Até lá, não publicar branches de tarefa; testes locais podem simular todo o ciclo.
## Destino previsto
Revisão normativa e, depois, Scope do fluxo Git.

## Histórico — Decisão 05 (2026-10-21; atualizada pela decisão 07)

**Origem:** aprovação explícita do responsável — opção A, preservar branches de tarefa e propor atualização normativa separada.
**Classificação:** PROPOSED. Esta decisão registra o objetivo de governança do aplicativo, mas não modifica nem revoga as regras de escrita atuais dos agentes.

- A proposta inicial previa branch para cada inferência. **Substituída:** a política desejada passa a utilizar uma branch por sessão, com continuidade entre inferências, conforme a decisão 07 abaixo.
- A futura política deverá limitar os prefixos e os nomes de branches permitidos, impedir referências arbitrárias, registrar a origem de cada alteração e verificar concorrência antes de publicar.
- O modelo inicial previa aprovações separadas de **aplicar alterações**, **criar commit** e **publicar remotamente**. **Substituído:** na política de produto atual, o commit validado é automático e a aprovação humana Git fica reservada à integração posterior à `develop`.
- Não autorizar gravação direta em `main`, merge, force-push ou elevação de permissão. O processo de revisão e eventual integração permanecerá sujeito a autorização independente.
- Preparar uma proposta de atualização do documento normativo de governança em `ahhvahh/doc-arquitetura` pelo processo próprio desse repositório, sem alterações normativas por este Charter e sem presumir que a mudança será aprovada.
- **Até a publicação autorizada da mudança normativa, não criar nem publicar branches de tarefa** pela automação documental; permitir somente a preparação local, a simulação, o diff e a revisão sem gravação em refs proibidas.
- A aplicação dessa futura política dependerá de validações em Scope/Implementation após os respectivos gates.

**Pendência impeditiva:** a norma `ProjectCharter/1.0`, na `main` de `ahhvahh/doc-arquitetura`, mantém a restrição de escrita exclusiva em `develop`. A decisão de produto não é exceção à norma. Manter a `OPEN-003` em `refinement` e seu bloqueio para publicação de branches de tarefa até alinhamento normativo verificável.

## Decisão 07 — Sessão contínua e integração humana (2026-10-21)
**Origem:** decisão explícita do responsável, aprovada para registro em Charter.

**Resultado da triagem:** decisão funcional aceita; execução remota continua **bloqueada** por incompatibilidade com a referência normativa de governança.

- Cada sessão de análise mantém uma branch de trabalho identificável, inicialmente derivada da `develop`; um novo prompt pode ser escolhido do catálogo ou apresentado pelo operador.
- Enquanto a sessão estiver ativa, novas inferências continuam na mesma branch, incluindo os commits já produzidos, sem gerar uma branch por inferência.
- Ao concluir e validar cada trabalho, a aplicação deverá criar automaticamente um commit rastreável nessa branch de sessão. Resultado inválido ou fora do escopo não gera commit.
- O operador não precisa aprovar cada commit. Somente a integração das alterações acumuladas à `develop` exige autorização humana explícita, após revisão do resultado e das diferenças.
- Nenhum modelo de linguagem controla comandos Git diretamente. A autorização para gastos de inferência permanece independente.
- A `main` do repositório e a referência normativa são sempre protegidas; a funcionalidade não autoriza merges executados por este agente.
- **Até haver alteração normativa formal autorizando essa modalidade, não criar, publicar nem atualizar branches de sessão por agentes documentais**. Somente simulações locais sem gravação de refs são admissíveis.

**Conflito ainda bloqueador:** [norma main](https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0) exige gravação exclusiva na branch `develop`. A política do aplicativo não pode revogá-la. A ação necessária é uma revisão normativa externa, por processo independente e autorizado, seguida de validação dos contratos em Scope/Implementation.

Relações: [OBJ-003](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md), [NOBJ-001](../non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md).
