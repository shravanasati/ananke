package html2md

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestDecimalCounter(t *testing.T) {
	start := 1
	step := 2
	c := newDecimalCounter(start, step)
	for i := 0; i < 78; i++ {
		got := c.next()
		expected := strconv.Itoa(start + i*step)
		if got != expected {
			t.Errorf("expected `%v`, got `%v`", expected, got)
		}
	}
}

func TestAlphabeticCounter(t *testing.T) {
	start := 1
	step := 1
	upperCounter := newAlphabetCounter(start, step, upper)
	lowerCounter := newAlphabetCounter(start, step, lower)

	upperExpectedValues := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
		"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
		"U", "V", "W", "X", "Y", "Z",
		"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ",
	}
	lowerExpectedValues := slices.Collect(mapIter(func(s string) string { return strings.ToLower(s) }, slices.Values(upperExpectedValues)))

	for i := 0; i < len(upperExpectedValues); i++ {
		got := upperCounter.next()
		expected := upperExpectedValues[i]
		if got != expected {
			t.Errorf("expected `%v`, got `%v`", expected, got)
		}
	}

	for i := 0; i < len(lowerExpectedValues); i++ {
		got := lowerCounter.next()
		expected := lowerExpectedValues[i]
		if got != expected {
			t.Errorf("expected `%v`, got `%v`", expected, got)
		}
	}

}

func TestRomanCounter(t *testing.T) {
	start := 1
	step := 1
	upperCounter := newRomanCounter(start, step, upper)
	lowerCounter := newRomanCounter(start, step, lower)

	upperExpectedValues := []string{
		"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
		"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
	}
	lowerExpectedValues := slices.Collect(mapIter(func(s string) string { return strings.ToLower(s) }, slices.Values(upperExpectedValues)))

	for i := 0; i < len(upperExpectedValues); i++ {
		got := upperCounter.next()
		expected := upperExpectedValues[i]
		if got != expected {
			t.Errorf("expected `%v`, got `%v`", expected, got)
		}
	}

	for i := 0; i < len(lowerExpectedValues); i++ {
		got := lowerCounter.next()
		expected := lowerExpectedValues[i]
		if got != expected {
			t.Errorf("expected `%v`, got `%v`", expected, got)
		}
	}
}
