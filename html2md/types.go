package html2md

type MarkdownElement interface {
	String() string
}

// type BoldText struct {
// 	children []
// }

type MarkdownElementType int

const (
	Bold = iota
	Italic
	H1
	H2
	H3
	H4
	H5
	H6
	Paragraph
	Unknown
)