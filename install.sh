#!/bin/bash

set -e

REPO="https://github.com/xenomech/portal.git"
INSTALL_DIR="$HOME/.local/bin"
BINARY_NAME="portal"

install() {
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
    cp dist/portal "$INSTALL_DIR/$BINARY_NAME"
    chmod +x "$INSTALL_DIR/$BINARY_NAME"

    # Cleanup
    cd - > /dev/null
    rm -rf "$TMP_DIR"

    echo "=>> Portal installed successfully!"
    echo ""
    echo "Make sure $INSTALL_DIR is in your PATH:"
    echo "  export PATH=\"\$HOME/.local/bin:\$PATH\""
    echo ""
    echo "Then run: portal --help"
}

uninstall() {
    echo "Uninstalling Portal..."

    if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
        rm "$INSTALL_DIR/$BINARY_NAME"
        echo "=>> Portal uninstalled successfully!"
        echo "Removed: $INSTALL_DIR/$BINARY_NAME"
    else
        echo "Portal is not installed at $INSTALL_DIR/$BINARY_NAME"
        exit 1
    fi
}

# Main
case "${1:-install}" in
    install|-i|--install)
        install
        ;;
    uninstall|-u|--uninstall)
        uninstall
        ;;
    *)
        echo "Usage: $0 {install|uninstall}"
        echo ""
        echo "Commands:"
        echo "  install      Install Portal (default)"
        echo "  uninstall    Remove Portal"
        exit 1
        ;;
esac
