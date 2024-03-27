#!/usr/bin/env bash
SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
BASE_DIR=$(cd "$SCRIPT_DIR/.." && pwd)

if [[ ! -e "$BASE_DIR/.bingo/variables.env" ]]; then
  echo "Plese run 'task deps' first"
  exit 1
fi

source "$BASE_DIR/.bingo/variables.env"

SERVER="$BASE_DIR/internal/server"
TEMPLATES="$BASE_DIR/internal/templates"
DEFINITIONS="$BASE_DIR/internal/definitions"

# Generate the code
mkdir -p "$SERVER"
$OTO -template "$TEMPLATES/echo.go.plush" \
  -out "$SERVER/server.gen.go" \
  -ignore Ignorer \
  -pkg server \
  "$DEFINITIONS"
$OTO -template "$TEMPLATES/models.go.plush" \
  -out "$SERVER/models.gen.go" \
  -ignore Ignorer \
  -pkg server \
  "$DEFINITIONS"

mkdir -p "$BASE_DIR/client"
$OTO -template "$TEMPLATES/client.go.plush" \
  -out "$BASE_DIR/client/client.gen.go" \
  -ignore Ignorer \
  -pkg client \
  "$DEFINITIONS"

goimports -w -local github.com/dashotv client/*.go "$SERVER"/*.go
