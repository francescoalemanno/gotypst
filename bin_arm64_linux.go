//go:build arm64 && linux

package gotypst

import _ "embed"

//go:embed bin/typst-arm64-linux
var typst_binary []byte
