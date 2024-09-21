package days

import (
	"fmt"
	"strings"
)

// Resolves Aoc 2023 day 04 challenges.
func Day04(input string) {
	var sum uint

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		sum += linePoints(line)
	}

	fmt.Printf("Day 04 solution 1 is %d\n", sum)
}

// linePoints counts the points gained on a line.
// It trims by itself the line Id.
func linePoints(line string) uint {
	line = trimId(line) // removes the "Id" part at the beginning.

	winNumbersString, numbersString, found := strings.Cut(line, "|") // separates the winning numbers from the numbers.
	if !found {
		panic("'|' separator not found")
	}

	winNumbersString = strings.TrimSpace(winNumbersString) // removes leading and trailing spaces.
	winNumbers := strings.Split(winNumbersString, " ")     // splits over spaces.

	numbersString = strings.TrimSpace(numbersString)
	numbers := strings.Split(numbersString, " ")

	var points uint = 0

	for _, n := range numbers {
		for _, wn := range winNumbers {
			if n == wn && n != "" {
				if points == 0 {
					points += 1
				} else {
					points *= 2
				}
			}
		}
	}

	return points
}
