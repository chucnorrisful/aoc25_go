package main

import (
	"strconv"
)

func d1() {
	input := readFileAsLines(1, 1)

	zCnt1 := 0
	zCnt2 := 0
	dial1 := 50
	dial2 := 50
	for _, in := range input {
		dir := -1
		if in[0] == 'R' {
			dir = 1
		}

		val, err := strconv.Atoi(in[1:])
		if err != nil {
			panic(err)
		}

		dial1 = (dial1 + val*dir) % 100
		if dial1 == 0 {
			zCnt1++
		}

		for val > 0 {
			dial2 = (dial2 + 1*dir) % 100
			val = val - 1
			if dial2 == 0 {
				zCnt2++
			}
		}

	}

	p(r, zCnt1)
	p(r, zCnt2)
}
