//go:build arm && linux

package gotypst

import _ "embed"

//go:embed bin/typst-arm-linux
var typst_binary []byte
