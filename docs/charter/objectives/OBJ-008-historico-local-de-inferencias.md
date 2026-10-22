# OBJ-008 — Histórico local e auditável de inferências

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
Sem persistir solicitações, respostas e métricas torna-se difícil revisitar experimentos, comparar modelos e compreender diferenças entre previsão e cobrança.

## Resultado esperado
Manter um histórico em banco de dados local para análise futura de solicitações e tentativas de inferência, com prompts e resultados completos e medições de consumo e valores previstos e realizados, quando conhecidos.

## Critérios verificáveis
- Cada solicitação e tentativa tem ID, horário, duração, estado, projeto-alvo, etapa, prompt selecionado, provedor, modelo, SHA normativo e baseline documental identificados.
- O registro contém o prompt efetivamente transmitido e a resposta integral recebida, preservando o vínculo ao contexto de origem e a sinalização de trechos omitidos por segurança.
- Tokens previstos de entrada e saída e tokens realizados de entrada e saída são independentes; categorias extras de uso podem ser registradas com sua origem.
- Custos previstos e realizados são campos distintos; tarifa, moeda, cotação, data e origem diferenciam valores confirmados por provedor dos calculados localmente.
- Medições não fornecidas pelo provedor são desconhecidas, não zero; falhas, cancelamentos e respostas incompletas são rastreáveis.
- O histórico permite consultas e comparações por tarefa, modelo, provedor, especialização, qualidade, período, duração, tokens e custos.

## Limites
Este objetivo não escolhe banco, esquema ou política de exclusão. A retenção e o acesso dependem de [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md).

## Relações
- [OBJ-007 — Comparação](OBJ-007-comparacao-controlada-de-modelos.md)
- [OBJ-005 — Orçamentos](OBJ-005-orcamento-previsivel-e-bloqueio.md)
- [CON-005 — Persistência local](../conditions/CON-005-persistencia-local-de-inferencias.md)
- [CON-002 — Credenciais](../conditions/CON-002-protecao-credenciais.md)

## Requisito de interoperabilidade — métricas JSON (2026-10-21)
As métricas de cada solicitação/tentativa deverão ser **persistidas ou disponibilizadas como objetos JSON estruturados e recuperáveis**, permitindo que aplicações consumidoras consultem os dados e reconstruam relatórios em outro ambiente, sem depender da interface de análise original.

**Critérios verificáveis:**
- Um registro recuperado em JSON identifica sua solicitação, sessão, tentativa, provedor/modelo, timestamps, resultado, tokens e custos conhecidos, com referências de origem e contexto suficientes para comparação.
- Métricas nativas de cada provedor e indicadores comparáveis preservam sua proveniência, sem perda silenciosa de categorias ou alteração de unidades.
- Ausência de medição, estimativa e cobrança confirmada são diferenciadas; não se inventa zero em campos sem informação.
- Uma aplicação independente consegue recuperar as métricas e reconstruir uma visão histórica por sessão, inferência ou provedor sem repetir as chamadas externas.

O Charter não fixa nomes de propriedades, modelo físico de banco ou transporte de consulta; essas definições pertencem a Scope/Implementation e deverão respeitar [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md) e [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md).

## Decisão aprovada — retenção até exclusão manual (2026-10-21)
- Manter prompts, respostas, avaliações e métricas JSON recuperáveis até exclusão manual do operador autorizado, sem prazo automático de vencimento.
- Consultas e reconstrução externa de relatórios são permitidas somente ao operador e a consumidores expressamente autorizados; consulta não autoriza remoção.
- Omissões ou mascaramentos necessários para impedir vazamento de segredos são explicitados, sem registrar os próprios segredos.
- O descarte manual deverá registrar de modo verificável quem solicitou, quando, o que foi afetado e o resultado, sem manter o conteúdo descartado no evento de auditoria.
- Informar o alcance real da exclusão no histórico, em cópias e em exportações; não afirmar remoção de cópias fora do controle da aplicação.

**Critérios de verificação propostos:** não ocorrer expiração automática; registros continuam recuperáveis até descarte manual; exclusões não autorizadas são negadas; consultas indevidas e vazamentos de segredos são impedidos; o descarte gera evidência suficiente de sua execução.

Decisão: [OPEN-006 — Retenção e acesso](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md). Formato, armazenamento, proteção e consultas técnicas serão definidos em Scope/Implementation.

