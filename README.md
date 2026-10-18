# Doc Arquitetura — Aplicação Debian

Ferramenta Go de análise experimental que baixa a referência normativa [doc-arquitetura/main](https://github.com/ahhvahh/doc-arquitetura/tree/main/ProjectCharter/1.0), gera um bundle DSL e ajuda a revisar a documentação de outros repositórios GitHub.

> Estado: MVP. A auditoria completa de Charter, Scope e Implementation e a gestão integrada de modelos ainda estão em desenvolvimento.

## Estrutura

- `src/main.go`, `src/sync.go`, `src/sync_test.go`: CLI Go e testes.
- `src/dsl/`: gramática e regras Charter existentes.
- `src/docs/`: arquitetura e plano de implementação do MVP.
- `src/package.yaml`: contrato de compilação, testes e pacote Debian (`schema: 1`).

A documentação normativa **não é distribuída no pacote**; `sync` obtém `ProjectCharter/1.0` da branch `main` do repositório de referência. Para acessar a referência privada, configure a autenticação Git no computador de execução.

## Compilação local

Requer Go 1.22 ou superior e Git:

```sh
cd src
go test ./...
go build -o doc-arquitetura .
./doc-arquitetura sync -target https://github.com/OWNER/REPO -workspace ./workspace
```

## Compatibilidade com A-Bioma Debian Package Builder

O script externo `atualizar-e-compilar.sh` exige um `projetos.conf` e o `build-package.sh` [do empacotador A-Bioma](https://github.com/ahhvahh/a-bioma-pkg-debian) no mesmo diretório. Acrescente a linha abaixo ao **arquivo de projetos do orquestrador**, não ao repositório da aplicação:

```text
https://github.com/ahhvahh/doc-arquitetura-app-debian.git
```

Execute `./atualizar-e-compilar.sh` no diretório do orquestrador; escolha o projeto, a branch `develop` e a arquitetura `amd64`. Ele localizará automaticamente `src/package.yaml`, executará compilação/testes pelo builder e salvará o `.deb` em `packages/`, com histórico em `logs/`. O script **gera** o pacote; não realiza a instalação.

Após a geração do `.deb`, instale-o explicitamente com `sudo apt install ./caminho/para/o/pacote.deb`.

O pacote distribui:
- `/usr/bin/doc-arquitetura` — CLI compilada com `CGO_ENABLED=0`.
- `/usr/share/doc-arquitetura-app/dsl/charter.pcdsl` — regras Charter da versão atual.

A opção `-dsl` permite fornecer um arquivo alternativo; a variável `DOC_ARQUITETURA_DSL` ajusta o caminho padrão. `llama-server` e modelos GGUF devem ser instalados separadamente.

## Limitações

A conversão DSL preserva hashes, mas não prova equivalência semântica. O comando `review-bundle` limita o conteúdo enviado ao modelo e **não** corresponde a uma auditoria completa. O projeto ainda não inclui `llama-server` ou modelos locais no pacote Debian. Não foi publicado um serviço systemd.
