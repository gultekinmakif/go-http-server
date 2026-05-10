#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

if ! command -v air >/dev/null; then
  echo "installing air..."
  go install github.com/air-verse/air@latest
fi

air
