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
- [x] MVP de chat no terminal, diff de propostas e aprovações distintas para aplicação local, commit e push.\n- [ ] Completar validações semânticas, gates e recuperação seletiva por documento para toda a árvore.
- [x] Menu e chat CLI local sem interface Web.\n- [ ] Teste integrado com GitHub real e llama-server; concluir testes por etapa e sessão.

**Gate:** READY_FOR_SCOPE so pode ser reportado apos revisar *toda a baseline normativa* pela regra charter-review.md e comprovar ausência de bloqueadores. Parecer isolado do modelo equivale a REVIEW_REQUIRED, nunca aprovação.

## Administração root — gate pré-serviço

- [ ] Implementar `doc-arquitetura admin` TUI somente `geteuid()==0`, sem sudo interno, setuid ou API administrativa exposta.
- [ ] Implementar esquema root-owned `/etc/doc-arquitetura-app/config.yaml`, gerenciamento de contas/grupos e allowlist verificável.
- [ ] Importar e validar GitHub App PEM ou SSH key de serviço sem imprimir/gravar segredos em logs/config.
- [ ] Implementar `doc-arquitetura-app.socket` systemd com `SocketGroup=docarq-clients`; daemon `User=docarq-svc`, `Group=docarq-runtime` e `SO_PEERCRED`.
- [ ] Impedir paths legados de execução Git/llama pelo cliente, preparar storage somente do daemon.
- [ ] Validar usuário fora do grupo, usuário fora da allowlist, operações entre usuários, acesso indevido ao socket privado, edição root-only e startup sem credencial.

A [especificação administrativa](ADMIN-SETUP.md) é requisito antes de ativar o daemon. Não há implementação funcional neste commit documental.
