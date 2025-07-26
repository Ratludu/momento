#!/bin/bash
# install.sh

set -e

# Default values
INSTALL_DIR="/usr/local/bin"
REPO="ratludu/momento"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) echo -e "${RED}Unsupported architecture: $ARCH${NC}" && exit 1 ;;
esac

case $OS in
    linux|darwin) ;;
    *) echo -e "${RED}Unsupported OS: $OS${NC}" && exit 1 ;;
esac

echo -e "${GREEN}Installing momento for $OS-$ARCH${NC}"

# Get latest release
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
    echo -e "${RED}Failed to get latest release${NC}"
    exit 1
fi

echo -e "${YELLOW}Latest version: $LATEST_RELEASE${NC}"

# Download URL
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/myapp_${OS^}_$ARCH.tar.gz"

# Create temp directory
TEMP_DIR=$(mktemp -d)
cd "$TEMP_DIR"

# Download and extract
echo -e "${YELLOW}Downloading from $DOWNLOAD_URL${NC}"
curl -sL "$DOWNLOAD_URL" | tar xz

# Install
if [ -w "$INSTALL_DIR" ]; then
    mv myapp "$INSTALL_DIR/"
else
    echo -e "${YELLOW}Installing to $INSTALL_DIR requires sudo${NC}"
    sudo mv myapp "$INSTALL_DIR/"
fi

# Cleanup
cd - > /dev/null
rm -rf "$TEMP_DIR"

echo -e "${GREEN}myapp installed successfully!${NC}"
echo -e "${GREEN}Run 'momento --help' to get started${NC}"
