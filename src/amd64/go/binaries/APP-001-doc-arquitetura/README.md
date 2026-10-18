# Doc Arquitetura — CLI Go (MVP migrado)

> Estado: protótipo experimental, sem empacotamento Debian nem auditoria integral de Charter, Scope e Implementation.

Este aplicativo consulta `ahhvahh/doc-arquitetura` na branch `main`, compila localmente os Markdown da árvore `ProjectCharter/1.0` em bundle DSL e compara com a documentação de um projeto GitHub na branch `develop`.

## Requisitos
Go 1.22+, Git e, opcionalmente, `llama-server` (llama.cpp) para modelos GGUF locais.

## Uso (neste diretório)
```sh
go test ./...
go build -o doc-arquitetura .
./doc-arquitetura sync -target https://github.com/OWNER/REPO -workspace ./workspace
# Instale separadamente llama.cpp e escolha o modelo GGUF:
./doc-arquitetura model -hf OWNER/MODEL-GGUF:Q4_K_M -port 8080
# Em outro terminal:
./doc-arquitetura review-bundle -workspace ./workspace -api http://127.0.0.1:8080
```

`sync` clona `doc-arquitetura/main` e o projeto-alvo `develop`, mantendo os originais no workspace. Gera `reference.pcdsl` e `client.pcdsl` com hashes SHA-256. O comando `-norm-ref develop` é experimental, não normativo.

Comandos herdados: `check`, `analyze`, `routes`, `clone`, `serve`, `convert`. `check` ainda cobre apenas regras Charter em `dsl/charter.pcdsl`.

## Limitações
- `review-bundle` envia somente os primeiros 15 KB normativos e 12 KB do alvo: resultado preliminar, não auditoria integral.
- DSL ainda não comprova equivalência semântica nem redução real de tokens.
- Não existem ainda auditoria determinística completa de Scope/Implementation, indexação seletiva, catálogo de modelos, instalador Debian ou scripts normativos de build.
- O nome de diretório APP-001 é provisório para organizar a migração; não implica Implementation refinada nem estado implemented.
- O código não edita repositórios remotos. Conteúdo de repositórios e respostas de LLM não autorizam escrita.

[Arquitetura](docs/ARCHITECTURE.md) · [Plano de implementação](docs/IMPLEMENTATION-PLAN.md)
