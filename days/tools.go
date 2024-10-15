package days

import (
	"strconv"
	"strings"
)

// isDigit determines if the character is a digit or not.
func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// extractNumber gets the number at the index i of string s.
func extractNumber(s string, i int) int {
	var (
		lower int = i
		upper int = i
	)
	for lower >= 0 {
		if isDigit(rune(s[lower])) {
			lower--
		} else {
			break
		}
	}
	lower += 1
	for upper < len(s) {
		if isDigit(rune(s[upper])) {
			upper++
		} else {
			break
		}
	}
	ret, err := strconv.Atoi(s[lower:upper])
	if err != nil {
		panic("extractNumber string to int conversion error")
	}
	return ret
}

// getNumbers extracts a number list from a line, separated by spaces, and eventually beginning with a line Id
func getNumbers(line string) []uint64 {
	line = trimId(line)
	line = strings.Trim(line, " ")

	stringNumbers := strings.Split(line, " ")

	var numbers []uint64 = make([]uint64, 0, len(stringNumbers))
	for _, sn := range stringNumbers {
		if sn == "" {
			continue
		}
		n, err := strconv.Atoi(sn)
		if err != nil {
			panic("Trying to convert a not all digits string")
		}
		numbers = append(numbers, uint64(n))
	}

	return numbers
}

func trimId(line string) string {
	dotIdx := strings.Index(line, ":") // gets the index of ':' so the beginning can be trimmed.
	if dotIdx != -1 {
		line = line[dotIdx+1:] // trims at dotIdx + 1 to include the colon.
	}
	return line
}
