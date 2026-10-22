# OBJ-003 — Propostas de alterações verificáveis

![Tipo: OBJ](../../assets/documentation-style/badges/badge-type-obj.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Motivação
O resultado da inferência deve ser auditável, não uma gravação direta de arquivos sem revisão.
## Resultado esperado
A aplicação apresenta conjuntos de alterações identificadas por destino e conteúdo, permitindo inclusão, edição, exclusão autorizada, validação e revisão anterior à publicação.
## Critérios verificáveis
- Toda proposta apresenta origem, caminho de destino, natureza da operação, identificação da baseline, nome proposto de branch, título e descrição do commit.
- A aplicação valida o nome de branch e aplica mudanças em branch local derivada da baseline verificada, sem executar instruções Git fornecidas pelo modelo.
- Modificações preservam as partes não alteradas e as referências válidas.
- Rejeições ou conflitos deixam o repositório inalterado.
- A futura integração das alterações à `develop` exige aprovação humana explícita e resultado auditável. A gravação intermediária em branch da sessão não poderá ser habilitada antes do alinhamento normativo indicado em [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).
- A entrega final prevê uma branch de tarefa distinta, nunca um push direto para `main`; sua publicação depende de atualização da política vigente.
## Dependências
- [NOBJ-001](../non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md).
- [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Entrega textual DSL por arquivo — decisão do responsável
- A resposta do agente deverá ser **texto DSL**, sem envelope JSON ou Markdown diretamente editável.
- O contrato planejado `BIOMA-DSL-CHANGESET/2.0` reúne: repositório, branch de origem `develop`, SHA do projeto-alvo, SHA normativo, sugestão de branch, título e descrição.
- Cada arquivo afetado aparecerá em bloco próprio, delimitado por `=== FILE ADD|MODIFY|DELETE caminho ===` e `=== END FILE ===`, com `BASE-SHA256` quando já existir. O resultado termina em `=== END CHANGESET ===`.
- A aplicação Go, e não o agente, valida o formato, caminhos, escopo, hashes e permissões, converte DSL em Markdown e preserva conteúdo fora das alterações.
- O agente apenas sugere nome de branch, título e descrição. A aplicação valida essas sugestões e, **quando a norma permitir**, administra a branch da sessão, aplica somente alterações válidas e realiza o commit automaticamente após cada trabalho concluído. A aprovação do operador incide sobre a integração posterior à `develop`, não sobre cada commit intermediário; este texto substitui a política anterior de três aprovações.
- Saída inconsistente, tentativa de violar a norma ou exclusão fora do escopo autorizado causa `BLOCKED`. Alterações destrutivas devem ser evidenciadas para revisão antes de sua eventual integração à `develop`; a autorização de mesclagem não dispensa critérios de validação.
- A sintaxe ilustrativa `DOC`, `SECTION`, `ITEM` **não constitui gramática homologada**. Escape, multiline, badges, links, tabelas, Mermaid, hashes, alterações por trecho e ida/volta sem perdas continuam pendentes em [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Decisão funcional — commits automáticos e continuidade da sessão (2026-10-21)
**Origem:** determinação explícita do responsável, posterior à decisão anterior de aprovações independentes.

**Resultado esperado:** na futura funcionalidade compatível com a norma, a aplicação administra uma branch de trabalho para cada **sessão**, não para cada inferência. Após validar cada resultado documental, cria automaticamente o commit correspondente nessa mesma branch, preservando o histórico incremental. A sessão permanece ativa para novas inferências com prompts novos ou reutilizados.

**Critérios verificáveis de produto:**
- Uma sessão identifica repositório, origem em `develop`, SHA de referência, branch associada, especialização e conjunto de tentativas.
- Ao executar novas inferências na sessão ativa, a aplicação reutiliza a branch e considera os commits anteriores; não abre outra branch simplesmente por receber novo prompt.
- Uma saída inválida não gera commit, não sobrescreve alterações anteriores nem publica conteúdo não verificado.
- Cada commit de resultado validado pode ser associado à inferência, ao prompt, ao modelo, à baseline usada e às mudanças observáveis.
- O operador consulta os commits e o diff acumulado antes de solicitar ou aprovar a integração à `develop`; a ausência de aprovação impede o merge.
- A autorização para **custos de inferência** continua independente da autorização de integração Git.

**Impedimento normativo obrigatório:** o comportamento acima descreve o produto pretendido, **não autoriza sua execução hoje**. A referência [ProjectCharter/1.0/main](https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0) proíbe gravações de agentes fora de `develop`. Até a resolução verificável de [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md), a execução deve permanecer em modo de simulação/preparação local, sem criar/atualizar a branch de sessão remotamente.

## Critério de aceite em três etapas — decisão 10 (2026-10-21)
Para a primeira entrega pretendida, cada proposta documental deve passar por **validação estrutural automática**, **avaliação crítica de outro LLM** e **revisão humana** antes de ser considerada aceita para integração.

- Rejeitar alterações com violação de escopo, links ou integridade, independentemente dos pareceres dos modelos.
- Conservar a identificação do LLM gerador e do LLM avaliador, a norma, o contexto analisado, as diferenças propostas, as evidências de validação e a decisão humana.
- Exibir à pessoa revisora tanto as divergências quanto a justificativa e as limitações apontadas pela segunda avaliação; o segundo LLM não aprova automaticamente o conteúdo.
- Se a segunda avaliação obrigatória não estiver disponível ou autorizada sob as condições de custo e segurança, o processo **não é concluído como validado**; registrar a pendência sem contornar bloqueios financeiros.
- A autorização humana do resultado é distinta da autorização para integrar à `develop`; a política de commits automáticos em branch de sessão permanece dependente da [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

Relações: [OPEN-002 — Decisão 10](../open-questions/OPEN-002-fidelidade-e-custo.md) e [OBJ-004 — Verificações](OBJ-004-evidencias-de-testes-automatizados.md). Isso não atesta implementação ou fidelidade homologada.

## Tratamento de divergências críticas na validação — decisão 12
Uma divergência relevante apontada pelo segundo LLM exige correção da proposta e **nova passagem pela validação estrutural e avaliação crítica de outro LLM**, antes de permitir a revisão humana de aceite. **Propostas com divergência relevante pendente não podem ser consideradas validadas nem gerar commit de resultado validado**.

Falhas estruturais, violações normativas e alterações fora de escopo permanecem bloqueadoras, ainda que os modelos concordem. Ausência de reavaliação, recursos financeiros autorizados ou evidências mantém o trabalho bloqueado, com histórico preservado.

A aprovação de integração Git permanece humana e separada; a execução de commits em branches de sessão ainda depende da resolução de [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md). Relação: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Reapresentação integral do contexto original a cada correção — decisão 13
Quando uma divergência relevante impedir o aceite, a próxima proposta deve ser obtida com o **mesmo conjunto original de documentos, prompt, tarefa, normas, identificadores e hashes/SHAs da primeira análise**, acrescido **da última proposta e da observação mais recente do revisor LLM**.

O ciclo permite **no máximo quatro iterações de correção e reavaliação após a tentativa inicial**, enquanto as condições financeiras, de segurança e de contexto permitirem. Cada proposta corrigida exige seus próprios testes estruturais e nova avaliação por modelo distinto do gerador. Propostas com divergência relevante ou erro normativo/estrutural não são declaradas válidas, não geram commit de resultado validado nem avançam à revisão humana.

Se o conteúdo integral não couber na janela de contexto do modelo, ou não houver autorização financeira para nova chamada, a execução dependente fica bloqueada sem simplificar silenciosamente a baseline. A autorização de integração Git continua humana e separada, e a publicação em branch de sessão depende de [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

Relação: [OPEN-002 — Decisão 13](../open-questions/OPEN-002-fidelidade-e-custo.md). Trata-se de comportamento pretendido, não implementado.

## Critério de parada após quatro correções — decisão 14
O fluxo pretendido admite a geração inicial e **até quatro versões corrigidas**, cada qual validada estruturalmente e submetida novamente a um LLM revisor distinto do gerador. Uma resposta válida encerra antecipadamente as correções e é encaminhada à revisão humana.

Se a quarta correção não resolver a divergência relevante, o trabalho fica **bloqueado**, preservando a proposta e os motivos, sem gerar quinta correção automática, commit de resultado validado ou integração Git. A continuidade também pode ser bloqueada antes por segurança, contexto insuficiente ou custo não autorizado. Relação: [OPEN-002 — Decisão 14](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Sessão com branch nominada e PR após encerramento — decisão 15
O fluxo de trabalho pretendido para a aplicação passa a considerar o **repositório de destino escolhido pelo operador**, a seleção do nível documental e um identificador próprio por sessão. Quando uma futura política normativa permitir, a sessão terá **`feature/doc/{session-id}` derivada de `develop` do repositório escolhido**, preservada para todas as inferências e commits intermediários verificados até seu encerramento.

**Critérios verificáveis de produto:**
- A sessão evidencia repositório escolhido, nível documental, ID próprio, SHA inicial de `develop`, branch prevista e histórico de tentativas e commits.
- Correções e resultados validados pertencentes à mesma sessão permanecem na mesma branch; propostas inválidas não geram commits declarados válidos.
- No fechamento, o usuário visualiza o diff **acumulado**, o histórico de verificações e aprova expressamente o conjunto de mudanças **antes da criação da PR** para `develop`.
- A PR identifica sessão, branch de origem e destino e resultados dos checks previstos em [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md). Falha normativa, estrutural ou de segurança impede considerar o trabalho apto para merge.
- **Aprovação para abrir PR não é aprovação para mesclar.** Mesmo com checks concluídos, integração à `develop` depende de decisão humana específica e das proteções efetivas da branch.

O procedimento está **bloqueado para execução remota por agentes**, porque a [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md) registra conflito com a referência normativa atual. Não autoriza operação Git fora de `develop`, criação de PR, alteração da norma ou execução de workflow nesta etapa.

## Redefinição autorizada quando o contexto integral não cabe — decisão 16
Quando o material original completo mais a última proposta e o parecer não couber no contexto comprovadamente disponível, o trabalho **não pode continuar nem ser publicado como validado**. A aplicação deverá comunicar a limitação e aguardar **redefinição explícita do escopo e de nova baseline pelo operador**, sem trocar o modelo, resumir ou truncar informações silenciosamente.

A tarefa redefinida deve possuir identidade e baseline próprias, com relação auditável à tentativa interrompida. A redefinição não transforma o resultado anterior em válido e não pode reiniciar automaticamente o teto de **quatro ciclos de correção** da mesma tarefa. Uma tarefa nova é submetida novamente a validação estrutural, avaliação crítica de LLM distinto e autorização humana final.

A aprovação de nova tarefa não equivale a autorização financeira de inferências nem de integração Git; a [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md) continua bloqueando operações em branches de sessão na norma atual. Relação: [OPEN-002 — Decisão 16](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Uma sessão contínua, configuração por inferência e Markdown gerado pela aplicação — decisão 17
A sessão mantém **repositório escolhido, nível de documentação, ID, histórico e branch `feature/doc/{session-id}` pretendida**; entre inferências dentro da mesma sessão, o usuário pode mudar **modelo autorizado, tamanho de contexto admissível e prompt**, sem abrir nova sessão ou reescrever automaticamente resultados previamente validados.

**Entrega esperada em cada inferência:**
- O LLM devolve, no formato textual DSL previsto, **título e descrição sugeridos para o commit, identificação de todos os arquivos afetados** e as operações/documentação correspondentes em DSL. A descrição, os caminhos e as operações são validados pela aplicação.
- A **aplicação**, não o LLM, converte a representação DSL em **Markdown completo**, preservando inclusive **tabelas, links, badges, blocos de código, diagramas Mermaid, representações de gráficos e recursos vinculados**. Elementos externos só são declarados preservados mediante verificação de suas referências e recursos reais.
- O conteúdo não modificado, os identificadores estáveis, a norma, os limites da especialização e as baselines verificadas são preservados; perda de informações, recursos gráficos sem representação ou conversão incompleta **bloqueiam aplicação/commit**.
- Cada tentativa associa a configuração efetivamente autorizada (modelo, contexto, prompt, custo), a resposta DSL original, os arquivos Markdown resultantes, as validações estruturais, o parecer secundário e o eventual commit de conteúdo válido.
- Divergência relevante exige correção e nova avaliação, limitadas a **quatro correções por inferência/tarefa inicialmente autorizada**; mudar configuração em inferências posteriores não concede um quinto ciclo de reparo disfarçado para a mesma tarefa bloqueada.

Os commits automáticos na branch de sessão, a criação de PR após o encerramento aprovado e a revisão humana específica antes de integrar à `develop` são **comportamentos futuros**, não permissões vigentes, conforme [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md). Relação de fidelidade: [OPEN-002 — decisão 17](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Retomada da sessão após checks reprovados — decisão 18
Uma reprovação dos checks obrigatórios da PR **não encerra definitivamente a sessão**: o operador poderá reabri-la para correção, preservando a identidade da sessão, a branch prevista e a mesma PR. Somente resultados documentais validados podem gerar novos commits na implementação futura. Após cada atualização, os checks são repetidos; para encerrar novamente exige-se **nova aprovação humana**, e o merge exige **autorização humana independente**.

O limite de quatro correções **por inferência/tarefa**, os controles financeiros e a revisão independente não são dispensados por essa retomada. A criação de branches e a atualização de PRs continuam bloqueadas para agentes pela norma atual. Relação: [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Limite da sincronização da sessão — decisão 19
A branch prevista `feature/doc/{session-id}` deriva de `develop` **somente na abertura da sessão** e permanece como referência de trabalho entre inferências. A aplicação não precisa importar mudanças feitas posteriormente em `develop` para essa branch. Ao constatar novas alterações **na mesma branch remota da sessão**, o objetivo é atualizar a cópia local sem perda de trabalho, quando seguro.

A [NOBJ-003](../non-objectives/NOBJ-003-nao-gerenciar-concorrencia-na-branch-da-sessao.md) exclui coordenação de edições simultâneas por terceiros e resolução automática de conflitos externos: uma divergência que inviabilize atualização segura **bloqueia a operação afetada**, sem sobrescrita. Isso não dispensa as validações de PR, checks obrigatórios e aprovação humana para integrar à `develop`. A habilitação de branch de sessão ainda depende da compatibilização normativa em [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Critério de materialidade na revisão secundária — decisão 20
Somente divergências **materiais** — perda/alteração de requisitos, mudança de significado, contradição, falha de integridade ou violação normativa — obrigam **corrigir e reavaliar** uma proposta. O parecer deve indicar local, evidência, classificação e encaminhamento. Melhoria puramente editorial, comprovadamente sem impacto no conteúdo e nas exigências documentais, é **registrada para revisão humana** e não determina por si nova rodada automática.

Falhas estruturais ou normativas continuam bloqueadoras mesmo sem divergência entre os modelos. A impossibilidade de comprovar que uma diferença é somente editorial impede classificá-la como inofensiva por conveniência. Até sanar uma falha material ou encerrar justificadamente a tarefa como bloqueada, **não produzir commit de resultado validado**. O limite é **quatro correções após a tentativa inicial**, com contexto integral, segunda avaliação independente, autorização financeira por chamada e aprovação humana de integração Git preservados. Relação: [OPEN-002 — Decisão 20](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Configuração inicial da sessão preservada — decisão 26
A sessão preserva **sua identidade, repositório, especialização, baseline documental, histórico e versão dos padrões iniciais de gerador e revisor**. Alterações na configuração geral realizadas depois da abertura **não** atualizarão automaticamente a sessão, não apagarão seu histórico e terão efeito apenas em **novas sessões**.

Na mesma sessão, o operador poderá **selecionar explicitamente por inferência** modelos/motores homologados, tamanho de contexto e prompts, conforme decisão 17. Registrar a origem e a versão do padrão inicial, a operação/papel, os valores resolvidos e os valores efetivos de cada tentativa. Uma escolha pontual **não** modifica o perfil geral ou o perfil inicial da sessão, nem substitui retroativamente escolhas anteriores.

A continuidade da sessão não significa aceite de configurações que perderam elegibilidade: se o gerador ou o revisor não atender aos requisitos atuais, **bloquear inferência remunerada ou aceite/commit de resultado validado** até decisão válida. A segunda avaliação distinta, as quatro correções máximas por tarefa, autorização financeira por chamada, revisão humana e restrições normativas de Git permanecem. Relações: [OBJ-001](OBJ-001-agentes-documentais-especializados.md), [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md), [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).

## Continuidade da mesma tarefa após falha da revisão — decisão 28
Após uma **geração autorizada e concluída**, se a revisão secundária falhar por indisponibilidade do revisor, o produto deverá **preservar a mesma sessão e tarefa**, a **proposta gerada sem modificação**, o material original integral, as baselines/hashes, os modelos e prompts usados, as validações realizadas, os custos e o histórico. O estado é **avaliação pendente**, não resultado aprovado.

**Critério positivo:** diante de **escolha expressa de revisor distinto, homologado e disponível**, integridade do contexto e da proposta comprovada e **nova autorização financeira** para a revisão, executar **somente a avaliação pendente**, sem repetir automaticamente a geração. Anexar parecer, identificação do revisor e custo ao mesmo histórico, mantendo as configurações fixadas na abertura da sessão.

**Critério negativo:** se a integridade documental, as baselines, a elegibilidade ou a autorização não puderem ser verificadas, **bloquear**, informar e solicitar uma decisão explícita; não trocar modelos, redefinir a tarefa ou gerar novamente de forma silenciosa. A avaliação pendente não autoriza commit de proposta ainda não validada, integração Git ou gasto adicional.

Uma revisão retomada não constitui automaticamente um **ciclo de correção**; caso haja divergência material, a correção seguirá o saldo existente do teto de **até quatro correções por tarefa** com nova revisão independente. Referências: [OPEN-002 — decisão 28](../open-questions/OPEN-002-fidelidade-e-custo.md), [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md) e [OBJ-007](OBJ-007-comparacao-controlada-de-modelos.md).

## Tarefa pendente recuperável após reinicialização — decisão 29
A aplicação futura deverá permitir ao **operador autorizado** recuperar uma **tarefa de avaliação pendente** após fechamento ou reinicialização do aplicativo, sem alterar **identificadores de sessão/tarefa**, repositório, especialização, **baseline**, proposta original, histórico de inferências, custos e saldo de **até quatro correções**. Encerrar o processo da aplicação não equivale a aprovar, descartar ou encerrar explicitamente uma sessão documental.

**Critério de continuidade:** para uma geração já concluída e íntegra cujo revisor falhou, a tarefa recuperada permite **retomar somente a revisão pendente**, após autenticação, verificação do contexto e da proposta, escolha/confirmacão expressa de revisor homologado e distinto, revalidação dos controles e **nova autorização financeira para a chamada remunerada**. O parecer deve se vincular à **mesma tarefa**. Não repetir a geração por padrão.

**Critério de bloqueio:** sem os dados originais ou evidência suficiente de integridade, permissão e condições de uso, a retomada não é possível; informar o motivo e solicitar tratamento explícito. Não usar como contexto atual alterações posteriores em `develop` para substituir silenciosamente a baseline da tarefa. A recuperação não autoriza commit de resultado não validado, nem branches/PRs fora da restrição normativa.

Os mecanismos de armazenamento e interface ficam fora do Charter. Relações: [OPEN-002 — decisão 29](../open-questions/OPEN-002-fidelidade-e-custo.md), [OBJ-008](OBJ-008-historico-local-de-inferencias.md), [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md) e [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Trabalho ativo suspenso com tarefas ainda pendentes — decisão 30
O operador poderá **encerrar o trabalho ativo** com tarefas não resolvidas, deixando a **mesma sessão suspensa e não concluída para publicação**. A suspensão não transforma uma proposta em resultado validado, não significa cancelamento de tarefa, não é aprovação final e **não desencadeia PR, commit de resultado pendente, merge ou gasto**.

O produto deverá conservar o identificador de sessão e das tarefas, o repositório/nível documental, baselines e hashes, as propostas intactas, avaliações realizadas ou pendentes, os custos autorizados, o histórico de interrupções e o saldo de até **quatro correções por tarefa**. Na reabertura, apresentar a situação real e permitir ao operador autorizado resolver as pendências na mesma sessão, **sem repetir automaticamente gerações**, respeitando a integridade da decisão 29 e a revisão por segundo LLM distinto.

A **conclusão formal para publicação** requer que as tarefas impeditivas sejam resolvidas por validação ou retirada/cancelamento **expressamente autorizado e rastreável**, sem promover o material cancelado a aprovado; exige ainda apresentação do conjunto final e **nova aprovação humana explícita**. Ela permanece separada da autorização posterior de integração Git e condicionada à regularização normativa em [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

A suspensão/reabertura não altera perfis inicialmente fixados, não autoriza chamadas pagas, não reinicia correções, não substitui documentos normativos nem escreve fora de `develop`. Relações: [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md), [OBJ-008](OBJ-008-historico-local-de-inferencias.md), [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md).

## Encerrar sessão sem entregas documentais publicáveis — decisão 31
O objetivo de conclusão da sessão inclui o caso em que o operador **cancelou expressamente todas as tarefas restantes**, **não há alterações documentais validadas para publicar** e aprova **formalmente o encerramento sem publicação**. A aplicação deverá reconhecer o desfecho da sessão como **"concluída sem publicação"**, distinto de suspensão, tarefa cancelada, sessão abandonada ou conclusão com alterações aptas à publicação.

**Critério positivo:** demonstrar o inventário de tarefas e o conjunto final vazio de alterações publicáveis, preservar a identidade da sessão, baseline, histórico, propostas descartadas para aceite, custos e limites das correções, registrar as decisões humanas de cancelamento e a aprovação final. **Não criar PR, commit de resultado documental, merge ou gasto adicional** como consequência do encerramento vazio.

**Critério negativo:** não tratar uma suspensão, avaliação pendente ainda não cancelada, cancelamento automático por inatividade ou alterações previamente validadas ainda destinadas à publicação como autorização para encerrar sem publicação. Quando existir trabalho válido publicável, aplicar a governança apropriada em vez de descartá-lo tacitamente.

O estado de uma sessão não altera os **status de documentos normativos** e não apaga os históricos de inferência, sob retenção até exclusão manual autorizada. Preservam-se controles de segurança, revisão por modelo distinto para qualquer resultado validado, até quatro correções por tarefa, autorização financeira por chamada e a restrição normativa de escrita Git. Relações: [OPEN-003 — decisão 31](../open-questions/OPEN-003-politica-de-branches.md), [OBJ-004](OBJ-004-evidencias-de-testes-automatizados.md) e [OBJ-008](OBJ-008-historico-local-de-inferencias.md).
