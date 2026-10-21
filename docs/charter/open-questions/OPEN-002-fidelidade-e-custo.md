# OPEN-002 — Fidelidade de DSL e custos

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Diretriz de validação aprovada, com critérios de comprovação de fidelidade e limites de custo ainda pendentes.
## Questão
**Diretriz de validação escolhida:** verificações estruturais obrigatórias, segunda avaliação por modelo de linguagem diferente do gerador e revisão humana final. Ainda é necessário comprovar a fidelidade semântica da DSL, preservar diagramas e referências e validar critérios mensuráveis de qualidade e tokens; o orçamento financeiro segue obrigatório em OBJ-005 e CON-004.
## Contexto
A compressão sintática existente ainda não possui prova de equivalência semântica. Diagramas, assets, critérios de aceite e exceções não podem desaparecer durante edição.
## Impacto
**Bloqueadora:** sim para declarar a inferência documental pronta para produção; não bloqueia planejamento e medição exploratória.
## Delimitação da decisão de armazenamento
A publicação manual da DSL em armazenamento consultável pode ser experimentada independentemente da inferência, conforme [OBJ-006](../objectives/OBJ-006-publicacao-manual-base-normativa.md). Upload e indexação bem-sucedidos não comprovam cobertura da norma, fidelidade semântica ou redução real de tokens. Essas dúvidas continuam bloqueadoras da inferência em produção.

## Próxima ação
Estabelecer evidências verificáveis para integridade estrutural, revisão crítica por segundo LLM e decisão humana; medir a fidelidade e os custos reais das avaliações, testar recuperação do original e impedimentos de publicação/commit quando houver falha. Não presumir equivalência semântica automática por concordância entre modelos.
## Destino previsto
Scope de extração, rastreabilidade e validação; Implementation de armazenamento e transporte.

## Pendências específicas do contrato DSL textual
A saída textual pretendida é `BIOMA-DSL-CHANGESET/2.0` com cabeçalho de tarefa e blocos `FILE` por arquivo, sem JSON. A decisão de formato **não** homologa a gramática interna. É necessário especificar caracteres reservados, delimitadores escapáveis, blocos multilinha, estrutura de metadados, badges, links, tabelas, código, Mermaid e regras de edição por seção, e demonstrar conversão DSL → Markdown que preserve todo conteúdo não tocado. O estado `BLOCKED` e sua representação textual também precisam ser formalizados.

Validar cenários de duplicidade de arquivos, hash divergente, alterações fora da etapa, symlinks, path traversal, tamanho excessivo e mudanças de baseline. Até lá, o conversor JSON/Markdown existente não representa implementação deste novo contrato.

## Proposta de regularização — Decisão 02 (2026-10-21)

**Origem:** decisão explícita do responsável — opção A, **DSL reversível com compactação seletiva**.
**Resultado da triagem:** aceita como diretriz para a regularização documental; a gramática DSL e os critérios de homologação permanecem pendentes.

- O material que puder resultar em criação ou alteração de arquivo deve preservar integralmente o conteúdo necessário à reconstrução sem perdas, inclusive Markdown original, cabeçalhos, badges, links, tabelas, blocos de código, diagramas Mermaid, metadados e separação por arquivo.
- Compactar seletivamente índices, catálogos e material de apoio à consulta quando for seguro, sem ocultar condições normativas obrigatórias, referências relevantes ou evidências necessárias à análise.
- Se a representação resumida não bastar para fundamentar uma análise ou alteração, consultar a fonte original íntegra antes de concluir; sem essa recuperação, bloquear a conclusão dependente.
- Manter por arquivo o caminho de origem, SHA da baseline e hash de integridade; identificar explicitamente conteúdo omitido, compactado e recuperado.
- A aplicação deve validar o ciclo Markdown → DSL → Markdown contra o documento original e impedir aplicação de mudanças que removam ou transformem conteúdo não solicitado.
- Medir tokens efetivamente consumidos e estimativa de custo, comparando o modo completo ao seletivo, sem considerar a redução de tokens comprovada antes de medição.

