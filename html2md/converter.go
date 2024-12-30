package html2md

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Converter struct {
	// todo add options
}

func NewConverter() *Converter {
	return &Converter{}
}

func markdownElementToCodeStart(elem MarkdownElement) string {
	switch elem.Type() {
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
	case Anchor:
		return "["
	default:
		return ""
	}
}

func markdownElementToCodeEnd(elem MarkdownElement) string {
	// For elements that need closing syntax
	switch elem.Type() {
	case Bold:
		return "**"
	case Italic:
		return "*"
	case H1, H2, H3, H4, H5, H6:
		return "\n"
	case Paragraph:
		return "\n\n"
	case Anchor:
		return fmt.Sprintf("](%v)", elem.(*AnchorTag).href)
	default:
		return ""
	}
}

func findHrefAttribute(node *html.Node) string {
	if node.Type != html.ElementNode && node.Data != "a" {
		panic("findHrefAttribute function called non-anchor tag")
	}
	for _, attr := range node.Attr {
		if attr.Key == "href" {
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
		href := findHrefAttribute(node)
		return NewAnchorTag(href)
	default:
		return NewUnknownTag()
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
	replacer := strings.NewReplacer(
		`\`, `\\`, `*`, `\*`, `_`, `\_`, `{`, `\{`, `}`, `\}`,
		`[`, `\[`, `]`, `\]`, `(`, `\(`, `)`, `\)`, `#`, `\#`,
		`+`, `\+`, `-`, `\-`, `!`, `\!`,
	)
	return replacer.Replace(text)
}
