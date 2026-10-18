# Plano de implementacao DSL/1.0

## MVP registrado
- [x] Projeto sob ProjectCharter/1.0/DSL na develop.
- [x] CLI Go com comandos sync, convert, model e review-bundle; preserva comandos preexistentes do MVP de verificacao.
- [x] Clone de repositório cliente develop e referencia main/develop selecionada.
- [x] Conversao de todos os Markdown de ProjectCharter/1.0 para bundle textual PCDSL-DOCS 1 com SHA-256 e status.
- [x] Integração opcional com llama-server usando modelo GGUF escolhido pelo usuario.

## Evolucao necessária
- [ ] Parser formal para bundle e prova de preservacao/equivalencia de conteúdo.
- [ ] Comparativo de tokens Markdown x DSL por tokenizador do modelo escolhido.
- [ ] Auditoria exaustiva: PCH/ACT, OBJ/NOBJ, CON/ASM, OPEN e DSC, badges, links, IDs e estados.
- [ ] Recuperação seletiva por documento/ID para economizar contexto sem truncar evidencias.
- [ ] Catalogo de modelos, quantizações, tamanhos em disco/RAM e verificacao do download.
- [ ] Perguntas iterativas, diffs e aprovacao antes de qualquer alteração documental.
- [ ] Interface local TUI/GUI e testes com repositório cliente real.

**Gate:** READY_FOR_SCOPE so pode ser reportado apos revisar *toda a baseline normativa* pela regra charter-review.md e comprovar ausência de bloqueadores. Parecer isolado do modelo equivale a REVIEW_REQUIRED, nunca aprovação.