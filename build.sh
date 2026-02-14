#!/bin/bash

set -e

PLUGIN_NAME="selectellinter"

echo "Building plugin: $PLUGIN_NAME..."

mkdir -p ./bin

CGO_ENABLED=1 go build -buildmode=plugin -trimpath -o ./bin/$PLUGIN_NAME.so ./plugin/plugin.go

echo "Plugin built successfully: ./bin/$PLUGIN_NAME.so"
echo "Go version used: $(go version)"