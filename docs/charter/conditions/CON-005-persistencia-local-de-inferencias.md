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
