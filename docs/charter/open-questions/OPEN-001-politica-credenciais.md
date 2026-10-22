# OPEN-001 — Política de credenciais e usuários

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Diretriz funcional aprovada para a primeira entrega: arquivos administrativos de segredos protegidos pelo Linux, entrega restrita ao serviço por `systemd` e execução com identidade exclusiva e privilégios mínimos, sem cofre dedicado obrigatório. Comprovação efetiva das permissões, da continuidade e dos controles financeiros permanece pendente.
## Questão
**Política vigente (2026-10-22, Segurança — decisão 1, opção A):** administrador Linux autorizado ou superusuário configura, substitui e revoga segredos mantidos em arquivos administrativos protegidos e separados da configuração geral; o serviço com identidade exclusiva e privilégios mínimos os recebe por mecanismo restrito de `systemd`, sem cofre dedicado obrigatório na primeira entrega. Usuários comuns e guardião não leem nem administram segredos. A execução independe da sessão do solicitante; grupo Linux habilita somente uso e consulta via aplicação, sem banco direto. Cada chamada paga requer autorização financeira individual, e falhas de credencial preservam tarefas bloqueadas até liberação autorizada. As decisões anteriores de cofre ficam preservadas como histórico, parcialmente superadas quanto ao armazenamento obrigatório.
## Contexto
A aplicação existente é executada em contexto de usuário; o daemon isolado está descrito, mas ainda não está implementado. Passar a chave na linha de comando pode expô-la em argumentos e histórico.
## Impacto
**Bloqueadora:** sim; afeta a segurança e a autorização da integração.
## Próxima ação
Comprovar que somente o administrador gerencia arquivos secretos, que a configuração geral não contém segredos, que o serviço não privilegiado recebe credenciais por `systemd` sem ler os arquivos de origem e que usuários comuns/guardiões não extraem chaves; testar acesso por grupo Linux via aplicação, ausência de acesso direto ao banco, execução sem sessão interativa, pré-autorização financeira por chamada, revogação e retomada autorizada de tarefas bloqueadas. Detalhes de unidade, caminhos, permissões, rotação e formatos pertencem às etapas responsáveis.
## Destino previsto
Scope de administração/credenciais, após decisão do Charter.

## Proposta de regularização — Decisão 03 (2026-10-21)

**Nota de vigência (2026-10-22):** o cofre individual por usuário e sua administração por operadores comuns foram **superados** pela política administrativa registrada adiante. O registro original é preservado como evidência histórica; continuam válidos sigilo, entrada oculta para administradores, revogação e proibição de exposição de segredos.

**Origem:** decisão explícita do responsável — opção A, cofre de credenciais Linux por usuário.
**Resultado da triagem:** proposta aceita para regularização do Charter. A definição não equivale à implantação dos controles de segurança.

- As credenciais necessárias aos provedores remotos de inferência e às integrações autorizadas devem ser protegidas no cofre de segredos do ambiente Linux, segregadas por identidade de usuário e acessíveis somente ao contexto autorizado da aplicação.
- O cadastro ou a atualização de credenciais deve solicitar entrada oculta, sem incluir o valor da chave em argumentos da CLI ou no histórico do shell. A opção legada `--chave`, se mantida, não poderá transportar o valor do segredo em texto.
- Nunca gravar segredos em Markdown, DSL, repositórios Git, arquivos de configuração legíveis, relatórios, diffs, logs ou saídas do terminal. Armazenar nos metadados operacionais somente referências não sensíveis às credenciais.
- A aplicação deverá recuperar o segredo exclusivamente quando necessário para a operação autorizada e impedir leitura cruzada entre usuários. Sem cofre acessível, identidade autorizada ou credencial válida, bloquear a operação dependente em vez de usar armazenamento em texto puro.
- Prever substituição, revogação e exclusão de credenciais com confirmação e sem divulgação do seu conteúdo. A necessidade de executar AWS CLI não deve ser pressuposta para o cliente de inferência.
- O futuro serviço Linux isolado continua sendo uma evolução independente: não presumir que o cofre de um usuário possa ser acessado por outro usuário ou pelo daemon. Os contratos de identidade, permissão e transferência de autorização permanecem para Scope/Implementation.

**Relação canônica:** [CON-002 — Proteção de credenciais](../conditions/CON-002-protecao-credenciais.md).

**Pendências e gate:** especificar e testar a integração com o cofre, permissões efetivas por usuário, compatibilidade da CLI, comportamento em ambiente sem sessão de cofre, rotação e a eventual operação em daemon. Manter a OPEN em `refinement` e a passagem Charter → Scope bloqueada enquanto as condições necessárias não forem regularizadas.

## Decisão 11 — Serviço coordenador com execução isolada por usuário (2026-10-21)

