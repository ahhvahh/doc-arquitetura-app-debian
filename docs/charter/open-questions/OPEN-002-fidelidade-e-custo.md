# OPEN-002 — Fidelidade de DSL e custos

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Informação faltante e decisão em aberto.
## Questão
Quais critérios objetivos comprovam fidelidade semântica da DSL, inclusão de diagramas e referências e limites aceitáveis de tokens por execução? O orçamento monetário obrigatório é requisito definido em OBJ-005 e CON-004, não uma autorização para prosseguir sem teto.
## Contexto
A compressão sintática existente ainda não possui prova de equivalência semântica. Diagramas, assets, critérios de aceite e exceções não podem desaparecer durante edição.
## Impacto
**Bloqueadora:** sim para declarar a inferência documental pronta para produção; não bloqueia planejamento e medição exploratória.
## Delimitação da decisão de armazenamento
A publicação manual da DSL em armazenamento consultável pode ser experimentada independentemente da inferência, conforme [OBJ-006](../objectives/OBJ-006-publicacao-manual-base-normativa.md). Upload e indexação bem-sucedidos não comprovam cobertura da norma, fidelidade semântica ou redução real de tokens. Essas dúvidas continuam bloqueadoras da inferência em produção.

## Próxima ação
Definir critérios verificáveis de fidelidade, consultas seletivas ao original e experiências comparativas com tokenização real, sem presumir equivalência semântica automática.
## Destino previsto
Scope de extração, rastreabilidade e validação; Implementation de armazenamento e transporte.

## Pendências específicas do contrato DSL textual
A saída textual pretendida é `BIOMA-DSL-CHANGESET/2.0` com cabeçalho de tarefa e blocos `FILE` por arquivo, sem JSON. A decisão de formato **não** homologa a gramática interna. É necessário especificar caracteres reservados, delimitadores escapáveis, blocos multilinha, estrutura de metadados, badges, links, tabelas, código, Mermaid e regras de edição por seção, e demonstrar conversão DSL → Markdown que preserve todo conteúdo não tocado. O estado `BLOCKED` e sua representação textual também precisam ser formalizados.

Validar cenários de duplicidade de arquivos, hash divergente, alterações fora da etapa, symlinks, path traversal, tamanho excessivo e mudanças de baseline. Até lá, o conversor JSON/Markdown existente não representa implementação deste novo contrato.
