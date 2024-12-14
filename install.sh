set -e

USER="devnazir"
REPO="ip-address-lang"
API_URL="https://api.github.com/repos/${USER}/${REPO}/releases/latest"
ARCHIVE_URL="https://github.com/${USER}/$REPO/archive/refs/tags"

echo "Downloading..."

# Get the latest release from GitHub API
latest_tag=$(curl -s $API_URL |  grep '"tag_name"' | cut -d'"' -f4 )

if [ -z "$latest_tag" ]; then
  echo "Failed to get latest release"
  exit 1
fi

# Download the latest release

ARCHIVE_URL="${ARCHIVE_URL}/${latest_tag}.zip"

curl -L -o /tmp/ip-address-lang.zip $ARCHIVE_URL

# Unzip the archive
unzip -o /tmp/ip-address-lang.zip -d /tmp

# Copy the binary to /usr/local/bin
sudo cp /tmp/ip-address-lang-${latest_tag}/ip-address-lang /usr/local/bin

# Clean up
rm -rf /tmp/ip-address-lang.zip /tmp/ip-address-lang-${latest_tag}

echo "Installed ip-address-lang ${latest_tag}"
