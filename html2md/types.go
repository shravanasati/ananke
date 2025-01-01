package html2md

import (
	"fmt"
	"strings"
)

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
	Image
	List // can be ordered as well unordered
	ListItem
	Blockquote
	InlineCode
	Unknown
)

type MarkdownElement interface {
	Type() MarkdownElementType
	StartCode() string
	EndCode() string
}

type H1Tag struct{}

func (h1 H1Tag) Type() MarkdownElementType {
	return H1
}
func (h1 H1Tag) StartCode() string {
	return "# "
}
func (h1 H1Tag) EndCode() string {
	return "\n"
}
func NewH1Tag() *H1Tag {
	return &H1Tag{}
}

type H2Tag struct{}

func (h2 H2Tag) Type() MarkdownElementType {
	return H2
}
func (h2 H2Tag) StartCode() string {
	return "## "
}
func (h2 H2Tag) EndCode() string {
	return "\n"
}
func NewH2Tag() *H2Tag {
	return &H2Tag{}
}

type H3Tag struct{}

func (h3 H3Tag) Type() MarkdownElementType {
	return H3
}
func (h3 H3Tag) StartCode() string {
	return "### "
}
func (h3 H3Tag) EndCode() string {
	return "\n"
}
func NewH3Tag() *H3Tag {
	return &H3Tag{}
}

type H4Tag struct{}

func (h4 H4Tag) Type() MarkdownElementType {
	return H4
}
func (h4 H4Tag) StartCode() string {
	return "#### "
}
func (h4 H4Tag) EndCode() string {
	return "\n"
}
func NewH4Tag() *H4Tag {
	return &H4Tag{}
}

type H5Tag struct{}

func (h5 H5Tag) Type() MarkdownElementType {
	return H5
}
func (h5 H5Tag) StartCode() string {
	return "##### "
}
func (h5 H5Tag) EndCode() string {
	return "\n"
}
func NewH5Tag() *H5Tag {
	return &H5Tag{}
}

type H6Tag struct{}

func (h6 H6Tag) Type() MarkdownElementType {
	return H6
}
func (h6 H6Tag) StartCode() string {
	return "###### "
}
func (h6 H6Tag) EndCode() string {
	return "\n"
}
func NewH6Tag() *H6Tag {
	return &H6Tag{}
}

type BoldTag struct{}

func (bold BoldTag) Type() MarkdownElementType {
	return Bold
}
func (bold BoldTag) StartCode() string {
	return "**"
}
func (bold BoldTag) EndCode() string {
	return "**"
}
func NewBoldTag() *BoldTag {
	return &BoldTag{}
}

type ItalicTag struct{}

func (italic ItalicTag) Type() MarkdownElementType {
	return Italic
}
func (italic ItalicTag) StartCode() string {
	return "*"
}
func (italic ItalicTag) EndCode() string {
	return "*"
}
func NewItalicTag() *ItalicTag {
	return &ItalicTag{}
}

// todo strikethrough tag (GFM)

type ParagraphTag struct{}

func (p ParagraphTag) Type() MarkdownElementType {
	return Paragraph
}
func (p ParagraphTag) StartCode() string {
	return ""
}
func (p ParagraphTag) EndCode() string {
	return "\n\n"
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
func (a AnchorTag) StartCode() string {
	return "["
}
func (a AnchorTag) EndCode() string {
	return fmt.Sprintf("](%v)", a.href)
}
func NewAnchorTag(href string) *AnchorTag {
	return &AnchorTag{href: href}
}

type ImageTag struct {
	src     string
	altText string
}

func (img ImageTag) Type() MarkdownElementType {
	return Image
}
func (img ImageTag) StartCode() string {
	return fmt.Sprintf("![%v](%v)", img.altText, img.src)
}
func (img ImageTag) EndCode() string {
	return "\n"
}
func NewImageTag(src, altText string) *ImageTag {
	return &ImageTag{src: src, altText: altText}
}

type ListOrdering uint

const (
	UnorderedList = iota
	OrderedList
)

type ListTag struct {
	type_ ListOrdering
	depth int
}

func (l ListTag) Type() MarkdownElementType {
	return List
}

func (l ListTag) StartCode() string {
	if l.depth == 0 {
		return ""
	}
	return "\n"
}
func (l ListTag) EndCode() string {
	return ""
}
func NewListTag(type_ ListOrdering, depth int) *ListTag {
	return &ListTag{type_: type_, depth: depth}
}

type ListItemTag struct {
	depth  int
	type_  ListOrdering
	number int
}

func (li ListItemTag) Type() MarkdownElementType {
	return ListItem
}

func (li ListItemTag) StartCode() string {
	if li.type_ == UnorderedList {
		return strings.Repeat("\t", li.depth) + "- "
	}
	return fmt.Sprintf("%v%v. ", strings.Repeat("\t", li.depth), li.number)
}
func (li ListItemTag) EndCode() string {
	return "\n"
}
func NewListItemTag(depth int, type_ ListOrdering, number int) *ListItemTag {
	return &ListItemTag{depth: depth, type_: type_, number: number}
}

type BlockquoteTag struct {}
func (bl BlockquoteTag) Type() MarkdownElementType {
	return Blockquote
}
func (bl BlockquoteTag) StartCode() string {
	return "> "
}
func (bl BlockquoteTag) EndCode() string {
	return "\n\n"
}
func NewBlockquoteTag() *BlockquoteTag {
	return &BlockquoteTag{}
}

type InlineCodeTag struct{}

func (ic InlineCodeTag) Type() MarkdownElementType {
	return InlineCode
}
func (ic InlineCodeTag) StartCode() string {
	return "`"
}
func (ic InlineCodeTag) EndCode() string {
	return "`"
}
func NewInlineCodeTag() *InlineCodeTag {
	return &InlineCodeTag{}
}


type UnknownTag struct {
	data string
}

func (u UnknownTag) Type() MarkdownElementType {
	return Unknown
}
func (u UnknownTag) StartCode() string { return "" }
func (u UnknownTag) EndCode() string   { return "" }
func NewUnknownTag(data string) *UnknownTag {
	return &UnknownTag{data: data}
}
