#!/bin/bash

set -e

REPO="https://github.com/xenomech/portal.git"
INSTALL_DIR="$HOME/.local/bin"

echo "Installing Portal..."

# Clone
TMP_DIR=$(mktemp -d)
git clone --depth 1 "$REPO" "$TMP_DIR"
cd "$TMP_DIR"

# Build
make build

# Install
echo "Installing to $INSTALL_DIR"
mkdir -p "$INSTALL_DIR"
cp dist/portal "$INSTALL_DIR/portal"
chmod +x "$INSTALL_DIR/portal"

# Cleanup
cd - > /dev/null
rm -rf "$TMP_DIR"

echo "Portal installed successfully!"
echo ""
echo "Make sure $INSTALL_DIR is in your PATH:"
echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
echo ""
echo "Then run: portal --help"
