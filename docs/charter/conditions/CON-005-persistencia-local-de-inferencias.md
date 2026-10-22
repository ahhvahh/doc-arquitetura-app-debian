# CON-005 — Persistência local e proteção do histórico de inferências

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
Solicitações, prompts, respostas, previsões, medições e resultados de inferência necessários à comparação futura devem ser persistidos em banco de dados local sob controle do operador, sem depender de armazenamento remoto como histórico canônico. As informações não podem expor credenciais ou segredos de outros projetos.

## Origem verificável
Exigência explícita do responsável em 2026-10-21: banco local para conservar solicitações, tokens previstos e realizados de entrada e saída, custos previstos e realizados quando disponíveis, prompts e resultados para análise posterior.

## Motivo e impacto
Assegura auditoria e comparabilidade independentemente da retenção do provedor externo. Os conteúdos podem conter dados confidenciais de repositórios analisados, exigindo controle de acesso e tratamento seguro.

## Forma de verificação
- O histórico permanece acessível localmente após encerramento da inferência e sem acesso à API externa.
- Tokens e custos previstos/realizados indicam sua origem; dados desconhecidos não são preenchidos artificialmente.
- Informações sensíveis não são expostas a usuários não autorizados, logs e exportações indevidas.
- Falhas de gravação e lacunas de evidência são identificáveis e não são tratadas como execução auditada com sucesso.

## Itens afetados
- [OBJ-008 — Histórico](../objectives/OBJ-008-historico-local-de-inferencias.md)
- [CON-002 — Credenciais](CON-002-protecao-credenciais.md)
- [OPEN-006 — Retenção](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md)

## Condição de revisão
Somente por decisão explícita do responsável quanto à necessidade de histórico local ou aos limites de privacidade. O mecanismo técnico de armazenamento pertence ao Scope.

## Recuperação interoperável das métricas
O histórico local deve preservar informações suficientes para disponibilizar as métricas como **objetos JSON recuperáveis** por outras aplicações autorizadas. Essa condição não obriga que toda a persistência interna use um único campo JSON e não define banco de dados nem esquema físico.

A consulta, extração ou reconstrução de relatórios fora da aplicação original deverá conservar identidade de sessão/tentativa, origem dos números e distinção entre previsão, medição e cobrança. O acesso deve observar [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md).

## Política de retenção, autorização e descarte (2026-10-21)
**Origem:** opção C escolhida em [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md).

A aplicação deve conservar prompts, respostas e métricas JSON no banco local **até exclusão manual**, sem eliminação automática por idade. Consulta e exportação são restritas ao operador responsável e a consumidores expressamente autorizados; acesso de leitura não concede automaticamente autorização de exclusão.

Credenciais e segredos não devem ser persistidos nem expostos em registros, relatórios ou exportações. A exclusão manual autorizada deve gerar evidências de autoria, data, identificação e alcance dos registros afetados e resultado, sem reproduzir dados apagados. Backups e cópias sob controle da aplicação devem ter tratamento de descarte identificável; cópias externas não podem ser declaradas eliminadas sem comprovação.

**Verificação:** ausência de expiração automática, bloqueio de leitura e exclusão indevidas, prevenção de vazamento de segredos e comprovação do alcance de cada descarte autorizado. Contratos técnicos são assuntos de Scope/Implementation.

## Restrição complementar — solicitações e pendências do serviço autônomo (2026-10-22)

**Resultado esperado:** o banco local deverá conservar como registro canônico as **solicitações, autorizações, pendências, estados, resultados e eventos de auditoria** necessários para acompanhar tarefas executadas pelo serviço mesmo após o encerramento do terminal ou da sessão Linux do solicitante.

**Verificações pretendidas:**
- Cada tarefa mantém identidade, solicitante, projeto, especialização, autorização financeira correspondente, estado, histórico de transições, justificativa de bloqueio e resultados conhecidos.
- Encerrar sessão interativa não implica conclusão, perda, cancelamento nem aprovação automática da tarefa. O serviço continua apenas as operações já elegíveis e autorizadas.
- **Regra de consulta vigente:** usuários Linux autenticados pertencentes ao grupo autorizado podem consultar o histórico compartilhado **somente pela aplicação e pelo serviço**, inclusive resultados de outros integrantes; não recebem acesso direto ao banco, permissão de exclusão, administração de cofre, nova despesa ou execução por essa razão.
- O guardião opcional consulta informações autorizadas sem adquirir poder executivo. Falha de credencial preserva a solicitação como impedida até regularização e liberação autorizada.
- O banco **não contém valores de credenciais**; a retenção até exclusão manual de [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) permanece vigente.

