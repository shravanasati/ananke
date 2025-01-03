package html2md

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type counter interface {
	next() string
}

type counterType uint

const (
	decimal counterType = iota
	roman
	alphabet
)

type decimalCounter struct {
	current int
	step    int
}

func newDecimalCounter(start, step int) *decimalCounter {
	return &decimalCounter{
		current: start,
		step:    step,
	}
}

func (c *decimalCounter) next() string {
	c.current += c.step
	return strconv.Itoa(c.current - c.step)
}

type casing uint

const (
	lower casing = iota
	upper
)

type alphabetCounter struct {
	current int
	step    int
	case_   casing
}

func newAlphabetCounter(start, step int, c casing) *alphabetCounter {
	return &alphabetCounter{
		current: start,
		step:    step,
		case_:   c,
	}
}

func reverseString(input string) string {
	n := 0
	runes := make([]rune, len(input))
	for _, r := range input {
		runes[n] = r
		n++
	}
	runes = runes[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// Convert back to UTF-8.
	output := string(runes)

	return output
}

func (ac *alphabetCounter) decimalToAlphabet() string {
	temp := ac.current
	value := ""
	var baseline int
	if ac.case_ == upper {
		baseline = 65
	} else if ac.case_ == lower {
		baseline = 97
	} else {
		panic(fmt.Sprintf("unknown casing passed to alphabet counter: %v", ac.case_))
	}
	for temp != 0 {
		rem := temp % 26
		if rem == 0 {
			// for 'Z'
			rem = 26
			temp -= 1
		}
		temp /= 26
		value += string(rune(baseline + rem - 1))
	}

	return reverseString(value)
}

func (ac *alphabetCounter) next() string {
	value := ac.decimalToAlphabet()
	ac.current += ac.step
	return value
}

type romanCounter struct {
	current int
	step    int
	case_   casing
}

func newRomanCounter(start, step int, c casing) *romanCounter {
	return &romanCounter{
		current: start,
		step:    step,
		case_:   c,
	}
}

func (rc *romanCounter) decimalToRoman() string {
	num := rc.current
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	if rc.case_ == lower {
		// convert symbol  to lower case
		symbol = slices.Collect(
			mapIter(
				func(s string) string { return strings.ToLower(s) },
				slices.Values(symbol),
			),
		)
	}
	value := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	for num > 0 {
		for i := range value {
			if num >= value[i] {
				result += symbol[i]
				num -= value[i]
				break
			}
		}
	}
	return result
}

func (rc *romanCounter) next() string {
	value := rc.decimalToRoman()
	rc.current += rc.step
	return value
}
