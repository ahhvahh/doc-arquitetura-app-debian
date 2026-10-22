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
- O solicitante e identidades/consumidores **expressamente autorizados** consultam somente o escopo permitido; consulta não confere permissão de exclusão, administração de cofre, nova despesa ou execução.
- O guardião opcional consulta informações autorizadas sem adquirir poder executivo. Falha de credencial preserva a solicitação como impedida até regularização e liberação autorizada.
- O banco **não contém valores de credenciais**; a retenção até exclusão manual de [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) permanece vigente.

**Limite:** nenhuma escolha de tecnologia de banco, esquema físico, fila, sessão ou transporte é feita no Charter. `DOCUMENTED`, não `IMPLEMENTED`. Relações: [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md), [CON-002](CON-002-protecao-credenciais.md) e [OBJ-008](../objectives/OBJ-008-historico-local-de-inferencias.md).
