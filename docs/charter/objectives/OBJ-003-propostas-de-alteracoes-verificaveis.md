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
