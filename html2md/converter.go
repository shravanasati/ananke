package html2md

import (
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

var languageRegex = regexp.MustCompile(`language-(\w+)`)
var ignoreTags = []string{"script", "style"}

type Converter struct {
	// todo add options
	listStack    *stack[*listEntry]
	processed    map[string]bool
	preTagCount  int
	output       *outputWriter
	codeTagCount int
}

func NewConverter() *Converter {
	stack := newStack[*listEntry]()
	return &Converter{
		listStack:    stack,
		processed:    map[string]bool{},
		preTagCount:  0,
		codeTagCount: 0,
		output:       newOutputWriter(),
	}
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

func findCodeLanguage(node *html.Node) string {
	if node.Type != html.ElementNode && node.Data != "code" {
		panic("attempt to find language in a non-code tag")
	}

	classList := findAttribute(node, "class")
	matches := languageRegex.FindStringSubmatch(classList)
	if len(matches) < 2 {
		return ""
	}

	return matches[1]
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
		title := findAttribute(node, "title")
		c.output.insideAnchor = true
		return NewAnchorTag(href, title)

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
			c.listStack.push(newUnorderedListEntry())
			c.processed[fingerprint] = true
			depth := c.listStack.size() - 1
			return NewListTag(UnorderedList, depth)
		}
		return NewUnknownTag(node.Data)
	case "ol":
		fingerprint := generateFingerprint(node)
		if _, ok := c.processed[fingerprint]; !ok {
			// this tag has not been processed before
			type_ := findAttribute(node, "type")
			cType, case_ := getOrderedListParams(type_)

			start := findAttribute(node, "start")
			if start == "" {
				start = "1"
			}
			startNum, err := strconv.Atoi(start)
			if err != nil {
				startNum = 1
			}
			c.listStack.push(newOrderedListEntry(startNum, cType, case_))
			c.processed[fingerprint] = true
			depth := c.listStack.size() - 1
			return NewListTag(OrderedList, depth)
		}
		return NewUnknownTag(node.Data)

	case "li":
		topmost, err := c.listStack.top()
		if err != nil {
			// if list items without a parent ol/ul tags are found,
			// insert an ul tag in the stack
			topmost = newUnorderedListEntry()
			c.listStack.push(topmost)
			// panic("li tag without parent ol/ul tag")
		}
		depth := c.listStack.size() - 1
		var number string
		if topmost.type_ == UnorderedList {
			number = "0"
		} else {
			number = topmost.counter.next()
		}
		return NewListItemTag(depth, topmost.type_, number)

	case "blockquote":
		c.output.addBlockquote()
		return NewBlockquoteTag(c.listStack.size() > 0)

	case "pre":
		c.preTagCount++
		return NewPreTag()

	case "code":
		// use fenced code block when inside a `pre` tag
		// similar implementation to list stacks
		// for fenced code blocks, language is important too
		c.codeTagCount++
		if c.preTagCount == 0 {
			return NewInlineCodeTag()
		}
		language := findCodeLanguage(node)
		return NewFencedCodeTag(language)

	case "br":
		return NewBRTag()

	case "hr":
		return NewHRTag()

	default:
		return NewUnknownTag(node.Data)
	}
}

func (c *Converter) convertNode(node *html.Node) {
	switch node.Type {
	case html.TextNode:
		// Write the text content, escaping special Markdown characters
		// * dont escape when inside a pre or fenced code tag
		text := node.Data
		if c.codeTagCount == 0 {
			text = escapeMarkdown(text)
		}
		c.output.WriteString(text)

	case html.ElementNode:
		if itemInSlice(node.Data, ignoreTags) {
			return
		}

		// Determine the Markdown type
		markdownElem := c.htmlNodeToMarkdownElement(node)

		// Write opening Markdown syntax
		c.output.WriteString(markdownElem.StartCode())

		// Recursively process child nodes
		for child := range node.ChildNodes() {
			c.convertNode(child)
		}

		// Write closing Markdown syntax
		endCode := markdownElem.EndCode()
		if markdownElem.Type() == Blockquote {
			// doing this before writing the endcode of blockquote
			// to prevent `>` in trailing newlines
			c.output.removeBlockquote()
		}
		c.output.WriteString(endCode)

		if markdownElem.Type() == Pre {
			c.preTagCount--
		} else if markdownElem.Type() == InlineCode || markdownElem.Type() == FencedCode {
			c.codeTagCount--
		} else if markdownElem.Type() == Anchor {
			c.output.insideAnchor = false
		} else if markdownElem.Type() == ListItem && node.NextSibling == nil {
			// last li tag in a list
			_, err := c.listStack.pop()
			if err != nil {
				// stack underflow
				// panic("no items in listStack to pop for the last li tag")
			}
			c.output.WriteString("\n") // write an extra newline when the list ends
		}

	}
}

func (c *Converter) ConvertString(input string) (string, error) {
	// Parse the HTML input into a document tree
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return "", err
	}

	// Start recursive conversion from the root node's children
	for node := doc.FirstChild; node != nil; node = node.NextSibling {
		c.convertNode(node)
	}

	return c.output.String(), nil
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
