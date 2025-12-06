package main

import (
	"strconv"
	"strings"
)

func d6() {
	input := readFileAsLines(6, 1)

	rows := make([][]int, len(input)-1)
	for i, _ := range rows {
		rows[i] = make([]int, 0, len(input[0]))
	}

	type opType int
	const (
		mul opType = iota
		add
	)
	ops := make([]opType, 0, len(input[0]))

	for i, line := range input[:len(input)-1] {
		for _, ele := range strings.Split(line, " ") {
			if ele == "" {
				continue
			}
			num, _ := strconv.Atoi(ele)
			rows[i] = append(rows[i], num)
		}
	}
	for _, ele := range strings.Split(input[len(input)-1], " ") {
		if ele == "" {
			continue
		}
		if ele == "*" {
			ops = append(ops, mul)
		} else {
			ops = append(ops, add)
		}
	}

	agg := 0
	for i, op := range ops {
		innerAgg := 0
		if op == mul {
			innerAgg = 1
		}
		for _, row := range rows {
			if op == mul {
				innerAgg *= row[i]
			} else {
				innerAgg += row[i]
			}
		}
		agg += innerAgg
	}

	p(r, agg)

	currentOp := mul
	opSet := false
	groupAgg := 0
	totalAgg := 0
	for col := 0; col < len(input[0]); col++ {
		numAgg := 0
		for row := 0; row < len(input)-1; row++ {
			digit := input[row][col]
			if digit != ' ' {
				numAgg = numAgg*10 + int(digit-'0')
			}
		}
		if !opSet {
			if input[len(input)-1][col] == '*' {
				currentOp = mul
				groupAgg = 1
			} else {
				currentOp = add
				groupAgg = 0
			}
			opSet = true
		}

		if numAgg == 0 {
			//next OP/gap found
			opSet = false
			totalAgg += groupAgg
			continue
		}

		if currentOp == mul {
			groupAgg *= numAgg
		} else {
			groupAgg += numAgg
		}
	}

	totalAgg += groupAgg

	p(r, totalAgg)
}
