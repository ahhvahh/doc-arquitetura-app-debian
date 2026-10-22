# CON-002 — Proteção de credenciais

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
Credenciais de acesso à OpenAI, AWS/Amazon Bedrock, GitHub e eventuais serviços autorizados serão armazenadas, na primeira entrega, em **arquivos administrativos protegidos pelo Linux**, separados da configuração comum e geridos apenas por administrador autorizado. **Não se exige cofre dedicado.** A entrega ao serviço deverá ser restrita por mecanismo nativo do `systemd`; a identidade técnica exclusiva executará com privilégios mínimos, sem leitura direta dos arquivos secretos de origem nem privilégios permanentes de superusuário. Usuários comuns e guardiões não podem consultar ou administrar segredos, nem acessá-los em documentos, banco de tarefas, argumentos, logs ou Git.
## Origem verificável
Exigência de segurança necessária ao cadastro de chave solicitado pelo responsável.
## Motivo e impacto
Exposição de credenciais possibilitaria acesso não autorizado e uso indevido dos serviços.
## Verificação
Registros, argumentos, histórico, relatórios, DSL, configuração geral, exportações e commits não revelam segredos. Comprovar que somente administrador autorizado gerencia arquivos secretos, que usuários/guardiões e serviço não acessam os arquivos administrativos de origem, e que a identidade exclusiva do serviço recebe por `systemd` somente as credenciais necessárias. Comprovar serviço sem privilégios permanentes de `root`, revogação efetiva e bloqueio da operação na ausência de credencial, autorização específica ou proteção financeira válida, preservando tarefas e rastreabilidade.
## Relações
- [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md).
## Condição de revisão
Não é permitida flexibilização sem avaliação e decisão formal documentada.

## Isolamento de usuário para serviço coordenador — decisão 11 (2026-10-21)

**Registro histórico parcialmente superado (2026-10-22):** deixam de valer cofre pessoal e execução de inferências sob a identidade do solicitante; permanecem válidos sigilo, isolamento dos dados, autorização financeira por solicitação e auditoria. Aplica-se a política atual descrita adiante e na [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md).
A escolha do responsável exige que o **serviço Linux somente organize tarefas**, enquanto cada inferência seja realizada no contexto autorizado do seu próprio usuário, sem compartilhar chaves entre operadores.

- O serviço organizador não recebe nem utiliza, como substituição, credenciais privadas armazenadas nos cofres pessoais.
- Credenciais e permissões devem ser acessíveis somente à identidade autorizada responsável pela tarefa, incluindo execução não interativa, quando comprovadamente suportada.
- Sem comprovação de isolamento, identidade, autorização e cofre acessível, negar a tarefa dependente, jamais executá-la sob usuário diferente ou privilegiado.
- Proibir exposição de segredos em argumentos de processo, histórico do terminal, logs, artefatos Git, relatórios e métricas exportadas.
- O isolamento de usuário **não reduz** os controles financeiros da [CON-004](CON-004-orcamento-obrigatorio.md), incluindo a verificação da proteção remota dos quatro projetos OpenAI independentes quando aplicável.

**Origem e relação:** opção C aprovada em [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md). Os contratos de identidade, autorização, execução e testes de isolamento são responsabilidades posteriores de Scope/Implementation. Nenhuma implementação é presumida.

## Restrição vigente — cofre administrativo e identidade exclusiva (2026-10-22)
**Nota de superação parcial — Segurança 1, opção A (2026-10-22):** esta seção permanece como histórico da custódia administrativa aprovada. A obrigatoriedade de “cofre” como mecanismo de armazenamento foi substituída por arquivos administrativos Linux protegidos e entrega restrita via `systemd`, conforme a restrição atualizada acima e o complemento final abaixo. Permanecem válidos sigilo, isolamento, identidade exclusiva, privilégios mínimos, bloqueios e regras financeiras.

**Origem:** decisões selecionadas pelo responsável na [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md).

