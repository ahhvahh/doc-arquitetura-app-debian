# CON-002 — Proteção de credenciais

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
Credenciais de acesso à OpenAI, à AWS/Amazon Bedrock, ao GitHub e aos eventuais serviços autorizados futuramente não podem ser transmitidas como contexto documental nem gravadas em commits, logs, argumentos da CLI, exportações ou arquivos de documentação. Os segredos devem permanecer em cofre administrado por identidade privilegiada autorizada e ser consumidos somente por identidade de serviço exclusiva, com acesso mínimo ao escopo necessário; usuários comuns e guardiões não podem consultar nem administrar essas credenciais.
## Origem verificável
Exigência de segurança necessária ao cadastro de chave solicitado pelo responsável.
## Motivo e impacto
Exposição de credenciais possibilitaria acesso não autorizado e uso indevido dos serviços.
## Verificação
Registros, argumentos, histórico, relatórios, DSL, exportações e commits não revelam segredos. Comprovar que apenas administrador autorizado gerencia credenciais, que a identidade exclusiva do serviço possui somente os acessos necessários e que usuários/guardiões não leem ou alteram o cofre. A tarefa deve bloquear sem credencial, autorização, isolamento de dados e proteção financeira válidos, preservando rastreabilidade.
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

**Origem:** decisões selecionadas pelo responsável na [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md).

- Somente administrador autorizado ou superusuário efetua configuração inicial, substituição, revogação e manutenção de credenciais; usuários comuns e guardião opcional não podem cadastrá-las, lê-las, exportá-las nem modificá-las.
- Uma **identidade exclusiva de serviço** consome somente os segredos necessários às operações autorizadas; não se presume necessidade de executar o serviço continuamente como `root`. Privilégios de administração não são conferidos aos solicitantes.
- Credenciais ficam fora de bancos de solicitações, Markdown, DSL, Git, logs, comandos de linha, argumentos de processo, relatórios e exportações. Metadados podem identificar apenas referências não sensíveis.
- A execução autônoma do serviço é distinta da **autoria e autorização** da tarefa; leitura de históricos de outras identidades e autorização de despesas alheias continuam proibidas.
- Ausência/revogação do segredo ou incerteza quanto às permissões impede a operação dependente, preserva a pendência e exige regularização e liberação autorizada antes de qualquer retomada. Os controles financeiros de [CON-004](CON-004-orcamento-obrigatorio.md) continuam cumulativos e não são flexibilizados.

**Verificação pendente:** demonstrar negação de leitura/escrita indevida, acesso mínimo, revogação efetiva, execução sem usuário conectado, continuidade do bloqueio e ausência de vazamento. Diretriz `DOCUMENTED`, não implementada; manter `refinement`.
