#!/bin/bash

set -e

PLUGIN_NAME="selectellinter"
PLUGIN_DIR="$HOME/.golangci-lint-plugins"
PLUGIN_PATH="$PLUGIN_DIR/$PLUGIN_NAME.so"

echo "Building plugin..."
bash build.sh

echo ""
echo "Installing plugin..."
mkdir -p "$PLUGIN_DIR"
cp "./bin/$PLUGIN_NAME.so" "$PLUGIN_PATH"

if [ -f "$PLUGIN_PATH" ]; then
    echo "Plugin installed at: $PLUGIN_PATH"
    echo ""
    echo "Add to .golangci.yml:"
    echo ""
    echo "linters:"
    echo "  custom:"
    echo "    selectellinter:"
    echo "      path: ~/.golangci-lint-plugins/selectellinter.so"
    echo "      description: Selectel linter for log messages"
    echo "  enable:"
    echo "    - selectellinter"
else
    echo "Plugin installation failed!"
    exit 1
fi
