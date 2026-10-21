# CON-002 — Proteção de credenciais

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
Credenciais de acesso à OpenAI, Anthropic, Google Gemini, GitHub e demais serviços autorizados não podem ser transmitidas como contexto documental nem gravadas em commits, logs, argumentos da CLI, exportações ou arquivos de documentação. O acesso deve respeitar o cofre Linux segregado por identidade de usuário; nenhuma instância de serviço pode substituir essa segregação.
## Origem verificável
Exigência de segurança necessária ao cadastro de chave solicitado pelo responsável.
## Motivo e impacto
Exposição de credenciais possibilitaria acesso não autorizado e uso indevido dos serviços.
## Verificação
Os registros e conjuntos documentais exportados não contêm segredos; os controles locais de acesso são auditados. É possível comprovar que usuário ou tarefa não acessa cofre, credenciais, dados nem permissões financeiros de outro usuário e que a inferência é bloqueada na ausência de identidade, cofre e autorização válidos.
## Relações
- [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md).
## Condição de revisão
Não é permitida flexibilização sem avaliação e decisão formal documentada.

## Isolamento de usuário para serviço coordenador — decisão 11 (2026-10-21)
A escolha do responsável exige que o **serviço Linux somente organize tarefas**, enquanto cada inferência seja realizada no contexto autorizado do seu próprio usuário, sem compartilhar chaves entre operadores.

- O serviço organizador não recebe nem utiliza, como substituição, credenciais privadas armazenadas nos cofres pessoais.
- Credenciais e permissões devem ser acessíveis somente à identidade autorizada responsável pela tarefa, incluindo execução não interativa, quando comprovadamente suportada.
- Sem comprovação de isolamento, identidade, autorização e cofre acessível, negar a tarefa dependente, jamais executá-la sob usuário diferente ou privilegiado.
- Proibir exposição de segredos em argumentos de processo, histórico do terminal, logs, artefatos Git, relatórios e métricas exportadas.
- O isolamento de usuário **não reduz** os controles financeiros da [CON-004](CON-004-orcamento-obrigatorio.md), incluindo a verificação da proteção remota dos quatro projetos OpenAI independentes quando aplicável.

**Origem e relação:** opção C aprovada em [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md). Os contratos de identidade, autorização, execução e testes de isolamento são responsabilidades posteriores de Scope/Implementation. Nenhuma implementação é presumida.
