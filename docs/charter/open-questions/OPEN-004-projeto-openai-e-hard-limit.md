# OPEN-004 — Associação de agentes a projetos OpenAI e controle de hard limits

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão de governança financeira parcialmente definida: segregação de projetos OpenAI e política conservadora de bloqueio aprovadas; comprovação dos mecanismos remotos continua pendente.

## Questão
**Decisões aprovadas:** quatro projetos OpenAI independentes, um por especialização, e **bloqueio de qualquer inferência paga quando não houver proteção remota suficiente e comprovada**. Resta demonstrar, por serviço e projeto aplicável, que os controles disponíveis conseguem interromper o consumo remunerado de modo compatível com o risco financeiro autorizado, sem comprometer o sigilo das credenciais.

## Contexto e origem
O responsável exige limites locais por agente e controles também na conta OpenAI. A existência e a eficácia de mecanismos de bloqueio financeiro no projeto/organização ou na conta exigem comprovação, e não equivalem automaticamente a uma cota bloqueante por agente. Não pressupor que alertas e orçamentos administrativos sejam hard limits efetivos.

## Impacto
**Bloqueadora para habilitar pagamentos em produção:** sim, para cada provedor e projeto aplicável enquanto não forem comprovadas a segregação exigida, as credenciais, a eficácia da proteção remota e os limites locais. A escolha da política está concluída; sua comprovação operacional não. Não impede testes simulados ou operações comprovadamente sem cobrança.

## Próxima ação verificável
Conferir, para cada um dos quatro projetos OpenAI e para a integração de agentes Amazon Bedrock (AWS), se há mecanismo remoto suficiente e efetivamente bloqueante, sua unidade de abrangência (inclusive recursos AWS cobrados), permissões de verificação e evidências de ativação. Anthropic e Gemini não são integrações diretas do escopo inicial conforme decisão 22 em [OPEN-005](OPEN-005-estrategia-de-provedores-e-modelos.md). Sem tal comprovação, manter inferências pagas bloqueadas. Especificações técnicas de verificação cabem a Scope e Implementation.

## Destino previsto
- [CON-004](../conditions/CON-004-orcamento-obrigatorio.md)
- [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)

## Proposta de regularização — Decisão 04 (2026-10-21)

**Origem:** decisão explícita do responsável — opção A, quatro projetos OpenAI independentes.
**Resultado da triagem:** proposta aceita para regularização do Charter, sem criar projetos, chaves ou limites na plataforma.

- Organizar a integração OpenAI em quatro projetos distintos: `charter`, `scope`, `implementation` e `source` (correspondente à etapa normativa `src`).
- Cada especialidade deve utilizar a identidade do projeto OpenAI e sua própria credencial de projeto, com segregação de permissões, consumo, rastreabilidade e orçamento. A aplicação não deve misturar credenciais ou contexto documental de especialidades distintas.
- Manter limites locais independentes por operação, por especialidade e um limite global da aplicação, além dos controles de orçamento que estiverem efetivamente disponíveis na plataforma para cada projeto.
- A inexistência de um teto remoto efetivamente bloqueante não pode ser ocultada nem tratada como garantia de cobrança máxima. Confirmar as capacidades reais do provedor e adotar a política local de falha fechada prevista em [CON-004](../conditions/CON-004-orcamento-obrigatorio.md).
- Antes de permitir inferência paga em produção, verificar o vínculo correto entre especialidade, projeto, credencial, orçamento autorizado, estimativa de consumo e mecanismo de acompanhamento. Sem essa verificação, bloquear a operação remunerada.
- A criação de projetos, a definição de orçamentos remotos e o provisionamento de credenciais serão ações administrativas explícitas e separadas da execução de análises documentais; nenhuma credencial será armazenada em Git, DSL ou logs.
- O controle e a auditoria das despesas deverão permitir consulta por especialidade e também consolidar o total utilizado na aplicação, incluindo provedores adicionais quando forem habilitados.