**Limite:** nenhuma escolha de tecnologia de banco, esquema físico, fila, sessão ou transporte é feita no Charter. `DOCUMENTED`, não `IMPLEMENTED`. Relações: [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md), [CON-002](CON-002-protecao-credenciais.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md).

## Restrição vigente — acesso à base somente pelo serviço (2026-10-22)

**Origem:** autorização do responsável para confiar no controle de identidade/grupo Linux e oferecer consultas por intermédio exclusivo da aplicação instalada.

- O **banco de dados local continua sendo o histórico canônico** de solicitações, autorizações, estados, inferências, métricas, pendências e auditoria; localização física e esquema não são redefinidos no Charter.
- **Usuários comuns e guardião opcional não podem acessar diretamente o banco.** A aplicação cliente não lhes disponibiliza credenciais do banco, acesso SQL, leitura direta dos arquivos ou meios de contornar o serviço.
- A **identidade técnica autorizada do serviço** consulta e modifica os dados necessários. O serviço verifica o acesso funcional da identidade Linux e a participação no grupo autorizado antes de entregar resultados.
- Membros do grupo podem consultar através da aplicação **dados compartilhados do serviço, inclusive de outros membros**. Não haverá cadastro de leitores, senhas ou concessões de leitura por tarefa mantidas em paralelo pela aplicação.
- Consulta não significa aprovação de gasto, permissão de exclusão, acesso ao cofre nem poder de modificar registros fora das operações autorizadas. Eventos mantêm identidade do solicitante e evidências suficientes de consulta/alteração.
- Falha na autenticação Linux, ausência do grupo ou tentativa de acesso direto por usuário comum devem resultar em **acesso negado**; não se presumem controles implementados sem comprovação.

**Preservação histórica:** na política de 2026-10-21 e na extensão anterior de 2026-10-22, referências à leitura somente pelo operador responsável ou por concessionários individuais estão **superadas exclusivamente quanto à autorização de consulta**. Preservam-se retenção até exclusão manual, sigilo, segregação de responsabilidades, descarte auditável e proibição de segredos no banco.

**Relações:** [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md), [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md). `DOCUMENTED`, não `IMPLEMENTED`; definição técnica em Scope/Implementation.

## Restrição consolidada — banco privado, backup e restauração (2026-10-22)

**Origem:** Segurança **1A/2A/3A**, Continuidade **1A/2A/3A/4A** e Backup/restauração **1A/2A/3A**.

- **Histórico canônico local:** somente o serviço da aplicação lê/escreve diretamente o banco privado. Usuário Linux do grupo consulta informações compartilhadas por meio de **um socket Unix local protegido** e do serviço; cliente e guardião não recebem arquivos, chaves, conexões ou permissões diretas ao banco.
- **Continuidade segura:** na falta de socket/serviço, negar nova operação e comunicar falha. Antes de repetir pedido de efeito incerto, consultar o estado persistido. Reinício automático do serviço preserva histórico e bloqueios, **sem retomar automaticamente uma sessão**, sem liberar vaga por suposição, sem cobrar ou aprovar novas operações.
- **Backup escolhido:** **manual, administrativo e autorizado**, com cópia consistente e recuperável em **diretório local protegido, separado dos dados ativos**. Não exigir agendamento nem cópia externa na primeira entrega; cliente comum não pode acessar a cópia. Registrar resultado, falha, integridade conhecida e abrangência.
- **Restauração escolhida:** **manual e supervisionada**, autorizada por administrador com verificação prévia da integridade e origem da cópia e dos riscos de substituição do estado. Restauração não cria aprovação financeira, publicação Git, reinício de tarefa ou retomada de sessão.
- **Banco corrompido:** negar as operações dependentes, preservar evidências e exigir regularização demonstrável antes de liberar o serviço. **Sem reparo, descarte ou restauração automáticos.**
- **Limites:** backup do banco não é snapshot documental de `/docs/backup/`. Retenção até exclusão manual e abrangência real de descarte em cópias seguem [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md). Cópia no mesmo servidor **não garante recuperação de perda física total**.

**Comprovação pendente:** consistência e recuperabilidade da cópia, acesso negado a usuários comuns, funcionamento da restauração manual e bloqueios em caso de corrupção. Tecnologia do banco e forma física da cópia não são decisões do Charter. `DOCUMENTED`, não `IMPLEMENTED`; status `refinement`. Relações: [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md), [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md), [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md).
