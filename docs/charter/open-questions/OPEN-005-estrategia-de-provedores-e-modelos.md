# OPEN-005 — Estratégia de provedores e modelos locais/remotos

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão sobre modalidade de inferência aceita; seleção inicial de serviços e validação de viabilidade ainda pendentes.

## Questão
**Modalidade decidida:** múltiplos provedores remotos (opção B). Resta escolher quais provedores e modelos serão habilitados na primeira entrega e comprovar as condições mínimas de uso, custo e segurança.

## Contexto e origem
O responsável quer experimentar LLMs de capacidades distintas. O protótipo possui integração administrativa com OpenAI e sua documentação informa a remoção do antigo processamento local e de benchmarks. Essa implementação anterior não determina automaticamente o escopo futuro.

## Impacto
**Bloqueadora parcial:** a escolha da modalidade está decidida; a viabilidade dos provedores concretos e os controles necessários ainda bloqueiam a homologação da inferência paga e o fechamento da baseline multimodelo.

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

## Decisão 06 — Provedores remotos distintos (2026-10-21)
**Origem:** opção B escolhida explicitamente pelo responsável e autorizada para documentação. **Resultado da triagem:** aceita quanto à modalidade remota; seleção concreta de provedores permanece pendente.

- A aplicação terá como objetivo trabalhar com modelos de **diferentes provedores remotos**, permitindo comparar ao menos dois provedores sobre a mesma tarefa e contexto documental.
- Preservar Charter, Scope, Implementation e Source como especializações normativas independentes dos provedores; modelos simples e complexos são escolhas de execução, não novas etapas.
- Identificar provedor/modelo/versão em cada tentativa e preservar as condições de comparação de qualidade, duração, tokens e custos, com histórico local em [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md).
- Modelos executados localmente ficam fora da primeira entrega, conforme [NOBJ-002](../non-objectives/NOBJ-002-sem-inferencia-local-na-primeira-entrega.md), sem eliminar estudos futuros.
- Manter autorização financeira, proteção de credenciais e bloqueio de operações fora do orçamento; a existência de limites remotos e a disponibilidade das métricas devem ser verificadas por provedor, sem pressupor equivalência com OpenAI.
- Não atribuir aos modelos permissão para aplicar mudanças Git autonomamente. A política de sessão/branch e de integração será tratada em [OPEN-003](OPEN-003-politica-de-branches.md).

**Próxima decisão:** quais provedores concretos serão inicialmente aceitos e com quais evidências de uso, preços, segurança e monitoramento de despesas. Contratos de integração e nomes de modelos específicos serão objeto de Scope depois do gate. Manter este documento em `refinement` até concluir os critérios necessários.