**Relações canônicas:** [OBJ-005 — Inferências com custo previsível e limitado](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [CON-004 — Orçamentos obrigatórios em duas camadas](../conditions/CON-004-orcamento-obrigatorio.md).

**Pendências:** validar no serviço OpenAI as garantias, características e permissões reais dos controles remotos de orçamento; especificar os contratos de identidade, autorização e verificação em Scope/Implementation após o gate. A OPEN permanece em `refinement` e continua bloqueadora para habilitar inferências pagas em produção.

## Decisão 09 — Bloqueio obrigatório sem proteção remota comprovada (2026-10-21)
**Origem:** opção A escolhida explicitamente pelo responsável e autorizada para registro no Charter da `develop`.

**Resultado da triagem:** política de proteção financeira aprovada; **não houve validação de mecanismos de cobrança ou bloqueio das APIs**.

- Para **todo provedor/serviço autorizado a executar operações pagas**, incluindo **OpenAI e agentes Amazon Bedrock (AWS)** no escopo vigente, proibir a inferência paga enquanto não existir **proteção remota suficiente, efetiva e verificável** contra consumo superior ao teto autorizado, além de controles locais válidos de orçamento por operação, especialização e global.
- **Alertas de gasto, limites apenas informativos e previsões locais não satisfazem isoladamente** a condição de proteção remota bloqueante.
- Preservar a **decisão 04:** OpenAI terá quatro projetos separados, associados a `charter`, `scope`, `implementation` e `source`, com credenciais e controle de consumo segregados. O requisito de eficácia do bloqueio remoto deve ser verificado **em cada projeto OpenAI habilitado**; não presumir que controles da conta sejam independentes por projeto.
- Outros provedores não são obrigados a reproduzir a organização administrativa da OpenAI; precisam, porém, demonstrar mecanismo remoto de proteção compatível com a mesma política de bloqueio.
- Sem comprovação suficiente, configuração, credencial ou limite válido, **negar a chamada paga**. Não conceder exceção pela aceitação do risco, nem apenas pela autorização de gasto do operador ou pela existência de limite local.
- Admitir **simulações e testes comprovadamente sem cobrança**, desde que não acionem acidentalmente uma operação paga e respeitem os controles de credenciais e dados.
- A decisão não constitui garantia absoluta sobre faturas de serviços externos nem comprova que os mecanismos remotos existem. A proteção só poderá ser declarada após evidência verificável da sua eficácia e abrangência.
- Cada verificação deve identificar **provedor, projeto/escopo aplicável, natureza do mecanismo remoto, evidência, validade e resultado**, permitindo suspender a inferência paga quando o controle perder efetividade.

**Destinação canônica:** [OBJ-005 — Orçamento](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [CON-004 — Restrições financeiras](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-005 — Provedores previstos](OPEN-005-estrategia-de-provedores-e-modelos.md).

**Pendências:** homologação técnica e administrativa dos controles em cada provedor/projeto e validação do custo/tokens/câmbio. A OPEN permanece em `refinement`, bloqueando habilitação de inferências pagas enquanto essas evidências forem desconhecidas.

## Escopo financeiro atualizado para OpenAI e agentes AWS — decisão 22
A escolha do responsável passou a prever **duas integrações iniciais**, conforme [OPEN-005](OPEN-005-estrategia-de-provedores-e-modelos.md). A política **não** foi flexibilizada: os quatro projetos OpenAI continuam segregados, e o uso de **agentes Amazon Bedrock (AWS)** também exige comprovação de proteção financeira remota **efetiva e bloqueante** sobre os custos relevantes, além de orçamento local e autorização explícita por chamada. Um serviço de orquestração pode envolver diferentes cobranças, que não devem ser omitidas.

O registro de Anthropic/Gemini nas decisões históricas não obriga implementar integrações diretas com eles agora; estão fora da primeira entrega conforme [NOBJ-004](../non-objectives/NOBJ-004-nao-integrar-outros-provedores-na-primeira-entrega.md). Nenhuma modalidade AWS ou mecanismo de controle foi homologado. Na dúvida sobre controle remoto, **bloquear operações pagas**.
