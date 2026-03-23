//go:build arm64 && darwin

package gotypst

import _ "embed"

//go:embed bin/typst-arm64-darwin
var typst_binary []byte
