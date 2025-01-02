package html2md

import "errors"

var errStackUnderflow = errors.New("cannot pop: number of elements in the stack is 0")

type stack[T any] struct {
	elems []T
}

func newStack[T any]() *stack[T] {
	return &stack[T]{
		elems: []T{},
	}
}

func (s *stack[T]) push(elements ...T) {
	s.elems = append(s.elems, elements...)
}

func (s *stack[T]) pop() (T, error) {
	var value T

	length := len(s.elems)
	if length == 0 {
		return value, errStackUnderflow
	}

	value = s.elems[length-1]
	s.elems = s.elems[:length-1]

	return value, nil
}

func (s *stack[T]) top() (T, error) {
	var value T
	if s.size() == 0 {
		return value, errStackUnderflow
	}
	value = s.elems[s.size()-1]
	return value, nil
}

func (s *stack[T]) size() int {
	return len(s.elems)
}
