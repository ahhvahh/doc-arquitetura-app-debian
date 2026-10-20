# CON-003 — Separação entre norma externa e contratos do aplicativo

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
A documentação normativa de `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0` deve permanecer externa e ser clonada localmente. Apenas os contratos particulares do `doc-arquitetura-app-debian` ficam em `/docs/charter`, `/docs/scope` e `/docs/implementation` do aplicativo, sem copiar os documentos normativos para seu Git.

## Origem verificável
Decisão explícita do responsável: manter os contratos do aplicativo em `/docs` e utilizar a norma externa como fonte de leitura e compilação DSL.

## Motivo e impacto
Evita documentos normativos duplicados e mantém a rastreabilidade da versão utilizada por cada agente e execução.

## Verificação
Os commits do aplicativo não contêm cópias da documentação normativa. A aplicação identifica URL, ref `main` e SHA do clone normativo, e usa clones transitórios não versionados para atualizar normas ou processar cada projeto, sem cache de contexto reutilizado entre inferências.

## Relações
- [CON-001](CON-001-referencia-normativa-main.md)
- [OBJ-001](../objectives/OBJ-001-agentes-documentais-especializados.md)

## Condição de revisão
Somente nova decisão do responsável, com análise de compatibilidade documental.
