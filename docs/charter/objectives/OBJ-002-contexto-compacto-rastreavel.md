# OBJ-002 — Contexto compacto e rastreável

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
Enviar todo o material bruto aumenta a quantidade de texto e dificulta localizar a origem dos requisitos.
## Resultado esperado
A aplicação fornece ao mecanismo de análise uma representação documental reduzida, mantendo rastreabilidade à versão e ao arquivo original.
## Critérios verificáveis
- Cada unidade documental processada aponta seu arquivo de origem e a versão verificada.
- O processo permite identificar quais documentos foram incluídos, omitidos ou descartados como backup.
- O estado local da branch `develop` do projeto-alvo é reconciliado com o remoto antes da extração, sem descartar trabalho local sem consentimento.
- Quando o resumo não for suficiente para revisão segura, o conteúdo original poderá ser consultado; não se perderá informação normativa silenciosamente.
- O custo e o volume real de contexto deverão ser mensuráveis.
## Delimitação da base normativa
A representação normativa destinada a consulta remota é compilada e atualizada apenas por solicitação do operador. O envio ao armazenamento não comprova equivalência semântica entre DSL e Markdown nem valida alterações no projeto-alvo. A fonte original permanece disponível para verificação. A sincronização manual é um objetivo próprio: [OBJ-006](OBJ-006-publicacao-manual-base-normativa.md).

## Dependências
- [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Tamanho de contexto configurável por inferência — decisão 17
A **sessão** mantém sua identidade, repositório e nível documental, mas a **janela/tamanho de contexto a utilizar é uma escolha do operador a cada inferência**, limitada pela capacidade efetiva comprovada do modelo selecionado e pelas reservas necessárias para resposta e revisão.

A aplicação deverá identificar tamanho solicitado, capacidade admissível, origem e baseline dos documentos efetivamente incluídos, material necessário à inferência e impacto de tokens e custos. A representação DSL pesquisável da norma não dispensa a verificação de conteúdo integral exigido pelo trabalho; **compactar é uma estratégia de consulta, não licença para perder texto, tabelas, blocos Mermaid, diagramas ou referências obrigatórias**.

Se não for possível manter todo o material requerido dentro do contexto disponível, interromper e seguir a recuperação com definição humana explícita prevista em [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), **sem truncar silenciosamente nem mudar o modelo automaticamente**. O usuário pode fazer outra escolha de configuração em uma nova inferência da mesma sessão, respeitando orçamento e identidade autorizada.

Esta diretriz não afirma capacidade técnica já implementada ou equivalência semântica comprovada da DSL. Relações: [OBJ-001](OBJ-001-agentes-documentais-especializados.md) e [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).

## Padrão prévio e escolha do contexto por inferência — decisão 24
A configuração persistente de motor/modelo e prompt por especialização, operação e papel é um **ponto de partida anterior à sessão**. Ao utilizar tais padrões, a aplicação deve apresentar o **tamanho de contexto realmente admissível** pelo modelo selecionado, considerando toda a entrada obrigatória e a reserva necessária à resposta e à revisão.

A **decisão 17 permanece vigente**: a janela/tamanho de contexto efetivamente solicitada é **escolhida ou confirmada pelo operador em cada inferência**, podendo diferir do padrão anterior. Registrar lado a lado a versão do perfil carregado, o tamanho solicitado, o limite efetivamente admissível, os documentos/baselines e o custo autorizado. Alterar o perfil persistente não altera retroativamente inferências já realizadas.

Não escolher automaticamente outro modelo, nem resumir, omitir ou truncar material obrigatório para fazê-lo caber; falta de capacidade implica **bloqueio e esclarecimento** conforme [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md). Relações: [OBJ-001](OBJ-001-agentes-documentais-especializados.md) e [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md). O formato do arquivo de configuração e a representação de limites pertencem às etapas seguintes.
