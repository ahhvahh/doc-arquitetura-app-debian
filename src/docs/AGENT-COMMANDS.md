# Modelos locais, referência normativa e DSL — CLI

Os comandos `models`, `reference` e `dsl` rodam sob o usuário que invocou o executável. Não exigem root e não instalam modelos no pacote Debian.

## Modelos

```sh
doc-arquitetura models catalog
doc-arquitetura models download qwen3.5-0.8b
doc-arquitetura models download qwen3-1.7b
doc-arquitetura models list
doc-arquitetura models assign triagem qwen3.5-0.8b
doc-arquitetura models assign analise qwen3-1.7b
# O motor será iniciado sob demanda no terminal (não use models serve)
# Porta/contexto opcionais:
# Configure DOC_ARQUITETURA_LLAMA_CONTEXT=4096, se necessário
```

`models download` consulta a API do Hugging Face para obter o SHA da revisão da fonte, fixa essa revisão na URL de download, verifica a assinatura GGUF e grava SHA-256, origem, tamanho e revisão. Um checksum **independente** pode ser exigido com `--sha256 HASH`; sem essa opção, o checksum registrado serve para detectar alterações locais, mas não prova autenticidade da fonte. `models list` verifica os arquivos instalados. Download interrompido deve ser repetido; retomada por Range ainda não é suportada. `llama-server` é fornecido pelo pacote e iniciado como subprocesso privado via socket Unix. `models assign` registra a associação do agente com um modelo; o pipeline do MVP ainda não inicia agentes automaticamente nem aplica essa associação a `review-bundle`.

## Referência e compilação DSL

```sh
doc-arquitetura reference update
doc-arquitetura reference status
doc-arquitetura dsl compile
doc-arquitetura dsl update
doc-arquitetura dsl status
```

Fonte imutável: `https://github.com/ahhvahh/doc-arquitetura.git`, branch `main`, pasta `ProjectCharter/1.0`. `reference update` consulta HEAD remoto, reutiliza cache íntegro ou clona uma cópia versionada. O SHA real obtido da clonagem é registrado, inclusive se a branch mudar enquanto é baixada. Necessita `git` e autenticação Git pré-configurada quando o repositório é privado. O comando `dsl compile` usa a função de conversão documental já existente no MVP e compila **todos os arquivos Markdown** sob `ProjectCharter/1.0`, inclusive Charter, Scope, Implementation e prompts. Grava `reference.pcdsl` e `manifest.json` contendo SHA de referência, checksum do bundle, contagem e horário. Só ativa a versão compilada após sucesso; caso falhe, a anterior permanece ativa. A compilação **não comprova equivalência semântica nem completude da auditoria**.

`dsl update` é a sequência `reference update` + `dsl compile`. `dsl status` verifica o checksum e compara com a referência **local**. `reference status` também consulta o HEAD remoto e informa quando a referência local está desatualizada.

## Diretórios

Padrões por usuário (XDG):

| Dados | Diretório |
|---|---|
| Modelos | `~/.local/share/doc-arquitetura/models/` |
| Configuração de agentes | `~/.config/doc-arquitetura/agents.json` |
| Git clone normativo | `~/.cache/doc-arquitetura/reference/<sha>/` |
| Bundle compilado | `~/.cache/doc-arquitetura/dsl/<sha>/` |
| Estado e relatórios futuros | `~/.local/state/doc-arquitetura/` |

Aceita `DOC_ARQUITETURA_DATA_DIR`, `DOC_ARQUITETURA_CACHE_DIR`, `DOC_ARQUITETURA_CONFIG_DIR`, `DOC_ARQUITETURA_STATE_DIR`, `DOC_ARQUITETURA_TEMP_DIR`, além das variáveis XDG convencionais. Diretórios passados por variáveis precisam ser absolutos. O diretório temporário configurado está reservado para a evolução do workspace de análises; o comando legado `sync` ainda usa `-workspace` (padrão `./workspace`).

## Limitações

- Runtime `llama.cpp` é compilado e incluído no `.deb`; os modelos GGUF são externos.
- A aplicação gerencia um subprocesso de inferência por sessão, sem daemon systemd ou porta TCP.
- O comando `review-bundle` ainda usa a limitação de contexto do MVP.
- Testes locais não dependem de rede; downloads reais e inferência devem ser validados no host de instalação.
- Não editar a branch `main` do repositório normativo.
## Chat documental no terminal (MVP)

```sh
doc-arquitetura
doc-arquitetura terminal --repo https://github.com/OWNER/REPO --branch develop --area scope 
```

O menu descobre branches remotas e permite escolher Charter, Scope ou Implementation. Os prompts são carregados da referência normativa `main` versionada pelo SHA da sessão. O agente envia propostas JSON para documentos Markdown na etapa escolhida; respostas da LLM jamais são comandos shell/Git executáveis.

Comandos no chat: `/help`, `/files`, `/status`, `/diff`, `/apply`, `/validate`, `/commit`, `/push`, `/exit`. Edições locais exigem confirmação via `/apply`. Commit e push possuem confirmações independentes. O push depende de confirmação de que `origin/develop` continua no SHA original e não usa force. Se a branch selecionada não for `develop`, a sessão é somente leitura e os comandos de escrita retornam erro.

