# Configuração administrativa root — CLI/TUI e credenciais

> **Atualização 2026-10-20:** este plano é histórico e ainda não implementado. O cliente atual não inclui Llama, GGUF, runtime de inferência ou socket de modelo; não use os caminhos abaixo como configuração atual. Qualquer inferência futura será especificada separadamente para OpenAI, com controle de gastos.

> **Status: refinement — contrato aprovado conceitualmente, implementação ainda pendente.** Esta página descreve o comportamento desejado; o binário atual **não** tem o comando `admin`, não cria usuários/grupos nem gerencia credenciais.
>
> Referência normativa exclusivamente de leitura: `ahhvahh/doc-arquitetura@main/ProjectCharter/1.0`. Projeto aplicativo: `ahhvahh/doc-arquitetura-app-debian@develop`. Git no repositório-alvo: **branch de tarefa criada de `develop`**, sem alterar diretamente `main` nem `develop`.

## 1. Janela administrativa — somente terminal

O operador inicia `sudo doc-arquitetura admin`. O próprio binário verifica `geteuid()==0` **antes de abrir o editor ou executar qualquer mutação**; sem root, informa como executar com `sudo` e termina. Não haverá executável setuid, endpoint root no socket de clientes nem elevação automática solicitada pelo daemon. O `sudo` do sistema faz a autenticação; futuramente `pkexec` pode ser opção, mas não é requisito.

```text
$ sudo doc-arquitetura admin
DOC ARQUITETURA — ADMINISTRAÇÃO (root)
Arquivo: /etc/doc-arquitetura-app/config.yaml

1. Identidade do serviço          docarq-svc / docarq-runtime
2. Grupo de acesso ao terminal    docarq-clients
3. Usuários autorizados           2 usuários cadastrados
4. Credencial GitHub              não configurada
5. Repositórios permitidos        0 configurados
6. Política de execução e Git     develop → docs/<etapa>/<nome-id>
7. Verificar segurança / testar configuração
8. Prévia das alterações e aplicar
0. Sair sem aplicar
```

A tela não reproduz segredo, senha ou chave privada. Alterações ficam em rascunho até `8. Prévia e aplicar`; se não forem confirmadas, não gravar. Após confirmar, validar todos os campos, instalar/chmod grupos e credenciais, gravar configuração atomicamente e recarregar/reiniciar o serviço com confirmação separada quando necessário. Se o serviço estiver com tarefas ativas, avisar e não interromper silenciosamente.

## 2. Arquivo de configuração (sem segredo)

Local obrigatório: **`/etc/doc-arquitetura-app/config.yaml`**. Proprietário `root:docarq-runtime`, modo `0640`. Diretório `/etc/doc-arquitetura-app/`: `root:docarq-runtime`, modo `0750`. `docarq-svc` **lê** o arquivo, mas nunca o escreve. Nenhum comando do cliente comum pode modificar a configuração. Exemplo ilustrativo:

```yaml
schema_version: 1
service:
  user: docarq-svc
  group: docarq-runtime
access:
  group: docarq-clients
  authorized_users:
    - operador1
    - operador2
github:
  provider: github.com
  authentication:
    type: github_app
    app_id: 123456
    installation_id: 987654
    credential_name: github-app-private-key
  allowed_repositories:
    - ahhvahh/a-bioma-sauva
git_policy:
  base_branch: develop
  task_branch_prefix: docs
  max_branch_name_bytes: 100
  forbidden_destinations:
    - refs/heads/main
    - refs/heads/develop
execution:
  concurrent_model_instances: 1
  runtime_context_tokens: 4096
  require_approval_to_apply: true
  require_approval_to_commit: true
  require_approval_to_push: true
```

Os nomes de usuário devem existir no sistema local, resolver via NSS e não ser `root` nem a própria conta de serviço. A conta técnica pode ser escolhida na tela **somente por root**, deve existir ou ser criada como *system user*, sem login interativo, sem shell e sem UID 0. A identidade selecionada deve coincidir com `User=` e `Group=` da unidade systemd, controlados por drop-in `/etc/systemd/system/doc-arquitetura-app.service.d/10-identity.conf` gerado pela administração root. Trocar essa identidade exige migração controlada da propriedade dos diretórios privados e restart aprovado; não usar `chown -R` indiscriminadamente.

Os parâmetros `base_branch=develop`, `forbidden_destinations=[main, develop]`, referência normativa `main` e destinos documentais `/docs/{charter,scope,implementation,backup}` são **invariantes de segurança**: configuração malformada ou tentativa de relaxá-los bloqueia a execução; o root não pode simplesmente mudar um campo para fazer o agente escrever em `main`.

## 3. Grupos, usuários e autorização

- Grupo **`docarq-clients`**: autoriza a conexão ao socket público de comandos; membros **não** leem workspace, GGUF, chave, logs privados nem socket do Llama.
- Grupo **`docarq-runtime`**: grupo técnico da conta de serviço. Operadores humanos nunca são adicionados pelo assistente administrativo a esse grupo. O serviço usa `Group=docarq-runtime`, não `docarq-clients`.
- Permitir cadastrar/remover usuários existentes no `docarq-clients` e na lista `authorized_users` de forma sincronizada. Nunca usar `usermod -G` que remova grupos preexistentes; preferir APIs do sistema, ou comandos `usermod -aG` e `gpasswd -d` com argumentos validados e `execve`, sem shell.
- **Dois controles cumulativos** no daemon: permissão do socket e autenticação por `SO_PEERCRED`; resolver UID para usuário e conferir **associação ao grupo + presença em `authorized_users`** em cada comando. `UID` informado no JSON pelo cliente é ignorado para autenticação.
- As operações `task.approve`, `task.commit`, `task.publish` só aceitam o UID que criou a tarefa, ou administrador explicitamente configurado em política separada. Não presumir que todos no grupo possam aprovar tarefas alheias.
- Revogação deve valer para **novos comandos imediatamente** após recarga; conexões já abertas precisam ser revalidadas por operação e encerradas ao revogar. Processos interativos podem necessitar novo login para obter grupo adicionado; a lista explícita impede que grupo ainda presente em sessão antiga mantenha acesso após revogação.
- O root pode executar a ferramenta administrativa; usuários não root, inclusive os clientes autorizados, usam somente a CLI comum conectada ao daemon. **O daemon nunca executa solicitações com privilégios root.**

