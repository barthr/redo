#!/bin/bash

# Set variables
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
TAG_WITHOUT_VERSION_PREFIX=$(echo "<tag>" | sed 's/^v//')

if [ "$OS" = "darwin" ]; then
  if [ "$ARCH" = "arm64" ]; then
    BINARY_NAME="redo_${TAG_WITHOUT_VERSION_PREFIX}_Darwin_arm64.tar.gz"
  else
    BINARY_NAME="redo_${TAG_WITHOUT_VERSION_PREFIX}_Darwin_x86_64.tar.gz"
  fi
elif [ "$OS" = "linux" ]; then
  if [ "$ARCH" = "arm64" ]; then
    BINARY_NAME="redo_${TAG_WITHOUT_VERSION_PREFIX}_Linux_arm64.tar.gz"
  else
    BINARY_NAME="redo_${TAG_WITHOUT_VERSION_PREFIX}_Linux_x86_64.tar.gz"
  fi
else
  echo "Unsupported OS/ARCH: $OS/$ARCH"
  exit 1
fi

RELEASE_URL="https://github.com/barthr/redo/releases/download/<tag>/${BINARY_NAME}"
INSTALL_PATH="/usr/local/bin"

# Download the binary
curl -L -o redo.tar.gz $RELEASE_URL

# Extract the binary and move it to the installation path
echo "Extracting and installing redo binary to $INSTALL_PATH..."
tar -zxvf redo.tar.gz redo
sudo mv redo "$INSTALL_PATH"

# Make the binary executable
sudo chmod +x $INSTALL_PATH/redo

#Cleanup
rm redo.tar.gz

# Print success message
echo "Successfully installed redo to $INSTALL_PATH"
