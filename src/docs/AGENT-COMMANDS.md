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
doc-arquitetura models serve qwen3-1.7b
# Porta/contexto opcionais:
doc-arquitetura models serve --port 8080 --context 2048 qwen3-1.7b
```

`models download` consulta a API do Hugging Face para obter o SHA da revisão da fonte, fixa essa revisão na URL de download, verifica a assinatura GGUF e grava SHA-256, origem, tamanho e revisão. Um checksum **independente** pode ser exigido com `--sha256 HASH`; sem essa opção, o checksum registrado serve para detectar alterações locais, mas não prova autenticidade da fonte. `models list` verifica os arquivos instalados. Download interrompido deve ser repetido; retomada por Range ainda não é suportada. `models serve` exige `llama-server` instalado separadamente, executa em primeiro plano e escuta somente `127.0.0.1`. `models assign` registra a associação do agente com um modelo; o pipeline do MVP ainda não inicia agentes automaticamente nem aplica essa associação a `review-bundle`.

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

- Runtime `llama.cpp` e modelos GGUF são externos ao `.deb`.
- Cada modelo é executado isoladamente; não há servidor de múltiplos modelos nem daemon systemd no MVP.
- O comando `review-bundle` ainda usa a limitação de contexto do MVP.
- Testes locais não dependem de rede; downloads reais e inferência devem ser validados no host de instalação.
- Não editar a branch `main` do repositório normativo.