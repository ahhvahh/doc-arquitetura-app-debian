# OPEN-005 — Estratégia de provedores e modelos locais/remotos

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão de produto atualizada: duas integrações previstas para a primeira entrega — OpenAI direta e Amazon Bedrock AgentCore (AWS) —, com modalidade AWS escolhida na decisão 23 e homologação ainda pendente. As escolhas anteriores estão preservadas no histórico.

## Questão
**Modalidade decidida:** integrações remotas. **Escopo vigente da primeira entrega: integração direta OpenAI e integração de agentes no ecossistema Amazon Bedrock (AWS)**, ambas **planejadas, não implementadas ou homologadas**. Anthropic e Google Gemini **não são integrações diretas previstas nesta entrega**. A oferta AWS foi definida como **Amazon Bedrock AgentCore (decisão 23)**; a viabilidade, elegibilidade da conta, segurança e proteção financeira permanecem sem comprovação.

## Contexto e origem
O responsável quer experimentar LLMs de capacidades distintas. O protótipo possui integração administrativa com OpenAI e sua documentação informa a remoção do antigo processamento local e de benchmarks. Essa implementação anterior não determina automaticamente o escopo futuro.

## Impacto
**Bloqueadora parcial:** as duas integrações da primeira entrega foram escolhidas, mas a habilitação depende de verificar disponibilidade real, oferta AWS elegível, modelos subjacentes, custos, métricas, credenciais, segurança e proteção efetiva dos gastos. Não presumir que um agente gerenciado da AWS equivale a um modelo de linguagem nem que os serviços tenham capacidades idênticas.

## Alternativas avaliadas
- Utilizar inicialmente apenas modelos OpenAI: não escolhida.
- Trabalhar com múltiplos provedores remotos: modalidade escolhida na decisão 06.
- Incluir execução local de LLMs na primeira entrega: fora do escopo conforme NOBJ-002.
- **Seleção nominal histórica (decisão 08, superada pela decisão 22):** OpenAI, Anthropic e Google Gemini eram os três provedores cogitados para a primeira entrega.
- **Seleção vigente (decisão 22):** integração direta OpenAI e integração de agentes Amazon Bedrock (AWS); integrações diretas Anthropic e Gemini excluídas da primeira entrega, conforme [NOBJ-004](../non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md).

## Próxima ação verificável
Verificar **separadamente OpenAI e Amazon Bedrock AgentCore (AWS)**: viabilidade do AgentCore, acesso à oferta/conta/região, identidade dos agentes e modelos subjacentes e da especialização, informações de uso/tokens, tarifas/moedas, limites e mecanismos de interrupção financeira, segurança e proteção de credenciais. Documentar evidências por integração antes de habilitar chamadas pagas; contratos e mecanismos técnicos pertencem a Scope/Implementation. Anthropic e Gemini não serão investigados como integrações diretas obrigatórias da primeira entrega.

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

**Histórico substituído:** a lista de três integrações deste registro foi substituída, para o escopo inicial, pela **decisão 22**. Os textos abaixo documentam a escolha passada e não constituem meta vigente da primeira entrega.

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
**Aplicabilidade vigente:** a lista nominal original abaixo era histórica; a exigência de controles financeiros é **transversal** e se aplica às integrações atualmente previstas, **OpenAI e agentes AWS**, conforme decisão 22.

**OpenAI, Anthropic e Google Gemini** permanecem como os três provedores previstos, porém **nenhum é automaticamente elegível à inferência paga**. Antes de habilitar qualquer provedor/modelo, exigir comprovação individual de mecanismo remoto bloqueante suficiente, além dos limites locais de orçamento e da autorização de gasto.

Alertas ou quotas apenas informativas não satisfazem o requisito. Na OpenAI, as evidências devem ser verificadas também para cada um dos quatro projetos independentes das especializações. Se algum serviço não oferecer proteção remota adequada, sua ativação paga fica bloqueada até nova decisão e comprovação, sem presumir que a API oferece hard limit. Relação: [OPEN-004 — Política aprovada](OPEN-004-projeto-openai-e-hard-limit.md).

## Decisão 21 — Entrega parcial identificada somente com provedores homologados (2026-10-21)
**Atualização de escopo:** a regra de habilitar somente integrações comprovadamente homologadas **permanece válida**, mas a meta nominal de três provedores indicada abaixo foi **substituída pela decisão 22**; a entrega completa agora depende das duas integrações planejadas, não da antiga lista.

**Origem:** escolha expressa do responsável — **opção B, liberar somente os provedores homologados** — autorizada apenas para registro no Charter da branch `develop`.

**Resultado da triagem:** a seleção de **OpenAI, Anthropic e Google Gemini** como três provedores pretendidos permanece válida, mas deixa de representar uma condição de simultaneidade para a liberação de uma **entrega parcial explicitamente identificada**. A meta completa de suporte aos três continua **pendente** enquanto houver provedor não homologado; uma entrega parcial **não deve ser apresentada como cumprimento integral** do objetivo original.

