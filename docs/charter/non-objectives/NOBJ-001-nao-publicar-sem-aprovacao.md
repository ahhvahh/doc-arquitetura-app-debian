# NOBJ-001 — Não publicar alterações autonomamente

![Tipo: NOBJ](../../assets/documentation-style/badges/badge-type-nobj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Exclusão
A solução não dará aos modelos de inferência autoridade para integrar alterações em `develop` nem para executar operações Git arbitrárias sem validação e governança da aplicação.
## Justificativa
Decisões geradas podem conter lacunas, alucinações ou operações incompatíveis com as regras do projeto.
## Efeito da exclusão
O responsável humano deve aprovar a **integração final à `develop`**. A criação de commits intermediários pela própria aplicação, após validação em uma branch de sessão, é um objetivo funcional condicionado à revisão normativa de [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md), não uma permissão atualmente vigente. Sem a aprovação de integração, a `develop` não recebe as alterações da sessão.
## Condição de revisão
Somente após nova decisão documentada no Charter e análise das consequências de segurança e auditoria.

## Delimitação da responsabilidade (2026-10-21)
A decisão anterior exigia aprovação humana de aplicação, commit e publicação como três fases. O responsável redefiniu a aprovação humana Git apenas para o merge à `develop`, preservando controles de validade, orçamento e proibição de comandos Git determinados diretamente por LLM. O commit automático em branch de sessão continua impedido na automação documental pela norma atual.

## Distinção entre fechamento, PR e integração Git — decisão 15
O operador poderá, no fluxo pretendido, aprovar o **encerramento de uma sessão**, após examinar o conjunto acumulado de alterações, autorizando assim a abertura de uma **Pull Request de `feature/doc/{session-id}` para `develop`**. A PR deverá passar pelas verificações obrigatórias previstas em [OBJ-004](../objectives/OBJ-004-evidencias-de-testes-automatizados.md).

A **criação da PR não integra alterações** em `develop`. A integração depende de checks bem-sucedidos, proteção efetiva da branch e **aprovação humana final distinta**. A solução não permitirá ao LLM aprovar PRs, contornar verificações ou escolher comandos Git arbitrários.

Esse comportamento é apenas projetado e permanece impedido pela norma externa atual, conforme [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Histórico — Exclusão de branch ao cancelar na decisão 33 (superada)

**Registro histórico superado pela decisão 34:** as orientações antigas de excluir a branch ao cancelar não estão mais vigentes. Os critérios atuais são os da decisão 34 abaixo. Não executar exclusão automática.

A aplicação não poderá interpretar **cancelamento expresso de sessão** como aprovação de conteúdo ou autorização para publicação. No cancelamento, a **exclusão futura da branch exclusiva da sessão** requer confirmação humana específica, vínculo inequívoco e checagem de segurança das refs, com registro de sucesso ou bloqueio; nunca permitir excluir `develop`, `main`, refs protegidas ou de terceiros. O histórico local permanece sujeito à retenção até exclusão manual autorizada.

Na outra alternativa, **somente após aprovação expressa do operador ao encerramento e conjunto documental validado** a aplicação deverá criar **PR da branch da sessão para `develop`**. Essa PR **não mescla automaticamente**. Checks obrigatórios, proteção real da `develop` e **autorização humana posterior e separada para integração** continuam necessários. Sem mudanças publicáveis, não abrir PR.

A decisão funcional pertence à [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md), e as operações de criação/exclusão de branch, PR e merge permanecem **impedidas pela norma vigente**. O conteúdo está **DOCUMENTED**, não `IMPLEMENTED`.

## Decisão 34 — Reprovação não elimina branch nem autoriza publicação
O acionamento de **"Cancelar"** solicita ao usuário escolher **encerrar ou continuar**; não gera exclusão automática de branch. A **reprovação expressa do conjunto** encerra o trabalho sem PR, sem exclusão da branch e sem outras operações sobre ela; a branch permanece intacta. **Aprovação expressa do conjunto validado** habilitará no produto futuro a criação automática de PR **para `develop`**, sem efetuar merge.

A PR continua submetida a checks, proteção da `develop` e **autorização humana posterior e separada para integração**. Não abrir PR sem alterações válidas. A intenção está registrada na [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md), mas a **norma atual proíbe** criar/atualizar branches de sessão, PR ou merge; nenhuma dessas operações está autorizada por esta decisão. `DOCUMENTED`, não `IMPLEMENTED`.
