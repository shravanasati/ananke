# html2md

[![Go Reference](https://pkg.go.dev/badge/github.com/shravanasati/ananke/html2md.svg)
](https://pkg.go.dev/github.com/shravanasati/ananke/html2md)

html2md is a Go library which provides a simple interface to convert HTML code to markdown. It aims to support the [CommonMark specification](https://spec.commonmark.org/0.31.2/) as well as GitHub-flavored markdown.

It powers [ananke](https://github.com/shravanasati/ananke), a CLI tool which converts HTML to Markdown.

### Usage

```go
package main

import (
	"fmt"
	"github.com/shravanasati/ananke/html2md"
)

func main() {
	converter := html2md.NewConverter()
	// pass options

	input := "<ul></ul>"
	output := converter.ConvertString()
}
```

The converter is **NOT thread-safe** and must be used **exactly once** for each input. 

The caller should ensure that the input is UTF-8 encoded.