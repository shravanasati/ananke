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

func markdownElementToCodeStart(elem MarkdownElementType) string {
	switch elem {
	case H1:
		return "# "
	case H2:
		return "## "
	case H3:
		return "### "
	case H4:
		return "#### "
	case H5:
		return "##### "
	case H6:
		return "###### "
	case Bold:
		return "**"
	case Italic:
		return "*"
	default:
		return ""
	}
}

func markdownElementToCodeEnd(elem MarkdownElementType) string {
	// For elements that need closing syntax
	switch elem {
	case Bold:
		return "**"
	case Italic:
		return "*"
	case H1, H2, H3, H4, H5, H6:
		return "\n"
	case Paragraph:
		return "\n\n"
	default:
		return ""
	}
}

func htmlNodeToMarkdownType(node *html.Node) MarkdownElementType {
	switch node.Data {
	case "h1":
		return H1
	case "h2":
		return H2
	case "h3":
		return H3
	case "h4":
		return H4
	case "h5":
		return H5
	case "h6":
		return H6
	case "b", "strong":
		return Bold
	case "i", "em":
		return Italic
	case "p":
		return Paragraph
	default:
		return Unknown
	}
}

func (c *Converter) convertNode(node *html.Node, output *strings.Builder) {
	switch node.Type {
	case html.TextNode:
		// Write the text content, escaping special Markdown characters
		output.WriteString((escapeMarkdown(node.Data)))
	case html.ElementNode:
		// Determine the Markdown type
		markdownElem := htmlNodeToMarkdownType(node)

		// Write opening Markdown syntax
		output.WriteString(markdownElementToCodeStart(markdownElem))

		// Recursively process child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			c.convertNode(child, output)
		}

		// Write closing Markdown syntax
		output.WriteString(markdownElementToCodeEnd(markdownElem))
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
	// const specialCharacters = `\*_{}[]()#+-.!`
	replacer := strings.NewReplacer(
		`\`, `\\`, `*`, `\*`, `_`, `\_`, `{`, `\{`, `}`, `\}`,
		`[`, `\[`, `]`, `\]`, `(`, `\(`, `)`, `\)`, `#`, `\#`,
		`+`, `\+`, `-`, `\-`, `!`, `\!`,
	)
	return replacer.Replace(text)
}
