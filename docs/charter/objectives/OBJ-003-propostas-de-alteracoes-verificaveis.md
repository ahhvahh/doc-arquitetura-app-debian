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
- Publicação requer aprovação explícita, com resultado auditável.
- A entrega final prevê uma branch de tarefa distinta, nunca um push direto para `main`; sua publicação depende de atualização da política vigente.
## Dependências
- [NOBJ-001](../non-objectives/NOBJ-001-nao-publicar-sem-aprovacao.md).
- [OPEN-003](../open-questions/OPEN-003-politica-de-branches.md).

## Entrega textual DSL por arquivo — decisão do responsável
- A resposta do agente deverá ser **texto DSL**, sem envelope JSON ou Markdown diretamente editável.
- O contrato planejado `BIOMA-DSL-CHANGESET/2.0` reúne: repositório, branch de origem `develop`, SHA do projeto-alvo, SHA normativo, sugestão de branch, título e descrição.
- Cada arquivo afetado aparecerá em bloco próprio, delimitado por `=== FILE ADD|MODIFY|DELETE caminho ===` e `=== END FILE ===`, com `BASE-SHA256` quando já existir. O resultado termina em `=== END CHANGESET ===`.
- A aplicação Go, e não o agente, valida o formato, caminhos, escopo, hashes e permissões, converte DSL em Markdown e preserva conteúdo fora das alterações.
- O agente apenas sugere nome de branch, título e descrição. A aplicação sanitiza o nome, cria branch local, apresenta diff e exige aprovações independentes para aplicar, commitar e publicar.
- Exclusão não deve ser executada sem aprovação específica; saída inconsistente ou ausência de norma obrigatória causa `BLOCKED`.
- A sintaxe ilustrativa `DOC`, `SECTION`, `ITEM` **não constitui gramática homologada**. Escape, multiline, badges, links, tabelas, Mermaid, hashes, alterações por trecho e ida/volta sem perdas continuam pendentes em [OPEN-002](../open-questions/OPEN-002-fidelidade-e-custo.md).
