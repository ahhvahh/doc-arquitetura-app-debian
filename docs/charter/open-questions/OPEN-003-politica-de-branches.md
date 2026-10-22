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

## Decisão 19 — Origem fixa em develop e exclusão de concorrência na branch de sessão (2026-10-21)
**Origem:** esclarecimento explícito do responsável: cada sessão futura terá **uma branch exclusiva derivada da `develop` no momento de abertura**. O usuário descartou a proposta de mesclar continuamente as mudanças posteriores de `develop` e decidiu não tratar edição concorrente por terceiros como capacidade do produto.

- **Ponto de partida:** na abertura, identificar a baseline/SHA de `develop` do repositório escolhido e criar conceitualmente a branch de trabalho única `feature/doc/{session-id}`, quando a política externa permitir. A sessão permanece nessa branch até seu encerramento.
- **Relação com `develop`:** alterações posteriores em `develop` **não são incorporadas automaticamente** à branch da sessão. Não executar merge, rebase ou sincronização periódica de `develop` para manter a sessão aberta. A comparação com a base de integração pode ser exigida no fechamento/PR para verificações e revisão humana, sem alterar a baseline automaticamente.
- **Mesma branch remota:** se a **própria `feature/doc/{session-id}` remota** tiver novos commits observados, a aplicação futura deverá atualizar a visão/cópia local a partir **dessa mesma branch**, somente se a atualização puder preservar o trabalho existente com segurança. Não confundir esta atualização com importar `develop`.
- **Concorrência fora do escopo:** não desenvolver coordenação multiusuário, merge automático de conflitos externos ou resolução por LLM de commits feitos por terceiros na branch da sessão. A expectativa é de uso exclusivo da branch pela sessão; essa expectativa **não é garantia técnica**. Se a divergência impedir atualização segura, interromper a ação afetada e informar a condição, **sem sobrescrever trabalho alheio, reset, force-push ou integração presumida**. A responsabilidade de tratar o incidente externo permanece com o operador.
- **Relações:** a exclusão deliberada está registrada em [NOBJ-003 — Não gerenciar concorrência na branch da sessão](../non-objectives/NOBJ-003-nao-gerenciar-concorrencia-na-branch-da-sessao.md). Comportamento futuro relacionado a [OBJ-003](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md).

**Restrição normativa preservada:** esta decisão é somente de Charter. A norma vigente de `main/ProjectCharter/1.0` autoriza agentes a escrever **apenas em `develop`**. Não criar nem atualizar branches de sessão, não executar sincronização Git, merge/rebase, PR ou workflow nesta tarefa. `OPEN-003` continua em `refinement`, pendente de mudança normativa externa e validação das proteções de integração.

## Decisão 30 — Suspender trabalho ativo com pendências, sem concluir sessão para publicação (2026-10-21)
**Origem:** o responsável aprovou expressamente a **opção B — permitir encerramento do trabalho ativo sem publicação**, mesmo com tarefas ainda aguardando revisão ou outra regularização.

**Distinção de eventos:**
- **Encerrar o trabalho ativo / suspender a sessão:** o operador poderá parar de trabalhar e deixar a sessão **suspensa, não concluída para publicação**, com pendências explicitamente identificadas. Essa ação **não equivale** à aprovação humana de encerramento para publicação, ao cancelamento de tarefas, à autorização de PR/merge ou a uma autorização financeira.
- **Reabrir a mesma sessão:** o operador autorizado deverá poder recuperar a **mesma identidade de sessão e tarefas**, seus estados, propostas, histórico, custos, baselines/hashes, perfis resolvidos na abertura e saldos dos ciclos de correção. As pendências precisarão ser verificadas e tratadas **na sessão reaberta**, sem criar outra sessão por padrão ou apagar o histórico.
- **Concluir formalmente para publicação:** somente após regularizar todas as tarefas que bloqueiam a conclusão, com **revisão independente válida e demais verificações** (ou decisão humana expressa e rastreável de cancelamento/retirada de uma tarefa pendente, sem considerar seu conteúdo aprovado), a sessão poderá ser apresentada ao operador para **aprovação explícita do conjunto final**. Essa aprovação é distinta de qualquer posterior autorização humana de integração Git.

**Consequências verificáveis:**
- É permitido suspender trabalho mesmo que uma tarefa esteja em **avaliação pendente**, preservando a proposta não validada exclusivamente no histórico controlado; **não é permitido declará-la concluída, aceita, publicada ou gerar commit de resultado validado** por esse motivo.
- Enquanto houver avaliação pendente não regularizada, a sessão permanece **não concluída para publicação**. O sistema deve exibir o **motivo impeditivo**, a relação de tarefas pendentes e a exigência de reabertura para prosseguir; nenhuma aprovação de encerramento ou publicação é inferida do ato de suspender.
- A reabertura respeita as **decisões 25 a 29**: herança controlada e padrões iniciais congelados, modelos distintos e homologados para nova geração validada, recuperação de **revisão já pendente sem refazer a geração** quando a integridade permitir, e persistência entre encerramentos/reinicializações.
- Antes de qualquer nova inferência paga, revisão retomada ou correção, **revalidar identidade/permissões, integridade do contexto e baseline, homologação, autorização financeira específica e proteção remota bloqueante comprovada**. Suspensão, reabertura e aprovação de encerramento não fornecem autorização implícita para gastar nem restauram credenciais.
- A suspensão e reabertura **não reiniciam o máximo de quatro correções por tarefa**, não dispensam avaliação por segundo LLM, verificação normativa/estrutural, auditoria ou revisão humana final. Não incorporar automaticamente atualizações de `develop` à baseline fixa da sessão.
- **Bloqueio normativo vigente:** a intenção futura de criar/atualizar `feature/doc/{session-id}`, abrir PR e integrar alterações permanece **incompatível com a regra atual de escrita exclusiva na branch `develop`**. Esta decisão registra comportamento de produto e **não autoriza** criar refs, PRs, merges, workflows ou gravar fora do Charter na `develop`.

