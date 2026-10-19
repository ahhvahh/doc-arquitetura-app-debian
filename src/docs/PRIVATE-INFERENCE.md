# Inferência local privada — llama.cpp

## Decisão

O `llama.cpp` é compilado **durante o build do pacote Debian** a partir da tag upstream `v0.7.0`, cujo commit fixado é `620d2225cc436ab0b001e3200e590e3b286ef9ab`. A compilação ocorre em um diretório temporário do empacotador, sem instalar executáveis globalmente. A revisão de origem é verificada antes do CMake. <https://github.com/ggml-org/llama.cpp/releases/tag/v0.7.0>

Artefatos do pacote:

- `/usr/bin/doc-arquitetura`: aplicação Go;
- `/usr/lib/doc-arquitetura-app/llama-server`: motor privado;
- `/usr/share/doc/doc-arquitetura-app/llama.cpp-LICENSE`: licença upstream;
- `/usr/share/doc/doc-arquitetura-app/llama.cpp-version.txt`: revisão da origem;
- modelos GGUF permanecem em `~/.local/share/doc-arquitetura/models`, **fora do .deb**.

### Empacotamento

O `src/package.yaml` chama `src/scripts/build-private-llama.sh`. O script baixa **o código-fonte fixado** no GitHub durante o build, verifica o SHA, compila com CMake em modo CPU (sem GPU) e inclui somente o binário resultante, licença e marcador de versão. Não depende de `llama-server` previamente instalado no host ou pelo usuário. Requer conectividade **durante o empacotamento** e ferramentas de compilação `git`, `cmake`, `c++`, `Go`. O download do código-fonte não ocorre na execução da aplicação.

O build de C/C++ é portátil (`GGML_NATIVE=OFF`), não ajustado a uma CPU específica. Otimização por arquitetura e build totalmente offline devem ser tratados separadamente. Em produção, testar os binários gerados com a distribuição-alvo.

### Isolamento

O Go inicia o executável do pacote como processo-filho com `--host <socket>.sock`, `--no-webui`, `--alias <modelo>`, `-ngl 0` e contexto padrão de 4096 tokens. A comunicação usa **HTTP por Unix Domain Socket**, com transporte `net.Dialer("unix",...)`: nenhuma porta TCP é aberta e o código não aceita `--api` TCP na sessão. O diretório efêmero do socket é criado com modo `0700`, protegido contra outros usuários comuns; processos do **mesmo usuário** ainda podem acessá-lo se obtiverem o caminho. O projeto não promete isolamento contra comprometimento do próprio usuário ou de root.

### Ciclo de vida

1. Usuário abre sessão GitHub, escolhe modelo GGUF e prompt.
2. Na primeira solicitação, `ensureRuntime` valida o GGUF, cria socket privado e inicia o subprocesso.
3. O Go aguarda `/health` e confere `/v1/models` por até 90 s.
4. Pedidos ao agente e benchmarks são enviados pelo transporte do socket, nunca pela rede.
5. Ao trocar de modelo, executar `/model stop`, sair da sessão ou ocorrer timeout de startup, o Go encerra o processo e remove o socket. O processo não persiste como daemon.

Comandos: `/model start`, `/model check`, `/model stop`, `/benchmark`, `/model select`. `/benchmark` mede `timings.predicted_per_second` e `timings.prompt_per_second` reais do motor. Sem execução local não é possível atribuir velocidade exata ao i5-7400.

Configurações opcionais:

- `DOC_ARQUITETURA_LLAMA_CONTEXT`: contexto 512..16384 tokens (padrão 4096). Um valor fora da faixa usa o padrão. Para mudar o contexto, encerre a sessão e reinicie o CLI com essa variável no ambiente antes de carregar o modelo.
- `DOC_ARQUITETURA_LLAMA_BIN`: caminho absoluto de binário alternativo, apenas para desenvolvimento/testes.
- `DOC_ARQUITETURA_LLAMA_BUILD_JOBS`: paralelismo na compilação C++ (padrão 2).

`doc-arquitetura models serve` e `doc-arquitetura terminal --api http://...` não são mais suportados no fluxo documental. Não é necessário instalar llama.cpp globalmente, ativar systemd, manter porta 8080 ou abrir navegador.

### Observações

A ferramenta ainda não implementa auditoria semântica integral, gates normativos automáticos ou migração destrutiva de arquivos. O agente documental edita apenas a área permitida em `/docs`, requer aprovação e publica somente em `develop`, com referência normativa `doc-arquitetura/main`.


## Contexto ampliado para auditorias (v0.2.2)

Para análise documental maior no servidor Debian, inicie uma nova instância do CLI com:

```bash
DOC_ARQUITETURA_LLAMA_CONTEXT=16384 doc-arquitetura
```

O runtime continua privado no socket Unix, mas passa a criar **apenas um slot de inferência** (`-np 1`) em vez de quatro slots automáticos. O contexto configurado é por slot; esta escolha evita quatro caches de contexto para um único chat e é adequada ao host CPU de quatro núcleos e memória limitada. A chamada HTTP de inferência dispõe agora de até 15 minutos, já que uma solicitação com ~9,4 mil tokens pode levar vários minutos para ler e produzir a resposta. Medir memória real e desempenho com um prompt representativo, pois benchmark curto não garante a capacidade de processar 16 mil tokens nem a mesma taxa de tokens/s.

**Limitação ainda aberta:** o agente envia conjuntamente regras, prompts, documentação e pedido. O aplicativo ainda **não** mede os tokens da solicitação antes de transmiti-la e **não** divide automaticamente documentos em lotes. Mesmo com 16.384 tokens, solicitações superiores ao limite ou saídas longas podem falhar. Aumentar contexto não substitui orçamento de tokens e execução por lotes.


## Configuração persistente do contexto (v0.2.3)

Além de `DOC_ARQUITETURA_LLAMA_CONTEXT`, o runtime aceita `~/.config/doc-arquitetura/config.json` com `{"llama":{"context_tokens":16384}}`. Limites: 512..32768 tokens. Use `doc-arquitetura config set-context 16384` para gravar o arquivo; use `/context set 16384` dentro do chat para atualizar e reiniciar o subprocesso. Erros de configuração não são convertidos silenciosamente ao valor padrão. Consulte [Configuração e sessões](LOCAL-CONFIG.md).
