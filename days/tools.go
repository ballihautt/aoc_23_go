package days

import "strconv"

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
