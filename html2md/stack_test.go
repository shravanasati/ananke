package html2md

import (
	"errors"
	"slices"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := newStack[int]()
	if stack.size() != 0 {
		t.Errorf("expected size 0, got %d", stack.size())
	}
}

func TestPush(t *testing.T) {
	stack := newStack[int]()
	stack.push(1, 2, 3)

	if stack.size() != 3 {
		t.Errorf("expected size 3, got %d", stack.size())
	}

	top, _ := stack.top()
	if top != 3 {
		t.Errorf("expected top element 3, got %d", top)
	}
}

func TestPop(t *testing.T) {
	stack := newStack[int]()
	stack.push(1, 2, 3)

	element, err := stack.pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if element != 3 {
		t.Errorf("expected popped element 3, got %d", element)
	}

	if stack.size() != 2 {
		t.Errorf("expected size 2, got %d", stack.size())
	}
}

func TestPopEmptyStack(t *testing.T) {
	stack := newStack[int]()
	_, err := stack.pop()
	if !errors.Is(err, errStackUnderflow) {
		t.Errorf("expected error %v, got %v", errStackUnderflow, err)
	}
}

func TestTop(t *testing.T) {
	stack := newStack[int]()
	stack.push(10, 20, 30)

	top, err := stack.top()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if top != 30 {
		t.Errorf("expected top element 30, got %d", top)
	}
}

func TestTopEmptyStack(t *testing.T) {
	stack := newStack[int]()
	_, err := stack.top()
	if !errors.Is(err, errStackUnderflow) {
		t.Errorf("expected error %v, got %v", errStackUnderflow, err)
	}
}

func TestSize(t *testing.T) {
	stack := newStack[int]()
	if stack.size() != 0 {
		t.Errorf("expected size 0, got %d", stack.size())
	}

	stack.push(1, 2, 3)
	if stack.size() != 3 {
		t.Errorf("expected size 3, got %d", stack.size())
	}
}

func TestPushAndPopMultiple(t *testing.T) {
	stack := newStack[string]()
	stack.push("a", "b", "c")

	element, err := stack.pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if element != "c" {
		t.Errorf("expected popped element 'c', got '%s'", element)
	}

	element, err = stack.pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if element != "b" {
		t.Errorf("expected popped element 'b', got '%s'", element)
	}

	element, err = stack.pop()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if element != "a" {
		t.Errorf("expected popped element 'a', got '%s'", element)
	}

	if stack.size() != 0 {
		t.Errorf("expected size 0, got %d", stack.size())
	}
}

func TestAll(t *testing.T) {
	s := newStack[int]()
	elems := []int{3,4,5,6,7}
	s.push(elems...)
	
	slices.Reverse(elems)
	i := 0
	for elem := range s.All() {
		if elem != elems[i] {
			t.Errorf("expected %v got %v", elems[i], elem)
		}
		i++
	}
}