**Nota de vigência (2026-10-22):** a execução obrigatória com credenciais pessoais e identidade de processo do solicitante foi **superada**. A responsabilidade da autorização, a segregação dos dados de usuários, os bloqueios de despesas e a auditoria permanecem exigíveis; a política substituta está registrada ao final desta OPEN.
**Origem verificável:** o responsável selecionou explicitamente a **opção C — serviço com execução isolada por usuário**, autorizando sua documentação apenas no Charter da branch `develop`.

**Resultado da triagem:** decisão de comportamento **aceita como intenção de produto**; não há comprovação de implantação de daemon, isolamento de processo ou acesso real ao cofre.

- O serviço Linux pretendido **organiza e acompanha tarefas**, sem assumir a identidade do usuário nem obter, concentrar, registrar ou compartilhar suas credenciais de inferência.
- Cada inferência deve ocorrer exclusivamente no **contexto autorizado do usuário responsável pela tarefa**, com acesso apenas às credenciais desse mesmo usuário e do provedor/escopo permitido. Identidade do serviço não equivale a autorização para utilizar o cofre pessoal do operador.
- A escolha mantém a **decisão 03**: segredos de OpenAI, Anthropic, Google Gemini e GitHub devem ser protegidos em cofre Linux **segregado por usuário**, com entrada oculta e sem exposição em argumentos da CLI, Markdown, DSL, logs, Git, arquivos legíveis ou exportações.
- Não executar tarefas de um usuário com credenciais, autorização financeira, identidade de faturamento, contexto documental ou resultados privados de outro usuário.
- Caso não seja possível confirmar o usuário efetivo, a autorização, o cofre acessível ou o isolamento, **bloquear a inferência dependente**, sem executar sob identidade privilegiada nem usar credencial de serviço como substituto.
- Cada solicitação remunerada, incluindo a segunda avaliação LLM, continua exigindo os controles locais e a **proteção remota financeira comprovada** de [CON-004](../conditions/CON-004-orcamento-obrigatorio.md) e [OPEN-004](OPEN-004-projeto-openai-e-hard-limit.md). A escolha do usuário não substitui a verificação individual dos **quatro projetos OpenAI por especialização**.
- A trilha de auditoria deve permitir relacionar a solicitação à identidade autorizada, à especialização, ao provedor/projeto, à autorização e ao resultado, **sem gravar segredos** nem expor dados de outro operador.
- O mecanismo concreto de delegação/execução, permissões do ambiente Linux, acesso ao cofre sem sessão interativa, recuperação de tarefas e revogação pertencem a Scope/Implementation e exigem evidências e testes específicos.

**Critérios verificáveis a detalhar:** impossibilidade de leitura cruzada entre usuários, falha fechada sem cofre/identidade autorizada, preservação da atribuição de custo ao contexto correto, revogação efetiva e ausência de credenciais em logs e argumentos.

**Pendências:** validar a execução isolada em ambiente real e sem sessão ativa, controle de identidade do serviço, mapeamento de usuários/autorização, integração ao cofre e proteção de credenciais nos provedores. Permanecer em `refinement`; esta decisão **não declara serviço implementado nem libera inferência paga**.

