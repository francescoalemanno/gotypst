# gotypst
[![Go Report Card](https://goreportcard.com/badge/github.com/francescoalemanno/gotypst)](https://goreportcard.com/report/github.com/francescoalemanno/gotypst)
[![API Reference](https://img.shields.io/badge/docs-API_Reference-blue)](https://pkg.go.dev/github.com/francescoalemanno/gotypst)

gotypst is a Go package that compiles Typst code into a PDF. It provides an easy-to-use function to pass Typst markup as bytes and receive the compiled PDF as bytes. This can be used to integrate Typst into Go projects, automate PDF generation, or add Typst support to your web services.
Features

- Convert Typst code into PDF on the fly
- Flexible options for customization
- Simple, minimal interface

## example
```go
package main

import (
	"fmt"
	"github.com/francescoalemanno/gotypst"
)

func main() {
    bts, err := gotypst.PDF([]byte("= hello"))
    if err!=nil {
        return
    }
    fmt.Println(bts)
}
```
