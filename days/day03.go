package days

import (
	"fmt"
	"strings"
)

// Resolves Aoc 2023 day 03 challenges.
func Day03(input string) {
	var (
		sum          int
		gearRatioSum uint
		numbers      []int
	)

	engineMap := strings.Split(input, "\n")

	for j := 0; j < len(engineMap); j++ {
		for i := 0; i < len(engineMap[0]); i++ { // engineMap has fixed length lines.
			if isSymbol(rune(engineMap[j][i])) {
				numbers = findNumbers(engineMap, i, j)
				for _, n := range numbers {
					sum += n
				}
				if rune(engineMap[j][i]) == '*' && len(numbers) > 1 {
					grs := 1
					for _, n := range numbers {
						grs *= n
					}
					gearRatioSum += uint(grs)
				}
			}
		}
	}

	fmt.Printf("Day 03 solution 1 is %d\n", sum)
	fmt.Printf("Day 03 solution 2 is %d\n", gearRatioSum)
}

// isSymbol returns true if the rune is not a digit nor a dot.
func isSymbol(c rune) bool {
	return c != '.' && (c < '0' || c > '9')
}

// findNumbers searches for number around a symbol.
func findNumbers(engineMap []string, x, y int) (numbers []int) {
	ymin := max(0, y-1)
	ymax := min(len(engineMap)-1, y+1)
	xmin := max(0, x-1)
	xmax := min(len(engineMap[0])-1, x+1)

	for j := ymin; j <= ymax; j++ {
		for i := xmin; i <= xmax; i++ {
			if isDigit(rune(engineMap[j][i])) {
				if i == xmin || !isDigit(rune(engineMap[j][i-1])) {
					numbers = append(numbers, extractNumber(engineMap[j], i))
				}
			}
		}
	}
	return
}