**Relações canônicas:** [CON-002 — Proteção de credenciais](../conditions/CON-002-protecao-credenciais.md), [OBJ-005 — Política financeira](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [PCH-001 — Contexto e participantes](../project-context/PCH-001-contexto-integracao-openai-dsl.md).

## Política vigente — cofre administrativo e serviço autônomo (2026-10-22)
**Nota de atualização da Segurança — decisão 1, opção A (2026-10-22):** esta subseção registra a escolha anterior de administração centralizada. A expressão “cofre” **não obriga mais um produto ou serviço de cofre dedicado** na primeira entrega; a política de armazenamento vigente é a de arquivos administrativos protegidos com entrega restrita via `systemd`, definida ao final desta OPEN. Continuam válidas identidade exclusiva, privilégios mínimos, sigilo, bloqueio e auditoria.

**Origem:** o responsável escolheu explicitamente as cinco alternativas: identidade de serviço **A**, guardião **A**, acesso aos resultados **B**, pré-autorização financeira **A** e tratamento de falha de credencial **B**. Decisão de produto **DOCUMENTED**, não `IMPLEMENTED`.

- **Administração:** somente administrador autorizado ou superusuário cadastra, substitui e revoga credenciais após a instalação. A entrada do segredo não deverá expor seu valor. Usuário comum não visualiza, cadastra, altera, extrai nem recebe chaves.
- **Identidade de execução:** serviço da aplicação opera com identidade exclusiva e acesso mínimo ao cofre e aos provedores necessários. A necessidade de privilégios administrativos na configuração **não autoriza** executar permanentemente como superusuário sem justificativa. O serviço processa solicitações já autorizadas mesmo sem sessão interativa do solicitante, mantendo autoria e permissão identificáveis.
- **Guardião opcional:** observa pendências e resultados visíveis ao usuário autorizado; não lê segredos, não executa inferências e não concede autorizações financeiras. Ausência do guardião não impede o serviço de registrar ou processar solicitações elegíveis.
- **Acesso a resultados — revisão de 2026-10-22:** a regra anterior de concessões individuais foi substituída pela autorização coletiva de consulta dos membros do grupo Linux habilitado, exclusivamente por intermédio da aplicação e do serviço. Permanecem distintas as permissões para consultar, excluir, administrar credenciais, autorizar gastos ou executar operações Git.
- **Custos:** chamadas remuneradas planejadas podem ser **individualmente pré-autorizadas** antes de enfileirar a tarefa, permitindo execução posterior sem usuário conectado. Cada chamada, inclusive revisão independente e nova correção, exige autorização específica e validação de custos, orçamento e proteção remota efetivamente bloqueante **no momento da execução**. Nenhuma autorização da tarefa libera despesa adicional.
- **Credencial indisponível:** segredo ausente, inválido, revogado, inacessível ou permissão insuficiente mantém a operação dependente **bloqueada**, com solicitação, justificativa e histórico preservados. Após regularização administrativa, a retomada exige **liberação autorizada** e revalidação das permissões, do custo e da integridade, sem execução automática nem troca silenciosa de chave.

**Superações expressas:** da decisão 03, somente a segregação do cofre por usuário e o cadastro pelo usuário comum; da decisão 11, somente a exigência de inferência em identidade/processo do solicitante e a vedação geral ao serviço executor consumir o cofre administrado. Preservam-se controle de acesso, privacidade por solicitante, responsabilidade financeira individual, rastreabilidade, ausência de segredos em registros e as proibições de operações Git incompatíveis com a norma.

**Critérios verificáveis futuros:** usuário comum/guardião não extrai segredos; administrador autorizado gerencia credenciais sem vazamento; serviço executa tarefa previamente autorizada sem login do solicitante; leitura fora do escopo é negada; cobrança sem autorização individual ou proteção financeira remota comprovada é negada; falha de cofre mantém pendência íntegra; retomada exige liberação e revalidação explícitas.

**Pendências:** comprovar acessos e execução autônoma segura no Linux, revogação, proteção do cofre, concessão de consulta, continuidade, fila de pendências e bloqueios financeiros. Permanecer em `refinement`; regras técnicas somente após o gate em Scope/Implementation.

**Relações:** [CON-002 — Credenciais](../conditions/CON-002-protecao-credenciais.md), [PCH-001 — Participantes](../project-context/PCH-001-contexto-integracao-openai-dsl.md), [CON-005 — Banco local](../conditions/CON-005-persistencia-local-de-inferencias.md), [OBJ-005 — Custos](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md), [OBJ-008 — Histórico](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-006 — Retenção](OPEN-006-retencao-e-protecao-do-historico.md).

## Política vigente complementar — acesso Linux e banco exclusivo do serviço (2026-10-22)

**Origem verificável:** o responsável aprovou autenticação pelo próprio Linux, associação a grupo Linux autorizado e consulta compartilhada somente através da aplicação instalada, sem acesso direto dos usuários comuns ao banco. Esta decisão **substitui somente a exigência anterior de concessões individuais de leitura por tarefa ou projeto**; não revoga políticas de retenção, sigilo, cofre, custos ou Git.

- O servidor é administrado. **Identidade Linux autenticada e participação no grupo autorizado** constituem a condição de acesso funcional à aplicação; a administração do servidor concede/revoga participação no grupo. A aplicação não criará cadastro paralelo de usuários, senhas ou autorizações individuais de leitura.
- Os membros desse grupo podem **consultar via aplicação** tarefas, estados, resultados e histórico disponibilizados pelo serviço, inclusive registros de outros membros. A identidade Linux do solicitante continua vinculada aos eventos para auditoria e responsabilização.
- A aplicação cliente **não concede conexão direta ao banco** e não distribui credenciais do banco. A identidade técnica autorizada do serviço efetua consultas e gravações; usuários comuns não possuem permissões diretas sobre os dados persistidos.
- A participação no grupo **não concede** acesso ao cofre administrativo, autorização financeira própria ou de terceiros, exclusão de histórico, administração do serviço, operações Git excepcionais ou aprovação de publicação. Cada chamada paga segue com autorização financeira individual validada e controles remotos comprovados.
- O guardião opcional pode acompanhar resultados permitidos **somente por meio da aplicação/serviço**, sem acesso direto ao banco ou ao cofre.
- O **banco local permanece histórico canônico**: esta decisão não autoriza migração para armazenamento remoto nem escolhe banco, esquema, transporte ou mecanismo de execução.

**Critérios futuros de comprovação:** membro do grupo utiliza a aplicação e consulta resultados compartilhados; não membro é rejeitado; usuário comum/guardião não obtém acesso direto ao banco ou segredos; remoção do grupo elimina acesso funcional; autoria, consulta, custos e exclusões permanecem auditáveis; ausência de autorização financeira individual bloqueia operação paga.

**Situação:** `DOCUMENTED`, não `IMPLEMENTED`; controle efetivo de grupo Linux, exposição do serviço, privilégios e proteção da base dependem de validação posterior. A `OPEN-001` permanece em `refinement`.

**Relações:** [CON-005 — Banco local](../conditions/CON-005-persistencia-local-de-inferencias.md), [OBJ-008 — Histórico](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-006 — Retenção e acesso](OPEN-006-retencao-e-protecao-do-historico.md), [PCH-001 — Atores](../project-context/PCH-001-contexto-integracao-openai-dsl.md), [CON-002 — Segredos](../conditions/CON-002-protecao-credenciais.md) e [OBJ-005 — Controle financeiro](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md).

## Decisão vigente — Segurança 1: arquivos Linux protegidos e entrega por systemd (2026-10-22)

**Origem:** opção **A** selecionada e autorizada pelo responsável para a primeira entrega do `doc-arquitetura-app-debian`. Trata-se de política de produto `DOCUMENTED`, não comprovação de implementação.

- **Custódia administrativa:** somente usuário Linux com privilégios administrativos autorizados poderá cadastrar, substituir ou revogar chaves e senhas de OpenAI, AWS/Amazon Bedrock, GitHub e integrações autorizadas. A operação administrativa deverá proteger a entrada e jamais exibir o valor secreto.
- **Proteção simples sem cofre obrigatório:** segredos permanecerão em **arquivos administrativos protegidos pelo Linux**, separados da configuração geral da aplicação. **Não é obrigatório implantar ou administrar cofre externo/dedicado** na primeira entrega. A escolha não autoriza arquivos legíveis por usuários comuns.
- **Entrega exclusiva:** a identidade técnica exclusiva do serviço deverá receber apenas os segredos necessários às operações elegíveis **por mecanismo restrito do `systemd`**, sem ganhar acesso direto de leitura aos arquivos administrativos de origem. A administração pode requerer privilégios elevados; o serviço contínuo não poderá depender de rodar como `root`.
- **Restrições de divulgação:** usuário comum ou guardião opcional não pode ler, cadastrar, exportar ou alterar segredos. Chaves/senhas não irão para argumentos da CLI, histórico do shell, logs, DSL, Markdown, Git, relatórios, exportações, arquivos de configuração comum ou banco de tarefas.
- **Isolamento já aprovado:** autenticação e participação no grupo Linux permitem usar a aplicação e consultar o histórico compartilhado **somente por meio da aplicação/serviço**. Nem associação ao grupo nem acesso de leitura conferem acesso direto ao banco, à origem dos segredos, à administração ou a gastos de terceiros.
- **Bloqueio seguro:** segredo ausente, revogado, inacessível ou inválido mantém a operação dependente **bloqueada**, com tarefa e histórico preservados. Depois da correção administrativa, a retomada exige liberação autorizada e revalidação de condições, sem credencial substituta automática.
- **Financeiro e Git inalterados:** continuam válidas pré-autorização **por chamada paga**, orçamento local, proteção remota bloqueante comprovada e restrição normativa de gravação exclusiva em `develop`, sem merges ou operações Git não permitidas.

**Registro de superação:** a decisão 03 (cofre individual) e o armazenamento em “cofre administrado” da política intermediária permanecem documentados como histórico. Ficam superadas **somente as exigências de cofre individual/dedicado e a execução sob credenciais pessoais**; **não** os requisitos de proteção, custódia, privilégios mínimos, autoria, sigilo, auditoria e orçamento. A decisão de acesso compartilhado por grupo Linux permanece íntegra.

**Evidências requeridas futuramente:** verificar permissões negativas para usuário comum/guardião e para serviço sobre arquivos de origem; entrega restrita via `systemd`; separação de configuração/segredos; identidade efetiva não-root; revogação e rotação sem vazamento; bloqueio de chamada paga sem autorização ou controles suficientes; continuidade segura das tarefas.

**Status da decisão:** `DOCUMENTED`. Manter o badge documental `refinement` até os critérios pertinentes serem verificados, sem afirmar `IMPLEMENTED`. Caminhos, permissões numéricas e configuração concreta do `systemd` pertencem a Scope/Implementation.

**Relações:** [CON-002 — Proteção de credenciais](../conditions/CON-002-protecao-credenciais.md), [CON-005 — Banco local](../conditions/CON-005-persistencia-local-de-inferencias.md), [OBJ-005 — Controle financeiro](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [OPEN-006 — Histórico](OPEN-006-retencao-e-protecao-do-historico.md).
