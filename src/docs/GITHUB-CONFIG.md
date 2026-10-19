# GitHub — conexão por arquivo do aplicativo (CLI legado)

> Disponível no código em `develop`; instalar nova build `.deb` para usar. A versão Debian **0.2.0 instalada anteriormente não contém este recurso**. Até o serviço systemd root-owned estar implementado, as credenciais são individuais por usuário Linux.

## Configurar no servidor Debian como usuário comum

A aplicação lê o arquivo `~/.config/doc-arquitetura/github.json` (ou o diretório definido por `DOC_ARQUITETURA_CONFIG_DIR`). Não precisa de `ssh-keygen`, alterações em `~/.ssh/config` ou em `git config --global`.

Crie o diretório:

```bash
mkdir -p ~/.config/doc-arquitetura
chmod 700 ~/.config/doc-arquitetura
```

### Opção A — configuração em apenas um arquivo (mais fácil)

Crie/edite `~/.config/doc-arquitetura/github.json` com editor de texto interativo (`nano`, por exemplo):

```json
{
  "github": {
    "authentication": "token",
    "username": "SEU_USUARIO_GITHUB",
    "token": "COLE_AQUI_TOKEN_FINE_GRAINED"
  }
}
```

```bash
chmod 600 ~/.config/doc-arquitetura/github.json
```

**Este modo guarda o token em texto puro no arquivo local.** Nunca incluir `github.json` em um repositório Git, upload, anexo ou backup compartilhado; não conceder acesso a outros usuários. O próprio aplicativo exige proprietário correspondente ao usuário em execução e modo restrito.

### Opção B — token em arquivo separado (preferencial)

`github.json`:

```json
{
  "github": {
    "authentication": "token",
    "username": "SEU_USUARIO_GITHUB",
    "token_file": "/home/usr01/.config/doc-arquitetura/github.token"
  }
}
```

Crie `github.token` com um editor no terminal; coloque **somente** o token, sem espaços. Proteja os dois arquivos:

```bash
chmod 600 ~/.config/doc-arquitetura/github.json ~/.config/doc-arquitetura/github.token
```

O `token_file` deve ser absoluto. A aplicação recusa links simbólicos, arquivos acessíveis por grupo/outros e proprietários incompatíveis.

## Gerar token GitHub

Use GitHub **Settings → Developer settings → Personal access tokens → Fine-grained tokens**, selecione somente o(s) repositório(s) necessários. `Contents: Read-only` é suficiente para `github check`, clonagem e referência; `Contents: Read and write` é necessário para criar commits remotos na branch permitida, quando essa etapa estiver disponível. Não conceda mais permissões que o necessário. Defina validade e política de rotação adequadas. Para repositórios de organizações, pode ser necessária aprovação do responsável.

## Testar sem publicar

```bash
doc-arquitetura github config
doc-arquitetura github check https://github.com/ahhvahh/a-bioma-sauva
```

O primeiro mostra apenas a conta configurada e o caminho, sem exibir o segredo. O segundo faz `git ls-remote` e confirma a branch `develop`, sem clone, commit, push ou alterações remotas. Ele confirma acesso de leitura, **não** prova permissão de escrita.

O uso normal permanece:

```bash
doc-arquitetura
```

O token é entregue ao Git por um processo auxiliar do mesmo aplicativo (`GIT_ASKPASS`) quando necessário, **não** por URL, argumento de linha de comando ou variável de ambiente contendo o segredo. O auxiliar só responde a prompts HTTPS destinados a `github.com`. As alterações usam `git` subprocessos controlados e não modificam as configurações globais do usuário. Caminhos de Git utilizados pelo terminal, referência normativa, sync e clone compartilham a configuração.

## Escopo de segurança

A autenticação é por usuário Linux, **não é isolamento forte entre programas do mesmo UID**. Qualquer programa executado como o mesmo usuário que controla `github.json` pode ler o arquivo e obter o token; permissões 0600 protegem contra **outros usuários comuns**, não contra o titular ou root. A futura implementação de serviço systemd e `sudo doc-arquitetura admin` substituirá esta credencial individual por segredo root-only injetado ao serviço. Não habilitar daemons com credenciais do usuário como atalho.

O recurso é retrocompatível: sem `github.json`, acesso anônimo a repositórios públicos continua funcionando.
