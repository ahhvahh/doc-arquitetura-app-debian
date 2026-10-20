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