**Critérios de aceitação propostos:** sessão com tarefa T em avaliação pendente pode ser suspensa sem erro, mas permanece **não concluída e inelegível para publicação**; após reabertura autorizada, T ainda está pendente com identidade, proposta, baseline, histórico e saldo intactos; T poderá ser revisada conforme [OPEN-002](OPEN-002-fidelidade-e-custo.md), decisão 29, sem nova geração quando íntegra; a conclusão formal exige remoção verificável dos impedimentos e **nova aprovação humana explícita**, sem efeitos de publicação automática.

**Pendências:** comprovar recuperação, segregação, estados funcionais e bloqueios de conclusão nos testes de fases posteriores; compatibilizar separadamente a política Git da OPEN-003 com a norma principal antes de habilitar branches/PRs. **DOCUMENTED**, não `IMPLEMENTED`. Documento permanece `refinement`. Relações: [OBJ-003](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md), [OBJ-004](../objectives/OBJ-004-evidencias-de-testes-automatizados.md), [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-002](OPEN-002-fidelidade-e-custo.md) e [OPEN-006](OPEN-006-retencao-e-protecao-do-historico.md).

## Decisão 31 — Conclusão formal sem alterações publicáveis e sem PR (2026-10-21)
**Origem:** o responsável aprovou expressamente a **opção A — concluir sem publicação**, quando, após reabrir ou revisar uma sessão, todas as tarefas restantes forem **canceladas expressamente** e **não houver qualquer alteração documental validada a publicar**.

**Regra de produto:** a aplicação deverá permitir que o operador autorizado **conclua formalmente essa sessão como "concluída sem publicação"**, em vez de obrigatoriamente classificá-la como sessão cancelada. **Nenhuma Pull Request, publicação, integração Git ou commit de resultado documental** será criado por essa conclusão vazia. Essa classificação descreve o desfecho da **sessão**, e **não altera os status normativos dos arquivos/documentos do Charter**.

**Condições cumulativas verificáveis para a conclusão:**
- Todas as tarefas pendentes foram **canceladas ou retiradas individualmente por decisão humana expressa**, com autoria, momento e razão/alcance rastreáveis; uma tarefa interrompida, suspensa ou sem revisão **não é cancelada automaticamente**.
- Foi conferido o inventário final da mesma sessão e **não existe modificação documental validada ainda destinada à publicação**, nem trabalho já aprovado pendente de decisão de publicação. Se houver entrega validada, tarefa sem resolução ou alteração não contabilizada, este encerramento vazio **não é permitido**; apresentar os impedimentos.
- O operador visualiza a **relação final de tarefas e seus destinos**, a inexistência de alterações publicáveis e as consequências da conclusão; exige-se **confirmação humana explícita do encerramento sem publicação**, distinta das autorizações de cancelamento de tarefas e de qualquer autorização Git.
- **Preservar o histórico completo** da sessão e das tarefas (inclusive propostas não validadas e canceladas, baseline, hashes quando existentes, versões dos padrões fixados na abertura, modelos, tentativas, validações, autorizações, custos reais e o contador/saldo do teto de **quatro correções por tarefa**), segundo [OPEN-006](OPEN-006-retencao-e-protecao-do-historico.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md). **Cancelar uma tarefa não exclui seus registros**; descarte só ocorre mediante exclusão manual autorizada pela política já definida.
- Após a confirmação, registrar estado terminal conceitual **"concluída sem publicação"** e inexistência de PR/commit/merge correspondentes; **não publicar nem recuperar automaticamente as propostas canceladas como válidas**. Os fatos históricos permanecem consultáveis ao operador autorizado conforme a política de retenção.

**Relação com as decisões 25–30:** mantém-se a herança controlada, os padrões de abertura congelados, a escolha por inferência, a pré-verificação de dois LLMs para geração validada, a recuperação de revisão pendente e a possibilidade de suspender/reabrir a sessão. Esta decisão **não dispensa avaliação independente para qualquer conteúdo que venha a ser validado**, nem amplia o máximo de quatro correções, concede gastos, expõe credenciais ou contorna a aprovação humana. **Nenhuma inferência remunerada é necessária para concluir uma sessão sem publicações.**

**Impedimento normativo inalterado:** o processo futuro que prevê `feature/doc/{session-id}` e PR permanece bloqueado até revisão própria da norma `main/ProjectCharter/1.0`; esta decisão **não cria autorização** para escrever em outras branches, criar PR, executar merges ou alterar configuração Git. Documento **DOCUMENTED**, não `IMPLEMENTED`, permanece em `refinement`.

**Critérios de aceite futuro:** sessão S com tarefas pendentes canceladas **expressamente**, sem alterações validadas publicáveis, pode ser marcada **concluída sem publicação** mediante confirmação final; a contagem de PRs, commits documentais e integrações resultantes é **zero**; sessão S conserva trilha de auditoria; suspensão sem cancelamento, pendência ativa ou resultado validado ainda destinado à publicação **bloqueiam** essa modalidade. Relações: [OBJ-003](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md), [OBJ-004](../objectives/OBJ-004-evidencias-de-testes-automatizados.md), [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-002](OPEN-002-fidelidade-e-custo.md) e [NOBJ-001](../non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md).
