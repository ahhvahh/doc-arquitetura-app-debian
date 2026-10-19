#!/usr/bin/env bash
# Gerencia pacotes Debian via SSH: instala .deb ou localiza e remove pacotes.
# No final, entrega uma sessão remota como o usuário comum.
# Senhas são solicitadas pelo OpenSSH/sudo e não são armazenadas pelo script.
set -Eeuo pipefail
# ANSI colors: terminal interativo apenas, NO_COLOR=1 ou TERM=dumb desativa.
if [[ -t 1 && -z "${NO_COLOR:-}" && "${TERM:-dumb}" != dumb ]]; then
    C_RESET=$'\033[0m'; C_BOLD=$'\033[1m'
    C_BLUE=$'\033[34m'; C_CYAN=$'\033[36m'; C_GREEN=$'\033[32m'
    C_YELLOW=$'\033[33m'; C_RED=$'\033[31m'; C_MAGENTA=$'\033[35m'
else
    C_RESET='' C_BOLD='' C_BLUE='' C_CYAN=''
    C_GREEN='' C_YELLOW='' C_RED='' C_MAGENTA=''
fi
heading() { printf '\n%s%s=== %s ===%s\n' "$C_BOLD" "$C_BLUE" "$1" "$C_RESET"; }
step() { printf '\n%s%s[%s]%s %s\n' "$C_BOLD" "$C_CYAN" "$1" "$C_RESET" "$2"; }
success() { printf '%s[OK]%s %s\n' "$C_GREEN" "$C_RESET" "$1"; }
warning() { printf '%s[ATENÇÃO]%s %s\n' "$C_YELLOW" "$C_RESET" "$1"; }
failure() { printf '%s[ERRO]%s %s\n' "$C_RED" "$C_RESET" "$1" >&2; }


for program in ssh scp dpkg-deb realpath mktemp sha256sum; do
    if ! command -v "$program" >/dev/null 2>&1; then
        printf 'ERRO: comando necessário não encontrado: %s\n' "$program" >&2
        exit 1
    fi
done

workdir="$(realpath -e -- "$PWD")"
control_dir=""
control_socket=""
remote_package=""
target=""

cleanup() {
    local original_status=$?
    trap - EXIT INT TERM

    if [[ -n "$remote_package" && -n "$target" && -S "$control_socket" ]]; then
        ssh -o BatchMode=yes -o ControlMaster=auto \
            -o ControlPath="$control_socket" "$target" \
            "rm -f -- '$remote_package'" >/dev/null 2>&1 || true
    fi
    if [[ -n "$target" && -S "$control_socket" ]]; then
        ssh -S "$control_socket" -O exit "$target" >/dev/null 2>&1 || true
    fi
    if [[ -n "$control_dir" && -d "$control_dir" ]]; then
        rm -rf -- "$control_dir"
    fi
    return "$original_status"
}
trap cleanup EXIT
trap 'exit 130' INT
trap 'exit 143' TERM

heading "Gerenciador remoto de pacotes Debian via SSH"
printf 'Diretório local de referência: %s\n\n' "$workdir"

read -r -p "${C_YELLOW}IP IPv4 do servidor: ${C_RESET}" server_ip
if [[ ! "$server_ip" =~ ^[0-9]{1,3}(\.[0-9]{1,3}){3}$ ]]; then
    printf 'ERRO: informe um endereço IPv4 válido.\n' >&2
    exit 2
