package main

import (
	"strconv"
)

func d1() {
	input := readFileAsLines(1, 1)

	zCnt := 0
	dial := 50
	for _, in := range input {
		dir := -1
		if in[0] == 'R' {
			dir = 1
		}

		val, err := strconv.Atoi(in[1:])
		if err != nil {
			panic(err)
		}

		for val > 0 {
			dial = (dial + 1*dir) % 100
			val = val - 1
			if dial == 0 {
				zCnt++
			}
		}

		//dial = (dial + val*dir) % 100
	}

	p(r, zCnt)
}
