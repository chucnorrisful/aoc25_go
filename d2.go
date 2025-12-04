package main

import (
	"math"
	"strconv"
	"strings"
)

func d2() {
	input := readFileAsLines(2, 1)

	agg := 0
	agg2 := 0
	for _, raw := range strings.Split(input[0], ",") {
		pair := strings.Split(raw, "-")

		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		for start <= end {

			if isSymmetricNumber(start) {
				agg += start
			}
			if isRepeatingNumber(start) {
				agg2 += start
			}

			start++
		}

	}

	p(r, agg)
	p(r, agg2)
}

func isSymmetricNumber(i int) bool {
	digits := int(math.Log10(float64(i)) + 1)
	if i == 0 || digits%2 != 0 {
		return false //only even number of digits can be symmetrical
	}

	left := i % int(math.Pow10(digits/2))
	right := i / int(math.Pow10(digits/2))

	return left == right
}

func isRepeatingNumber(i int) bool {
	digits := int(math.Log10(float64(i)) + 1)
	if i == 0 {
		return false
	}

groupLoop:
	for groupSize := 1; groupSize <= digits/2; groupSize++ {
		//skip all group sizes that are non-dividers of digit count
		if digits%groupSize != 0 {
			continue
		}
		num := i / int(math.Pow10(groupSize))
		t := i % int(math.Pow10(groupSize)) // last group-size digits
		for num > 0 {
			t2 := num % int(math.Pow10(groupSize))
			if t2 != t {
				continue groupLoop
			}
			num = num / int(math.Pow10(groupSize))
		}
		return true
	}

	return false
}
