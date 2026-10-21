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