**Limites da decisão:** esta proposta não define delimitadores, escape, formato de campos, recuperação remota, limite numérico de tokens ou algoritmo de compactação. O contrato `BIOMA-DSL-CHANGESET/2.0` continua sujeito às validações descritas nesta OPEN. Essas definições deverão ser formalizadas em Scope/Implementation, seguindo os gates documentais.

**Impacto no gate:** a diretriz de arquitetura foi escolhida, mas a `OPEN-002` permanece em `refinement` e continua bloqueadora para declarar inferência documental pronta para produção até comprovar fidelidade e segurança dos casos de teste.

## Extensão aprovada — comparação objetiva entre modelos (2026-10-21)
Uma comparação válida deverá usar a mesma tarefa, prompt efetivo, documentação e SHAs de referência, ou identificar explicitamente variações. Avaliar integridade do Markdown, links, diagramas, validade do changeset, aderência à norma, correções humanas, tempo, tokens e custos previstos e realizados, quando mensuráveis.

Pesos, escala de qualidade e limites numéricos de aprovação ainda exigem definição e validação. Registrar dados experimentais não comprova automaticamente superioridade, fidelidade ou equivalência semântica.

Relações: [OBJ-007](../objectives/OBJ-007-comparacao-controlada-de-modelos.md), [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-005](OPEN-005-estrategia-de-provedores-e-modelos.md).

## Decisão 10 — Dupla avaliação de qualidade com revisão humana (2026-10-21)
**Origem:** o responsável selecionou explicitamente a opção B — **validação automática + segunda análise por LLM + revisão humana** — e aprovou sua documentação exclusivamente na `develop`.

**Resultado da triagem:** modalidade de controle de qualidade **aceita como objetivo funcional**, sem afirmar implementação, confiabilidade homologada ou disponibilidade financeira dos modelos.

- **Primeira barreira:** a aplicação realiza verificações automatizadas obrigatórias da integridade das alterações propostas, abrangendo escopo e hashes de origem, conteúdo não solicitado, links, badges, tabelas, blocos de código, referências e diagramas pertinentes. Falhas estruturais impedem aceitar e aplicar o resultado; uma segunda análise por LLM não dispensa nem substitui esses testes.
- **Segunda avaliação:** após a validação estrutural, **outro modelo de linguagem**, distinto do que gerou a alteração, analisa criticamente o resultado frente à tarefa, às evidências e à versão normativa. Seu parecer deverá apontar qualidade, omissões, contradições, riscos e trechos sem sustentação, com identificação do provedor/modelo e contexto avaliado.
- **Decisão humana:** apresentar as alterações, os testes e o parecer secundário ao operador para revisão humana antes do aceite final e de qualquer integração Git que venha a ser normativamente permitida. Parecer favorável do revisor LLM não constitui aprovação automática.
- **Rastreabilidade:** relacionar inferência geradora, avaliação secundária, parecer, métricas de uso/custo, eventuais divergências, ajustes e decisão humana ao projeto, sessão e baselines. Não tratar concordância entre LLMs como prova da equivalência semântica da DSL.
- **Orçamento:** a revisão secundária também é uma operação de inferência e obedece a [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md). Sem proteção remota comprovada ou autorização financeira exigida para chamada paga, **não executar a revisão paga nem considerar o fluxo de validação concluído**.
- **Governança Git:** a aprovação humana final de qualidade não substitui a autorização explícita para integração à `develop`; commits em branch de sessão continuam condicionados à [OPEN-003](OPEN-003-politica-de-branches.md). Nenhuma gravação fora de `develop` é autorizada por esta decisão.

**Pendências de comprovação:** critérios de equivalência Markdown → DSL → Markdown, parâmetros de qualidade e escala, comportamento para discordâncias, limites numéricos de tokens, validação dos modelos elegíveis e testes de custos/segurança. Definições técnicas pertencem a Scope/Implementation após os gates. Manter a OPEN em `refinement` enquanto tais evidências faltarem.

**Relações canônicas:** [OBJ-003 — Alterações verificáveis](../objectives/OBJ-003-propostas-de-alteracoes-verificaveis.md), [OBJ-004 — Evidências de testes](../objectives/OBJ-004-evidencias-de-testes-automatizados.md), [OBJ-007 — Avaliação de modelos](../objectives/OBJ-007-comparacao-controlada-de-modelos.md).
