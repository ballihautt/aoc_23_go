package day01

import (
	"fmt"
	"strings"
)

// Resolve Aoc 2023 day 01 challenge
func Solve(input string) {
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
		if firstDigit < 0 && isDigit(rune(line[i])) {
			firstDigit = int(line[i]) - 48 // convert the digit byte into the equivalent int.
		}
		j := length - 1 - i
		if lastDigit < 0 && isDigit(rune(line[j])) {
			lastDigit = int(line[j]) - 48 // convert the digit byte into the equivalent int.
		}
		if firstDigit >= 0 && lastDigit >= 0 {
			break
		}
	}
	return firstDigit*10 + lastDigit
}

// isDigit determines if the character is a digit or not.
func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}
