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
	counter counter
}

// _newListEntry creates and returns a new *listEntry with the given list ordering and start.
// `start` parameter is relevant only when `type_` is `Ordered`.
func _newListEntry(orderType ListOrdering, start int, counterType counterType, case_ casing) *listEntry {
	entry := &listEntry{type_: orderType}
	if orderType == OrderedList {
		switch counterType {
		case decimal:
			entry.counter = newDecimalCounter(start, 1)
		case roman:
			entry.counter = newRomanCounter(start, 1, case_)
		case alphabet:
			entry.counter = newAlphabetCounter(start, 1, case_)
		default:
			panic(fmt.Sprintf("unknown counter type: %v", counterType))
		}
	}

	return entry
}

func newUnorderedListEntry() *listEntry {
	return _newListEntry(UnorderedList, 1, decimal, lower) // last 3 params are irrelevant
}

func newOrderedListEntry(start int, counterType counterType, case_ casing) *listEntry {
	return _newListEntry(OrderedList, start, counterType, case_)
}

// This function uses the "type" attribute for ol tag to figure out
// the counter type and casing to be used.
func getOrderedListParams(typeString string) (counterType, casing) {
	var countType counterType
	var case_ casing

	if typeString == "i" {
		countType = roman
		case_ = lower
	} else if typeString == "I" {
		countType = roman
		case_ = upper
	} else if typeString == "a" {
		countType = alphabet
		case_ = lower
	} else if typeString == "A" {
		countType = alphabet
		case_ = upper
	} else {
		countType = decimal
	}

	return countType, case_
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
