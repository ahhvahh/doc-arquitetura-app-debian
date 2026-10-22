# OBJ-001 — Agentes documentais especializados

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
O uso de regras indistintas pode misturar responsabilidades das etapas e gerar propostas inconsistentes.
## Resultado esperado
A aplicação oferece quatro especializações independentes — Charter, Scope, Implementation e Source — aderentes à norma versionada, com identificação da especialização e baseline utilizada.
## Critérios verificáveis
- A interface apresenta exatamente as quatro especializações solicitadas.
- Cada análise indica a especialização, a referência normativa persistida no agente e o commit-base.
- A atualização das instruções normativas de cada agente ocorre apenas em comando explícito; a inferência comum não as reenvia.
- Uma especialização não executa silenciosamente a responsabilidade de outra.
## Dependências
- [CON-001](../conditions/CON-001-referencia-normativa-main.md).
- [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Interface simplificada e seleção de prompts — decisão do responsável
- Quatro entradas principais: `charter`, `scope`, `implementation` e `source` (`source` corresponde à etapa normativa `src`).
- Uma única árvore seleciona `ETAPA → OPERAÇÃO → AÇÃO → PROMPT especializado`; reutilizar o catálogo existente sem manter árvores paralelas.
- Cada etapa usa um agente especializado e um acervo normativo próprio, identificado por SHA; o projeto-alvo analisado é contexto transitório e nunca substitui a fonte normativa.
- O agente persistente deve consultar as regras relevantes recuperadas da base DSL e obedecer ao prompt escolhido. A aplicação controla o acesso à recuperação quando a API remota não permitir vincular diretamente a ferramenta ao agente.
- A configuração de agente deve ser um fluxo explícito integrado: selecionar a etapa, atualizar a referência normativa, preparar DSL, sincronizar acervo remoto, confirmar indexação, criar/atualizar o agente, registrar IDs e testar recuperação. O modo padrão é prévia; a publicação exige `--apply`.
- A interface proposta `agent setup <charter|scope|implementation|source|all> --model MODELO [--apply]` não equivale a comando já implementado.
- A configuração do agente não executa tarefas do projeto-alvo nem autoriza publicação de alterações.

## Extensão aprovada — seleção de modelos independente da especialização (2026-10-21)
As quatro especializações Charter, Scope, Implementation e Source permanecem separadas. A aplicação deverá permitir escolher modelos simples ou avançados sem duplicar o catálogo normativo por modelo. O modelo/provedor é identificado em cada análise; a referência normativa da especialização permanece rastreável por SHA.

A administração de agentes OpenAI já prevista não determina exclusividade de provedor. **À época da decisão 08**, a lista incluía **OpenAI, Anthropic e Google Gemini**; essa seleção foi **substituída pela decisão 22**: a primeira entrega agora prevê **OpenAI e agentes Amazon Bedrock (AWS)**, condicionados a verificações individuais em [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md). Modelos locais estão fora da primeira entrega. A comparabilidade esperada está em [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md); não se declara equivalência entre recursos administrativos de provedores distintos.

## Definição adicional — múltiplos provedores e prompts por sessão (2026-10-21)
A escolha aprovada é trabalhar com **provedores remotos diferentes**, não somente modelos diferentes da OpenAI. O operador poderá selecionar novo prompt ou reutilizar prompt do catálogo da etapa durante uma sessão, sem duplicar os objetivos normativos dos agentes.

**Verificação esperada:** cada execução associa explicitamente etapa, prompt selecionado (com identificação de versão/origem), provedor e modelo, conservando o histórico entre inferências da mesma sessão.

Relações: [OPEN-005 — modalidade remota](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e [OBJ-003 — continuidade das alterações](OBJ-003-propostas-de-alteracoes-verificaveis.md).

## Modelo e prompt escolhidos por solicitação — decisão 17
Durante uma sessão de documentação aberta, o operador não fica preso a um único LLM ou prompt. **Cada inferência** permite selecionar um **modelo/provedor previamente elegível** e um **prompt do catálogo ou novo/editado**, sempre associado à especialização normativa da sessão. A referência `main/ProjectCharter/1.0` e as restrições aplicáveis não podem ser alteradas pelo prompt do operador.

**Evidência de produto:** cada inferência registra origem, texto/versão do prompt, especialização, provedor/modelo/versão, identidade do usuário autorizado, janela de contexto solicitada e efetivamente admissível, custo autorizado e baselines; a nova escolha não apaga as configurações usadas nas inferências anteriores.

A seleção de modelo não habilita automaticamente chamadas pagas: mantêm-se [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md), [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md) e a segregação de credenciais. Relações: [OBJ-002](OBJ-002-contexto-compacto-rastreavel.md), [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md) e [OPEN-002 — decisão 17](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Integrações inicialmente previstas: OpenAI e agentes AWS — decisão 22
As quatro especializações Charter, Scope, Implementation e Source continuarão independentes da plataforma utilizada. O **escopo atual da primeira entrega** prevê **duas integrações a desenvolver**, uma direta **OpenAI** e outra para **agentes Amazon Bedrock (AWS)**, sem afirmar que ambas já foram implementadas, homologadas ou habilitadas. A modalidade AWS foi definida como **Amazon Bedrock AgentCore** na [OPEN-005 — decisão 23](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md), mas sua viabilidade e homologação permanecem pendentes.

A seleção de modelo/contexto/prompt por inferência e o histórico de sessões devem permanecer rastreáveis, mas capacidades distintas de gerência de agentes, autenticação e acesso normativo não podem ser supostas equivalentes. **Anthropic e Google Gemini como integrações diretas** foram retirados da entrega inicial, conforme [NOBJ-004](../non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md). Contratos detalhados e validação de compatibilidade pertencem às etapas posteriores.

## Padrões persistentes de geração e revisão antes da sessão — decisão 24 (2026-10-21)
**Resultado funcional aprovado:** antes de iniciar uma sessão documental, a aplicação deverá disponibilizar uma **tela de configuração prévia** e manter uma **configuração persistente, editável e reutilizável**, organizada conceitualmente por **especialização** (`Charter`, `Scope`, `Implementation`, `Source`), **operação/tipo de análise** e **papel da inferência** (`gerador` ou `revisor`). A aplicação deverá carregar esses padrões ao abrir uma sessão, depois de o operador escolher o repositório GitHub e o nível documental.

**Para cada papel e operação**, o operador poderá estabelecer previamente o **motor/integração autorizado** (OpenAI direta ou Amazon Bedrock AgentCore), a **identificação do modelo elegível** e o **prompt de tarefa padrão** identificado por origem e versão. O modelo gerador realiza a proposta documental; o modelo revisor fornece avaliação crítica independente. Quando um papel estiver associado ao AgentCore, **a identidade do agente/motor não substitui a identificação verificável do modelo de linguagem efetivamente utilizado**.

**Critérios verificáveis:**
- É possível recuperar e apresentar os padrões registrados para cada combinação de especialização, operação e papel, distinguindo explicitamente **geração** e **revisão**; uma especialização/operação não herda silenciosamente o modelo de outra.
- Ao abrir a sessão, o operador escolhe **repositório-alvo e nível documental** e recebe os padrões correspondentes; o registro histórico distingue a versão dessa configuração dos valores realmente autorizados e usados em cada inferência.
- **Durante a sessão**, o operador continua podendo **alterar explicitamente modelo/motor elegível, tamanho de contexto e prompt a cada inferência**, conforme a decisão 17. A configuração prévia é o **valor inicial sugerido**, não um bloqueio permanente nem uma ordem para substituir automaticamente modelos ou prompts.
- O sistema deve mostrar condição de elegibilidade e indisponibilidades de cada papel **antes de prometer uma análise validada**. **Não considerar válido** um resultado sem segundo LLM distinto e homologado. Um rascunho, se permitido em uma operação autorizada, não equivale a aceite ou commit de resultado validado.
- A persistência de preferências **não contém credenciais ou segredos**; identidades, cofres, autorização financeira individual, orçamento, proteção remota bloqueante comprovada e integridade normativa permanecem exigências independentes. O prompt configurável da tarefa não pode substituir os prompts normativos obrigatórios.

**Limite de etapa:** o requisito de manter preferências em **arquivo de configuração** e de apresentar **uma tela anterior à sessão** é uma intenção funcional aceita; **caminho, sintaxe, estrutura do arquivo, tecnologia, comportamento visual e contratos** serão especificados em Scope/Implementation, não neste Charter. Relações: [OBJ-002](OBJ-002-contexto-compacto-rastreavel.md), [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md), [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md) e [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md).
