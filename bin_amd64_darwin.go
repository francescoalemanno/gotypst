//go:build amd64 && darwin

package gotypst

import _ "embed"

//go:embed bin/typst-amd64-darwin
var typst_binary []byte
