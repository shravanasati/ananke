package html2md

import (
	"testing"
)

func TestOutputWriter(t *testing.T) {
	testCases := []struct {
		inputs   []string
		expected string
	}{
		{
			inputs:   []string{"Hello\n\n", "\nWorld", "\n\n\nAnother Line"},
			expected: "Hello\n\nWorld\n\nAnother Line",
		},
		{
			inputs:   []string{"- ", "Item 1", "\n", " - Subitem 1", "\n", "\n", "\n", "- Item 2", "\n\n"},
			expected: "- Item 1\n - Subitem 1\n\n- Item 2\n\n",
		},
	}

	for _, tc := range testCases {
		writer := newOutputWriter()
		for _, input := range tc.inputs {
			writer.WriteString(input)
		}

		got := writer.String()
		if got != tc.expected {
			t.Errorf("got=%v\nexpected=%v", replaceNewline(got), replaceNewline(tc.expected))
		}
	}

}
