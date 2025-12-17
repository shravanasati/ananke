package html2md

import (
	"strings"
	"testing"
)

// to better view the test output, \n -> <newline>
func replaceNewline(s string) string {
	return strings.ReplaceAll(s, "\n", "<newline>")
}

func TestConvertString(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple Heading",
			input:    `<h1>Hello World</h1>`,
			expected: "# Hello World\n",
		},
		{
			name:     "Paragraph with Bold and Italic",
			input:    `<p>This is <strong>bold</strong> and <em>italic</em>.</p>`,
			expected: "This is **bold** and *italic*.\n\n",
		},
		{
			name:     "Nested Bold and Italic",
			input:    `<p><strong>Bold and <em>Italic</em></strong></p>`,
			expected: "**Bold and *Italic***\n\n",
		},
		{
			name:     "Multiple Headings",
			input:    `<h1>Title</h1><h2>Subtitle</h2>`,
			expected: "# Title\n## Subtitle\n",
		},
		{
			name:     "Heading with Bold text",
			input:    `<h4><strong>Important</strong> heading</h4>`,
			expected: `#### **Important** heading` + "\n",
		},
		{
			name:     "Escaping Special Characters",
			input:    `<p>*Markdown* needs escaping: [link]</p>`,
			expected: `\*Markdown\* needs escaping: \[link\]` + "\n\n",
		},
		{
			name:     "More escaping",
			input:    `<h2># Heading #</h2><p># failed heading #hashtag</p>`,
			expected: `## \# Heading \#` + "\n" + `\# failed heading \#hashtag` + "\n\n",
		},
		{
			name:     "Unordered List",
			input:    `<ul><li>Item 1</li><li>Item 2</li></ul>`,
			expected: "- Item 1\n- Item 2\n\n",
		},
		{
			name:     "Ordered List",
			input:    `<ol><li>First</li><li>Second</li></ol>`,
			expected: "1. First\n2. Second\n\n",
		},
		{
			name:     "Nested List",
			input:    `<ul><li>Item 1<ul><li>Subitem 1</li></ul></li><li>Item 2</li></ul>`,
			expected: "- Item 1\n\t- Subitem 1\n\n- Item 2\n\n",
		},
		{
			name:     "Orphan List Items",
			input:    `<li>hello</li><li>world</li>`,
			expected: "- hello\n- world\n\n",
		},
		{
			name:     "Hyperlink",
			input:    `<p>Visit <a href="https://example.com">example</a>.</p>`,
			expected: "Visit [example](https://example.com).\n\n",
		},
		{
			name: "Hyperlink with line breaks and formatting",
			input: `<a href="/post">Line 1<br/>
<strong>Line 2</strong><br/>
Line 3
</a>`,
			expected: `[Line 1  
\
**Line 2**  
\
Line 3
](/post)`,
		},
		{
			name:     "Complex Document",
			input:    `<h1>Welcome</h1><p>This is <strong>bold</strong>, <em>italic</em>, and <a href="https://example.com">a link</a>.</p><ul><li>Item 1</li><li>Item 2</li></ul>`,
			expected: "# Welcome\nThis is **bold**, *italic*, and [a link](https://example.com).\n\n- Item 1\n- Item 2\n\n",
		},
		{
			name:     "Image Without Alt Text",
			input:    `<img src="https://example.com/image.png" />`,
			expected: "![image](https://example.com/image.png)\n",
		},
		{
			name:     "Image With Alt Text",
			input:    `<img src="https://example.com/image.png" alt="An example image" />`,
			expected: "![An example image](https://example.com/image.png)\n",
		},
		{
			name:     "Image Inside Paragraph",
			input:    `<p>Here is an image: <img src="https://example.com/image.png" alt="An example image" /></p>`,
			expected: "Here is an image: ![An example image](https://example.com/image.png)\n\n",
		},
		{
			name:     "Complex Document with Images",
			input:    `<h1>Gallery</h1><p>Check this out: <img src="https://example.com/cat.png" alt="A cat" /></p><img src="https://example.com/dog.png" />`,
			expected: "# Gallery\nCheck this out: ![A cat](https://example.com/cat.png)\n\n![image](https://example.com/dog.png)\n",
		},
		{
			name: "README",
			input: `<h1>ananke</h1>

<p>
A HTML to markdown converter.

Powered by <a href="https://github.com/shravanasati/ananke/tree/master/html2md">html2md</a>.
</p>

<h3>Usage</h3>
<p>
ananke can read.
</p>
`,
			expected: "# ananke\n\nA HTML to markdown converter.\n\nPowered by [html2md](https://github.com/shravanasati/ananke/tree/master/html2md).\n\n### Usage\n\nananke can read.\n\n",
		},
		{
			name:     "Code Block",
			input:    `<pre><code>fmt.Println("Hello, World!")</code></pre>`,
			expected: "```\nfmt.Println(\"Hello, World!\")\n```\n",
		},
		{
			name:     "Inline Code",
			input:    `<p>This is an example of <code>inline code</code> in a paragraph.</p>`,
			expected: "This is an example of `` inline code `` in a paragraph.\n\n",
		},
		{
			name:     "Blockquote",
			input:    `<blockquote>This is a blockquote.</blockquote>`,
			expected: "> This is a blockquote.\n\n",
		},
		{
			name:     "Nested Blockquote",
			input:    `<blockquote><p>This is a nested blockquote.</p><p>Are you kidding me?</p><blockquote>And this is another level.</blockquote></blockquote>`,
			expected: "> This is a nested blockquote.\n> \n> Are you kidding me?\n> \n> > And this is another level.\n> \n> ",
			// * kind of hack here, blockquotes with just newlines are equivalent to not having them at all
		},
		{
			name:     "Single Line Break",
			input:    `<p>First line<br>Second line</p>`,
			expected: "First line  \nSecond line\n\n",
		},
		{
			name:     "Multiple Line Breaks",
			input:    `<p>First line<br><br>Second line</p>`,
			expected: "First line  \n  \nSecond line\n\n",
		},
		{
			name:     "Line Break Between Tags",
			input:    `<p>First line</p><br><p>Second line</p>`,
			expected: "First line\n\n  \nSecond line\n\n",
		},
		{
			name:     "Horizontal Rule",
			input:    `<p>Before the line</p><hr /><p>After the line</p>`,
			expected: "Before the line\n\n---\n\nAfter the line\n\n",
		},
		{
			name:     "Standalone Horizontal Rule",
			input:    `<hr />`,
			expected: "---\n\n",
		},
		{
			name:     "Multiple Horizontal Rules",
			input:    `<hr /><hr /><p>After multiple lines</p>`,
			expected: "---\n\n---\n\nAfter multiple lines\n\n",
		},
		{
			name:     "Nested list with blockquote",
			input:    `<ul><li>Simple List</li><li><p>Someone once said:</p><blockquote>My famous quote</blockquote><span>by someone</span></li></ul>`,
			expected: "- Simple List\n- Someone once said:\n\n > My famous quote\n\nby someone\n\n",
		},
		{
			name:     "Ordered list with start",
			input:    `<ol start="9"><li>Nine</li><li>Ten</li><li>Eleven<ul><li>Nested</li></ul></li></ol>`,
			expected: "9. Nine\n10. Ten\n11. Eleven\n\t- Nested\n\n",
		},
		{
			name:     `Ordered list with type="A"`,
			input:    `<ol type="A" start="4"><li>First</li><li>Second</li><li>Third</li></ol>`,
			expected: "D. First\nE. Second\nF. Third\n\n",
		},
		{
			name:     `Ordered list with type="a"`,
			input:    `<ol type="a"><li>First</li><li>Second</li><li>Third</li></ol>`,
			expected: "a. First\nb. Second\nc. Third\n\n",
		},
		{
			name:     `Ordered list with type="I"`,
			input:    `<ol type="I" start="89"><li>First</li><li>Second</li><li>Third</li></ol>`,
			expected: "LXXXIX. First\nXC. Second\nXCI. Third\n\n",
		},
		{
			name:     `Ordered list with type="i"`,
			input:    `<ol type="i" start="v"><li>First</li><li>Second</li><li>Third</li></ol>`,
			expected: "i. First\nii. Second\niii. Third\n\n",
		},
		{
			name:     `Ordered list with type="5"`,
			input:    `<ol type="5"><li>Five</li><li>Six</li><li>Seven</li></ol>`,
			expected: "1. Five\n2. Six\n3. Seven\n\n",
		},

		{
			name:  "Blockquote with Heading, Ordered List, and Nested Blockquote",
			input: `<blockquote><h2>Heading</h2><ol><li>List</li><li>List</li></ol><blockquote><p>Another Quote</p><p>by someone</p></blockquote></blockquote>`,
			expected: `> ## Heading
> 1. List
> 2. List
> 
> > Another Quote
> > 
> > by someone
> > 
> > `,
		},
		{
			name:     "Inline Code",
			input:    `<p>Output a message: <br/><code>console.log("hello")</code></p>`,
			expected: "Output a message:   \n`` console.log(\"hello\") ``\n\n",
		},
		{
			name:     "Code with Backticks",
			input:    "<code>with `` backticks</code>",
			expected: "`` with `` backticks ``",
		},
		{
			name:     "Variable in Backticks",
			input:    "<code>`variable`</code>",
			expected: "`` `variable` ``",
		},
		{
			name:     "Code Block with Language Tag",
			input:    `<pre><code class="language-js">This ` + "``\ntotally ``` works!\n" + `</code></pre>`,
			expected: "```js\nThis ``\ntotally ``` works!\n\n```\n",
		},
		{
			name:     "link with title attribute",
			input:    `<a href="/about.html" title="title text">About</a>`,
			expected: `[About](/about.html "title text")`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			converter := NewConverter()
			output, err := converter.ConvertString(test.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if output != test.expected {
				t.Errorf("unexpected output:\nGot:      %s\nExpected: %s", replaceNewline(output), replaceNewline(test.expected))
			}
		})
	}
}
