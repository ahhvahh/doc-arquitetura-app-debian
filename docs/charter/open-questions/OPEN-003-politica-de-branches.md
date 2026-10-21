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

## Decisão 15 — Sessão por repositório e nível documental, PR e validação antes de integração (2026-10-21)
**Origem:** definição expressa do responsável sobre o fluxo funcional do aplicativo `doc-arquitetura-app-debian`: o repositório de documentação trabalhado em cada sessão será escolhido pelo usuário, distinto do repositório da aplicação e da referência normativa fixa. Decisão aceita como **intenção funcional**, sem autorização para contornar a norma existente.

- **Início da sessão:** o operador escolhe o repositório GitHub de trabalho e o **nível documental** a revisar (Charter, Scope, Implementation ou Source). A aplicação identifica a sessão e a baseline correspondente da branch `develop` do **repositório escolhido pelo usuário**, e associa a sessão a esse projeto e nível, com origem rastreável.
- **Branch prevista:** cada sessão terá uma única branch de trabalho denominada **`feature/doc/{session-id}`**, derivada de `develop` do repositório de trabalho selecionado. O identificador da sessão diferencia os trabalhos e a branch não deve ser reutilizada para outra sessão. A identidade e a autoridade da aplicação não substituem as permissões do repositório do operador.
- **Execução da sessão:** novas inferências, inclusive correções e mudanças validadas, continuam na **mesma branch**, com commits rastreáveis da aplicação para cada conjunto de alterações **validado**; conteúdo inválido ou bloqueado não gera commit de resultado válido. Modelos não executam comandos Git nem decidem permissões.
- **Encerramento humano:** ao concluir as iterações, apresentar ao operador **todo o conjunto de alterações acumuladas e evidências da sessão**. A aplicação somente solicita a criação de uma **Pull Request da branch `feature/doc/{session-id}` para `develop`** após o operador aprovar explicitamente o encerramento e o conjunto final de alterações.
- **Validação da PR:** a PR deverá disparar **verificações automatizadas de integridade e aderência documental**, incluindo organização de arquivos e escopo do nível selecionado, referências/links, identificadores e status documentais, norma e versão, alterações não solicitadas e testes de repositório pertinentes. Falhas relevantes tornam a PR **não apta à integração**; logs e resultados devem ser observáveis e vinculados à sessão.
- **Proteção da `develop`:** o aceite de encerramento **não é aprovação automática de merge**. A integração depende de **checagens obrigatórias bem-sucedidas e autorização humana específica**, associadas a controles de proteção de branch efetivamente configurados e verificados. O simples acionamento de GitHub Actions não comprova que a `develop` esteja protegida.
- **Separação de decisões:** encerramento e abertura de PR, validação de conteúdo, aprovação para merge e **autorização financeira de inferências** são eventos distintos; não converter qualquer deles em autorização tácita de gasto ou escrita fora da política.
- **Repositórios distintos:** `ahhvahh/doc-arquitetura-app-debian` contém a aplicação; `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0` é a referência normativa **somente leitura**; o repositório-alvo da **sessão** é indicado pelo operador e não é inferido do endereço da aplicação.

**Impedimento normativo preservado:** a referência normativa `ProjectCharter/1.0` em `main` ainda determina que agentes documentais escrevam **somente em `develop`**. Esta decisão **não autoriza criar, atualizar ou publicar branches `feature/doc/*`, abrir PRs em nome de agentes ou alterar a norma**. Para habilitar o fluxo, exigir revisão normativa em processo próprio, fora deste trabalho, seguida de aprovação e comprovação das políticas de proteção da `develop`. Até lá, admite-se somente planejamento, diffs e simulação local sem gravação nas refs vedadas.

**Pendências antes da habilitação:** mudança normativa externa aprovada; controle da identidade da sessão, de concorrência e de permissões do repositório de trabalho; demonstração da proteção de branch e dos checks obrigatórios; definição posterior de contratos, condições de falha e auditoria em Scope/Implementation. `OPEN-003` continua em `refinement`, bloqueadora da publicação automatizada em branches de sessão.

**Relações:** [OBJ-003 — Propostas verificáveis](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md), [OBJ-004 — Evidências automatizadas](../objectives/OBJ-004-evidencias-de-testes-automatizados.md), [NOBJ-001 — Sem publicação arbitrária](../non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md) e [PCH-001 — Contexto](../project-context/PCH-001-contexto-integracao-openai-dsl.md).

## Complemento da decisão 15 — Configuração por inferência sem mudar a sessão (decisão 17, 2026-10-21)
A intenção funcional da **sessão única por repositório e nível documental** foi refinada: durante a mesma sessão e branch prevista `feature/doc/{session-id}`, **cada nova inferência pode utilizar modelo autorizado, tamanho de contexto e prompt diferentes**, escolhidos pelo operador. Essas mudanças devem constar do histórico por inferência, **sem criar nova branch ou encerrar a sessão**.

O LLM retorna a proposta em **DSL com título/descrição sugeridos para commit e arquivos alterados**; a aplicação converte integralmente os documentos para **Markdown**, verifica integridade (incluindo tabelas, Mermaid, gráficos e recursos vinculados) e aceita commit intermediário apenas do resultado validado. A regra de **quatro correções após a primeira tentativa**, as análises críticas independentes, a política financeira e a revisão humana de integração permanecem.

**Impedimento normativo inalterado:** a referência externa autoriza agentes a gravar apenas em `develop`. Este complemento **não autoriza criar ou atualizar `feature/doc/*`, gerar PR nem mesclar branches** antes do processo formal separado de revisão normativa e de efetiva proteção da branch `develop`. Relações: [OBJ-001](../objectives/OBJ-001-agentes-documentais-especializados.md), [OBJ-003](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md) e [OPEN-002](OPEN-002-fidelidade-e-custo.md).

## Decisão 18 — Retomada da mesma sessão após reprovação dos checks da PR (2026-10-21)
**Decisão aprovada — opção A:** se os checks obrigatórios de uma PR da sessão falharem, a PR **não pode ser integrada** e a sessão retorna ao estado conceitual **em correção**, preservando ID, repositório, nível documental, histórico, baseline e a branch futura `feature/doc/{session-id}`. O operador recebe as falhas e pode solicitar novas inferências, ajustes e commits **somente após validação de cada resultado**. A mesma PR é atualizada pelos commits válidos e seus checks obrigatórios devem ser executados novamente sobre a versão atualizada. **Checks aprovados na versão antiga não valem como aprovação do novo conteúdo.**

Para voltar a encerrar a sessão é necessária **nova aprovação humana explícita do encerramento**, seguida de checks satisfatórios da PR. A **autorização humana de merge** continua distinta da aprovação de encerramento. Persistindo a reprovação, a sessão permanece em correção e a integração bloqueada. As tentativas e resultados anteriores não são apagados; cada inferência continua sujeita a limite de quatro correções por tarefa, revisão independente, controles de credenciais e autorização financeira por chamada paga.

**Bloqueio normativo:** trata-se apenas de comportamento futuro documentado. A referência vigente permite escrita de agentes **exclusivamente em `develop`**; esta decisão **não autoriza** criar/atualizar `feature/doc/*`, abrir/atualizar PR ou fazer merge. Sua habilitação depende de mudança normativa externa e validação da proteção de `develop`, sem modificar workflows ou permissões nesta etapa. `OPEN-003` permanece `refinement`.
