#!/usr/bin/env bash
# llama.cpp is built inside the isolated Debian package workspace, never installed globally.
set -Eeuo pipefail
dest="${1:?usage: build-private-llama.sh OUTPUT_DIRECTORY}"
for exe in git cmake c++ install; do
    command -v "$exe" >/dev/null 2>&1 || { echo "Dependency required for build: $exe" >&2; exit 2; }
done

version="v0.7.0"
pinned_sha="620d2225cc436ab0b001e3200e590e3b286ef9ab"
upstream="https://github.com/ggml-org/llama.cpp.git"
temporary="$(mktemp -d "${TMPDIR:-/tmp}/doc-arq-llama-build.XXXXXXXX")"
trap 'rm -rf -- "$temporary"' EXIT

echo "[INFO] Clonando llama.cpp $version ($pinned_sha) apenas no workspace do build"
git clone --depth 1 --single-branch --branch "$version" "$upstream" "$temporary/llama.cpp"
actual="$(git -C "$temporary/llama.cpp" rev-parse HEAD)"
if [[ "$actual" != "$pinned_sha" ]]; then
    echo "[ERRO] Revisão llama.cpp inesperada: $actual (esperado $pinned_sha)" >&2
    exit 3
fi

# Generic, CPU-only amd64 build; static libraries within the executable.
# No external runtime, GPU SDK, app web UI, HTTP download support or OpenMP runtime.
cmake -S "$temporary/llama.cpp" -B "$temporary/build" \
    -DCMAKE_BUILD_TYPE=Release \
    -DBUILD_SHARED_LIBS=OFF \
    -DLLAMA_BUILD_SERVER=ON \
    -DLLAMA_BUILD_EXAMPLES=OFF \
    -DLLAMA_BUILD_APP=OFF \
    -DLLAMA_BUILD_UI=OFF \
    -DGGML_NATIVE=OFF \
    -DGGML_OPENMP=OFF \
    -DGGML_CUDA=OFF \
    -DLLAMA_CURL=OFF

jobs="${DOC_ARQUITETURA_LLAMA_BUILD_JOBS:-2}"
if ! [[ "$jobs" =~ ^[1-9][0-9]?$ ]]; then
    echo "[ERRO] DOC_ARQUITETURA_LLAMA_BUILD_JOBS deve ser 1..99" >&2
    exit 2
fi
cmake --build "$temporary/build" --target llama-server --parallel "$jobs"

test -x "$temporary/build/bin/llama-server" || { echo "[ERRO] llama-server não produzido" >&2; exit 4; }
mkdir -p "$dest"
install -m 0755 "$temporary/build/bin/llama-server" "$dest/llama-server"
install -m 0644 "$temporary/llama.cpp/LICENSE" "$dest/llama.cpp-LICENSE"
printf 'llama.cpp version=%s commit=%s\n' "$version" "$actual" > "$dest/llama.cpp-version.txt"
echo "[INFO] Runtime privado pronto: $dest/llama-server"
