#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GLOBAL_AGENTS="$HOME/.agents/skills"
TARGET_DIR="$HOME/src"

# 1. Install global agents
mkdir -p "$HOME/.agents/skills"
for dir in "$SCRIPT_DIR"/skills/*/; do
  [ -d "$dir" ] || continue
  cp -r "$dir" "$HOME/.agents/skills/"
done

# 2. Recursively symlink into all git repos under TARGET_DIR
find "$TARGET_DIR" -type d -name .git -prune | while read -r gitdir; do
  repo="${gitdir%/.git}"
  if [ ! -e "$repo/.agents" ]; then
    ln -s "$GLOBAL_AGENTS" "$repo/.agents"
  fi
done

# 3. Install post-checkout git hook
mkdir -p "$HOME/.git-hooks"
cat > "$HOME/.git-hooks/post-checkout" << 'EOF'
#!/bin/sh
# Only run on branch checkout, not file checkout
[ "$3" = "0" ] && exit 0

REPO_ROOT=$(git rev-parse --show-toplevel)
AGENTS_LINK="$REPO_ROOT/.agents"
GLOBAL_AGENTS="$HOME/.agents/skills"

if [ ! -e "$AGENTS_LINK" ]; then
  ln -s "$GLOBAL_AGENTS" "$AGENTS_LINK"
  echo ".agents symlinked from global config"
fi
EOF
chmod +x "$HOME/.git-hooks/post-checkout"
