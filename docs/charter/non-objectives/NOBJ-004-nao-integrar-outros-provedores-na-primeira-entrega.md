# NOBJ-004 — Não integrar outros provedores diretos na primeira entrega

![Tipo: NOBJ](../../assets/documentation-style/badges/badge-type-nobj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Exclusão

A primeira entrega **não incluirá integrações diretas independentes com Anthropic ou Google Gemini**, nem com outros provedores não selecionados. As **únicas integrações previstas para desenvolvimento inicial** são **OpenAI direta** e **agentes Amazon Bedrock (AWS)**. Esta exclusão não proíbe, antecipadamente, modelos de terceiros disponibilizados por um serviço AWS validado; o escopo é a **integração direta adicional**.

## Justificativa

O responsável restringiu o escopo inicial a duas integrações, substituindo a antiga seleção nominal de OpenAI, Anthropic e Gemini, para concentrar os esforços de desenvolvimento e homologação nas ofertas efetivamente escolhidas.

## Efeito da exclusão

O operador não contará na primeira entrega com conectores diretos específicos para Anthropic, Gemini ou outros serviços, nem com garantias de comparação entre tais integrações. A eventual compatibilidade de modelos através da AWS dependerá de homologação do serviço e não representará implementação de uma integração direta excluída.

## Responsabilidade remanescente

O projeto deverá viabilizar e homologar separadamente a **OpenAI** e os **agentes Amazon Bedrock**, sujeitos a disponibilidade e políticas de segurança, orçamento, credenciais e aprovação humana. O responsável poderá decidir mais adiante pela adição de outros provedores, sem declarar sua implementação agora. Relação: [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).

## Condição de revisão

Revisar somente mediante decisão expressa de ampliar o escopo inicial ou planejar entrega posterior, sem reativar automaticamente as decisões históricas de três provedores. Documento em `refinement`; não comprova integração implementada.
