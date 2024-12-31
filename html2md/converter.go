package html2md

import (
	"strings"

	"golang.org/x/net/html"
)

type Converter struct {
	// todo add options
	listStack *stack[*listEntry]
	processed map[string]bool
}

func NewConverter() *Converter {
	stack := newStack[*listEntry]()
	return &Converter{listStack: stack, processed: map[string]bool{}}
}

// performs a linear search for the given attribute in a html node
func findAttribute(node *html.Node, key string) string {
	for _, attr := range node.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}

func (c *Converter) htmlNodeToMarkdownElement(node *html.Node) MarkdownElement {
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

	case "ul":
		fingerprint := generateFingerprint(node)
		if _, ok := c.processed[fingerprint]; !ok {
			// this tag has not been processed before
			c.listStack.push(newListEntry(UnorderedList))
			c.processed[fingerprint] = true
			depth := c.listStack.size() - 1
			return NewListTag(UnorderedList, depth)
		}
		return NewUnknownTag(node.Data)
	case "ol":
		fingerprint := generateFingerprint(node)
		if _, ok := c.processed[fingerprint]; !ok {
			// this tag has not been processed before
			c.listStack.push(newListEntry(OrderedList))
			c.processed[fingerprint] = true
			depth := c.listStack.size() - 1
			return NewListTag(OrderedList, depth)
		}
		return NewUnknownTag(node.Data)

	case "li":
		topmost, err := c.listStack.top()
		if err != nil {
			panic("li tag without parent ol/ul tag")
		}
		depth := c.listStack.size() - 1
		var number int
		if topmost.type_ == UnorderedList {
			number = 0
		} else {
			number = topmost.counter.next()
		}
		return NewListItemTag(depth, topmost.type_, number)

	default:
		return NewUnknownTag(node.Data)
	}
}

func (c *Converter) convertNode(node *html.Node, output *outputWriter) {
	switch node.Type {
	case html.TextNode:
		// Write the text content, escaping special Markdown characters
		output.WriteString((escapeMarkdown(node.Data)))
	case html.ElementNode:
		// Determine the Markdown type
		markdownElem := c.htmlNodeToMarkdownElement(node)

		// Write opening Markdown syntax
		output.WriteString(markdownElem.StartCode())

		// Recursively process child nodes
		for child := range node.ChildNodes() {
			c.convertNode(child, output)
		}

		// Write closing Markdown syntax
		endCode := markdownElem.EndCode()
		output.WriteString(endCode)

		if markdownElem.Type() == ListItem && node.NextSibling == nil {
			// last li tag in a list
			_, err := c.listStack.pop()
			if err != nil {
				// stack underflow
				panic("no items in listStack to pop for the last li tag")
			}
			output.WriteString("\n")  // write an extra newline when the list ends
		}
	}
}

func (c *Converter) ConvertString(input string) (string, error) {
	output := newOutputWriter()

	// Parse the HTML input into a document tree
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return "", err
	}

	// Start recursive conversion from the root node's children
	for node := doc.FirstChild; node != nil; node = node.NextSibling {
		c.convertNode(node, output)
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
