# OBJ-009 — Execução automática de tarefas elegíveis

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação

Permitir o processamento contínuo de instruções complexas dentro de uma sessão, evitando confirmações desnecessárias entre tarefas, sem comprometer autorização financeira, segurança ou qualidade documental.

## Resultado esperado

Ao receber uma instrução vinculada a uma sessão autorizada, o serviço organiza o trabalho em um plano rastreável de tarefas e executa automaticamente todas as tarefas elegíveis, sem solicitar aprovação geral do plano.

O plano e seu progresso ficam disponíveis ao operador. Tarefas impedidas permanecem pendentes ou bloqueadas, com motivos identificáveis.

## Critérios verificáveis

- A instrução pertence a uma sessão ativa e autorizada, com repositório, especialização e baseline identificados.
- O planejamento identifica tarefas, objetivos, dependências e resultados esperados, sem ultrapassar o escopo autorizado.
- O serviço valida o plano antes de executá-lo; a saída do planejador não constitui autorização para executar comandos arbitrários.
- Tarefas elegíveis iniciam automaticamente sem aprovação humana adicional do plano.
- Toda chamada paga, inclusive planejamento, geração, revisão ou correção, exige autorização financeira individual previamente válida, orçamento local e proteção remota comprovadamente bloqueante.
- A aprovação financeira geral da sessão ou do plano não substitui autorizações específicas por chamada.
- Uma tarefa que dependa de resultado pendente, bloqueado ou inválido não é iniciada como se a dependência estivesse concluída.
- Resultados documentais continuam sujeitos a validação estrutural, segundo LLM efetivamente distinto e homologado e, quando necessário, até quatro correções por tarefa.
- O serviço registra plano, tarefas, estados, resultados, custos, autorizações e bloqueios, permitindo acompanhamento e retomada autorizada.
- A conclusão da instrução apresenta relatório consolidado, inclusive quando houver execução parcial.
- A conclusão da instrução não encerra automaticamente a sessão.
- Executar tarefas automaticamente não concede autorização para gravação documental, criação de branches, commits, PRs ou integração Git fora da política vigente.

## Limites e dependências

- [OBJ-001](OBJ-001-agentes-documentais-especializados.md)
- [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md)
- [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [OBJ-008](OBJ-008-historico-local-de-inferencias.md)
- [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md)
- [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md)
- [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md)
- [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md)

## Atores relacionados

Operador autorizado e serviço executor local, conforme [PCH-001](../project-context/PCH-001-contexto-integracao-openai-dsl.md).

## Origem da decisão

Escolha expressa do responsável em 2026-10-22: opção B — executar automaticamente tarefas elegíveis com autorizações financeiras prévias, sem aprovação geral do plano.

Decisão funcional aceita, ainda não implementada. Contratos de planejamento, estados e tratamento de dependências pertencem a Scope; mecanismos e testes de execução pertencem a Implementation.

## Limite de tokens por inferência e orçamento acumulado por sessão — 2026-10-22

**Resultado esperado complementar:** o serviço utiliza limite padrão configurável de **1.000.000 de tokens por inferência**, independentemente de quantas inferências integrem a mesma sessão. A sessão acumula os tokens consumidos e o custo monetário das operações realizadas, sob **teto local padrão configurável de US$ 10,00 por sessão**.

**Critérios verificáveis:**
- O serviço acompanha separadamente o consumo de cada inferência e os totais acumulados da sessão para planejamento, geração, revisão independente e correções; o limite de tokens de uma inferência não é confundido com limite global de tokens da sessão.
- A elegibilidade de uma tarefa remunerada pressupõe saldo de sessão suficiente para reserva conservadora da operação, autorização financeira individual válida e todas as proteções de [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md) e [CON-004](../conditions/CON-004-orcamento-obrigatorio.md); o serviço não aceita uma tarefa paga por causa apenas do valor padrão configurado.
- Uma inferência acima do limite efetivo de contexto, de saída ou de tokens do modelo é bloqueada para redefinição autorizada, sem truncamento silencioso nem alteração automática da baseline.
- Ao esgotar o orçamento da sessão, o serviço deixa pendentes ou bloqueadas novas tarefas que gerariam cobrança, preserva resultados anteriores, histórico, saldo de correções e a sessão para retomada autorizada.
- A reabertura, o reinício do serviço ou novas instruções na mesma sessão não reiniciam os acumuladores nem concedem despesa adicional implicitamente.

**Relações:** [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md), [OBJ-008](OBJ-008-historico-local-de-inferencias.md), [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), [OPEN-004](../open-questions/OPEN-004-projeto-openai-e-hard-limit.md).

**Situação:** comportamento esperado `DOCUMENTED`, não `IMPLEMENTED`; contratos detalhados de contabilização e execução pertencem às próximas etapas.
