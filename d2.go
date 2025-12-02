package main

import (
	"math"
	"strconv"
	"strings"
)

func d2() {

	input := readFileAsLines(2, 1)

	agg := 0
	for _, raw := range strings.Split(input[0], ",") {
		pair := strings.Split(raw, "-")

		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		for start <= end {

			if isSymmetricNumber(start) {
				agg += start
			}

			start++
		}

	}

	p(r, agg)
}

func isSymmetricNumber(i int) bool {
	digits := int(math.Log10(float64(i)) + 1)
	if i != 0 && digits%2 != 0 {
		return false //only even number of digits can be symmetrical
	}

	left := i % int(math.Pow10(digits/2))
	right := i / int(math.Pow10(digits/2))

	return left == right
}
