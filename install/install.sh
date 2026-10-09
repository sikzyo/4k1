#!/bin/bash

set -e

REPO="sikzyo/4k1"
BIN="4k1"

OS="$(uname -s | tr '[:upper:]' '[:lower:]')"

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

curl -fsSL "https://github.com/$REPO/releases/latest/download/$BIN" -o "$TMP/$BIN" && chmod +x "$TMP/$BIN"

"$TMP/$BIN"
