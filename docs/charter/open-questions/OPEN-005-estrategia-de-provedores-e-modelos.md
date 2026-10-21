# OPEN-005 — Estratégia de provedores e modelos locais/remotos

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão de produto aprovada sobre modalidade remota e três provedores previstos; validação da viabilidade e homologação ainda pendentes.

## Questão
**Modalidade decidida:** múltiplos provedores remotos (opção B). **Provedores previstos para a primeira entrega:** OpenAI, Anthropic e Google Gemini (opção C da seleção de provedores). A disponibilidade e os controles específicos de cada um ainda precisam ser comprovados antes da habilitação.

## Contexto e origem
O responsável quer experimentar LLMs de capacidades distintas. O protótipo possui integração administrativa com OpenAI e sua documentação informa a remoção do antigo processamento local e de benchmarks. Essa implementação anterior não determina automaticamente o escopo futuro.

## Impacto
**Bloqueadora parcial:** a modalidade e os três provedores pretendidos foram escolhidos; a homologação depende de verificar disponibilidade, preços, métricas, identificação de modelos, credenciais, segurança e controles de gastos para cada serviço, sem presumir que capacidades entre eles sejam equivalentes.

## Alternativas avaliadas
- Utilizar inicialmente apenas modelos OpenAI: não escolhida.
- Trabalhar com múltiplos provedores remotos: modalidade escolhida na decisão 06.
- Incluir execução local de LLMs na primeira entrega: fora do escopo conforme NOBJ-002.
- **Seleção nominal aprovada:** OpenAI, Anthropic e Google Gemini para a primeira entrega, condicionada à verificação de viabilidade.

## Próxima ação verificável
Verificar separadamente em OpenAI, Anthropic e Google Gemini as APIs disponíveis, os modelos/versões elegíveis, informações de uso e tokens, tarifas/moedas, limites e ferramentas de monitoramento de gastos, proteção de credenciais, condições de tratamento de conteúdo e compatibilidade com as especializações normativas. Registrar evidências e bloqueios por serviço antes de habilitar inferências pagas; detalhar os contratos somente em Scope/Implementation.

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

**Decisão posterior:** a seleção nominal dos provedores foi feita na decisão 08 abaixo. Ainda são pendentes as evidências de uso, preços, segurança, métricas e monitoramento de despesas. Contratos de integração e modelos/versões específicos serão avaliados nas etapas seguintes. Manter este documento em `refinement` até concluir os critérios necessários.

## Decisão 08 — Três provedores remotos previstos (2026-10-21)

**Origem verificável:** responsável escolheu explicitamente a **opção C: OpenAI + Anthropic + Google Gemini**, autorizando o registro no Charter do repositório-alvo na branch `develop`.

**Resultado da triagem:** seleção nominal **aceita como intenção de produto**, sem comprovação de disponibilidade técnica, integração implementada, tarifas, métricas ou garantias comerciais.

- **OpenAI:** provedor remoto previsto para a primeira entrega.
- **Anthropic:** provedor remoto previsto para a primeira entrega.
- **Google Gemini:** provedor remoto previsto para a primeira entrega.
- As quatro especializações Charter, Scope, Implementation e Source mantêm sua finalidade normativa, independentemente do provedor e do modelo escolhido.
- Cada serviço só poderá ser habilitado após comprovar sua utilização autorizada, identificação do modelo e versão, estimativas e medições de uso/tokens quando disponibilizadas, preços/câmbio, medidas de segurança e controle efetivo de despesas.
- Se algum serviço não atender a requisitos obrigatórios, **bloquear sua ativação** e registrar a lacuna; sua substituição ou retirada do objetivo inicial depende de nova decisão do responsável, não de inferência automática.
- Não presumir equivalência de tokenização, métricas, faturamento, capacidades de pesquisa ou limites financeiros entre provedores.
- A comparação entre provedores deve conservar a procedência de métricas e os mesmos contextos de tarefa, conforme [OBJ-007](../objectives/OBJ-007-comparacao-controlada-de-modelos.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md).
- A proteção de credenciais e a política de orçamento são obrigatórias para os três serviços, conforme [CON-002](../conditions/CON-002-protecao-credenciais.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md).

**Pendências bloqueadoras:** realizar verificação documentada individual de disponibilidade, segurança, condições de uso, tarifas e métricas de cada serviço antes da liberação de inferência paga. Nomes específicos de modelos, interfaces técnicas e contratos pertencem a Scope/Implementation; este registro não os antecipa.

## Condição financeira da primeira entrega — decisão 09
**OpenAI, Anthropic e Google Gemini** permanecem como os três provedores previstos, porém **nenhum é automaticamente elegível à inferência paga**. Antes de habilitar qualquer provedor/modelo, exigir comprovação individual de mecanismo remoto bloqueante suficiente, além dos limites locais de orçamento e da autorização de gasto.

Alertas ou quotas apenas informativas não satisfazem o requisito. Na OpenAI, as evidências devem ser verificadas também para cada um dos quatro projetos independentes das especializações. Se algum serviço não oferecer proteção remota adequada, sua ativação paga fica bloqueada até nova decisão e comprovação, sem presumir que a API oferece hard limit. Relação: [OPEN-004 — Política aprovada](OPEN-004-projeto-openai-e-hard-limit.md).
