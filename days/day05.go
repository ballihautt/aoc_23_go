package days

import (
	"fmt"
	"strconv"
	"strings"
)

// Resolves Aoc 2023 day 05 challenges.
func Day05(input string) {
	blocks := strings.Split(input, "\n\n")

	numbers := getSeeds(blocks[0]) // first block is the line of seeds.

	for i, block := range blocks[1:] {
		fmt.Printf("Block %d", i)
		numbers = transform(numbers, block)
	}

	minimum := numbers[0]
	for _, n := range numbers {
		if n < minimum {
			minimum = n
		}
	}

	fmt.Printf("Day 05 solution 1 is %d\n", minimum)
}

// getSeeds extracts the seeds from the first block of Almanac.
func getSeeds(block string) []uint {
	words := strings.Split(block, " ") // line has format : `seed: xx xxx xx xxxx x xx`
	seeds := words[1:]                 // removes leading `seed:`, so only seed numbers remain.
	var seedsUint []uint
	for i, seed := range seeds {
		if i%2 == 1 {
			continue
		}
		s, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(seeds[i+1])
		if err != nil {
			panic(err)
		}
		for k := s; k < s+r; k++ {
			seedsUint = append(seedsUint, uint(k))
		}
	}
	return seedsUint
}

// transform maps all source numbers to their corresponding destination number.
func transform(sourceNumbers []uint, block string) []uint {
	var (
		newNumbers []uint = make([]uint, 0, len(sourceNumbers))
	)
	lines := strings.Split(block, "\n")
	lines = lines[1:] // Removing section title

srcNumLoop:
	for _, n := range sourceNumbers {
		for _, line := range lines {
			transformation := strings.Split(line, " ")
			dstart, err1 := strconv.Atoi(transformation[0])
			sstart, err2 := strconv.Atoi(transformation[1])
			ran, err3 := strconv.Atoi(transformation[2])
			if err1 != nil || err2 != nil || err3 != nil {
				panic("line Atoi conversion fail")
			}
			d, s, r := uint(dstart), uint(sstart), uint(ran)

			if n >= s && n <= s+r {
				newNumbers = append(newNumbers, d+(n-s)) // add offset of n from s to d
				continue srcNumLoop                      // a destination number have been found
			}
		}
		newNumbers = append(newNumbers, n) // no match has been found for this one, it must be kept.
	}

	return newNumbers
}
