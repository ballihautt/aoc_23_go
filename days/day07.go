package days

import (
	"fmt"
	"strings"
)

var cards = map[rune]uint64{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

// Resolves Aoc 2023 day 07 challenges.
func Day07(input string) {
	lines := strings.Split(input, "\n")

	var (
		llen       = len(lines)
		handValues = make([]uint16, llen)
		minIndex   int
		minValue   uint16
		minBid     int

		sum uint64
	)

	for i, l := range lines {
		words := strings.Split(l, " ")
		handValues[i] = handLevel(words[0])
	}

	minValue = handValues[0]
	minBid = extractNumber(lines[0], 6)
	for i := 1; i < llen+1; i++ {
		for j := 0; j < llen; j++ {
			if handValues[j] < minValue || (handValues[j] == minValue && isHandLower(lines[minIndex][:5], lines[j][:5])) {
				minValue = handValues[j]
				minBid = extractNumber(lines[j], 6)
				minIndex = j
			}
		}
		sum += uint64(i * minBid)
		handValues[minIndex] = 99 + uint16(minIndex) // High value to not be counted again. Adding index as salt to not compare ghost hand.
		if i == llen {
			break
		}
		minValue = handValues[i]
		minBid = extractNumber(lines[i], 6)
		minIndex = i
	}

	fmt.Printf("Day 07 solution 1 is %d\n", sum)
}

// handLevel returns the strength of the hand. Max is 7, for a five of a kind
func handLevel(hand string) uint16 {
	if len(hand) != 5 {
		return 0 // Invalid hand.
	}
	var chars map[rune]uint16 = make(map[rune]uint16, 5)

	for _, c := range hand {
		chars[c] += 1
	}

	var max uint16 = 0
	for c, e := range chars {
		if e == max && max == 2 {
			if chars['J'] != 0 {
				break // If one is a J, a special rule applies.
			}
			return 3 //	Elements are in 1-5 range, if there are 2 pairs, it is the only possibility.
		}
		if e > max && c != 'J' { // 'J' should not be the max
			max = e
		}
	}

	max += chars['J']
	delete(chars, 'J')

	switch max {
	case 5:
		fallthrough
	case 4:
		return max + 2 // 	Only one possibility for each 4 and 5 outcomes, and rank 6 and 7.
	case 3:
		if len(chars) == 2 { // 2 different types, meaning one type has 3 cards, while the other 2. This is a full house. Ranks 5.
			return max + 2
		} else { // 3 different types, with one type with 3 cards, three of a kind. Ranks 4.
			return max + 1
		}
	default:
		return max // Only one possibility for 1, 2 as the two pairs outcome is detected in the max search loop. Rank as 1 and 2.
	}
}

// isHandLower returns true if the second hand is smaller than the first.
func isHandLower(a string, b string) bool {
	if len(a) != 5 || len(b) != 5 {
		panic("Non standard hand length")
	}

	for i := 0; i < 5; i++ {
		if cards[rune(b[i])] < cards[rune(a[i])] {
			return true
		} else if cards[rune(b[i])] > cards[rune(a[i])] {
			return false
		}
	}

	return false
}
