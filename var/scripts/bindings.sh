#!/usr/bin/env bash
set -euo pipefail

BUILD_DIR=var/builds
OUT_DIR=internal/common/service/client/blockchain
PKG_NAME=blockchain  

for abi in "$BUILD_DIR"/*.abi; do
  name=$(basename "$abi" .abi)
  bin="$BUILD_DIR/$name.bin"
  out="$OUT_DIR/${name,,}.go"   # имя файла в lowercase
  echo "Generating Go binding for $name → $out"
  abigen \
    --bin="$bin" \
    --abi="$abi" \
    --pkg="$PKG_NAME" \
    --out="$out"
done
