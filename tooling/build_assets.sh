#!/bin/bash
set -e
cd "$(dirname "$0")/.."
rm -rf bin
mkdir bin
ver="v0.14.2"
curl -LO "https://github.com/typst/typst/releases/download/$ver/typst-aarch64-apple-darwin.tar.xz"
curl -LO "https://github.com/typst/typst/releases/download/$ver/typst-aarch64-unknown-linux-musl.tar.xz"
curl -LO "https://github.com/typst/typst/releases/download/$ver/typst-armv7-unknown-linux-musleabi.tar.xz"
curl -LO "https://github.com/typst/typst/releases/download/$ver/typst-x86_64-apple-darwin.tar.xz"
curl -LO "https://github.com/typst/typst/releases/download/$ver/typst-x86_64-unknown-linux-musl.tar.xz"

tar -xf typst-aarch64-apple-darwin.tar.xz --strip-components=1 "*/typst"
mv typst bin/typst-arm64-darwin

tar -xf typst-aarch64-unknown-linux-musl.tar.xz --strip-components=1 "*/typst"
mv typst bin/typst-arm64-linux

tar -xf typst-armv7-unknown-linux-musleabi.tar.xz --strip-components=1 "*/typst"
mv typst bin/typst-arm-linux

tar -xf typst-x86_64-apple-darwin.tar.xz --strip-components=1 "*/typst"
mv typst bin/typst-amd64-darwin

tar -xf typst-x86_64-unknown-linux-musl.tar.xz --strip-components=1 "*/typst"
mv typst bin/typst-amd64-linux

tar -xf typst-x86_64-unknown-linux-musl.tar.xz --strip-components=1 "*/LICENSE"
mv LICENSE bin/LICENSE

tar -xf typst-x86_64-unknown-linux-musl.tar.xz --strip-components=1 "*/NOTICE"
mv NOTICE bin/NOTICE

rm -f typst-*.tar.xz
