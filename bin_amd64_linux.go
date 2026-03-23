//go:build amd64 && linux

package gotypst

import _ "embed"

//go:embed bin/typst-amd64-linux
var typst_binary []byte
