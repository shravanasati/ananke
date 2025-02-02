package html2md

import "iter"

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