- Somente administrador autorizado ou superusuário efetua configuração inicial, substituição, revogação e manutenção de credenciais; usuários comuns e guardião opcional não podem cadastrá-las, lê-las, exportá-las nem modificá-las.
- Uma **identidade exclusiva de serviço** consome somente os segredos necessários às operações autorizadas; não se presume necessidade de executar o serviço continuamente como `root`. Privilégios de administração não são conferidos aos solicitantes.
- Credenciais ficam fora de bancos de solicitações, Markdown, DSL, Git, logs, comandos de linha, argumentos de processo, relatórios e exportações. Metadados podem identificar apenas referências não sensíveis.
- A execução autônoma do serviço é distinta da **autoria e autorização** da tarefa. **Consulta compartilhada** de tarefas e históricos de outros membros autenticados do grupo Linux é permitida **somente via aplicação/serviço**; conexão direta de usuário comum ao banco, acesso a segredos e autorização de despesas alheias continuam proibidos.
- Ausência/revogação do segredo ou incerteza quanto às permissões impede a operação dependente, preserva a pendência e exige regularização e liberação autorizada antes de qualquer retomada. Os controles financeiros de [CON-004](CON-004-orcamento-obrigatorio.md) continuam cumulativos e não são flexibilizados.

**Verificação pendente:** demonstrar negação de leitura/escrita indevida, acesso mínimo, revogação efetiva, execução sem usuário conectado, continuidade do bloqueio e ausência de vazamento. Diretriz `DOCUMENTED`, não implementada; manter `refinement`.

## Restrição vigente complementar — arquivos administrativos e entrega de credenciais (2026-10-22)

**Origem:** [OPEN-001 — Segurança, decisão 1, opção A](../open-questions/OPEN-001-politica-credenciais.md), aprovada para a primeira entrega.

- **Armazenamento:** segredos apenas em arquivos administrativos protegidos pelo Linux, sob custódia de administrador autorizado. Configuração geral da aplicação não conterá chaves, senhas ou valores de credenciais. Não há obrigação de contratar, instalar ou manter cofre dedicado.
- **Limite de leitura:** usuários comuns e guardião opcional não obtêm acesso aos segredos. O serviço com identidade técnica exclusiva recebe as credenciais necessárias pelo mecanismo restrito de `systemd`, **sem permissão para ler diretamente a origem administrativa**. O serviço não roda permanentemente com `root` por conveniência.
- **Não divulgação:** é proibida a exposição de segredos em argumentos, histórico do terminal, arquivos de configuração comum, logs, Markdown, DSL, Git, banco de tarefas, resultados, relatórios, eventos de auditoria e exportações.
- **Proteções independentes:** o grupo Linux habilita somente o uso da aplicação e a consulta compartilhada **via serviço**, sem acesso direto ao banco nem às credenciais. Concessão de leitura não significa exclusão de dados, execução de trabalho pago, gestão de segredos ou operações Git extras.
- **Falha fechada:** segredo ausente, inválido, revogado ou indisponível bloqueia a tarefa dependente e preserva a pendência. A correção exige intervenção administrativa; retomada depende de liberação autorizada e nova verificação, incluindo autorização financeira por chamada e proteção remota comprovadamente bloqueante.

**Critérios de verificação:** negar leitura de arquivos secretos por cliente, guardião e identidade do serviço; comprovar entrega de escopo mínimo por `systemd` sem vazamentos; bloquear operações sem credencial/permissão; rastrear revisão e revogação; comprovar que apenas o serviço medeia o banco e não realiza despesa sem autorização. Caminhos físicos e detalhes de unidade `systemd` ficam para etapas posteriores.

**Estado:** `DOCUMENTED`, não `IMPLEMENTED`. Badge `refinement` preservado. A [CON-004](CON-004-orcamento-obrigatorio.md) e a norma de escrita Git continuam integralmente válidas.
