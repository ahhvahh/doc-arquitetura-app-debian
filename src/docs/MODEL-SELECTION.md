# Seleção do modelo por sessão e atividade — Terminal

No projeto-alvo, o terminal segue este fluxo:

1. Informar URL do repositório GitHub e escolher branch de origem;
2. Selecionar a etapa documental (Charter, Scope ou Implementation);
3. Clonar o repositório e atualizar a referência normativa **somente** de `ahhvahh/doc-arquitetura/main/ProjectCharter/1.0`;
4. **Escolher o modelo GGUF para esta sessão**, entre os instalados, baixar um modelo do catálogo após confirmação ou continuar sem modelo;
5. Escolher o prompt normativo especializado e acessar o chat.

A escolha fica registrada no JSON da sessão (`model`), independente das associações globais em `agents.json`. Quando o usuário retoma a sessão, o modelo escolhido é revalidado e restaurado. Para outras tarefas, pode alterá-lo pelo chat.

```sh
doc-arquitetura
# ou informar o modelo já instalado na inicialização:
doc-arquitetura terminal --repo https://github.com/OWNER/REPO --branch develop --area scope --model qwen3-1.7b
```

Dentro do chat:

```text
/model          Mostra o modelo da sessão e os comandos para atendê-lo
/model select   Permite baixar ou trocar o modelo antes da próxima ação
/model clear    Remove a seleção; operação documental sem LLM continua disponível
/status        Inclui o modelo selecionado
```

A lista de instalados é baseada no catálogo atual e nos arquivos locais GGUF/metadados; a **ativação verifica SHA-256 e assinatura GGUF**, recusando arquivos ausentes ou adulterados. A opção de download exige confirmação. Se nenhum modelo estiver instalado, o usuário pode baixá-lo ou continuar sem inferência; o chat exige seleção de modelo antes de processar pedidos em linguagem natural.

**Comportamento atual:** a escolha define o modelo da sessão e o Go inicia e interrompe automaticamente seu processo `llama-server` privado. Não é necessário instalá-lo ou iniciá-lo em outro terminal.

```sh
# O servidor será iniciado quando necessário pelo próprio aplicativo.
```

Ao trocar de modelo no chat, o processo anterior é encerrado e o próximo modelo será iniciado na solicitação seguinte. O catálogo embutido ainda oferece dois modelos (`qwen3.5-0.8b` e `qwen3-1.7b`); catálogo dinâmico e importação de GGUF externos são evoluções futuras. A seleção não autoriza alterações em `main`; edição documental e push continuam restritos a `develop` com aprovações independentes.


## Diagnóstico e benchmark de modelos (atualizado)

O aplicativo informa explicitamente quando a API local não está ativa, quando o modelo em `/v1/models` não é o escolhido ou quando a memória ainda está carregando. O pacote inclui `llama-server` em diretório privado, e o Go o inicia via socket Unix e verifica `/health` por até 90 segundos. Não é necessário instalar o runtime separadamente. O processo lançado pela CLI é encerrado ao sair da sessão; outros servidores existentes não são encerrados.

Comandos interativos:

```text
/model start   Inicia o runtime embutido por socket Unix privado (contexto padrão 4096)
/model check   Confere GET /health e o modelo anunciado em /v1/models
/model stop    Encerra apenas o processo iniciado por esta sessão
/benchmark     Mede a geração e o processamento do prompt no próprio servidor (tokens/s)
```

O comando `models serve` foi descontinuado; o runtime privado define `--alias <id>`, com contexto padrão 4096 (ou `DOC_ARQUITETURA_LLAMA_CONTEXT`). Benchmarks usam os campos `timings.prompt_per_second` e `timings.predicted_per_second` do llama.cpp; a medida não é uma estimativa genérica de CPU. Velocidade de prefill é diferente de velocidade de geração. Os resultados mudam conforme a CPU, quantidade de threads, memória, contextos, quantização e concorrência.

Além dos dois Qwen anteriores, o catálogo possui `gemma3-1b` (revisão textual leve) e `qwen3.5-4b` (tarefas mais exigentes). O Qwen3.5 4B exige mais RAM, pode ser lento em CPU e requer versão atual do llama.cpp que suporte sua arquitetura GGUF.

**Atenção às branches e ao legado:** uma sessão aberta em outra branch, como `ajusteDoc`, é somente leitura. Para modificar documentos e publicar, selecione `develop`. Charter de projetos legados também é lido como fonte histórica de diagnóstico quando a etapa escolhida é Charter, independentemente do prompt; os originais não são sobrescritos sem autorização. O contexto enviado ao modelo é limitado por tamanho, logo uma auditoria integral de vários documentos deve ser segmentada por assunto.


**Integração privada atual:** consulte [Inferência local privada](PRIVATE-INFERENCE.md) para build/revisão/lifecycle.