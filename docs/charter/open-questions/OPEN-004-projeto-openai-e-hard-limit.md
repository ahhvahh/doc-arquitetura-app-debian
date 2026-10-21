# OPEN-004 — Associação de agentes a projetos OpenAI e controle de hard limits

![Tipo: OPEN](../../assets/documentation-style/badges/badge-type-open.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Tipo
Decisão de implantação ainda não encerrada.

## Questão
Os quatro agentes utilizarão projetos OpenAI separados (quatro tetos mensais remotos e quatro credenciais de projeto) ou compartilharão um projeto com um teto agregado? Como comprovar a configuração remota sem armazenar credenciais administrativas no aplicativo?

## Contexto e origem
O responsável exige limites locais por agente e controles também na conta OpenAI. O hard limit na plataforma é configurado no projeto/organização; não equivale a uma cota mensal inerente ao objeto Agent.

## Impacto
**Bloqueadora para habilitar pagamentos em produção:** sim, até que o isolamento de orçamento remoto e as credenciais aplicáveis estejam documentados e verificados. Não bloqueia compilação DSL ou testes locais simulados.

## Próxima ação verificável
Selecionar política de projetos e chaves OpenAI, definir a verificação de hard limits e documentar falha fechada quando limites indisponíveis. Registrar a decisão em Charter e depois especificar contrato de API em Scope e Implementation.

## Destino previsto
- [CON-004](../conditions/CON-004-orcamento-obrigatorio.md)
- [OBJ-005](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md)

## Proposta de regularização — Decisão 04 (2026-10-21)

**Origem:** decisão explícita do responsável — opção A, quatro projetos OpenAI independentes.
**Resultado da triagem:** proposta aceita para regularização do Charter, sem criar projetos, chaves ou limites na plataforma.

- Organizar a integração OpenAI em quatro projetos distintos: `charter`, `scope`, `implementation` e `source` (correspondente à etapa normativa `src`).
- Cada especialidade deve utilizar a identidade do projeto OpenAI e sua própria credencial de projeto, com segregação de permissões, consumo, rastreabilidade e orçamento. A aplicação não deve misturar credenciais ou contexto documental de especialidades distintas.
- Manter limites locais independentes por operação, por especialidade e um limite global da aplicação, além dos controles de orçamento que estiverem efetivamente disponíveis na plataforma para cada projeto.
- A inexistência de um teto remoto efetivamente bloqueante não pode ser ocultada nem tratada como garantia de cobrança máxima. Confirmar as capacidades reais do provedor e adotar a política local de falha fechada prevista em [CON-004](../conditions/CON-004-orcamento-obrigatorio.md).
- Antes de permitir inferência paga em produção, verificar o vínculo correto entre especialidade, projeto, credencial, orçamento autorizado, estimativa de consumo e mecanismo de acompanhamento. Sem essa verificação, bloquear a operação remunerada.
- A criação de projetos, a definição de orçamentos remotos e o provisionamento de credenciais serão ações administrativas explícitas e separadas da execução de análises documentais; nenhuma credencial será armazenada em Git, DSL ou logs.
- O controle e a auditoria das despesas deverão permitir consulta por especialidade e também consolidar o total utilizado na aplicação, incluindo provedores adicionais quando forem habilitados.

**Relações canônicas:** [OBJ-005 — Inferências com custo previsível e limitado](../objectives/OBJ-005-orcamento-previsivel-e-bloqueio.md) e [CON-004 — Orçamentos obrigatórios em duas camadas](../conditions/CON-004-orcamento-obrigatorio.md).

**Pendências:** validar no serviço OpenAI as garantias, características e permissões reais dos controles remotos de orçamento; especificar os contratos de identidade, autorização e verificação em Scope/Implementation após o gate. A OPEN permanece em `refinement` e continua bloqueadora para habilitar inferências pagas em produção.
