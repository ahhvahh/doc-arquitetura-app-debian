#!/usr/bin/env bash
set -Eeuo pipefail
script="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd -P)/instalar-deb-ssh.sh"
tmp="$(mktemp -d)"
trap 'rm -rf -- "$tmp"' EXIT
mkdir -p "$tmp/packages/doc-arquitetura-app-debian/develop/amd64/2026-10-19_14-21-33" "$tmp/packages/doc-arquitetura-app-debian/develop/amd64/2026-10-19_16-11-40"
printf '%s\n' 'https://github.com/ahhvahh/doc-arquitetura-app-debian.git' > "$tmp/projetos.conf"
debroot="$tmp/pkg"; mkdir -p "$debroot/DEBIAN"
cat > "$debroot/DEBIAN/control" <<'EOF'
Package: doc-arquitetura-app
Version: 0.2.0
Architecture: amd64
Maintainer: tests <test@example.com>
Description: pacote de teste
EOF
dpkg-deb --build "$debroot" "$tmp/packages/doc-arquitetura-app-debian/develop/amd64/2026-10-19_14-21-33/doc-arquitetura-app_0.2.0_amd64.deb" >/dev/null
result="$(printf '1\n1\n' | ABIOMA_ROOT="$tmp" NO_COLOR=1 bash "$script" --list-builds)"
[[ "$result" == *"0.2.0_amd64.deb"* ]] || { printf "%s\n" "$result"; exit 1; }
[[ "$result" == *"falhou ou não gerou .deb"* ]] || { printf "%s\n" "$result"; exit 1; }
mkdir -p "$tmp/packages/doc-arquitetura-app-debian/develop/amd64/2026-10-19_17-00-00"
sed -i 's/Version: 0.2.0/Version: 0.2.1/' "$debroot/DEBIAN/control"
dpkg-deb --build "$debroot" "$tmp/packages/doc-arquitetura-app-debian/develop/amd64/2026-10-19_17-00-00/doc-arquitetura-app_0.2.1_amd64.deb" >/dev/null
result="$(printf '1\n\n' | ABIOMA_ROOT="$tmp" NO_COLOR=1 bash "$script" --list-builds)"
[[ "$result" == *"[MAIS RECENTE]"* && "$result" == *"0.2.1_amd64.deb"* ]] || { printf "%s\n" "$result"; exit 1; }
printf 'OK: projetos.conf, build falho ignorado e último pacote válido escolhido\n'