fi
IFS='.' read -r a b c d <<< "$server_ip"
for octet in "$a" "$b" "$c" "$d"; do
    if (( 10#$octet > 255 )); then
        printf 'ERRO: octeto IPv4 fora do intervalo: %s\n' "$octet" >&2
        exit 2
    fi
done

read -r -p "${C_YELLOW}Usuário do servidor (com sudo): ${C_RESET}" server_user
if [[ ! "$server_user" =~ ^[a-z_][a-z0-9_-]*\$?$ ]] || [[ "$server_user" == root ]]; then
    printf 'ERRO: nome de usuário inválido ou root. Use um usuário normal com sudo.\n' >&2
    exit 2
fi

target="${server_user}@${server_ip}"
control_dir="$(mktemp -d "${TMPDIR:-/tmp}/doc-arq-ssh.XXXXXXXX")"
chmod 700 -- "$control_dir"
control_socket="$control_dir/ssh.sock"

step "SSH" "Autenticando em $target..."
printf 'O próprio SSH solicitará a senha, se necessária. Ela não será armazenada.\n'
ssh -M -N -f -S "$control_socket" \
    -o ConnectTimeout=15 \
    -o ServerAliveInterval=30 \
    -o ServerAliveCountMax=3 \
    "$target"

if ! ssh -S "$control_socket" -O check "$target" >/dev/null 2>&1; then
    printf 'ERRO: não foi possível estabelecer a conexão SSH.\n' >&2
    exit 1
fi
success 'Conexão SSH estabelecida.'

# Navegador de arquivos: inicia no PWD e nunca sobe para fora dele.
# A navegação é local: não precisa de permissões adicionais no servidor.
relative_display() {
    local path="$1" rel
    rel="$(realpath --relative-to="$workdir" -- "$path")" || return 1
    if [[ "$rel" == . ]]; then printf '.\n'; else printf './%s\n' "$rel"; fi
}

choose_from_recursive_search() {
    local -a results=()
    local i input
    mapfile -d '' -t results < <(find "$workdir" -type f -name '*.deb' -print0 | sort -z)
    if (( ${#results[@]} == 0 )); then
        printf '\nNenhum arquivo .deb encontrado em %s ou nas subpastas.\n' "$workdir"
        return 1
    fi
    printf '\nArquivos .deb encontrados (%d):\n' "${#results[@]}"
    for i in "${!results[@]}"; do
        printf '  %2d) %s\n' "$((i + 1))" "$(relative_display "${results[$i]}")"
    done
    while true; do
        if ! read -r -p 'Número do pacote (v para voltar): ' input; then return 1; fi
        [[ "$input" == [vV] ]] && return 1
        if [[ "$input" =~ ^[0-9]+$ ]] && (( 10#$input >= 1 && 10#$input <= ${#results[@]} )); then
            package_path="${results[$((10#$input - 1))]}"
            return 0
        fi
        printf 'Opção inválida. Selecione um número da lista.\n'
    done
}

choose_package() {
    local current="$workdir" input i count selected candidate
    local -a folders=() debs=() options=()

step "ARQUIVOS" 'Escolha um pacote .deb nas pastas locais.'
    printf 'Diretório inicial: %s\n' "$workdir"
    while true; do
        mapfile -d '' -t folders < <(find "$current" -mindepth 1 -maxdepth 1 -type d -print0 | sort -z)
        mapfile -d '' -t debs < <(find "$current" -mindepth 1 -maxdepth 1 -type f -name '*.deb' -print0 | sort -z)
        options=()
        printf '\nPasta atual: %s\n' "$(relative_display "$current")"
        printf '%s\n' '-------------------------------'
        if (( ${#folders[@]} == 0 && ${#debs[@]} == 0 )); then
            printf '(Não há subpastas nem arquivos .deb nesta pasta.)\n'
        fi
        for candidate in "${folders[@]}"; do
            options+=("$candidate")
            printf '  %2d) [PASTA] %s/\n' "${#options[@]}" "$(basename -- "$candidate")"
        done
        for candidate in "${debs[@]}"; do
            options+=("$candidate")
            printf '  %2d) [PACOTE] %s\n' "${#options[@]}" "$(basename -- "$candidate")"
        done
        printf '\n  B) Buscar todos os .deb nas subpastas\n'
        if [[ "$current" != "$workdir" ]]; then printf '  V) Voltar à pasta anterior\n'; fi
        printf '  M) Informar caminho relativo manualmente\n'
        printf '  S) Voltar ao menu principal\n'
        if ! read -r -p "${C_BOLD}${C_YELLOW}Escolha: ${C_RESET}" input; then return 1; fi
        case "$input" in
            [bB])
                if choose_from_recursive_search; then return 0; fi
                ;;
            [vV])
                if [[ "$current" == "$workdir" ]]; then
                    printf 'Já estamos na pasta inicial.\n'
                else
                    current="$(dirname -- "$current")"
                fi
                ;;
            [mM])
                if ! read -r -p 'Caminho relativo a partir da pasta inicial: ' selected; then return 1; fi
                if [[ -z "$selected" || "$selected" == /* ]]; then
                    printf 'Informe um caminho relativo não vazio.\n'
                    continue
                fi
                if ! candidate="$(realpath -e -- "$workdir/$selected" 2>/dev/null)"; then
                    printf 'Arquivo não encontrado: %s\n' "$selected"
                    continue
                fi
                case "$candidate" in
                    "$workdir"/*) ;;
                    *) printf 'Caminho fora da pasta inicial não permitido.\n'; continue ;;
                esac
                if [[ -f "$candidate" && "$candidate" == *.deb ]]; then
                    package_path="$candidate"
                    return 0
                fi
                printf 'Selecione um arquivo regular com extensão .deb.\n'
                ;;
            [sS])
                printf 'Operação cancelada.\n'
                return 1
                ;;
            *)
                if [[ "$input" =~ ^[0-9]+$ ]] && (( 10#$input >= 1 && 10#$input <= ${#options[@]} )); then
                    candidate="${options[$((10#$input - 1))]}"
                    if [[ -d "$candidate" ]]; then
                        current="$candidate"
                    else
                        package_path="$candidate"
                        return 0
                    fi
                else
                    printf 'Opção inválida. Escolha uma pasta, um pacote ou uma letra do menu.\n'
                fi
                ;;
        esac
    done
}

# As operações são explícitas e executam sempre como usuário comum.
# Apenas o apt usa sudo, sem abrir shell root permanente.

ssh_remote() {
    ssh -T -o BatchMode=yes -o ControlMaster=auto \
        -o ControlPath="$control_socket" "$target" "$@"
}

ssh_remote_tty() {
    ssh -tt -o BatchMode=yes -o ControlMaster=auto \
        -o ControlPath="$control_socket" "$target" "$@"
}


# Verificação antes da instalação: sem remoção automática.
check_previous_installation() {
    local wanted="$1" newversion="$2" family data status name ver label response idx chosen oldstatus
    local -a candidates=() versions=() kinds=()
    family="${wanted%-*}"
    step "VERSÕES" "Consultando versões já instaladas no servidor..."
    if ! data="$(ssh_remote "dpkg-query -W -f='\${db:Status-Abbrev}\\t\${binary:Package}\\t\${Version}\\n'")"; then
        failure "Não foi possível consultar o dpkg; nova instalação bloqueada."
        return 1
    fi
    while IFS=$'\t' read -r status name ver; do
        [[ "$status" == ii* ]] || continue
        [[ "$name" =~ ^[a-z0-9][a-z0-9.+-]*(:[a-z0-9]+)?$ ]] || continue
        label=""
        if [[ "${name%%:*}" == "$wanted" ]]; then
            label="EXATO"
        elif [[ "${name%%:*}" == "$wanted-"* ]] || [[ ${#family} -ge 8 && "${name%%:*}" == "$family-"* ]]; then
            label="PARECIDO"
        fi
        if [[ -n "$label" ]]; then
            candidates+=("$name");versions+=("$ver");kinds+=("$label")
        fi
    done <<< "$data"
    if (( ${#candidates[@]} == 0 )); then
        success "Nenhum pacote igual ou semelhante instalado."
        return 0
    fi
    warning "Encontrados ${#candidates[@]} pacote(s) antes de instalar $wanted ($newversion):"
    for idx in "${!candidates[@]}"; do
        if [[ "${kinds[$idx]}" == EXATO ]]; then
            printf '  %s%2d) [IGUAL]%s %-36s versão %s\n' "$C_YELLOW" "$((idx+1))" "$C_RESET" "${candidates[$idx]}" "${versions[$idx]}"
        else
            printf '  %s%2d) [PARECIDO]%s %-34s versão %s\n' "$C_MAGENTA" "$((idx+1))" "$C_RESET" "${candidates[$idx]}" "${versions[$idx]}"
        fi
    done
    printf '  %s1)%s Manter/atualizar normalmente (padrão recomendado)\n' "$C_GREEN" "$C_RESET"
    printf '  %s2)%s Desinstalar um pacote antes de instalar\n' "$C_YELLOW" "$C_RESET"
    printf '  %s3)%s Cancelar instalação\n' "$C_RED" "$C_RESET"
    while true; do
        read -r -p "Escolha [1/2/3, padrão 1]: " response || return 2
        case "$response" in
            ""|1) success "APT atualizará sem desinstalação prévia.";return 0 ;;
            3) warning "Instalação cancelada.";return 2 ;;
            2) break ;;
            *) warning "Opção inválida." ;;
        esac
    done
    while true; do
        read -r -p "Número do pacote a remover (V para voltar): " response || return 2
        [[ "$response" == [vV] ]] && return 2
        if [[ "$response" =~ ^[0-9]{1,3}$ ]] && (( 10#$response>=1 && 10#$response<=${#candidates[@]} )); then
            chosen="${candidates[$((10#$response-1))]}";break
        fi
        warning "Número inválido."
    done
    warning "Se a nova instalação falhar, $chosen NÃO será restaurado automaticamente."
    warning "APT mostrará o plano e pedirá sua confirmação, sem -y."
    read -r -p "Digite REMOVER para confirmar a desinstalação de $chosen: " response || return 2
    [[ "$response" == REMOVER ]] || { warning "Remoção cancelada.";return 2; }
    if ! ssh_remote_tty "sudo apt-get remove -- '$chosen'; result=\$?; sudo -k; exit \$result"; then
        failure "Remoção falhou; nova instalação cancelada.";return 1
    fi
    oldstatus="$(ssh_remote "dpkg-query -W -f='\${db:Status-Status}' -- '$chosen' 2>/dev/null || true")" || return 1
    if [[ "$oldstatus" == installed ]]; then
        failure "O pacote ainda está instalado; instalação cancelada.";return 1
    fi
    success "$chosen removido; continuando para a nova instalação."
    return 0
}

install_package() {
    local package_path="" package_name package_version package_arch confirmation
    local local_sum local_hash remote_sum remote_hash
    if ! choose_package; then
        printf 'Retornando ao menu principal.\n'
        return 2
    fi

    # Proteger contra resolução para fora do diretório inicial.
    if ! package_path="$(realpath -e -- "$package_path")" || \
       [[ ! -f "$package_path" || "$package_path" != *.deb ]]; then
        printf 'ERRO: selecione um arquivo .deb regular.\n' >&2
        return 1
    fi
    case "$package_path" in
        "$workdir"/*) ;;
        *) printf 'ERRO: pacote fora do diretório inicial.\n' >&2; return 1 ;;
    esac
    printf '\nPacote selecionado: %s\n' "$(relative_display "$package_path")"
    if ! package_name="$(dpkg-deb -f "$package_path" Package)" || \
       ! package_version="$(dpkg-deb -f "$package_path" Version)" || \
       ! package_arch="$(dpkg-deb -f "$package_path" Architecture)"; then
        printf 'ERRO: não foi possível ler os metadados do .deb.\n' >&2
        return 1
    fi
    printf 'Nome: %s | Versão: %s | Arquitetura: %s\n' \
        "$package_name" "$package_version" "$package_arch"
    if ! read -r -p 'Confirma envio e instalação no servidor? [s/N] ' confirmation; then return 2; fi
    case "$confirmation" in [sS]|[sS][iI][mM]) ;; *) printf 'Cancelado.\n'; return 2 ;; esac

step 'TRANSFERÊNCIA' 'Reservando arquivo temporário no servidor...'
    if ! remote_package="$(ssh_remote 'mktemp /tmp/doc-arquitetura-app.XXXXXXXX.deb')"; then
        printf 'ERRO: não foi possível criar arquivo temporário remoto.\n' >&2
        return 1
    fi
    if [[ ! "$remote_package" =~ ^/tmp/doc-arquitetura-app\.[a-zA-Z0-9]{8}\.deb$ ]]; then
        printf 'ERRO: nome temporário remoto inesperado: %q\n' "$remote_package" >&2
        remote_package=""
        return 1
    fi
    printf 'Enviando para %s:%s\n' "$target" "$remote_package"
    if ! scp -o BatchMode=yes -o ControlMaster=auto \
        -o ControlPath="$control_socket" -- "$package_path" "$target:$remote_package"; then
        printf 'ERRO: transferência falhou.\n' >&2
        return 1
    fi

    local_sum="$(sha256sum -- "$package_path")"
    local_hash="${local_sum%% *}"
    if ! remote_sum="$(ssh_remote "sha256sum -- '$remote_package'")"; then
        printf 'ERRO: não foi possível verificar o hash remoto.\n' >&2
        return 1
    fi
    remote_hash="${remote_sum%% *}"
    if [[ ! "$remote_hash" =~ ^[0-9a-f]{64}$ || "$remote_hash" != "$local_hash" ]]; then
        printf 'ERRO: integridade SHA-256 divergente. Instalação cancelada.\n' >&2
        return 1
    fi
success "SHA-256 confirmado: $local_hash"

    # mktemp cria modo 0600. O _apt necessita ler o .deb local para baixar
    # de modo não privilegiado. chmod 0644 evita fallback como root.
    # /tmp é sticky (1777): apenas o dono/root pode apagar este arquivo.
    # O .deb não deve conter chaves, tokens ou outros segredos.
    if ! ssh_remote "chmod 0644 -- '$remote_package' && test -r '$remote_package'"; then
        printf 'ERRO: falha ao ajustar leitura do arquivo temporário.\n' >&2
        return 1
    fi
success 'Permissão 0644 aplicada: _apt pode ler o pacote.'
    if ! check_previous_installation "$package_name" "$package_version"; then
        warning "Instalação cancelada; removendo o pacote temporário de /tmp."
        if ssh_remote "rm -f -- '$remote_package'"; then remote_package=""; fi
        return 2
    fi

step 'INSTALAÇÃO' 'O sudo pedirá a senha, se necessário.'
    if ! ssh_remote_tty \
        "sudo apt-get install -y -- '$remote_package'; result=\$?; sudo -k; exit \$result"; then
        printf 'ERRO: instalação remota falhou.\n' >&2
        return 1
    fi
    printf '\nInstalação concluída. Apagando pacote transitório de /tmp...\n'
    if ! ssh_remote "rm -f -- '$remote_package'"; then
        printf 'ATENÇÃO: instalação concluída, mas não foi possível limpar o .deb temporário.\n' >&2
        # Mantém remote_package preenchido para nova tentativa no trap EXIT.
        return 0
    fi
    remote_package=""
    printf 'O pacote instalado é controlado por APT; não precisa do .deb em /tmp para remover depois.\n'
    return 0
}

remove_package() {
    local prefix raw status name version input choice confirmation selected="" selected_version=""
    local -a names=() versions=()
    local query
heading 'Pesquisar pacotes instalados no servidor'
    while true; do
        if ! read -r -p 'Começo do nome (ex.: doc-arq; V para voltar): ' prefix; then return 2; fi
        [[ "$prefix" == [vV] ]] && return 2
        if [[ ${#prefix} -ge 2 && ${#prefix} -le 80 && "$prefix" =~ ^[a-z0-9][a-z0-9.+-]*$ ]]; then
            break
        fi
        printf 'Informe pelo menos 2 caracteres minúsculos, números, hífen, ponto ou +.\n'
    done

    # Consultar informações reais do dpkg no Debian remoto. O prefixo é
    # filtrado LOCALMENTE, evitando interpolá-lo como comando shell remoto.
    query="$(cat <<'REMOTE_QUERY'
dpkg-query -W -f='${db:Status-Abbrev}\t${binary:Package}\t${Version}\n'
REMOTE_QUERY
)"
    if ! raw="$(ssh_remote "$query")"; then
        printf 'ERRO: consulta de pacotes no servidor falhou.\n' >&2
        return 1
    fi
    while IFS=$'\t' read -r status name version; do
        [[ "$status" == ii* ]] || continue
        [[ "$name" == "$prefix"* ]] || continue
        if [[ ! "$name" =~ ^[a-z0-9][a-z0-9.+-]*(:[a-z0-9]+)?$ ]]; then
            continue
        fi
        names+=("$name")
        versions+=("$version")
    done <<< "$raw"
    if (( ${#names[@]} == 0 )); then
        printf 'Nenhum pacote instalado começa com "%s".\n' "$prefix"
        return 2
    fi
    printf '\nPacotes instalados encontrados (%d):\n' "${#names[@]}"
    for input in "${!names[@]}"; do
        printf '  %2d) %-42s versão %s\n' "$((input+1))" "${names[$input]}" "${versions[$input]}"
    done
    while true; do
        if ! read -r -p 'Número do pacote para desinstalar (V para voltar): ' choice; then return 2; fi
        [[ "$choice" == [vV] ]] && return 2
        if [[ "$choice" =~ ^[0-9]+$ ]] && (( 10#$choice >= 1 && 10#$choice <= ${#names[@]} )); then
            selected="${names[$((10#$choice-1))]}"
            selected_version="${versions[$((10#$choice-1))]}"
            break
        fi
        printf 'Opção inválida.\n'
    done
    printf '\nATENÇÃO: remover %s, versão %s, do servidor %s.\n' \
        "$selected" "$selected_version" "$server_ip"
    printf 'APT mostrará dependências afetadas; não usamos -y na remoção.\n'
    if ! read -r -p 'Digite REMOVER para continuar (outro valor cancela): ' confirmation; then return 2; fi
    if [[ "$confirmation" != REMOVER ]]; then
        printf 'Remoção cancelada.\n'
        return 2
    fi
    if ! ssh_remote_tty \
        "sudo apt-get remove -- '$selected'; result=\$?; sudo -k; exit \$result"; then
        printf 'ERRO: a remoção não foi concluída.\n' >&2
        return 1
    fi
    printf '\nRemoção processada pelo APT. Arquivos de configuração podem permanecer; não foi usado purge.\n'
    return 0
}

show_diagnostics() {
heading 'Testes de leitura no servidor'
    ssh -T -o BatchMode=yes -o ControlMaster=auto \
        -o ControlPath="$control_socket" "$target" 'sh -s' <<'REMOTE_TEST'
printf '\n[Pacote instalado]\n'
dpkg-query -W -f='${binary:Package} ${Version} (${db:Status-Status})\n' doc-arquitetura-app || true
printf '\n[Comando CLI]\n'
command -v doc-arquitetura || true
printf '\n[Runtime privado]\n'
if [ -x /usr/lib/doc-arquitetura-app/llama-server ]; then
    /usr/lib/doc-arquitetura-app/llama-server --version 2>&1 | head -n 4
else
    printf 'Runtime não encontrado ou não executável.\n'
fi
printf '\n[Rotas da DSL]\n'
if command -v doc-arquitetura >/dev/null 2>&1; then
    doc-arquitetura routes || true
    printf '\n[Modelos disponíveis para o usuário atual]\n'
    doc-arquitetura models list || true
fi
printf '\n[Fim das verificações somente leitura]\n'
REMOTE_TEST
}

open_remote_shell() {
    printf '\nAbrindo terminal em %s como usuário comum %s.\n' "$server_ip" "$server_user"
    printf 'Para voltar ao desktop, execute: exit\n\n'
    ssh -tt -o BatchMode=yes -o ControlMaster=auto \
        -o ControlPath="$control_socket" "$target"
}

while true; do
heading "Gerenciamento de pacotes em $server_ip"
    printf '  %s1)%s Enviar e instalar um pacote .deb\n' "$C_GREEN" "$C_RESET"
    printf '  %s2)%s Pesquisar e desinstalar pacote instalado (por prefixo)\n' "$C_YELLOW" "$C_RESET"
    printf '  %s3)%s Verificar o Doc Arquitetura instalado\n' "$C_CYAN" "$C_RESET"
    printf '  %s0)%s Sair sem realizar operação\n' "$C_RED" "$C_RESET"
    if ! read -r -p "${C_BOLD}${C_YELLOW}Opção: ${C_RESET}" selected_operation; then exit 0; fi
    case "$selected_operation" in
        1) if install_package; then open_remote_shell; exit $?; fi ;;
        2) if remove_package; then open_remote_shell; exit $?; fi ;;
        3) if show_diagnostics; then open_remote_shell; exit $?; fi ;;
        0) printf 'Encerrado.\n'; exit 0 ;;
        *) printf 'Opção inválida.\n' ;;
    esac
done
