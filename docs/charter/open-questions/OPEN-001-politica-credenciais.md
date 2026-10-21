# OPEN-001 — Política de credenciais e usuários

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão em aberto.
## Questão
Como conciliar a CLI solicitada `openai --chave`, o cadastro em banco privado e o futuro daemon com usuários autorizados, garantindo sigilo da chave no processo e no armazenamento?
## Contexto
A aplicação existente é executada em contexto de usuário; o daemon isolado está descrito, mas ainda não está implementado. Passar a chave na linha de comando pode expô-la em argumentos e histórico.
## Impacto
**Bloqueadora:** sim; afeta a segurança e a autorização da integração.
## Próxima ação
Definir experiência principal de entrada oculta, compatibilidade de `--chave`, gestão de chave mestra, rotação, isolamento e modelo de acesso.
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