- **Habilitação independente:** a aplicação futura poderá disponibilizar **apenas os provedores e modelos individualmente homologados** para o cenário de uso autorizado, informando ao operador o conjunto efetivamente habilitado, o escopo e as limitações da entrega parcial.
- **Não homologado significa desabilitado:** todo provedor ou modelo sem evidências suficientes de disponibilidade, identidade e versão, segurança, segregação de credenciais por usuário, preços e métricas necessários, autorização financeira e **proteção remota efetiva e bloqueante** permanece indisponível para inferências remuneradas. A indisponibilidade não autoriza migração silenciosa, substituição automática nem exceção aos controles.
- **Proteção financeira cumulativa:** para **cada chamada paga** de geração, revisão secundária ou correção, permanecem obrigatórias a prova da proteção remota aplicável ao provedor/projeto, a verificação de orçamento local, o custo previsto e a autorização explícita da operação. Alertas, quotas apenas informativas e a existência de uma assinatura **não equivalem** a bloqueio remoto comprovado. Na OpenAI, preservar quatro projetos separados por especialização e verificar a eficácia do controle em cada um deles, conforme [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md) e [CON-004](../conditions/CON-004-orcamento-obrigatorio.md).
- **Identidade e segregação:** nenhuma habilitação de provedor modifica as regras de execução no contexto autorizado do usuário ou permite compartilhar segredos ou credenciais, conforme [OPEN-001](OPEN-001-politica-credenciais.md) e [CON-002](../conditions/CON-002-protecao-credenciais.md).
- **Evidências da entrega:** identificar quais provedores e modelos foram homologados, para qual especialização e uso, por quais evidências, qual é a validade da homologação e quais permanecem desabilitados e por quê. Apresentar comparação somente quando houver resultados válidos comparáveis; não inventar métricas de provedores ausentes nem contar indisponibilidade como desempenho zero.
- **Reavaliação:** um provedor pendente poderá ser habilitado posteriormente **apenas após homologação individual comprovada** e atualização explícita das informações da entrega. Se perder qualquer condição obrigatória, suspender seu uso remunerado; não pressupor disponibilidade contínua. Se **nenhum** provedor atender aos controles, não habilitar inferências pagas.
- **Limites do registro:** a decisão não comprova existência de hard limits nos provedores, de testes concluídos ou de integração implementada; também não dispensa a validação documental, a revisão independente, a norma Git ou a aprovação humana de integração, conforme [OPEN-002](OPEN-002-fidelidade-e-custo.md) e [OPEN-003](OPEN-003-politica-de-branches.md).

**Pendências:** homologação operacional e financeira individual, viabilidade e evidência de capacidades de cada provedor/modelo, critérios verificáveis para comparação e habilitação por especialização. `OPEN-005` continua em `refinement`; a entrega completa com os três provedores não está atestada.

## Decisão 22 — Duas integrações iniciais: OpenAI e agentes Amazon Bedrock (2026-10-21)
**Origem:** confirmação expressa do responsável, após solicitar restringir as implementações planejadas inicialmente a **OpenAI** e a **AWS, especificamente integração com agentes Amazon Bedrock**. Esta decisão **substitui a seleção nominal** das decisões 08 e 21 para a primeira entrega, sem revogar a obrigação de homologação individual e sem apagar seu histórico.

**Resultados pretendidos do produto:**
- **OpenAI:** uma integração direta que permita a execução autorizada de agentes/modelos para as quatro especializações documentais, respeitando a estrutura já aprovada de **quatro projetos separados por especialização**. A existência de funções anteriores no protótipo não comprova que esta integração funcional planejada esteja concluída.
- **AWS / agentes Amazon Bedrock:** uma segunda integração independente para utilizar a capacidade de **agentes gerenciados do ecossistema Amazon Bedrock**, com a identidade do agente, seus modelos subjacentes e recursos sujeitos à validação. **Contexto histórico da decisão 22:** naquela ocasião a oferta específica ainda não estava escolhida. A **decisão 23** posterior selecionou **Amazon Bedrock AgentCore**, sem comprovar capacidades, disponibilidade ou elegibilidade.
- **Não objetivos da primeira entrega:** não implementar integrações diretas separadas com **Anthropic** ou **Google Gemini**. A exclusão consta em [NOBJ-004](../non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md); futuras incorporações exigem nova decisão. A exclusão não define antecipadamente quais modelos de fundação a AWS poderá disponibilizar por sua própria plataforma.
- **Homologação e liberação:** ambos são **PLANNED / PROPOSED**, e não `implemented` ou habilitados. Um provedor/serviço só poderá operar quando demonstrar disponibilidade e controles de segurança, autorização, segregação de credenciais, orçamento local, proteção financeira remota **comprovadamente bloqueante** e compatibilidade documental. Se apenas um estiver homologado, sua disponibilidade isolada pode compor entrega parcial identificada, sem declarar a segunda integração entregue.
- **Revisão independente:** permanecem obrigatórios **dois modelos distintos** para geração e avaliação documental. **Dois serviços ou agentes diferentes não comprovam, por si só, dois modelos distintos**; é necessário identificar os modelos/versões subjacentes e validar sua elegibilidade. Falta de segundo modelo autorizado bloqueia o aceite do resultado conforme [OPEN-002](OPEN-002-fidelidade-e-custo.md).
- **Norma e Git:** a seleção de integrações não modifica a referência `main/ProjectCharter/1.0`, nem autoriza branches de sessão, PRs, merge ou commits remotos além das operações normativamente permitidas em `develop`.

