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

`routes`, `check`, `analyze`, `clone`, `sync`, `convert`, `model` e `review-bundle` são comandos do MVP. `check` aplica as regras Charter locais; `review-bundle` ainda é preliminar. Modelos GGUF não integram o pacote; `llama-server` é compilado e instalado privadamente.

Consulte [arquitetura](docs/ARCHITECTURE.md) e [plano de implementação](docs/IMPLEMENTATION-PLAN.md).

## Novos comandos locais

Consulte [Modelos e DSL](docs/AGENT-COMMANDS.md) para instalar modelos GGUF, associar modelos a agentes, consultar a referência `main` e compilar DSL versionada.

```sh
go run . models catalog
go run . reference update
go run . dsl update
go run . dsl status
```

## Sessões documentais pelo terminal

Execute `doc-arquitetura` para iniciar um menu local ou `doc-arquitetura terminal --repo https://github.com/OWNER/REPO --branch develop --area charter` para pular a seleção. O aplicativo consulta branches do GitHub, escolhe Charter, Scope ou Implementation, atualiza a referência normativa em `main` e abre um chat com o `llama-server` local. O servidor de interface Web anterior foi removido.

Comandos internos do chat: `/status`, `/files`, `/diff`, `/apply`, `/validate`, `/commit`, `/push`, `/help` e `/exit`.

**Publicação:** alterações do modelo são propostas; `/apply` exige autorização e edita somente o clone local. `/commit` exige validação estrutural inicial e confirmação; `/push` exige outra confirmação, uma única revisão local e que `develop` remoto não tenha sido modificado desde a clonagem. Outras branches são consultáveis, mas não permitem escrita. Não execute comandos Git diretamente a partir de respostas do modelo.

As sessões ficam em `DOC_ARQUITETURA_STATE_DIR` (ou `~/.local/state/doc-arquitetura`), com clones em `workspaces/`; os históricos são preservados para retomada. O modelo é executado automaticamente por socket Unix privado; não existe porta TCP nem parâmetro externo `--api`. A validação implementada é apenas estrutural; ainda é necessário completar a auditoria semântica e o gate normativo antes de declarar a documentação refinada ou implementada.


## Estrutura documental do projeto-alvo

A documentação do projeto-alvo deve estar diretamente em `/docs/charter/`, `/docs/scope/`, `/docs/implementation/` e `/docs/assets/`, **sem** `docs/ProjectCharter/1.0/`. A referência `ProjectCharter/1.0` existe somente na `main` do repositório normativo `ahhvahh/doc-arquitetura`.

O editor de terminal exige esses diretórios e bloqueia o commit se encontrar a etapa selecionada em uma das antigas localizações conhecidas. A migração requer aprovação, atualização de links e atualmente deve ser realizada fora do chat (não há migrador automático). Consulte [regras e comandos](docs/AGENT-COMMANDS.md).


## Seleção de prompts e migração com backup

O aplicativo oferece um seletor de prompts após o clone do projeto e os comandos `/prompts`, `/prompt select`, `/backup plan`, `/backup add <caminho>`, `/backup create` e `/migration plan`. O histórico de migração fica **exclusivamente em `/docs/backup/`**, acompanhado de `manifest.json`, SHA-256 e cópias fiéis dos documentos. Backups são criados com aprovação, incluídos no commit documental autorizado e excluídos da DSL ativa. A alteração dos documentos ativos depende de `/apply`, seguida de validação e aprovações independentes de commit/push. O agente não apaga automaticamente originais legados.

## Modelo por sessão GitHub

Ao iniciar um projeto GitHub, a CLI solicita o modelo GGUF a utilizar **antes de selecionar o prompt e abrir o chat**. O modelo fica registrado na sessão e pode ser trocado com `/model select`; veja [Seleção de modelo](docs/MODEL-SELECTION.md). O `llama-server` é um subprocesso privado iniciado sob demanda pela aplicação.


## Diagnóstico de Charter e benchmark local

O chat inicia o subprocesso privado e verifica `/health` e `/v1/models` no socket Unix antes de solicitar inferência. Disponibiliza `/model start`, `/model check`, `/model stop` e `/benchmark` para inicialização/diagnóstico e velocidade real em tokens/s. O runtime privado usa contexto padrão 4096 e alias do modelo, acessível apenas via socket Unix. A seleção de branch privilegia `develop` para edição; outras branches continuam disponíveis somente para leitura. Quando o Charter canônico ainda não existe em `/docs/charter/`, o agente recebe evidência parcial dos documentos legados `/docs/01-Project_Charter/charter/` para revisão, sem movimentar originais automaticamente.

**Integração privada atual:** consulte [Inferência local privada](docs/PRIVATE-INFERENCE.md) para build/revisão/lifecycle.

## Próxima arquitetura — serviço isolado (refinement)

Está registrada a proposta de execução exclusiva pelo daemon systemd, com dois grupos de acesso distintos, runtime `llama.cpp` privado, socket Unix de controle, auditoria e branches de trabalho individuais derivadas de `develop`. A documentação está em [Plano de serviço isolado](docs/SERVICE-ISOLATION-PLAN.md) e [Fluxo de tarefa e branch](docs/TASK-BRANCH-LIFECYCLE.md). **Ainda não está em execução neste binário**: os comandos legados de edição/commit para `develop` deverão ser desativados quando o daemon for implementado; não presumir segurança multiusuário no MVP atual.
