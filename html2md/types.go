package html2md

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
	Anchor
	Unknown
)

type MarkdownElement interface {
	Type() MarkdownElementType
	// String() string
}

type H1Tag struct{}
func (h1 H1Tag) Type() MarkdownElementType {
	return H1
}
func NewH1Tag() *H1Tag {
	return &H1Tag{}
}

type H2Tag struct{}
func (h2 H2Tag) Type() MarkdownElementType {
	return H2
}
func NewH2Tag() *H2Tag {
	return &H2Tag{}
}

type H3Tag struct{}
func (h3 H3Tag) Type() MarkdownElementType {
	return H3
}
func NewH3Tag() *H3Tag {
	return &H3Tag{}
}

type H4Tag struct{}
func (h4 H4Tag) Type() MarkdownElementType {
	return H4
}
func NewH4Tag() *H4Tag {
	return &H4Tag{}
}

type H5Tag struct{}
func (h5 H5Tag) Type() MarkdownElementType {
	return H5
}
func NewH5Tag() *H5Tag {
	return &H5Tag{}
}

type H6Tag struct{}
func (h6 H6Tag) Type() MarkdownElementType {
	return H6
}
func NewH6Tag() *H6Tag {
	return &H6Tag{}
}

type BoldTag struct{}
func (bold BoldTag) Type() MarkdownElementType {
	return Bold
}
func NewBoldTag() *BoldTag {
	return &BoldTag{}
}

type ItalicTag struct{}
func (italic ItalicTag) Type() MarkdownElementType {
	return Italic
}
func NewItalicTag() *ItalicTag {
	return &ItalicTag{}
}

type ParagraphTag struct{}
func (p ParagraphTag) Type() MarkdownElementType {
	return Paragraph
}
func NewParagraphTag() *ParagraphTag {
	return &ParagraphTag{}
}

type AnchorTag struct {
	href string
}
func (a AnchorTag) Type() MarkdownElementType {
	return Anchor
}
func NewAnchorTag(href string) *AnchorTag {
	return &AnchorTag{href: href}
}

type UnknownTag struct {}
func (u UnknownTag) Type() MarkdownElementType {
	return Unknown
}
func NewUnknownTag() *UnknownTag {
	return &UnknownTag{}
}