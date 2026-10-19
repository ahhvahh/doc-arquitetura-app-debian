# Scripts — Doc Arquitetura

## Instalador remoto Debian

Execute no **desktop Debian**, dentro da pasta onde estão os pacotes ou suas subpastas:

```bash
bash /caminho/do/projeto/src/scripts/instalar-deb-ssh.sh
```

A ferramenta solicita IP e usuário remoto, autentica via SSH e oferece:

- Instalar um `.deb` encontrado em pastas locais ou busca recursiva;
- Listar/remover pacotes instalados pelo prefixo;
- Consultar o estado do Doc Arquitetura e abrir um terminal remoto comum.

**Antes de instalar**, a aplicação envia o pacote para `/tmp`, valida SHA-256 e libera leitura do arquivo ao usuário `_apt` com `0644`. Só então consulta o `dpkg-query` remoto para localizar o pacote **idêntico** e nomes **semelhantes da mesma família**, exibindo as versões.

Ao encontrar uma versão anterior, oferece três escolhas: atualizar **sem remoção** (padrão recomendado), escolher um pacote para remover **com confirmação explícita**, ou cancelar. O comando de remoção é `sudo apt-get remove` **sem `-y`**, permitindo conferir as dependências afetadas. Uma falha ou recusa de remoção bloqueia a nova instalação. O arquivo temporário em `/tmp` é apagado; a sessão final SSH não fica como root.

**Aviso:** se optar por remover antes e a instalação nova falhar, a versão antiga **não é restaurada automaticamente**. Para atualizar o mesmo pacote, normalmente é melhor preservar o instalado e permitir que o APT atualize. Não envie arquivos `.deb` contendo segredos, pois o arquivo temporário passa a ser legível durante a operação.

As cores ANSI usam azul nos cabeçalhos, ciano nas etapas, verde para sucesso, amarelo para avisos, vermelho para erros e magenta para nomes semelhantes. Elas são desativadas em terminais não interativos ou com `NO_COLOR=1`.

Validação local (sem rede):

```bash
bash -n src/scripts/instalar-deb-ssh.sh
bash src/scripts/tests/test-instalar-deb-ssh.sh
```

O teste cobre seis decisões simuladas. A instalação real requer SSH configurado e usuário remoto com autorização `sudo`.

## Runtime privado

`build-private-llama.sh` constrói o executável privado do llama.cpp durante o empacotamento Debian e não é utilizado no desktop para instalação remota.

## Seleção por projetos.conf e histórico do build

O instalador procura `projetos.conf` e `packages/` na pasta do próprio instalador/build, depois no diretório atual e em diretórios ancestrais. Pode informar `ABIOMA_ROOT=/caminho/abioma` se necessário. A opção de instalação lista apenas os repositórios cadastrados que possuem saída em `packages/` e examina `packages/<projeto>/develop/amd64/<YYYY-MM-DD_HH-MM-SS>/`. Ordena por data descendente e valida os metadados de cada `.deb` com `dpkg-deb`; diretórios vazios, arquivos truncados ou builds falhos não são tratados como artefato instalável.

Se o diretório mais recente não contiver `.deb` válido, o instalador avisa que as versões disponíveis são anteriores e não as seleciona automaticamente. Para conferir localmente sem conexão SSH: `bash instalar-deb-ssh.sh --list-builds`. Se o arquivo não estiver na mesma pasta de `projetos.conf`/`packages`, execute `ABIOMA_ROOT=/home/usuario/Documentos/Projetos/abioma bash instalar-deb-ssh.sh --list-builds`.