## 4. GitHub — credenciais separadas

A opção 4 pede **tipo da credencial**, sem solicitar que o usuário digite segredo em uma linha de comando:

- **GitHub App (preferencial para múltiplos repositórios)**: coletar `app_id`, `installation_id` e caminho de um PEM privado existente, validar permissão e formato, apresentar apenas fingerprint, copiar para `/etc/doc-arquitetura-app/secrets/github-app.pem` e restringir a `root:root 0600`. O systemd usa `LoadCredential=github-app-private-key:/etc/doc-arquitetura-app/secrets/github-app.pem` para fornecer ao daemon uma cópia temporária privada; o serviço obtém token de instalação de curta duração apenas quando necessário e solicita só permissões `contents:read/write` e metadados indispensáveis.
- **Chave SSH dedicada (alternativa)**: informar caminho para chave privada de **identidade exclusiva do serviço**, separada da chave pessoal, instalada no mesmo diretório root restrito. Configurar repositórios autorizados, host key do GitHub confiável e verificação estrita; só funcionarão repositórios onde a chave pública correspondente tiver acesso de escrita. Chaves `deploy key` normalmente são limitadas a um repositório; usar credenciais distintas conforme a organização.
- **Não** colocar token, PEM ou senha dentro de `config.yaml`, URLs Git, argumento de processo, histórico de shell, variáveis impressas, eventos de progresso ou `/docs/backup/`. Sem arquivo de credencial válida, operações privadas e push ficam **BLOCKED**.

O diretório `/etc/doc-arquitetura-app/secrets/` fica `root:root 0700`, nunca acessível a `docarq-clients` ou ao grupo técnico. O serviço só recebe credencial pelo mecanismo privado do systemd, nunca lê o PEM na pasta de origem. A validação testa autenticação **sem enviar commits**, com resultado `OK / sem acesso / sem permissão de escrita`, sem revelar segredos.

## 5. Socket público e socket privado

Usar **socket activation systemd**, sem permitir que o daemon precise de privilégios root para mudar o grupo do socket:

```ini
# /usr/lib/systemd/system/doc-arquitetura-app.socket
[Unit]
Description=Doc Arquitetura — entrada IPC restrita

[Socket]
ListenStream=/run/doc-arquitetura-app/control.sock
SocketUser=root
SocketGroup=docarq-clients
SocketMode=0660
DirectoryMode=0755
RemoveOnStop=yes

[Install]
WantedBy=sockets.target
```

O diretório pai do socket público é acessível para travessia, mas o arquivo do socket aceita somente root e `docarq-clients`; o daemon valida `SO_PEERCRED` e a lista de usuários também. A unidade `doc-arquitetura-app.service` deve consumir o FD de socket herdado do systemd e executar com `User=docarq-svc` e `Group=docarq-runtime`, sem pertencer ao grupo de clientes.

Separar o diretório privado para evitar conceder `docarq-clients` acesso ao Llama: `/run/doc-arquitetura-app-private/`, dono `docarq-svc:docarq-runtime`, modo `0700`. Socket interno `llama.sock` com modo `0600`. O binário `/usr/lib/doc-arquitetura-app/llama-server` terá dono `root:docarq-runtime` e modo `0750`. Nenhum arquivo do runtime, clone ou modelo em pasta de usuário final. Ao ligar o serviço, o cliente comum **não** pode mais criar processo de inferência diretamente — desativar os comandos legados de execução e escrita no processo cliente.

## 6. Aplicação segura e verificação

O assistente root deverá:
1. Ler `config.yaml` atual com proteção contra symlink e verificar modo/proprietário, sem seguir caminhos suspeitos;
2. Confirmar existência da identidade do serviço, grupos e usuários autorizados;
3. Importar chaves somente de caminhos explícitos e verificar permissões; não imprimir o conteúdo;
4. Mostrar plano `antes → depois` sem secretos, com alterações de grupos e restart;
5. Fazer backup root-only **do arquivo de configuração anterior** fora do repositório, escrever novo arquivo via criação temporária `0600`, validação, `fsync` e rename atômico; ajustar para `0640` somente após owner/group;
6. Aplicar grupos e unidades de forma transacional quando possível; em falhas, interromper habilitação do daemon e reportar quais mudanças de sistema exigem reconciliação manual (criação de conta e grupo não é uma transação atômica);
7. Rodar verificações administrativas: UID real do daemon, modos DAC, conexão negada a usuário externo, cliente permitido pelo grupo+lista, credenciais carregáveis pelo serviço sem exposição, destino Git sempre na branch da tarefa;
8. Iniciar/habilitar o serviço somente após passar os testes; nunca executar `chmod 777`, `NOPASSWD` genérico, instalar `llama-server` setuid ou passar segredos para o processo do cliente.

**Dependências pendentes:** implementação da CLI `admin`, parser e esquema da configuração, testabilidade da migração UID/GID, daemon com `SO_PEERCRED`, suporte a socket activation, autenticação GitHub App/SSH, unidade systemd e empacotamento. O pacote atual **não deve criar os usuários e credenciais automaticamente** antes desses gates serem atendidos.
