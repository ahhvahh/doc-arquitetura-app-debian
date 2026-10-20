#!/usr/bin/env bash
# Structural documentation validation for GitHub Actions. No network calls or secrets.
set -euo pipefail

root="${CHARTER_DOCS_ROOT:-docs/charter}"
failures=0
documents=0
links=0
badge_paths=0

fail() {
  local file="$1" message="$2"
  printf '::error file=%s::%s\n' "$file" "$message"
  printf 'FAIL %s: %s\n' "$file" "$message"
  failures=$((failures + 1))
}

if [[ ! -d "$root" ]]; then
  fail "$root" 'Charter directory is missing'
  exit 1
fi

while IFS= read -r -d '' file; do
  documents=$((documents + 1))
  title="$(sed -n '1p' "$file")"
  if [[ "$title" != '# '* ]]; then
    fail "$file" 'First line must be an H1 heading'
  fi

  # Exactly three metadata badges on line 3, in their required order.
  badges="$(sed -n '3p' "$file")"
  count="$(printf '%s\n' "$badges" | grep -oE '!\[[^]]*\]\([^)]*badge-(type|status|version)-[^)]*\.svg\)' | wc -l || true)"
  count="${count//[[:space:]]/}"
  if [[ "$count" != 3 ]]; then
    fail "$file" "Expected three badges on line 3; found ${count:-0}"
  else
    # Check both order and labels. Do not assume every GUIDE README uses one specific badge type.
    if [[ ! "$badges" =~ badge-type-[^[:space:]]+\.svg.*badge-status-[^[:space:]]+\.svg.*badge-version-[^[:space:]]+\.svg ]]; then
      fail "$file" 'Badge order must be Type, Status, Version'
    fi
    if [[ ! "$badges" =~ \!\[Tipo: ]] || [[ ! "$badges" =~ \!\[Status: ]] || [[ ! "$badges" =~ \!\[Versão: ]]; then
      fail "$file" 'Badge alternate labels must identify Type, Status, Version'
    fi
  fi

  # Validate local link and image destinations, including SVG metadata assets.
  # Links to external destinations and anchors in the same document are not filesystem paths.
  while IFS= read -r url; do
    [[ -n "$url" ]] || continue
    case "$url" in
      http://*|https://*|mailto:*|tel:*|data:*|\#*) continue ;;
    esac
    target="${url%%#*}"
    target="${target%%\?*}"
    [[ -n "$target" ]] || continue
    links=$((links + 1))
    if [[ "$target" == *'.svg' ]]; then
      badge_paths=$((badge_paths + 1))
    fi
    if [[ ! -f "$(dirname "$file")/$target" ]] && [[ ! -d "$(dirname "$file")/$target" ]]; then
      fail "$file" "Missing relative target: $target"
    fi
  done < <(grep -oE '\]\([^)]*\)' "$file" | sed -e 's/^](//' -e 's/)$//' || true)
done < <(find "$root" -type f -name '*.md' -print0 | sort -z)

if (( documents == 0 )); then
  fail "$root" 'No Charter Markdown documents found'
fi
printf 'Charter audit: %d documents, %d local references (%d SVG), %d failure(s)\n' "$documents" "$links" "$badge_paths" "$failures"
if [[ -n "${GITHUB_STEP_SUMMARY:-}" ]]; then
  printf '### Charter audit\n\n- Documents: %d\n- Local references: %d\n- SVG references: %d\n- Failures: %d\n' "$documents" "$links" "$badge_paths" "$failures" >> "$GITHUB_STEP_SUMMARY"
fi
if (( failures > 0 )); then
  exit 1
fi
