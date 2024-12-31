package html2md

import "testing"

func TestCounter(t *testing.T) {
	start := 1
	step := 2
	c := newCounter(start, step)
	for i := 0; i < 78; i++ {
		got := c.next()
		expected := start + i*step
		if got != expected {
			t.Errorf("expected `%v`, got `%v`", expected, got)
		}
	}
}