**Evidência externa sobre viabilidade da oferta AWS:** a [documentação oficial da AWS sobre o modo de manutenção dos Bedrock Agents Classic](https://docs.aws.amazon.com/bedrock/latest/userguide/agents-classic-maintenance-mode.html) informa que **Agents Classic não aceita novos clientes desde 2026-07-30** e indica **Bedrock AgentCore** como alternativa para novos cenários. Registrar essa restrição como **OBSERVED**, sem inferir que o usuário possui acesso a Classic, nem escolher AgentCore sem decisão do responsável.

**Modalidade agora escolhida na decisão 23:** Amazon Bedrock AgentCore. Ainda será necessário verificar a elegibilidade da conta, a disponibilidade e os controles aplicáveis antes de definir qualquer contrato técnico. Homologação financeira individual e compatibilidade com [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-002](../conditions/CON-002-protecao-credenciais.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md) continuam pendentes. `OPEN-005` permanece em `refinement`.

## Decisão 23 — Amazon Bedrock AgentCore selecionado para a integração AWS (2026-10-21)
**Origem e autorização:** o responsável escolheu expressamente a **opção B — Amazon Bedrock AgentCore**, para o repositório `ahhvahh/doc-arquitetura-app-debian`, com gravação autorizada exclusivamente no Charter da branch `develop`.

**Decisão de escopo vigente:** a primeira entrega mantém **exatamente duas integrações planejadas**: (1) **OpenAI direta** e (2) **Amazon Bedrock AgentCore (AWS)**. A decisão 22 permanece como histórico da seleção de duas integrações, mas sua **indefinição entre Agents Classic e AgentCore fica superada** por esta decisão 23. **Agents Classic não faz parte do planejamento inicial**; integrações diretas adicionais permanecem excluídas conforme [NOBJ-004](../non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md).

**Resultados esperados e limites:**
- O AgentCore é a **modalidade AWS escolhida para futura implementação**, não uma implementação já disponível, autorizada ou homologada. A especialização documental e os prompts continuam subordinados à mesma referência normativa, sem conferir poder de publicação Git a agentes externos.
- A aplicação deve permitir identificar, por execução, a **integração AWS/AgentCore** e, quando demonstrável, o **modelo de linguagem subjacente** efetivamente utilizado. Um agente/orquestrador não comprova sozinho a existência de um segundo modelo distinto exigido na revisão independente da [OPEN-002](OPEN-002-fidelidade-e-custo.md).
- **Viabilidade UNKNOWN:** antes da habilitação, verificar disponibilidade efetiva para a conta e região AWS utilizadas, capacidades necessárias às tarefas documentais, identidade e permissões da execução, tratamento de dados, recursos envolvidos e eventuais custos por operação. Nenhuma garantia de recursos específicos da AWS está homologada por este Charter.
- **Segurança e finanças não flexibilizadas:** cada operação paga, inclusive componentes eventualmente cobrados durante a execução do AgentCore, permanece sujeita à **segregação de credenciais**, autorização individual por chamada, limites locais válidos e **mecanismo remoto de bloqueio financeiro suficiente e comprovadamente efetivo** no escopo de cobrança. Na ausência de evidência, **bloquear**; alertas ou limites meramente informativos não bastam. Relações: [CON-002](../conditions/CON-002-protecao-credenciais.md), [CON-004](../conditions/CON-004-orcamento-obrigatorio.md), [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md).
- **Governança inalterada:** seleção do AgentCore não cria autorização para criar branches de sessão, PRs, merges, nem editar a referência `main/ProjectCharter/1.0`. A integração OpenAI continua prevista com seus quatro projetos separados por especialização e verificações individuais.

**Próxima ação verificável:** realizar a investigação separada de **viabilidade funcional e financeira do AgentCore** e de sua aptidão para as quatro especializações; registrar evidências sem inventar configuração técnica no Charter. A homologação e a implementação permanecem **UNKNOWN**. Manter `OPEN-005` em `refinement`.
