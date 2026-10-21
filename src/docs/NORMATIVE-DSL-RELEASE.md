# Atualização da referência normativa via GitHub Release

O comando `doc-arquitetura dsl update` consulta os Releases publicados no repositório privado `ahhvahh/doc-arquitetura`, seleciona o mais recente com prefixo `projectcharter-1.0-` e baixa os assets `projectcharter-1.0.zip` e `projectcharter-1.0.zip.sha256`.

O download usa as credenciais já configuradas em `github.json` (token com permissão de leitura em Contents para o repositório normativo). O conteúdo é validado com SHA-256 do arquivo ZIP, hashes do manifesto, contrato do DSL e correspondência do SHA do Release. Caso não exista Release válido, o comando falha: não substitui automaticamente por uma cópia da `develop`, nem recompila Markdown.

Os arquivos `charter.dsl`, `scope.dsl`, `implementation.dsl`, `source.dsl`, `prompts.dsl` e `manifest.json` ficam em `~/.cache/doc-arquitetura/dsl/<SHA>/` (ou na configuração XDG/cache da aplicação). O arquivo `release-current.json` identifica a versão selecionada. `dsl status` informa o cache de Release instalado.

`dsl compile` permanece como ferramenta de compilação legada local de `PCDSL-DOCS 1`, enquanto o formato do Release é `PROJECTCHARTER-DSL/1`. São contratos distintos: consumidores da DSL normativa precisam interpretar o formato do Release. Os downloads não modificam documentação nem branches dos repositórios GitHub.

O comando `reference update` permanece disponível para operações que precisem dos arquivos Markdown originais. A mudança neste ciclo se limita ao download distribuído de DSL, sem alterar fluxo de `sync` nem inferência.
