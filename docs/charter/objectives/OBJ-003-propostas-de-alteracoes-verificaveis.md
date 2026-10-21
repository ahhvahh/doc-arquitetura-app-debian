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