**Limitações:** a validação antes do commit cobre alterações restritas a Markdown na etapa escolhida, H1, badges e checagem `git diff --check`; não substitui revisão normativa completa de Charter, Scope e Implementation. Arquivos maiores que o contexto admitido não são enviados integralmente e não podem ser substituídos pelo agente. Testes de integração com GitHub remoto e modelos reais não realizados nesta entrega.

## Destino obrigatório da documentação do projeto-alvo

O repositório normativo é `ahhvahh/doc-arquitetura`, branch **main**, diretório **`ProjectCharter/1.0/`**. Esse caminho é a **origem das regras**, e **não deve ser reproduzido no projeto-alvo**.

No projeto-alvo, os documentos ficam diretamente em `/docs/charter/`, `/docs/scope/`, `/docs/implementation/` e `/docs/assets/`, com índice em `/docs/README.md`. Nunca gere o nível intermediário `docs/ProjectCharter/1.0/`. O aplicativo aplica essa regra no prompt de sistema e na validação do caminho proposto pelo modelo; `main` do projeto-alvo nunca é editada.

Se uma etapa possuir documentos nas estruturas legadas `docs/ProjectCharter/1.0/`, `docs/01-Project_Charter/` ou `ProjectCharter/1.0/`, o terminal avisa e bloqueia o commit até que os documentos sejam migrados com matriz origem → destino, links corrigidos e autorização explícita. **A migração automática ainda não está implementada**. A análise e a apresentação de propostas não removem arquivos legados.

O comando `sync` prioriza `docs/` quando encontra as etapas nesse diretório; estruturas legadas são lidas como fallback para diagnóstico. A referência normativa do comando `sync` aceita somente `--norm-ref main`.

## Escolher prompts e reutilizar documentação no terminal

Após selecionar um repositório e uma das etapas `charter`, `scope` ou `implementation`, o programa apresenta os prompts `.md` especializados da referência versionada **`ahhvahh/doc-arquitetura/main/ProjectCharter/1.0/<etapa>/prompts/`**. A seleção fica associada à sessão e é incluída na mensagem do modelo local com o AGENT-PROMPT da etapa.

```text
/prompts                    # catálogo da etapa
/prompt select              # selecionar um prompt
/prompt show                # ver prompt selecionado
/prompt reset               # limpar seleção
/backup plan                # listar arquivos elegíveis e tamanhos
/backup add docs/legado     # incluir uma fonte documental adicional
/backup create              # gerar snapshot confirmado
/migration plan             # mostrar destinos sugeridos e conflitos
/diff                       # mostrar alterações propostas
/apply                      # aprovar edição dos documentos ativos
/validate
/commit
/push
```

**Backup obrigatório:** `/docs/backup/<UTC>-<sha-base-curto>/`, com `manifest.json` e `original/<caminho relativo de origem>`. A cópia preserva bytes, caminhos, hashes SHA-256 e referência do commit-base. A ferramenta só cria backup após confirmação expressa e quando a branch é `develop` e o workspace está limpo. A verificação do manifesto/hash impede que o agente edite o histórico. O diretório é incluído no commit autorizado, mas excluído da compilação DSL do projeto e dos gates de documentação ativa. Limites iniciais: 300 arquivos, 4 MiB por arquivo e 32 MiB por snapshot.

A função de inventário inclui as pastas canônicas e legadas conhecidas da etapa, além das fontes adicionadas pelo usuário. **O operador deve conferir se existem segredos/dados sensíveis antes de aprovar o backup ou push**. Migração e substituição de documentos são propostas pelo agente especializado; sem snapshot aprovado, `/apply` recusa propostas de prompt `legacy-*`. `/migration plan` não move nem apaga automaticamente os originais; a limpeza do legado exige operação própria futura e autorização independente.

A referência normativa é sempre `main` (somente leitura). Um prompt novo publicado apenas em `doc-arquitetura/develop` não aparece até ser promovido à `main`.

## Modelo por sessão

Após iniciar o projeto GitHub, escolha um modelo instalado, baixe um do catálogo ou continue sem inferência. A escolha persiste por sessão e pode ser alterada com `/model select`; `/model` mostra o estado e `/model clear` limpa a seleção. É preciso iniciar o `llama-server` separadamente com o modelo escolhido. Consulte [o fluxo completo](MODEL-SELECTION.md).


## Diagnóstico do chat e velocidade de inferência

Se o agente não responder, use `/model check`: o subprocesso `llama-server` privado será iniciado automaticamente na primeira solicitação, sem TCP. Use `/model check` para confirmar que a API responde e que o modelo realmente carregado corresponde à sessão. Para medir o desempenho CPU local, execute `/benchmark`; ele imprime tokens/s de geração, tokens/s do prompt e tempo total quando disponíveis.

A escolha de uma branch diferente de `develop` agora enfatiza que a sessão é somente leitura. Os documentos legados de Charter disponíveis em `docs/01-Project_Charter/charter/` podem ser consultados em diagnóstico, mesmo usando um prompt não migratório; a aplicação não confunde esse caminho com o destino canônico `docs/charter/`.

Catálogo complementar: `gemma3-1b` e `qwen3.5-4b` além de `qwen3.5-0.8b` e `qwen3-1.7b`. Consulte [Seleção de modelo](MODEL-SELECTION.md).


**Integração privada atual:** consulte [Inferência local privada](PRIVATE-INFERENCE.md) para build/revisão/lifecycle.