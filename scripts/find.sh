#!/usr/bin/env bash
SOURCE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
"$SOURCE_DIR/api.sh" "PageService.Show" "{\"id\": \"$1\"}" | jq
