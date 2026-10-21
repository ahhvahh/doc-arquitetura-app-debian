# NOBJ-003 — Não gerenciar concorrência na branch da sessão

![Tipo: NOBJ](../../assets/documentation-style/badges/badge-type-nobj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Exclusão

A aplicação não terá como objetivo **coordenar usuários diferentes fazendo commits simultâneos na mesma branch de sessão** `feature/doc/{session-id}`, nem resolver automaticamente conflitos causados por essas alterações externas, com ou sem LLM.

## Justificativa

Cada sessão foi concebida para trabalhar na sua própria branch, criada conceitualmente a partir de `develop` quando for aberta. Não se pretende oferecer colaboração concorrente na mesma branch; alterações posteriores da `develop` não são sincronizadas automaticamente com a sessão.

## Efeito da exclusão

O operador não contará com reconciliação multiusuário ou solução automática de commits conflitantes feitos por terceiros na branch da sessão. A aplicação poderá reconhecer novas alterações **na própria branch remota da sessão** e atualizar a cópia local somente quando for seguro preservar o trabalho. Se isso não puder ser assegurado, a ação fica bloqueada, com aviso, sem sobrescrita nem descarte de mudanças.

## Responsabilidade remanescente

Os envolvidos deverão evitar compartilhar a branch de sessão para gravações concorrentes; em incidente de conflito externo, a decisão e a resolução fora do produto cabem ao operador/responsável pelo repositório. A aplicação permanece responsável por **não destruir trabalho**, manter histórico rastreável e respeitar verificações, credenciais e permissões.

## Condição de revisão

Reabrir este não objetivo somente se houver decisão explícita de suportar colaboração simultânea na mesma branch. Qualquer execução futura de branches de sessão depende antes do alinhamento da referência normativa fixa, conforme [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md). O registro não comprova implementação.
