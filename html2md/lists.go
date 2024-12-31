package html2md

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type listEntry struct {
	type_   ListOrdering
	counter *counter
}

func newListEntry(type_ ListOrdering) *listEntry {
	entry := &listEntry{type_: type_}
	if type_ == OrderedList {
		entry.counter = newCounter(1, 1)
	}

	return entry
}

// * This fingerprint method is introduced to uniquely identify a HTML node.
// * This is because while processing lists, for each li tag, its parent ol/ul tag
// * would also appear in the traversal. The code written previously assumed
// * that ol/ul would appear only once. To encounter this incorrect assumption,
// * the converter keeps a track of list tags it has seen. Thus, even if it appears 
// * again, the converter would just ignore it and NOT add it to the list stack.

// Generate a fingerprint for an HTML node
func generateFingerprint(node *html.Node) string {
	var builder strings.Builder

	// Include tag type
	builder.WriteString(node.Data)

	// Include all attributes
	for _, attr := range node.Attr {
		builder.WriteString(fmt.Sprintf("%s=%s;", attr.Key, attr.Val))
	}

	// Include parent tag type if it exists
	if node.Parent != nil {
		builder.WriteString(fmt.Sprintf("parent=%s;", node.Parent.Data))
	}

	// Include position among siblings
	position := 0
	for sibling := node.PrevSibling; sibling != nil; sibling = sibling.PrevSibling {
		position++
	}
	builder.WriteString(fmt.Sprintf("position=%d;", position))

	// Hash the constructed string
	hash := sha256.Sum256([]byte(builder.String()))
	return hex.EncodeToString(hash[:])
}
