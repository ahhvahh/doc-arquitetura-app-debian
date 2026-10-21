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
