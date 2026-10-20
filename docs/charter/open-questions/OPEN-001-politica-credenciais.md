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
