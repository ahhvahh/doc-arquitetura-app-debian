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

**Limitação atual:** a escolha define o modelo da sessão e o ID enviado à API, mas não inicia nem troca automaticamente o processo `llama-server`. Ele deve estar instalado e ser iniciado em **outro terminal**, com o mesmo modelo escolhido:

```sh
doc-arquitetura models serve qwen3-1.7b
```

Se trocar de modelo no chat, encerre o servidor anterior e inicie outro com o novo GGUF. O catálogo embutido ainda oferece dois modelos (`qwen3.5-0.8b` e `qwen3-1.7b`); catálogo dinâmico e importação de GGUF externos são evoluções futuras. A seleção não autoriza alterações em `main`; edição documental e push continuam restritos a `develop` com aprovações independentes.
