#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
command -v jq >/dev/null
command -v go >/dev/null
catalog=test/groups.json
mode="${1:---all}"
if [[ "$mode" == --list ]]; then
  jq -r 'to_entries[] | [.key, .value.expected] | @tsv' "$catalog"
  exit 0
fi
sandbox_dir=$(mktemp -d /tmp/doc-tests.XXXXXXXX)
trap 'rm -rf "$sandbox_dir"' EXIT
unset OPENAI_API_KEY GH_TOKEN GITHUB_TOKEN
unset DOC_ARQUITETURA_CONFIG_DIR DOC_ARQUITETURA_DATA_DIR DOC_ARQUITETURA_CACHE_DIR DOC_ARQUITETURA_STATE_DIR DOC_ARQUITETURA_TEMP_DIR DOC_ARQUITETURA_DSL DOC_ARQUITETURA_LLAMA_CONTEXT DOC_ARQUITETURA_LLAMA_BIN
export XDG_CONFIG_HOME="$sandbox_dir/config" XDG_DATA_HOME="$sandbox_dir/data"
export XDG_CACHE_HOME="$sandbox_dir/cache" XDG_STATE_HOME="$sandbox_dir/state"
export TMPDIR="$sandbox_dir/tmp" GIT_CONFIG_GLOBAL=/dev/null GIT_CONFIG_NOSYSTEM=1
mkdir -p "$XDG_CONFIG_HOME" "$XDG_DATA_HOME" "$XDG_CACHE_HOME" "$XDG_STATE_HOME" "$TMPDIR"
go test -list '^Test' ./... > "$sandbox_dir/discovered"
grep '^Test' "$sandbox_dir/discovered" | LC_ALL=C sort -u > "$sandbox_dir/actual"
jq -r '.[].tests[] | split("/")[0]' "$catalog" | LC_ALL=C sort -u > "$sandbox_dir/expected"
if ! diff -u "$sandbox_dir/expected" "$sandbox_dir/actual"; then
  echo 'Test inventory differs from test/groups.json' >&2
  exit 1
fi
if [[ "$mode" == --check ]]; then exit 0; fi
if [[ "$mode" == --all ]]; then
  mapfile -t groups < <(jq -r 'keys[]' "$catalog")
elif [[ "$mode" == --group && $# == 2 ]]; then
  jq -e --arg g "$2" 'has($g)' "$catalog" >/dev/null
  groups=("$2")
else
  echo 'Usage: bash test.sh [--all|--list|--check|--group NAME]' >&2
  exit 2
fi
overall=0
for group in "${groups[@]}"; do
  result_dir="test-results/$group"
  mkdir -p "$result_dir"
  jq -r --arg g "$group" '.[$g].tests[]' "$catalog" > "$result_dir/expected.txt"
  expression=$(jq -r --arg g "$group" '.[$g].tests | map(split("/") | map("^" + . + "$") | join("/")) | join("|")' "$catalog")
  echo "Running $group"
  set +e
  CGO_ENABLED=1 go test -race -count=1 -timeout=5m -covermode=atomic \
    -coverprofile="$result_dir/coverage.out" -json -run "$expression" ./... \
    2>"$result_dir/stderr.log" | tee "$result_dir/tests.json"
  codes=("${PIPESTATUS[@]}")
  set -e
  status=0
  if jq -se 'any(.[]; .Action == "skip")' "$result_dir/tests.json" >/dev/null; then status=1; fi
  if (( codes[0] != 0 || codes[1] != 0 )); then status=1; fi
  jq -r 'select(.Test and .Action == "pass") | .Test' "$result_dir/tests.json" | LC_ALL=C sort -u > "$result_dir/passed.txt"
  while IFS= read -r expected; do
    if ! grep -Fxq "$expected" "$result_dir/passed.txt"; then
      echo "Expected test did not pass (failed, skipped or absent): $expected" | tee -a "$result_dir/stderr.log"
      status=1
    fi
  done < "$result_dir/expected.txt"
  if [[ -s "$result_dir/coverage.out" ]]; then
    go tool cover -func="$result_dir/coverage.out" > "$result_dir/coverage.txt" || status=1
  fi
  {
    echo "## $group"
    jq -r --arg g "$group" '.[$g].expected' "$catalog"
    echo
    echo "Exit status: $status"
    echo
    echo '| Test | Result |'
    echo '|---|---|'
    jq -r 'select(.Test and (.Action == "pass" or .Action == "fail" or .Action == "skip")) | "| \(.Test) | \(.Action) |"' "$result_dir/tests.json"
    echo
    echo 'Coverage is for this group over the whole package; it is not global coverage.'
    if [[ -f "$result_dir/coverage.txt" ]]; then tail -n 1 "$result_dir/coverage.txt"; fi
  } > "$result_dir/summary.md"
  cat "$result_dir/summary.md"
  if [[ -n "${GITHUB_STEP_SUMMARY:-}" ]]; then cat "$result_dir/summary.md" >> "$GITHUB_STEP_SUMMARY"; fi
  if (( status != 0 )); then overall=1; fi
done
exit "$overall"
