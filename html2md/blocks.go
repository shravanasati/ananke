package html2md

import (
	"strings"
)

type blockFormatter interface {
	transform(string) string
}

type blockquoteFormatter struct{}

func (bf *blockquoteFormatter) transform(s string) string {
	s = strings.ReplaceAll(s, "\n", "\n> ")
	return s
}

type listItemFormatter struct{
	isLast bool
}

func (lf *listItemFormatter) transform(s string) string {
	if !lf.isLast {
		s = strings.ReplaceAll(s, "\n", "\n    ")
	}
	return s
}
