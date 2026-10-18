# Arquitetura — avaliador local

1. O operador informa o repositório do cliente. O Go clona **somente develop** sem submodulos, para workspace/client.
2. O Go clona ahhvahh/doc-arquitetura **main** (ou develop apenas quando selecionado como experimento), para workspace/reference.
3. O conversor percorre todos os .md sob ProjectCharter/1.0, produzindo reference.pcdsl e client.pcdsl. Cada DOC inclui caminho relativo, hash SHA-256 do original, status e conteudo textual. Arquivos originais permanecem nos clones.
4. O servidor llama.cpp, opcional, baixa por -hf o modelo GGUF indicado e o atende apenas em 127.0.0.1.
5. O comando review-bundle usa a API local /v1/chat/completions para uma **analise preliminar**. O comando analyze do MVP permite carregar um prompt especializado original.
6. Auditoria deterministica completa, indexacao semantica, perguntas ao usuario, diffs aprovados, interface grafica e gate READY_FOR_SCOPE estao planejados, nao implementados.

**Seguranca:** sem escrita remota; sem execução de scripts do cliente; evitar credenciais em URLs; impedir dependencia de rede externa no servidor local. O LLM não tem permissao de editar. Modelos remotos devem ser escolhidos e suas licencas verificadas pelo operador.

**Restricao importante:** a compactação estrutural é parcial e nao substitui a leitura integral da documentacao normativa. O clone de main e a referencia; se desenvolver regras em develop, registrar REFERENCE_DIVERGENCE.