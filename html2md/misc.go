package html2md

import "iter"

// collapseWhitespace reduces runs of HTML whitespace (space, tab, newline,
// carriage return, form feed) to a single space, except runs containing two or
// more newlines which are reduced to exactly two newlines. This mimics HTML's
// whitespace collapsing while still allowing explicit paragraph breaks to
// survive.
func collapseWhitespace(s string) string {
	if s == "" {
		return s
	}

	var out []rune
	spacePending := false
	newlineCount := 0

	flushSpace := func() {
		if newlineCount >= 2 {
			out = append(out, '\n', '\n')
		} else if spacePending {
			out = append(out, ' ')
		}
		spacePending = false
		newlineCount = 0
	}

	for _, r := range s {
		switch r {
		case ' ', '\t', '\f', '\v':
			spacePending = true
		case '\r':
			spacePending = true
		case '\n':
			spacePending = true
			newlineCount++
		default:
			flushSpace()
			out = append(out, r)
		}
	}

	flushSpace()

	return string(out)
}

// mapIter returns an iterator over f applied to seq.
func mapIter[In, Out any](f func(In) Out, seq iter.Seq[In]) iter.Seq[Out] {
	return func(yield func(Out) bool) {
		for in := range seq {
			if !yield(f(in)) {
				return
			}
		}
	}
}

func itemInSlice[T comparable](item T, slice []T) bool {
	for _, val := range slice {
		if val == item {
			return true
		}
	}
	return false
}
