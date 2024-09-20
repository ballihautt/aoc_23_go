package days

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

// Resolve Aoc 2023 day 2 challenges.
func Day02(input string) {
	var (
		validGamesCount uint
		powerSum        uint
	)

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		if isGamePossible(line) {
			validGamesCount += uint(i) + 1
		}
		powerSum += power(line)
	}
	fmt.Printf("Day 02 solution 1 is %d\n", validGamesCount)
	fmt.Printf("Day 02 solution 2 is %d\n", powerSum)

}

// power calculates the power of the minimum set.
func power(line string) uint {
	dotIdx := strings.Index(line, ":") // gets the index of ':' so the beginning can be trimmed.
	if dotIdx != -1 {
		line = line[dotIdx+1:] // trims at dotIdx + 1 to include the colon.
	}

	sets := strings.Split(line, ";")
	red, green, blue := minSet(sets)

	minPower := red * green * blue
	return minPower
}

// minSet returns the minimum set possible.
func minSet(sets []string) (red, green, blue uint) {
	for _, set := range sets {
		r, g, b := values(set)
		red = max(red, r)
		green = max(green, g)
		blue = max(blue, b)
	}
	return
}

// values gets the value of the set.
func values(set string) (red, green, blue uint) {
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
	return
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
