#!/usr/bin/env bash
set -Eeuo pipefail
SCRIPT="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)/instalar-deb-ssh.sh"
bash -n "$SCRIPT"
helper="$(sed -n '/^check_previous_installation()/,/^install_package()/p' "$SCRIPT" | sed '$d')"
[[ -n "$helper" ]] || { echo "Falha: função de pré-instalação ausente"; exit 1; }

case_test() (
    local label="$1" records="$2" input="$3" want_status="$4" want_remove="$5"
    local removed=0 rc=0
    C_RESET="" C_GREEN="" C_YELLOW="" C_RED="" C_CYAN="" C_MAGENTA=""
    warning() { :; }; success() { :; }; failure() { :; }; step() { :; }
    ssh_remote() {
        if [[ "$1" == *"Status-Status"* ]]; then
            if (( removed )); then printf 'config-files\n'; else printf 'installed\n'; fi
        else
            printf '%s\n' "$records"
        fi
    }
    ssh_remote_tty() { removed=1; }
    eval "$helper"
    if check_previous_installation doc-arquitetura-app 0.2.1 <<< "$input" > /dev/null; then rc=0; else rc=$?; fi
    if [[ "$rc" -ne "$want_status" || "$removed" -ne "$want_remove" ]]; then
        printf 'Falha %s: rc=%s removed=%s (esperado %s %s)\n' "$label" "$rc" "$removed" "$want_status" "$want_remove" >&2
        return 1
    fi
    printf 'OK: %s\n' "$label"
)
same=$'ii \tdoc-arquitetura-app\t0.2.0\nii \tvim\t9.0'
similar=$'ii \tdoc-arquitetura-app\t0.2.0\nii \tdoc-arquitetura-cli\t0.1.0'
case_test "manter versão" "$same" "1" 0 0
case_test "remover pacote exato" "$same" $'2\n1\nREMOVER' 0 1
case_test "remover nome parecido" "$similar" $'2\n2\nREMOVER' 0 1
case_test "rejeitar remoção" "$same" $'2\n1\nNAO' 2 0
case_test "cancelar" "$same" "3" 2 0
case_test "instalação limpa" $'ii \tvim\t9.0' "" 0 0
echo "Seis cenários passaram (sem SSH nem APT real)."