## Recuperabilidade de tarefas pendentes entre execuções — decisão 29
Além da retenção dos registros de inferência até exclusão manual já aprovada, o produto deverá **preservar informações suficientes para localizar e reconstruir conceitualmente uma tarefa que ficou aguardando revisão**, mesmo quando a aplicação for encerrada ou reiniciada. O operador autorizado deverá conseguir identificar **sessão/tarefa, estado de avaliação pendente, proposta integral gerada, contexto original, baseline e hashes, modelos/prompts usados, autorizações e custos já incorridos, falha do revisor, correções realizadas e limite remanescente**.

**Critério positivo:** depois de reiniciar, a tarefa continua associada à sua identidade e à **mesma proposta**; verificação de integridade e permissão permite **retomar só a revisão independente** com revisor escolhido/confirmado, homologado e distinto e **nova autorização financeira própria**. Eventos novos são acrescentados sem alterar os registros anteriores.

**Critério negativo:** falha de acesso, integridade, vínculo com baseline, disponibilidade e homologação, ou **exclusão manual autorizada de conteúdo necessário** impede recuperar validamente a revisão. Os segredos continuam ausentes do histórico e das exportações; o operador recebe a causa do bloqueio. Não recuperar automaticamente dados legitimamente eliminados, nem usar valores desconhecidos de custo como zero.

**Limites:** a política canônica de retenção, proteção, acesso e exclusão continua em [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) e [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md). Recuperação entre execuções **não é prova de persistência implementada** nem determina banco, formato, esquema, interface ou mecanismo técnico. Relações: [OPEN-002 — decisão 29](../open-questions/OPEN-002-fidelidade-e-custo.md) e [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).

## Histórico da suspensão e reabertura da sessão — decisão 30
O registro histórico deverá distinguir **suspensão do trabalho ativo**, **sessão reaberta**, **tarefas ainda pendentes** e **sessão formalmente concluída para publicação**. A suspensão é um evento autorizado distinto de cancelamento de tarefa, revisão concluída, publicação ou integração Git.

**Critérios verificáveis:** o operador autorizado consegue recuperar, após suspensão e eventual reinicialização, a mesma sessão, os vínculos entre tarefas e propostas, as baselines normativas/documentais, os modelos e configurações efetivos, as revisões ausentes, falhas anteriores, custos e **contadores já consumidos de correção**. Os eventos de suspensão e reabertura devem ser distinguíveis no histórico sem sobrescrever fatos anteriores, e a condição de **não concluída para publicação** deve continuar explícita enquanto existirem impedimentos.

