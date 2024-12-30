package html2md

import (
	"strings"

	"golang.org/x/net/html"
)

type Converter struct {
	// todo add options
}

func NewConverter() *Converter {
	return &Converter{}
}

func findAttribute(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}

func htmlNodeToMarkdownElement(node *html.Node) MarkdownElement {
	switch node.Data {
	case "h1":
		return NewH1Tag()
	case "h2":
		return NewH2Tag()
	case "h3":
		return NewH3Tag()
	case "h4":
		return NewH4Tag()
	case "h5":
		return NewH5Tag()
	case "h6":
		return NewH6Tag()
	case "b", "strong":
		return NewBoldTag()
	case "i", "em":
		return NewItalicTag()
	case "p":
		return NewParagraphTag()
	case "a":
		href := findAttribute(node, "href")
		return NewAnchorTag(href)
	case "img":
		src := findAttribute(node, "src")
		alt := findAttribute(node, "alt")
		if alt == "" {
			alt = "image"
		}
		return NewImageTag(src, alt)
	default:
		return NewUnknownTag(node.Data)
	}
}

func (c *Converter) convertNode(node *html.Node, output *strings.Builder) {
	switch node.Type {
	case html.TextNode:
		// Write the text content, escaping special Markdown characters
		output.WriteString((escapeMarkdown(node.Data)))
	case html.ElementNode:
		// Determine the Markdown type
		markdownElem := htmlNodeToMarkdownElement(node)

		// Write opening Markdown syntax
		output.WriteString(markdownElem.StartCode())

		// Recursively process child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			c.convertNode(child, output)
		}

		// Write closing Markdown syntax
		endCode := markdownElem.EndCode()
		if isBlockLevelElem(htmlNodeToMarkdownElement(node.Parent)) {
			// this is to prevent extra newlines
			endCode = strings.TrimSuffix(endCode, "\n")
		}
		output.WriteString(endCode)
	}
}

func (c *Converter) ConvertString(input string) (string, error) {
	var output strings.Builder

	// Parse the HTML input into a document tree
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return "", err
	}

	// Start recursive conversion from the root node's children
	for node := doc.FirstChild; node != nil; node = node.NextSibling {
		c.convertNode(node, &output)
	}

	return output.String(), nil
}

func escapeMarkdown(text string) string {
	// Escape special Markdown characters
	replacer := strings.NewReplacer(
		`\`, `\\`, `*`, `\*`, `_`, `\_`, `{`, `\{`, `}`, `\}`,
		`[`, `\[`, `]`, `\]`, `(`, `\(`, `)`, `\)`, `#`, `\#`,
		`+`, `\+`, `-`, `\-`, `!`, `\!`,
	)
	return replacer.Replace(text)
}
