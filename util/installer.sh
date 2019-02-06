#!/bin/sh
# script borrowed from https://github.com/getantibody/installer
set -e
DOWNLOAD_URL="https://github.com/talal/bonclay/releases/download"

last_version() {
  curl -s https://raw.githubusercontent.com/talal/homebrew-tap/master/Formula/bonclay.rb |
    grep version |
    cut -f2 -d'"'
}

download() {
	version="$(last_version)" || true
  test -z "$version" && {
    echo "Unable to get bonclay version."
    exit 1
  }
  echo "Downloading bonclay v$version for $(uname -s)..."
  rm -f /tmp/bonclay /tmp/bonclay.tar.gz
  curl -s -L -o /tmp/bonclay.tar.gz \
    "$DOWNLOAD_URL/v$version/bonclay-$version-$(uname -s)_amd64.tar.gz"
}

extract() {
  tar -xf /tmp/bonclay.tar.gz -C /tmp
}

main() {
	download
	extract || true
	sudo mv -f /tmp/bonclay /usr/local/bin/bonclay
	rm -f /tmp/bonclay.tar.gz
	echo "bonclay v$version installed in $(which bonclay)"
}

main
