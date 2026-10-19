# Configuração do runtime e encerramento de sessões — v0.2.3

## Configuração persistente

A CLI lê `~/.config/doc-arquitetura/config.json` (ou `$DOC_ARQUITETURA_CONFIG_DIR/config.json` caso o diretório tenha sido personalizado). Este **não** é o `github.json`: credenciais de GitHub continuam separadas.

```json
{
  "llama": {
    "context_tokens": 16384
  }
}
```

A própria CLI pode criar ou alterar o arquivo mantendo permissões 0600:

```bash
doc-arquitetura config set-context 16384
doc-arquitetura config show
```

O contexto admite **512 a 32768 tokens**, inclusive. Escolha um valor de acordo com o modelo GGUF e a RAM disponível. Um valor inválido, arquivo malformado, link simbólico ou arquivo legível por outros usuários resulta em erro explícito — a aplicação não volta silenciosamente a 4096. A ordem é:

1. `DOC_ARQUITETURA_LLAMA_CONTEXT`, quando definida, prevalece;
2. `config.json`, quando apresenta `context_tokens`;
3. 4096 tokens como valor padrão.

A alteração é aplicada **quando o runtime privado inicia**, não enquanto o Llama está carregado. No chat, use `/context` para exibir a configuração. O comando `/context set 16384` grava a preferência, encerra o subprocesso atual e prepara sua reinicialização. Depois utilize `/model start`; a próxima mensagem de análise também pode iniciar o runtime sob demanda. Não é necessário editar `github.json` nem reiniciar o Debian.

**Atenção ao consumo de memória:** 16384 ou 32768 tokens não garantem uma inferência rápida ou possível no hardware disponível. No servidor CPU i5 com memória limitada, prefira começar com 16384 e observar `free -h`, swap e latência de uma auditoria real. O processo ainda usa **um único slot** e socket Unix privado, sem TCP.

**Limitação ainda aberta:** a aplicação continua montando parte da documentação em um prompt único e **não** divide automaticamente contextos maiores que a janela. Erros HTTP 400 ainda são possíveis; recomenda-se enviar documentos em lotes e reservar tokens para a resposta em próxima implementação.

## Sessões: sair versus encerrar

- `/exit`: sai do chat, libera o processo privado Llama, **preserva a sessão aberta** para continuar no futuro;
- `/session close`: pede confirmação `SIM`, marca a sessão atual como encerrada, libera o Llama e retorna ao menu principal;
- Menu **Encerrar sessão salva**: escolhe uma sessão aberta anterior e pede `SIM` para encerrá-la, sem precisar entrar no chat.

O arquivo JSON de sessão recebe `closed_at` (UTC RFC3339). Sessões encerradas são filtradas do menu **Continuar sessão**. O workspace/clones e registros em disco **não são apagados**; não se tenta criar commit, fazer push ou apagar arquivos. Propostas não aplicadas em memória são descartadas ao encerrar; as alterações já feitas no workspace permanecem. O encerramento não pode ser revertido pelo menu nesta versão.

O fluxo de branches exclusivas por operação e o serviço systemd de acesso restrito continuam **planejados**, não foram ativados nesta atualização. O projeto-alvo deve seguir as regras de permissão atuais da sessão.

## Teste no servidor

Após empacotar e instalar o `.deb` versão 0.2.3:

```bash
dpkg-query -W doc-arquitetura-app
doc-arquitetura config set-context 16384
doc-arquitetura config show
doc-arquitetura
```

No chat: `/context`, `/model start` e `/benchmark`. Procure no log `contexto 16384 (config.json)`, `n_slots = 1` e `n_ctx_slot = 16384`. Para encerrar definitivamente, execute `/session close` e confirme. O menu de retomada não deve mais listar essa sessão.

**Conflito com variável anterior:** ao usar `/context set 16384` no chat, a aplicação remove `DOC_ARQUITETURA_LLAMA_CONTEXT` somente do processo atual após salvar o JSON; assim o próximo subprocesso Llama lê 16384 do arquivo. No shell externo a variável continua como estava. O comando externo `doc-arquitetura config set-context 16384` avisa quando a variável estiver definida; abra o aplicativo sem ela para utilizar o arquivo.