**Proteções:** acesso às informações depende de autorização; não persistir segredos; fechamento ativo não faz limpeza automática, não autoriza gasto nem permite publicar propostas pendentes. A política de **retenção até exclusão manual** permanece em [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) e [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md). Recuperabilidade funcional não comprova implementação de persistência. Relações: [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md), [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md) e [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Histórico da conclusão sem publicação — decisão 31
A aplicação deverá conservar a diferença entre **sessão suspensa**, **tarefa cancelada individualmente**, **sessão formalmente concluída com publicações** e **sessão formalmente concluída sem publicação**. A última situação decorre da ausência **verificada** de alterações documentais validadas publicáveis, do cancelamento **expresso** de todas as tarefas pendentes e da confirmação humana final, sem criação de PR.

**Critérios de rastreabilidade:** registrar identidade da sessão e do operador autorizado, confirmação de conclusão, tarefas canceladas e suas justificativas, inventário final de alterações, referências da baseline, propostas e respostas geradas, autorizações e custos efetivos, versões de configurações e contagem dos ciclos de correção. A conclusão **não apaga nem substitui** registros passados, não converte propostas canceladas em aprovadas e **não é uma exclusão manual de dados**.

A consulta e eventual descarte permanecem sujeitos à política de retenção e acesso de [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) e [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md), com proteção de segredos. Nenhum formato físico, estrutura de banco ou código é definido aqui. Relações: [OPEN-003 — decisão 31](../open-questions/OPEN-003-politica-de-branches.md) e [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).

## Sessão concluída preservada como histórico; novo trabalho em nova sessão — decisão 32
O histórico deverá distinguir claramente **sessão suspensa**, **sessão formalmente concluída com publicação** e **sessão formalmente concluída sem publicação**. Uma conclusão formal é **terminal para novos trabalhos** na identidade daquela sessão, mas não equivale a exclusão dos seus registros nem impede consulta autorizada.

Quando o operador iniciar um novo trabalho no mesmo repositório, a aplicação deverá registrar uma **nova sessão**, com identidade e baseline próprias, preservando integralmente os eventos e cancelamentos da sessão anterior sem reativar suas tarefas. As credenciais, autorizações financeiras, homologações e avaliações não são consideradas transferidas ou automaticamente válidas na nova sessão.

**Critério verificável:** recuperar S1 como histórico após sua conclusão formal, impedir inserção de novas tarefas em S1, iniciar S2 distintamente e permitir reconstruir as decisões/custos anteriores sem misturar suas tentativas com S2. A eventual exclusão de histórico requer o procedimento manual autorizado já definido em [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) e [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md). Contratos de armazenamento e esquema são assuntos posteriores. Relação: [OPEN-003 — decisão 32](../open-questions/OPEN-003-politica-de-branches.md).

## Histórico — Exclusão de branch prevista na decisão 33 (superada)

**Registro histórico superado pela decisão 34:** as orientações antigas de excluir a branch ao cancelar não estão mais vigentes. Os critérios atuais são os da decisão 34 abaixo. Não executar exclusão automática.

A trilha da sessão deverá distinguir eventos de **suspensão**, **cancelamento expresso da sessão**, **cancelamento/retirada individual de tarefas**, **conclusão sem publicação** e **encerramento aprovado com PR**. No cancelamento da sessão, registrar operador autorizado, identidade da sessão e da branch, impactos sobre conteúdo não integrado, decisão expressa e **resultado verificável da solicitação de exclusão da branch**; uma falha na exclusão permanece explícita.

Na conclusão aprovada, registrar o **conjunto final de alterações**, a baseline de origem e branch de sessão, identificação da aprovação humana de encerramento, eventual **PR para `develop`**, seu destino e o estado de checks e de autorização de integração, **sem presumir merge**. Não transformar inferência ou aprovação de PR em autorização de gasto.

O cancelamento de sessão e a exclusão da **ref Git de trabalho** não equivalem a descarte do histórico local, que segue protegido e retido **até exclusão manual autorizada** conforme [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md) e [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md). Não alegar exclusão de cópias remotas ou audit trails além do controle da aplicação. Os eventos são requisitos documentais, **não implementação**. Relações: [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md), [OBJ-003](OBJ-003-propostas-de-alteracoes-verificaveis.md).

## Decisão 34 — Rastrear reprovação com branch preservada
Distinguir no histórico **pedido de cancelar**, **decisão de continuar**, **encerramento aprovado com PR solicitada**, **encerramento reprovado sem PR**, suspensão e conclusão sem publicação. Na reprovação registrar operador, sessão, branch, SHA da baseline, propostas, custos, revisões e **ausência de exclusão da branch e de operações posteriores**. Na aprovação registrar o conjunto validado, PR para `develop`, checks e autorização de merge posterior, sem considerá-lo automaticamente integrado.

Todos os registros continuam sujeitos à retenção até exclusão manual em [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md); manter a branch de uma sessão encerrada **não** reativa a sessão ou as tarefas. A intenção de criar branch/PR continua impedida pela norma vigente conforme [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md). `DOCUMENTED`, não `IMPLEMENTED`.

## Rastreabilidade da espera por PR para liberar abertura — decisão 35
Antes de iniciar nova sessão, a aplicação deverá registrar de maneira auditável as **condições verificadas para abertura**: repositório, destino `develop`, detecção de PRs abertas, decisão do operador (**aguardar ou cancelar**), eventuais verificações seguintes, resultado do período de espera e, **somente quando elegível**, a baseline SHA da `develop` utilizada para a sessão.

Falhas de consulta, cancelamento de inicialização, PR ainda aberta e encerramento por esgotamento da espera **não criam sessão ativa nem inferências remuneradas**. A evidência histórica deve distinguir PR **integrada** de PR **fechada sem merge** e não atribuir alteração à baseline sem verificação.

**Limite:** nenhuma evidência de ausência de PR aberta substitui prova de exclusividade de sessão ativa, cuja política está pendente em [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md). A proteção dos dados e eventual descarte permanecem regidos por [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md); nenhuma implementação é presumida.

## Rastreabilidade da exclusividade de sessão — decisão 36 (2026-10-21)

**Resultado esperado adicional:** registrar eventos de abertura, tentativa bloqueada, suspensão, retomada e encerramento de sessões, mantendo a **exclusividade de uma sessão ativa por repositório**, conforme [OPEN-003 — decisão 36](../open-questions/OPEN-003-politica-de-branches.md).

**Critérios verificáveis:**
- Identificar repositório-alvo, operador autorizado, sessão/tentativa, momento, ação, estado observado, decisão e motivo do bloqueio, sem registrar segredos.
- Distinguir sessão efetivamente ativada de tentativa negada; uma tentativa bloqueada não cria sessão ativa, branch ou consumo de inferência.
- Preservar a continuidade e identidade de sessões suspensas e registrar por que uma retomada foi permitida ou bloqueada por outra sessão ativa do mesmo repositório.
- Distinguir bloqueio por **outra sessão ativa** de bloqueio por **PR aberta para `develop`** ou falha de verificação. A ausência de PR não equivale a prova de exclusividade.
- Permitir reconstruir se duas ativações concorrentes no mesmo repositório foram corretamente serializadas ou bloqueadas e se repositórios distintos puderam manter sessões ativas independentes.

A forma de persistência e de consulta desses eventos pertence a Scope/Implementation. A retenção e o controle de acesso continuam sujeitos a [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md). **DOCUMENTED**, não `IMPLEMENTED`; nenhuma gravação em branches de sessão ou PR é autorizada por esta decisão.

## Auditoria da recuperação supervisionada — decisão 37 (2026-10-21)

**Resultado esperado adicional:** o histórico local deve distinguir interrupção inesperada, recuperação pendente, verificação de execução remanescente, decisão humana de retomada, suspensão liberadora de vaga e bloqueios, associados **à identidade da mesma sessão e ao repositório original**.

**Critérios verificáveis:**
- Registrar operador autorizado, repositório, sessão/tarefa, evento e horário, causa conhecida ou desconhecida da interrupção, estado anterior e posterior, evidências/resultado da verificação de ausência de execução, alternativa oferecida, escolha expressa e efeito sobre a vaga.
- Preservar baselines e hashes, branch vinculada, propostas e pareceres, pendências, número de correções, histórico de autorizações e custos, sem reescrever registros anteriores.
- Distinguir retomada autorizada da **mesma sessão** de criação de nova sessão e distinguir suspensão confirmada de encerramento, cancelamento, liberação por suposição e mera perda de conexão.
- Registrar falhas e estados inconclusivos como **bloqueados**, sem marcar vaga liberada ou trabalho concluído; não registrar gasto ou autorização de inferência inexistentes.
- Permitir auditoria de tentativas concorrentes: duas ativações no mesmo repositório não podem ser registradas como simultaneamente permitidas, enquanto repositórios diferentes podem operar de modo independente.
- Eventos de recuperação ou suspensão não equivalem a commit, PR, merge ou aprovação financeira; preservar a rastreabilidade da decisão 35 sobre PR aberta para `develop`.

**Relações:** [OPEN-003 — decisão 37](../open-questions/OPEN-003-politica-de-branches.md), [OPEN-002 — decisão 29](../open-questions/OPEN-002-fidelidade-e-custo.md) e [OPEN-006 — retenção e proteção](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md). O formato dos registros e os controles técnicos ficam para Scope/Implementation. **DOCUMENTED**, não `IMPLEMENTED`; as restrições atuais de escrita Git permanecem inalteradas.

## Rastreabilidade do período de espera por PR — decisão 38 (2026-10-22)
**Resultado esperado adicional:** registrar a espera autorizada por Pull Requests abertas para `develop`, respeitando a decisão 35 e os parâmetros definidos na [OPEN-003 — decisão 38](../open-questions/OPEN-003-politica-de-branches.md).

**Critérios verificáveis:**
- Identificar repositório, tentativa de abertura, operador autorizado e escolha inicial (**aguardar** ou **cancelar**), sem criar identidade de sessão ativa por mera espera.
- Registrar início e término de **cada período autorizado de até 5 minutos**, verificações previstas **a cada 15 segundos**, estado observado das PRs e eventuais erros ou consultas inconclusivas.
- Distinguir esgotamento da espera, **nova autorização expressa de outros 5 minutos**, cancelamento e liberação por ausência confirmada de PR aberta; jamais registrar renovação ou autorização por presunção.
- Registrar resultado da verificação da baseline/SHA de `develop` e da exclusividade de sessão por repositório quando a abertura passar a ser elegível; distinguir PR integrada de PR fechada sem merge.
- Tentativas bloqueadas ou canceladas não podem aparecer como sessões efetivamente ativas nem como operações de inferência, commit, PR ou merge ocorridas.
- Proteger dados e aplicar a retenção e o acesso já previstos na [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md), sem armazenar segredos.

**Situação:** `DOCUMENTED`, não `IMPLEMENTED`; formato dos registros e mecanismos técnicos pertencem a Scope/Implementation. A restrição normativa de escrita exclusiva em `develop` continua em vigor.

## Rastreabilidade de critérios eliminatórios e avaliação comparativa de qualidade (2026-10-22)

**Resultado esperado adicional:** permitir recuperar o motivo pelo qual uma tentativa foi declarada **elegível ou não elegível à pontuação**, distinguindo verificações eliminatórias, avaliações qualitativas e eventual pontuação quantitativa homologada.

**Critérios verificáveis:**
- Relacionar a tarefa, tentativa, modelo gerador efetivo, segundo revisor distinto, suas versões, prompts, contexto e SHAs documentais e normativos ao resultado de cada **critério obrigatório**, com fonte, data, evidência e justificativa da eventual falha.
- Registrar os resultados individuais de **completude, clareza, consistência, rastreabilidade e necessidade/importância de correção**, a natureza material/editorial dos apontamentos e a decisão humana associada, sem substituir registros anteriores.
- Distinguir **não elegível por falha obrigatória**, **pontuação não calibrada/indisponível** e **pontuação calculada com método homologado**. Campos ausentes ou indeterminados não serão transformados em zero ou aprovação presumida.
- Quando houver pontuação numérica, manter referência à versão do método, critérios, parâmetros, evidências, validade da comparação e eventual nota efetivamente apurada; **pesos, fórmula, escala e limiares não estão estabelecidos nesta decisão**.
- Preservar **custos, tokens, duração, quantidade de correções e aprovações humanas** como campos separados de qualquer indicador de qualidade. Não misturar resultados de tarefas ou baselines incompatíveis em um ranking.
- Vincular revisões posteriores ao registro de origem e permitir reconstrução de avaliações antigas de acordo com seus métodos históricos, sem alterar retroativamente fatos ou registrar execução de teste não ocorrida.

**Limites:** a retenção, o controle de acesso e a proteção de segredos permanecem sujeitos à [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md). O mecanismo técnico de representação pertence às etapas posteriores. `DOCUMENTED`, não `IMPLEMENTED`. Relações: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md) e [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md).

