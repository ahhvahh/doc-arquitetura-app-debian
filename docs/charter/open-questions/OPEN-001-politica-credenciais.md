# OPEN-001 — Política de credenciais e usuários

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Diretriz de proteção de credenciais e isolamento por usuário aprovada; sua viabilidade, permissões e contratos de execução ainda exigem comprovação.
## Questão
**Decisões funcionais:** credenciais mantidas no cofre Linux de cada usuário, com entrada oculta; o futuro serviço coordena as tarefas, mas cada inferência ocorre exclusivamente sob o contexto autorizado do usuário correspondente, sem acesso cruzado a segredos. Permanecem pendentes os critérios verificáveis de identidade, permissão, isolamento e disponibilidade de cofre para a CLI e a execução coordenada.
## Contexto
A aplicação existente é executada em contexto de usuário; o daemon isolado está descrito, mas ainda não está implementado. Passar a chave na linha de comando pode expô-la em argumentos e histórico.
## Impacto
**Bloqueadora:** sim; afeta a segurança e a autorização da integração.
## Próxima ação
Comprovar o fluxo de execução no contexto do usuário autorizado, com segregação de credenciais e orçamentos, acesso ao cofre somente pelo próprio contexto, bloqueio sem identidade ou permissão válida e auditoria sem exposição de segredos. Detalhar posteriormente em Scope/Implementation as permissões, disponibilidade de sessão/cofre, interface CLI, revogação e isolamento.
## Destino previsto
Scope de administração/credenciais, após decisão do Charter.

## Proposta de regularização — Decisão 03 (2026-10-21)

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
