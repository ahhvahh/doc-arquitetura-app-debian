# CON-002 — Proteção de credenciais

![Tipo: CON](../../assets/documentation-style/badges/badge-type-con.svg) ![Status: refinement](../../assets/documentation-style/badges/badge-status-refinement.svg) ![Versão: v1](../../assets/documentation-style/badges/badge-version-v1.svg)

## Restrição
Credenciais de acesso à OpenAI e ao GitHub não podem ser transmitidas como contexto documental nem gravadas em commits, logs ou arquivos de documentação.
## Origem verificável
Exigência de segurança necessária ao cadastro de chave solicitado pelo responsável.
## Motivo e impacto
Exposição de credenciais possibilitaria acesso não autorizado e uso indevido dos serviços.
## Verificação
Os registros e conjuntos documentais exportados não contêm segredos; os controles locais de acesso são auditados.
## Relações
- [OPEN-001](../open-questions/OPEN-001-politica-credenciais.md).
## Condição de revisão
Não é permitida flexibilização sem avaliação e decisão formal documentada.
