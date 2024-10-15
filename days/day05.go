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
	var ranges []urange = make([]urange, 0, len(numbers)/2)
	for i := 0; i < len(numbers); i += 2 {
		ranges = append(ranges, urange{numbers[i], numbers[i] + numbers[i+1]})
	}

	for i, block := range blocks[1:] {
		fmt.Printf("Block %d\n", i)
		ranges = transform(ranges, block)
	}

	minimum := ranges[0].min
	for _, r := range ranges {
		if r.min < minimum {
			minimum = r.min
		}
	}

	fmt.Printf("Day 05 solution 1 is %d\n", minimum)
}

// getSeeds extracts the seeds from the first block of Almanac.
func getSeeds(block string) []uint64 {
	words := strings.Split(block, " ") // line has format : `seed: xx xxx xx xxxx x xx`
	seeds := words[1:]                 // removes leading `seed:`, so only seed numbers remain.

	var seedsUint []uint64 = make([]uint64, len(seeds))
	for i, seed := range seeds {
		s, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seedsUint[i] = uint64(s)
	}
	return seedsUint
}

// transform maps all source ranges to their corresponding destination range.
func transform(sourceRanges []urange, block string) []urange {
	if len(sourceRanges) == 0 {
		return nil
	}
	var (
		newRanges       []urange = make([]urange, 0, len(sourceRanges))
		untouchedRanges []urange = make([]urange, 0, len(sourceRanges))
	)
	lines := strings.Split(block, "\n")
	lines = lines[1:] // Removing section title

srcNumLoop:
	for _, r := range sourceRanges {
		for _, line := range lines {
			transformation := strings.Split(line, " ")
			dstart, err1 := strconv.Atoi(transformation[0])
			sstart, err2 := strconv.Atoi(transformation[1])
			ran, err3 := strconv.Atoi(transformation[2])
			if err1 != nil || err2 != nil || err3 != nil {
				panic("line Atoi conversion fail")
			}
			d, s, ra := uint64(dstart), uint64(sstart), uint64(ran)

			untouched, newRange := applyTransformation(d, s, ra, r)
			if !newRange.isNul() {
				newRanges = append(newRanges, newRange)
				if len(untouched) > 0 {
					untouchedRanges = append(untouchedRanges, untouched...)
				}
				continue srcNumLoop
			}
		}
		newRanges = append(newRanges, r) // no match has been found for this one, it must be kept.
	}

	newRanges = append(newRanges, transform(untouchedRanges, block)...)

	return newRanges
}

// applyTransformation applies the transformation given by d,s and ran, to r.
// It returns a list of urange containing range of untouched numbers (2 ranges max as transform can apply inside min and max),
// and a range of transformed numbers. If no transformation have been applied, second parameter is nul urange.
func applyTransformation(d, s, ran uint64, r urange) ([]urange, urange) {
	if s+ran <= r.min || s >= r.max { // both bounds are either smaller (s < s+ran < r.min) or bigger (r.max >= s > s+ran) than the range r.
		return []urange{r}, urange{0, 0}
	}

	var untouched []urange = make([]urange, 0, 1)

	if s < r.min { // s < r.min but s+ran >= r.min, so applying to [r.min:s+ran]
		bis := r.split(s + ran)
		if !bis.isNul() { // if bis is nul, s+ran is bigger than r.max, transformation on all the range, nothing to return.
			untouched = append(untouched, bis) // bis is not nul, so there is an untouched part.
		}
		r.transform(s, d)
		return untouched, r
	} else if s+ran >= r.max { // s+r > r.max but s < r.max, so applying to [s:r.max]
		bis := r.split(s)
		if !bis.isNul() { // if bis is nul, s is lower than r.min, transformation on all the range, nothing to return in untouched[].
			bis.transform(s, d)
			if !r.isNul() {
				untouched = append(untouched, r) // bis is not nul, so there is an untouched part.
			}
			return untouched, bis
		} else {
			r.transform(s, d)
			return untouched, r
		}
	} else { // s and s+ran are in the range, so splitting twice to get the transformed one in the middle
		bis := r.split(s)
		if !r.isNul() {
			untouched = append(untouched, r)
		}
		ter := bis.split(s + ran)
		if !ter.isNul() {
			untouched = append(untouched, ter)
		}
		bis.transform(s, d)
		return untouched, bis
	}
}

// urange is a type containing 2 uint64, describing a range between them
type urange struct {
	min, max uint64
}

// isNul returns true if both bounds are equal to zero.
func (r *urange) isNul() bool {
	return r.min == 0 && r.max == 0
}

// isIn tests if k is in the range and return the result.
func (r *urange) isIn(k uint64) bool {
	return k >= r.min && k < r.max
}

// split splits the range on k. If k is not in the range (k < min or k >= max), it returns a nul range.
// k becomes the max of the actual range, and so not included, and the minimum of the returned range.
func (r *urange) split(k uint64) (bis urange) {
	if !r.isIn(k) { // k is not in range
		return bis
	} else { // k is in range
		bis.min, bis.max = k, r.max
		r.max = k
		if r.min == r.max { // empty range
			r.min, r.max = 0, 0
		}
		return
	}
}

// transform applies the offset of |d-s| by adding to d, r.bound - s as s <= r.bound
func (r *urange) transform(s, d uint64) {
	r.min = d + (r.min - s) // add offset or r.min from s to d to get corresponding transformed number.
	r.max = d + (r.max - s)
}
