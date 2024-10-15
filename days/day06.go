package days

import (
	"fmt"
	"strings"
)

// Resolves Aoc 2023 day 06 challenges.
func Day06(input string) {
	lines := strings.Split(input, "\n")

	times := getNumbers(lines[0])
	dists := getNumbers(lines[1])

	var (
		i    uint64
		k    uint64
		ways uint64
	)
	ways = 1

	for i = 0; i < uint64(len(times)); i++ {
		for k = 1; k < times[i]; k++ {
			if (k * (times[i] - k)) > dists[i] {
				break
			}
		}
		ways *= times[i] + 1 - 2*k // +1 as times[i] is included in possibility test
	}

	fmt.Printf("Day 06 solution 1 is %d\n", ways)
}
