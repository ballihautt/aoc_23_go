package day02

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	redCubes   = 12 // is the number of cubes of this colour in the bag, given by statement
	greenCubes = 13 // --
	blueCubes  = 14 // --
)

func Solve(input string) {
	var (
		validGamesCount uint
	)

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if isGamePossible(line) {
			validGamesCount += uint(i) + 1
		}
	}
	fmt.Printf("Day 02 solution is %d\n", validGamesCount)

}

// isGamePossible determines if this game is a possible play.
func isGamePossible(line string) bool {

	dotIdx := strings.Index(line, ":") // gets the index of ':' so the beginning can be trimmed.
	if dotIdx != -1 {
		line = line[dotIdx+1:] // trims at dotIdx + 1 to include the colon.
	}

	sets := strings.Split(line, ";")

	for _, subset := range sets {
		if !isSetPossible(subset) {
			return false
		}
	}
	return true
}

// isSetPossible validates or invalidates a given set.
func isSetPossible(set string) bool {
	var (
		red   uint // counts the number of cubes of this colour in set.
		green uint // --
		blue  uint // --
	)

	for _, subset := range strings.Split(set, ",") {
		subset = subset[1:] // trims the leading space which is present after each ponctuation.
		parts := strings.Split(subset, " ")
		if len(parts) != 2 {
			panic("Unsupported non 2-length subset.")
		}
		n, _ := strconv.Atoi(parts[0])
		switch parts[1] {
		case "red":
			red += uint(n)
		case "green":
			green += uint(n)
		case "blue":
			blue += uint(n)
		}
	}

	if red > redCubes || green > greenCubes || blue > blueCubes {
		return false
	} else {
		return true
	}
}
