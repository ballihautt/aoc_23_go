package days

import (
	"fmt"
	"strings"
)

// Resolves Aoc 2023 day 04 challenges.
func Day04(input string) {
	var (
		sum    uint
		points uint

		matching        uint
		scratchcards    []uint
		scratchcardsSum uint
	)

	scratchcards = make([]uint, 256) //	large enough to contain the 200+ lines

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		scratchcards[i] += 1
		points, matching = linePoints(line)
		sum += points

		for j := uint(1); j <= matching; j++ {
			scratchcards[uint(i)+j] += scratchcards[i]
		}
	}

	for i := 0; i < len(lines); i++ {
		scratchcardsSum += scratchcards[i]
	}

	fmt.Printf("Day 04 solution 1 is %d\n", sum)
	fmt.Printf("Day 04 solution 2 is %d\n", scratchcardsSum)

}

// linePoints counts the points gained on a line.
// It trims by itself the line Id.
func linePoints(line string) (points, matchingNumbers uint) {
	line = trimId(line) // removes the "Id" part at the beginning.

	winNumbersString, numbersString, found := strings.Cut(line, "|") // separates the winning numbers from the numbers.
	if !found {
		panic("'|' separator not found")
	}

	winNumbersString = strings.TrimSpace(winNumbersString) // removes leading and trailing spaces.
	winNumbers := strings.Split(winNumbersString, " ")     // splits over spaces.

	numbersString = strings.TrimSpace(numbersString)
	numbers := strings.Split(numbersString, " ")

	for _, n := range numbers {
		for _, wn := range winNumbers {
			if n == wn && n != "" {
				matchingNumbers++
				if points == 0 {
					points += 1
				} else {
					points *= 2
				}
			}
		}
	}

	return
}
