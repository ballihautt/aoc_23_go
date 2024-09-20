package days

import (
	"fmt"
	"strings"
)

const (
	digitLength = 5
)

var digits map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// Resolve Aoc 2023 day 01 challenge.
func Day01(input string) {
	var total int

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		total += lineNumber(line)
	}

	fmt.Printf("Day 01 solution is %d\n", total)

}

// lineNumber extracts the number of the line.
func lineNumber(line string) int {
	var (
		firstDigit int
		lastDigit  int
	)
	length := len(line)
	firstDigit = -1
	lastDigit = -1

	for i := 0; i < length; i++ {
		if firstDigit < 0 {
			if isDigit(rune(line[i])) {
				firstDigit = int(line[i]) - 48 // convert the digit byte into the equivalent int.
			} else {
				if n := isNumber(line[i:]); n != -1 { // if the analysed section start with written digit, sets firstDigit variable.
					firstDigit = n //	condition can be factorised on one line, as a negative value for isNumber is a negative value for firstDigit too.
				}
			}

		}
		j := length - 1 - i
		if lastDigit < 0 {
			if isDigit(rune(line[j])) {
				lastDigit = int(line[j]) - 48 // convert the digit byte into the equivalent int.
			} else {
				if n := isNumber(line[j:]); n != -1 {
					lastDigit = n
				}
			}
		}
		if firstDigit >= 0 && lastDigit >= 0 {
			break
		}
	}
	return firstDigit*10 + lastDigit
}

// isNumber determines the string passed in parameters start with a digit.
func isNumber(s string) int {
	if len(s) < 3 {
		return -1
	}
	sLen := min(len(s), digitLength) // cap max size to digitLength as digits are not longer than 5

	for sLen >= 3 { // from sLen (3-5) to 3 to test all 3 size possibilities.
		d, ok := digits[s[:sLen]]
		if ok {
			return d
		}
		sLen--
	}
	return -1
}

// isDigit determines if the character is a digit or not.
func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}
