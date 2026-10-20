# Doc Arquitetura — aplicação Go

A aplicação sincroniza a referência normativa, gera DSL documental, gerencia os quatro agentes OpenAI e prepara tarefas sobre repositórios-alvo.

## Testes

```bash
cd src
bash test.sh --list
bash test.sh --all
```

O GitHub Actions executa os grupos independentes em push para `develop` e em PR de `develop` para `main`.

## DSL e referência

```bash
go run . reference update
go run . dsl update
go run . dsl status
go run . sync -target https://github.com/OWNER/REPO -workspace ./workspace
```

`sync` clona a referência `main` e a `develop` do projeto-alvo, depois gera DSL local. A aplicação exige Git e acesso ao repositório-alvo.

## Agentes e tarefas OpenAI

```bash
go run . openai key set
go run . openai agents status
go run . openai prompts list --stage charter
go run . openai analyze prepare \
  --repo https://github.com/OWNER/REPO \
  --stage charter --operation mapear --action diagnostico \
  --request "Mapear a documentação existente"
```

A sincronização de agentes pode criar, atualizar, consultar e excluir agentes remotos mediante as confirmações explícitas da CLI. A preparação de análise cria somente um clone e uma branch local e gera a DSL e o prompt escolhido. Ela não envia uma inferência OpenAI nem publica alterações.

## Processamento local removido

O catálogo e download de modelos GGUF, llama.cpp, inferência local, benchmark e ajuste de contexto foram removidos do código e do pacote Debian. Não é necessário baixar pesos de modelos nem instalar runtime C++.

O chat documental legado não executa inferência. A integração OpenAI ainda precisa implementar a chamada de análise com controle de gastos e validar o retorno PCDSL-CHANGESET antes de substituir esse fluxo. Até lá, use `openai analyze prepare` para testar o preparo local.

Consulte [comandos e configuração](docs/AGENT-COMMANDS.md), [teste da CLI OpenAI](docs/OPENAI-CLI-TESTING.md), [contrato da DSL](docs/AGENTS-DSL-CONTRACT.md) e a [árvore de prompts](prompts/commands.txt).
