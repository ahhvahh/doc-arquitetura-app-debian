# Casos de teste — conversão de Markdown para DSL

Esta pasta possui **20 documentos Markdown de teste**, distribuídos em quatro áreas com cinco documentos cada:

- [Charter](docs/charter/): propósito, contexto PCH, OBJ, CON e OPEN
- [Scope](docs/scope/): SCP, TYPE, ALG e FLW
- [Implementation](docs/implementation/): LIB, DEP, SVC e APP
- [Development / Source](docs/src/): LIB, SVC, APP e CMD

As entradas e as expectativas estão registradas em [manifest.json](manifest.json).
Os badges SVG nas amostras são marcadores de texto sintéticos e seus destinos não são imagens instaladas: devem ser removidos pelo compactador, sem consulta à rede.
Os documentos exemplificam estados diferentes (backlog, refinement, refined, implemented, cancelled), mas não validam nenhum gate semântico real.

## Executar

A partir da pasta `src`:

```sh
go test ./... -run 'TestStageMarkdownDSLGolden|TestWholeProjectMarkdownDSLCoversFourStages|TestStageMarkdownDSLDetectsModifiedContent|TestFixtureMarkdownDSLExcludesBackup' -v
```

O GitHub Actions já executa `go test -race -count=1 -covermode=atomic ... ./...` em cada commit na `develop` e em PRs `develop` → `main`.

## Regras verificadas

1. Cada etapa contém exatamente cinco entradas Markdown.
2. A conversão de cada etapa em `PCDSL-DOCS 1` preserva origem, ordem determinística, caminho do arquivo, SHA-256 do Markdown bruto e status extraído.
3. O texto de saída é comparado integralmente ao resultado esperado, calculado de forma independente com base no manifesto (inclusive quantidade de bytes UTF-8).
4. Os badges SVG não entram na DSL compactada.
5. A conversão da raiz de quatro áreas inclui os 20 documentos.
6. Repetir a conversão produz bytes idênticos.
7. Alterar o badge de status muda o resultado.
8. Documentos em `backup/` ficam fora da DSL ativa.

**Limite:** o programa atual usa o formato preliminar `PCDSL-DOCS 1`. Estes testes verificam esse conversor existente; ainda não validam cadastro de agentes OpenAI, seleção remota de prompts nem geração de changesets.
