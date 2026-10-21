# OPEN-005 — Estratégia de provedores e modelos locais/remotos

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão em aberto sobre cobertura de modelos e viabilidade de provedores.

## Questão
Quais provedores e classes de modelos estarão disponíveis inicialmente? Serão utilizados modelos locais, serviços remotos ou ambos? Como combinar modelos simples e avançados com as quatro especializações documentais?

## Contexto e origem
O responsável quer experimentar LLMs de capacidades distintas. O protótipo possui integração administrativa com OpenAI e sua documentação informa a remoção do antigo processamento local e de benchmarks. Essa implementação anterior não determina automaticamente o escopo futuro.

## Impacto
**Bloqueadora:** sim para fechar a baseline de inferência multimodelo; não bloqueia o registro dos objetivos nem testes sem inferência paga.

## Alternativas a decidir
- Múltiplos modelos do provedor já previsto antes de ampliar integrações.
- Mais de um provedor remoto desde a primeira entrega.
- Inclusão de modelos locais em fase específica mediante avaliação de viabilidade.

## Próxima ação verificável
Definir a primeira seleção de provedores e modalidades, verificar disponibilidade, identificação de versão, informações de uso/tokens, preços e controles de gasto de cada serviço. Somente após decisão detalhar integrações em Scope.

## Relações
- [OBJ-001 — Especializações](../objectives/OBJ-001-agentes-documentais-especializados.md)
- [OBJ-007 — Comparação](../objectives/OBJ-007-comparacao-controlada-de-modelos.md)
- [OBJ-005 — Orçamento](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OPEN-002 — Fidelidade e custo](OPEN-002-fidelidade-e-custo.md)
