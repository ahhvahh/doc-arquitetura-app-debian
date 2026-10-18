# Doc Arquitetura — aplicação Go

A CLI, seus testes e os dados locais estão diretamente em `src/`, conforme a organização solicitada para o A-Bioma Debian Package Builder.

## Utilização

```sh
cd src
go test ./...
go build -o doc-arquitetura .
./doc-arquitetura routes
./doc-arquitetura sync -target https://github.com/OWNER/REPO -workspace ./workspace
```

O modo `sync` baixa a referência `doc-arquitetura/main` e a branch `develop` do repositório-alvo, convertendo os Markdown da árvore `ProjectCharter/1.0` em `reference.pcdsl` e `client.pcdsl`. Os clones ficam no workspace local. Autenticação Git é necessária caso o acesso ao repositório seja privado.

## Distribuição Debian

`package.yaml` é o contrato de build consumido pelo `build-package.sh` externo. O modo `package` valida os testes Go e publica o CLI em `/usr/bin/doc-arquitetura`, junto da DSL Charter em `/usr/share/doc-arquitetura-app/dsl/charter.pcdsl`.

`routes`, `check`, `analyze`, `clone`, `serve`, `sync`, `convert`, `model` e `review-bundle` são comandos do MVP. `check` aplica as regras Charter locais; `review-bundle` ainda é preliminar. Modelos GGUF e `llama-server` não integram o pacote.

Consulte [arquitetura](docs/ARCHITECTURE.md) e [plano de implementação](docs/IMPLEMENTATION-PLAN.md).

## Novos comandos locais

Consulte [Modelos e DSL](docs/AGENT-COMMANDS.md) para instalar modelos GGUF, associar modelos a agentes, consultar a referência `main` e compilar DSL versionada.

```sh
go run . models catalog
go run . reference update
go run . dsl update
go run . dsl status
```
