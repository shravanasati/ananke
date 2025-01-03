package html2md

import (
	"strings"
)

// outputWriter is a wrapper around strings.Builder.
// It ensures no more than 2 consecutive trailing newlines are written, even across multiple writes.
type outputWriter struct {
	writer           *strings.Builder
	trailingNewlines int
	blockquoteCount  int
	insideAnchor     bool // this is not a count because nested anchors are invalid in html
}

// newOutputWriter creates a new instance of outputWriter.
func newOutputWriter() *outputWriter {
	writer := new(strings.Builder)
	return &outputWriter{
		writer:           writer,
		trailingNewlines: 0,
		blockquoteCount:  0,
		insideAnchor:     false,
	}
}

func (w *outputWriter) addBlockquote() {
	w.blockquoteCount++
}

func (w *outputWriter) removeBlockquote() {
	if w.blockquoteCount == 0 {
		panic("remove blockquote called with 0 blockquoteCount")
	}
	w.blockquoteCount--
}

// countLeadingNewlines counts the number of leading newlines in a string.
func countLeadingNewlines(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '\n' {
			break
		}
		count++
	}
	return count
}

// countTrailingNewlines counts the number of trailing newlines in a string.
func countTrailingNewlines(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '\n' {
			break
		}
		count++
	}
	return count
}

// WriteString writes the string to the outputWriter, ensuring no more than 2 consecutive newlines.
func (w *outputWriter) WriteString(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	// fmt.Println("writing:", strings.ReplaceAll(s, "\n", "<newline>"))
	leadingNewlines := countLeadingNewlines(s)

	totalNewlines := w.trailingNewlines + leadingNewlines

	if totalNewlines > 2 {
		trimPrefix := totalNewlines - 2
		s = s[trimPrefix:]
	}

	trailingNewlines := countTrailingNewlines(s)

	if trailingNewlines > 2 {
		s = s[:len(s)-(trailingNewlines-2)]
		trailingNewlines = 2
	}

	// if the text only contains of new lines, then increment trailingNewLines
	if trailingNewlines != len(s) {
		w.trailingNewlines = trailingNewlines
	} else {
		w.trailingNewlines += trailingNewlines
	}

	if w.blockquoteCount > 0 {
		s = strings.ReplaceAll(s, "\n", "\n"+strings.Repeat("> ", w.blockquoteCount))
	}

	if w.insideAnchor && w.trailingNewlines == 1 {
		s = strings.ReplaceAll(s, "\n", "\n\\")
	}

	return w.writer.WriteString(s)
}

// String returns the complete string from the outputWriter.
func (w *outputWriter) String() string {
	return w.writer.String()
}
