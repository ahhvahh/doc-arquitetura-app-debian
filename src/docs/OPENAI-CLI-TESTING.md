# Execução experimental — catálogo, agentes OpenAI e DSL

**Estado:** protótipo para validação por etapas na branch `develop`. Os contratos Charter, Scope e Implementation do aplicativo ainda não foram refinados. Nenhuma funcionalidade desta página promove documentação para `refined` ou `implemented`.

## 1. Compilar e executar testes sem custos

```bash
cd src
go test ./... -v
go run . openai prompts list
go run . openai agents status
go run . openai key status
```

O catálogo está versionado em `src/prompts/commands.json`, incorporado ao executável por `go:embed`. Tem 34 caminhos `<etapa>/<operação>/<ação>`, mapeados para os prompts `.txt` reais da `main` de `ahhvahh/doc-arquitetura`. Uma ação marcada `executable=false`, como o padrão de script de build ainda PENDING, não é executada.

## 2. Token OpenAI e gestão de agentes

Instale `libsecret-tools` e mantenha Secret Service/keyring disponível no Linux. A CLI não grava chaves em config.json nem no repositório:

```bash
go run . openai key set
go run . openai key status
# Para remover somente a chave guardada no cofre:
go run . openai key delete
```

Também funciona com `OPENAI_API_KEY` já presente no ambiente. O valor do token não deve ser incluído em argumentos, logs ou histórico de comandos.

Cada agente é independente. Primeiro verifique a prévia das instruções normativas; somente com `--apply` haverá chamada à API para criar/atualizar o agente remoto:

```bash
go run . openai agents sync --stage charter --model MODELO
go run . openai agents sync --stage charter --model MODELO --apply
go run . openai agents sync --stage scope --model MODELO --apply
go run . openai agents sync --stage implementation --model MODELO --apply
go run . openai agents sync --stage src --model MODELO --apply
go run . openai agents status
```

`MODELO` deve ser um identificador válido e disponível em seu projeto OpenAI. A aplicação fixa `main` normativa e registra `id`, `model`, `normative_sha` e `instructions_sha256` em `~/.config/doc-arquitetura/openai-agents.json` (arquivo privado 0600). Repetir um sync de instruções idênticas com ID cadastrado consulta o objeto remoto e evita atualização redundante. O agente não recebe automaticamente os 34 prompts na sua configuração persistente.

Para remover, há prévia sem alteração e exclusão definitiva somente com ID explícito:

```bash
go run . openai agents delete --stage charter
go run . openai agents delete --stage charter --apply --confirm ID_EXATO
```

**Atenção:** a criação/atualização/exclusão de agentes depende de permissão na conta OpenAI. Não há nenhum agent remoto criado no GitHub Actions. Testes usam `httptest`, com token fictício e sem contato com a API real.

## 3. Preparação de tarefa num projeto real

```bash
go run . openai prompts list --stage charter
go run . openai analyze prepare \
  --repo https://github.com/OWNER/REPO \
  --stage charter \
  --operation mapear \
  --action diagnostico \
  --request "Catalogar os objetivos existentes e indicar lacunas"
```

A aplicação:

1. atualiza a referência normativa `main` e compila sua DSL;
2. seleciona apenas `prompts/charter/project-intake-review.txt`, validando o diretório;
3. clona a `develop` do projeto-alvo numa pasta de estado privada;
4. cria uma **branch apenas local** `docs/charter/mapear-diagnostico-<id>` a partir do commit verificado;
5. compacta os Markdown ativos do projeto como `target.pcdsl` com caminhos, hashes e status;
6. escreve `task.json` contendo instruções e DSL para revisão, sem enviar à OpenAI nem publicar no GitHub.

Veja o caminho absoluto de `task.json` exibido pela CLI. A preparação requer acesso Git válido ao repositório, mas não requer chave OpenAI.

## 4. Recepção de changeset e preservação de Markdown

O contrato de retorno é `PCDSL-CHANGESET/1`. O modelo entrega objetos estruturados por arquivo, nunca comandos Git ou shell. Consulte `src/test/markdown/changeset-example.json` (placeholders de SHA precisam ser substituídos antes de aplicar).

```bash
# Verificação segura: não altera arquivo algum
go run . openai analyze apply \
  --task-dir /CAMINHO/DA/TAREFA \
  --changeset /CAMINHO/DA/RESPOSTA.json

# Apenas após comparar conteúdo, caminho e hash:
go run . openai analyze apply \
  --task-dir /CAMINHO/DA/TAREFA \
  --changeset /CAMINHO/DA/RESPOSTA.json \
  --apply --confirm TASK_ID
```

O validador exige o mesmo `task_id`, `stage`, `prompt`, `normative_sha` e `target_develop_sha` da tarefa original. `replace_section` só altera uma seção H2/H3 existente se o SHA-256 do Markdown completo for idêntico; mantém as demais seções. `create_file` verifica diretório, H1 e badges. Caminhos `../`, saída fora de `docs/<etapa>/`, symlinks e exclusões são bloqueados. Para a etapa `src`, alterações de código-fonte necessitam de um contrato separado e ainda não são aplicáveis por este decodificador Markdown.

Quando autorizado, o resultado fica **somente no clone local** e a aplicação gera `updated.pcdsl`. Não executa commit, push ou merge. Repositório normativo `main` e branch remota `develop` permanecem sem alterações.

## 5. Restrições deliberadas

- **Inferência OpenAI paga permanece bloqueada**: faltam o preflight de preço, câmbio, teto por operação/agente, reserva de orçamento e verificação de hard limit no projeto.
- **Publicação remota de branches permanece bloqueada** até resolução de `OPEN-003` e ajuste da norma vigente.
- A reconstrução é por **patch Markdown validado contra SHA**, não decodificação inversa do `PCDSL-DOCS 1`. A DSL retira elementos e não é reversível byte a byte.
- Não há garantia de rollback transacional multiarquivo em caso de falha de disco durante aplicação local: revisões com múltiplos arquivos requerem mecanismos adicionais antes de produção.
- A interface de administração multiusuário, credenciais de serviço, auditoria semântica completa e geração final de código `src` ainda são pendências.
- Antes de declarar produção, refinar o Charter, registrar contratos funcionais em Scope e contratos técnicos/testes em Implementation.

## 6. Referências

- [Guia original de integração](OPENAI-INTEGRATION-PLAN.md)
- [Árvore de prompts](../prompts/commands.txt)
- [Contrato e retorno DSL](AGENTS-DSL-CONTRACT.md)
- [Limites de gasto](OPENAI-SPEND-LIMITS.md)

## Inferência local removida

O executável e o pacote não incluem llama.cpp, catálogo/download GGUF, chat local, benchmarks ou configuração de contexto. `terminal` permanece para inspeção e ações documentais locais; mensagens livres não são enviadas a nenhum modelo. A inferência OpenAI ainda não está implementada e permanece bloqueada pelos controles de gasto pendentes. Use `openai analyze prepare` para validar a preparação sem custos.
