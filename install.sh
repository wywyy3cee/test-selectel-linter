#!/bin/bash

set -e

PLUGIN_NAME="selectellinter"
PLUGIN_DIR="$HOME/.golangci-lint-plugins"
PLUGIN_PATH="$PLUGIN_DIR/$PLUGIN_NAME.so"

bash build.sh

if [ -f "$PLUGIN_PATH" ]; then
    echo "Plugin installed at: $PLUGIN_PATH"
    echo "Add to .golangci.yml:"
    echo ""
    echo "  plugins:"
    echo "    - module: github.com/wywyy3cee/test-selectel-linter"
    echo "      import: plugin"
    echo "      plugin_name: selectellinter"
    echo ""
    echo "And enable in linters:"
    echo "  enable:"
    echo "    - selectellinter"
else
    echo "Plugin build failed!"
    exit 1
fi
