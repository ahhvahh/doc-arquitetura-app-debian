# OBJ-001 — Agentes documentais especializados

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
O uso de regras indistintas pode misturar responsabilidades das etapas e gerar propostas inconsistentes.
## Resultado esperado
A aplicação oferece quatro especializações independentes — Charter, Scope, Implementation e Source — aderentes à norma versionada, com identificação da especialização e baseline utilizada.
## Critérios verificáveis
- A interface apresenta exatamente as quatro especializações solicitadas.
- Cada análise indica a especialização, a referência normativa persistida no agente e o commit-base.
- A atualização das instruções normativas de cada agente ocorre apenas em comando explícito; a inferência comum não as reenvia.
- Uma especialização não executa silenciosamente a responsabilidade de outra.
## Dependências
- [CON-001](../conditions/CON-001-referencia-normativa-main.md).
- [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Interface simplificada e seleção de prompts — decisão do responsável
- Quatro entradas principais: `charter`, `scope`, `implementation` e `source` (`source` corresponde à etapa normativa `src`).
- Uma única árvore seleciona `ETAPA → OPERAÇÃO → AÇÃO → PROMPT especializado`; reutilizar o catálogo existente sem manter árvores paralelas.
- Cada etapa usa um agente especializado e um acervo normativo próprio, identificado por SHA; o projeto-alvo analisado é contexto transitório e nunca substitui a fonte normativa.
- O agente persistente deve consultar as regras relevantes recuperadas da base DSL e obedecer ao prompt escolhido. A aplicação controla o acesso à recuperação quando a API remota não permitir vincular diretamente a ferramenta ao agente.
- A configuração de agente deve ser um fluxo explícito integrado: selecionar a etapa, atualizar a referência normativa, preparar DSL, sincronizar acervo remoto, confirmar indexação, criar/atualizar o agente, registrar IDs e testar recuperação. O modo padrão é prévia; a publicação exige `--apply`.
- A interface proposta `agent setup <charter|scope|implementation|source|all> --model MODELO [--apply]` não equivale a comando já implementado.
- A configuração do agente não executa tarefas do projeto-alvo nem autoriza publicação de alterações.

## Extensão aprovada — seleção de modelos independente da especialização (2026-10-21)
As quatro especializações Charter, Scope, Implementation e Source permanecem separadas. A aplicação deverá permitir escolher modelos simples ou avançados sem duplicar o catálogo normativo por modelo. O modelo/provedor é identificado em cada análise; a referência normativa da especialização permanece rastreável por SHA.

A administração de agentes OpenAI já prevista não determina exclusividade de provedor. A decisão atual prevê **OpenAI, Anthropic e Google Gemini**, condicionados a verificações individuais em [OPEN-005](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md). Modelos locais estão fora da primeira entrega. A comparabilidade esperada está em [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md); não se declara equivalência entre recursos administrativos de provedores distintos.

## Definição adicional — múltiplos provedores e prompts por sessão (2026-10-21)
A escolha aprovada é trabalhar com **provedores remotos diferentes**, não somente modelos diferentes da OpenAI. O operador poderá selecionar novo prompt ou reutilizar prompt do catálogo da etapa durante uma sessão, sem duplicar os objetivos normativos dos agentes.

**Verificação esperada:** cada execução associa explicitamente etapa, prompt selecionado (com identificação de versão/origem), provedor e modelo, conservando o histórico entre inferências da mesma sessão.

Relações: [OPEN-005 — modalidade remota](../open-questions/OPEN-005-estrategia-de-provedores-e-modelos.md) e [OBJ-003 — continuidade das alterações](OBJ-003-propostas-de-alteracoes-verificaveis.md).
