#!/bin/bash

set -e

USER="devnazir"
REPO="ip-address-language"
API_URL="https://api.github.com/repos/${USER}/${REPO}/releases/latest"
ARCH=$(uname -m)
OS=$(uname -s)
BIN_NAME="ipl"

echo "Downloading..."

# Get the latest release from GitHub API
latest_tag=$(curl -s $API_URL | grep 'tag_name' | cut -d\" -f4)
tag_without_v=$(echo $latest_tag | cut -c 2-) # Remove the leading 'v' from the tag

echo "Latest release: $latest_tag"

if [ -z "$latest_tag" ]; then
  echo "Failed to get latest release"
  exit 1
fi

# Construct the correct file name based on OS and architecture
case $OS in
  Darwin)
    case $ARCH in
      x86_64)
        ARCH_BIN="${REPO}_${tag_without_v}_darwin_amd64.tar.gz"
        ;;
      arm64)
        ARCH_BIN="${REPO}_${tag_without_v}_darwin_arm64.tar.gz"
        ;;
      *)
        echo "Unsupported architecture: $ARCH on macOS"
        exit 1
        ;;
    esac
    ;;
  Linux)
    case $ARCH in
      x86_64)
        ARCH_BIN="${REPO}_${tag_without_v}_linux_amd64.tar.gz"
        ;;
      i386|i686)
        ARCH_BIN="${REPO}_${tag_without_v}_linux_386.tar.gz"
        ;;
      aarch64)
        ARCH_BIN="${REPO}_${tag_without_v}_linux_arm64.tar.gz"
        ;;
      armv7l)
        ARCH_BIN="${REPO}_${tag_without_v}_linux_armv7.tar.gz"
        ;;
      *)
        echo "Unsupported architecture: $ARCH on Linux"
        exit 1
        ;;
    esac
    ;;
  CYGWIN*|MINGW*|MSYS*)
    # For Windows (using Cygwin, MinGW, or MSYS)
    case $ARCH in
      x86_64)
        ARCH_BIN="${REPO}_${tag_without_v}_windows_amd64.tar.gz"
        ;;
      i386|i686)
        ARCH_BIN="${REPO}_${tag_without_v}_win_386.tar.gz"
        ;;
      *)
        echo "Unsupported architecture: $ARCH on Windows"
        exit 1
        ;;
    esac
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

# Construct the URL to download the archive
ARCHIVE_URL="https://github.com/${USER}/${REPO}/releases/download/${latest_tag}/${ARCH_BIN}"

echo "Downloading $ARCHIVE_URL"

# Download the latest release tar.gz archive
curl -L -o /tmp/$REPO.tar.gz $ARCHIVE_URL

# Extract the archive
tar -xzvf /tmp/$REPO.tar.gz -C /tmp

# Copy the binary to /usr/local/bin (Linux/macOS) or equivalent on Windows
if [[ "$OS" == "Darwin" || "$OS" == "Linux" ]]; then
  sudo cp /tmp/$REPO /usr/local/bin/$BIN_NAME
elif [[ "$OS" == "CYGWIN"* || "$OS" == "MINGW"* || "$OS" == "MSYS"* ]]; then
  cp /tmp/$REPO.exe /c/Program\ Files/$REPO/$BIN_NAME.exe
else
  echo "Unsupported OS: $OS"
  exit 1
fi

# Clean up
rm -rf /tmp/$REPO.tar.gz /tmp/${REPO}-${latest_tag}

echo "Installed ip-address-lang ${latest_tag}"
