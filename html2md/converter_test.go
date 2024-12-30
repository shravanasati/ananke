package html2md

import "testing"

func TestConvertString(t *testing.T) {
	converter := &Converter{}

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
			name:     "Escaping Special Characters",
			input:    `<p>*Markdown* needs escaping: [link]</p>`,
			expected: `\*Markdown\* needs escaping: \[link\]\n\n`,
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
			expected: "- Item 1\n  - Subitem 1\n- Item 2\n\n",
		},
		{
			name:     "Hyperlink",
			input:    `<p>Visit <a href="https://example.com">example</a>.</p>`,
			expected: "Visit [example](https://example.com).\n\n",
		},
		{
			name:     "Complex Document",
			input:    `<h1>Welcome</h1><p>This is <strong>bold</strong>, <em>italic</em>, and <a href="https://example.com">a link</a>.</p><ul><li>Item 1</li><li>Item 2</li></ul>`,
			expected: "# Welcome\n\nThis is **bold**, *italic*, and [a link](https://example.com).\n\n- Item 1\n- Item 2\n\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := converter.ConvertString(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if output != test.expected {
				t.Errorf("unexpected output:\nGot:      %s\nExpected: %s", output, test.expected)
			}
		})
	}
}
