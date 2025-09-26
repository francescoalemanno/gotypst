rm -R assets assets.zip
mkdir assets
ver="v0.13.1"
wget "https://github.com/typst/typst/releases/download/$ver/typst-aarch64-apple-darwin.tar.xz"
wget "https://github.com/typst/typst/releases/download/$ver/typst-aarch64-unknown-linux-musl.tar.xz"
wget "https://github.com/typst/typst/releases/download/$ver/typst-armv7-unknown-linux-musleabi.tar.xz"
wget "https://github.com/typst/typst/releases/download/$ver/typst-x86_64-apple-darwin.tar.xz"
wget "https://github.com/typst/typst/releases/download/$ver/typst-x86_64-unknown-linux-musl.tar.xz"

tar -xf typst-aarch64-apple-darwin.tar.xz --wildcards --no-anchored 'typst' --strip-components=1
mv typst assets/arm64-darwin

tar -xf typst-aarch64-unknown-linux-musl.tar.xz --wildcards --no-anchored 'typst' --strip-components=1
mv typst assets/arm64-linux

tar -xf typst-armv7-unknown-linux-musleabi.tar.xz --wildcards --no-anchored 'typst' --strip-components=1
mv typst assets/arm-linux

tar -xf typst-x86_64-apple-darwin.tar.xz --wildcards --no-anchored 'typst' --strip-components=1
mv typst assets/amd64-darwin

tar -xf typst-x86_64-unknown-linux-musl.tar.xz --wildcards --no-anchored 'typst' --strip-components=1
mv typst assets/amd64-linux


tar -xf typst-x86_64-unknown-linux-musl.tar.xz --wildcards --no-anchored 'LICENSE' --strip-components=1
mv LICENSE assets/LICENSE

tar -xf typst-x86_64-unknown-linux-musl.tar.xz --wildcards --no-anchored 'NOTICE' --strip-components=1
mv NOTICE assets/NOTICE

tar -xf typst-x86_64-unknown-linux-musl.tar.xz --wildcards --no-anchored 'README.md' --strip-components=1
mv README.md assets/README.md

zip -r assets.zip assets

rm -R assets