## Acompanhamento de solicitações do serviço autônomo (2026-10-22)

**Resultado esperado adicional:** registros locais permitem acompanhar solicitações e resultados independentemente de sessões interativas, sem expor segredos e com acessos delimitados por identidade.

**Critérios verificáveis futuros:**
- Distinguir solicitação **pendente, em execução, concluída, interrompida ou bloqueada**, com causa, instante, projeto, tarefa e transições auditáveis; manter autoria da solicitação separada da identidade de serviço executora e da identidade que autorizou a despesa.
- O solicitante consulta suas tarefas; pessoas/consumidores explicitamente autorizados veem **somente o escopo concedido**. Toda consulta protegida é verificável; a autorização de consulta não concede exclusão, acesso ao cofre nem gasto.
- O **guardião opcional** apenas recupera e apresenta estados/resultados visíveis, não executa inferência nem atua como administrador de segredos. Sua ausência não elimina ou invalida tarefas.
- Após logout/reinício do usuário, solicitações, estados e histórico continuam disponíveis ao serviço; indisponibilidade ou revogação da credencial produz registro de **bloqueio com preservação do contexto**, não execução presumida.
- Regularização da credencial e **liberação autorizada** geram evento vinculado à tarefa, com revalidação de integridade, identidade e autorização financeira específica. Nunca inventar sucesso, cobrança zero ou retomada automática.
- Preserve-se a retenção até exclusão manual e o acesso restrito de [OPEN-006](../open-questions/OPEN-006-retencao-e-protecao-do-historico.md), sem persistir valores do cofre no histórico.

**Limites:** formato de armazenamento e interface de consulta não são definidos neste objetivo. Política `DOCUMENTED`, não `IMPLEMENTED`. Relações: [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md), [CON-002](../conditions/CON-002-protecao-credenciais.md), [CON-005](../conditions/CON-005-persistencia-local-de-inferencias.md), [OBJ-005](OBJ-005-orcamento-previsivel-e-bloqueio.md).
