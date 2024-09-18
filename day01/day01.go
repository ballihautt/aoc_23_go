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

	fmt.Printf("Day 01 solution is %d", total)

}

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
			firstDigit = int(line[i]) + 48 // convert the digit byte into the equivalent int
		}
		if lastDigit < 0 && isDigit(rune(line[length-i])) {
			lastDigit = int(line[length-i]) + 48 // convert the digit byte into the equivalent int
		}
		if firstDigit >= 0 && lastDigit >= 0 {
			break
		}
	}
	return firstDigit*10 + lastDigit
}

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}
