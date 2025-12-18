# html2md

[![Go Reference](https://pkg.go.dev/badge/github.com/shravanasati/ananke/html2md.svg)
](https://pkg.go.dev/github.com/shravanasati/ananke/html2md)

html2md is a Go library which provides a simple interface to convert HTML code to markdown. It aims to support the [CommonMark specification](https://spec.commonmark.org/0.31.2/).

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

	input := `
	<p>This is a <strong>sample</strong> paragraph with <em>emphasis</em>.</p>
	<ul>
		<li>First item</li>
		<li>Second item</li>
	</ul>
	`
	output, err := converter.ConvertString(input)
	if err != nil {
		// err is non-nil when the html is malformed
		fmt.Println("failed to convert html to markdown:", err)
	}

	fmt.Println("converted markdown:\n", output)
}
```

The converter is **NOT thread-safe** and must be used **exactly once** for each input. 

The caller should ensure that the input is UTF-8 encoded.
