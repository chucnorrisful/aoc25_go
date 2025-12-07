package main

import "strings"

func d7() {
	input := readFileAsLines(7, 1)

	grid := make([][]int, len(input))
	for i, _ := range grid {
		grid[i] = make([]int, len(input[0]))
	}

	startInd := strings.Index(input[0], "S")
	tBank := make([]int, len(input[0]))
	tBank[startInd] = 1
	totalSplitCnt := 0
	for _, line := range input {
		for col, cell := range line {
			if tBank[col] == 0 {
				continue
			}

			if cell == '^' {
				totalSplitCnt++
				if col != 0 {
					tBank[col-1] += tBank[col]
				}
				if col < len(line) {
					tBank[col+1] += tBank[col]
				}

				tBank[col] = 0
			}
		}
	}

	timelineCnt := 0
	for _, tls := range tBank {
		timelineCnt += tls
	}

	p(r, totalSplitCnt)
	p(r, timelineCnt)
}
