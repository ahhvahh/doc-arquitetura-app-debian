# Comandos e configuração

A aplicação mantém comandos locais para DSL e sincronização, e comandos OpenAI para gerenciar agentes e preparar tarefas.

## Referência e DSL

```bash
doc-arquitetura reference update
doc-arquitetura reference status
doc-arquitetura dsl update
doc-arquitetura dsl status
```

`sync` clona a referência normativa `main` e a branch `develop` do projeto informado e gera os arquivos DSL para revisão local. `convert` converte Markdown local em DSL.

## Agentes OpenAI e tarefas

```bash
doc-arquitetura openai key set
doc-arquitetura openai agents status
doc-arquitetura openai prompts list --stage charter
doc-arquitetura openai analyze prepare --repo https://github.com/OWNER/REPO --stage charter --operation mapear --action diagnostico --request "Mapear a documentação existente"
```

A gestão remota de agentes requer uma chave OpenAI e só faz chamadas à API com `--apply`. A preparação de análise clona `develop`, cria uma branch local e gera a DSL e a tarefa. Ainda não faz inferência OpenAI nem publica alterações.

## Processamento de modelos

O catálogo GGUF, downloads Hugging Face, execução local de llama.cpp, benchmark e configurações locais de contexto foram removidos. Os comandos antigos `models`, `model`, `config set-context` e `review-bundle` não fazem mais parte da CLI. O pacote Debian não compila nem instala llama.cpp e não depende de bibliotecas C++ para esse runtime.

O antigo chat terminal não envia pedidos a um modelo. Para análise remota, ainda é necessário implementar a inferência OpenAI com seus limites de gasto e saída PCDSL-CHANGESET; até lá, use o comando de preparação para validar clones, prompts e DSL sem inferência.

O programa continua exigindo Git para clonar os repositórios. A CLI não armazena modelos GGUF.
