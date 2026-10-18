# Arquitetura — avaliador local

1. O operador informa a URL GitHub no terminal. A ferramenta lista branches e clona a escolhida em sessão isolada, sem submódulos. Somente `develop` permite edição, commit local e push; outras branches ficam em leitura.
2. O Go clona ahhvahh/doc-arquitetura **main** (ou develop apenas quando selecionado como experimento), para workspace/reference.
3. O conversor percorre todos os .md sob ProjectCharter/1.0, produzindo reference.pcdsl e client.pcdsl. Cada DOC inclui caminho relativo, hash SHA-256 do original, status e conteudo textual. Arquivos originais permanecem nos clones.
4. O servidor llama.cpp, opcional, baixa por -hf o modelo GGUF indicado e o atende apenas em 127.0.0.1.
5. O comando review-bundle usa a API local /v1/chat/completions para uma **analise preliminar**. O comando analyze do MVP permite carregar um prompt especializado original.
6. Menu de terminal, chat com llama-server, propostas de arquivos Markdown por etapa, revisão por diff, aplicação aprovada e comandos separados de commit/push disponíveis como MVP. Auditoria normativa completa, recuperação seletiva por documento/ID e gate READY_FOR_SCOPE permanecem pendentes.

**Segurança:** o modelo não executa comandos Git nem shell. Propostas Markdown passam por validação de caminho e symlinks antes de alterar arquivos locais; commit e push exigem aprovação explícita. O push somente para develop verifica o HEAD remoto e usa avanço sem force. Credenciais não devem ser incorporadas à URL do repositório. Modelos remotos devem ser escolhidos e suas licencas verificadas pelo operador.

**Restricao importante:** a compactação estrutural é parcial e nao substitui a leitura integral da documentacao normativa. O clone de main e a referencia; se desenvolver regras em develop, registrar REFERENCE_DIVERGENCE.