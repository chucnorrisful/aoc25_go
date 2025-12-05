package main

import (
	"strconv"
	"strings"
)

func d5() {
	input := readFileAsLines(5, 1)

	ins, outs := make([]int, 0), make([]int, 0)
	ings := make([]int, 0)

	isIng := false
	maxNum := 0
	for _, line := range input {
		if line == "" {
			isIng = true
			continue
		}

		if !isIng {
			minInd := strings.Index(line, "-")

			left, _ := strconv.Atoi(line[:minInd])
			right, _ := strconv.Atoi(line[minInd+1:])

			if maxNum < right {
				maxNum = right
			}

			ins = append(ins, left)
			outs = append(outs, right)
		} else {
			ing, _ := strconv.Atoi(line)
			ings = append(ings, ing)
		}
	}

	v := new(ArrPrioQ)
	v.Init(maxNum, false)

	p(d, "init prioq")

	inOutLk := make(map[int]int)
	for _, in := range ins {
		inOutLk[in] += 1
	}
	for _, out := range outs {
		inOutLk[out+1] += -1
	}

	for _, in := range ins {
		v.Insert(in)
	}
	for _, out := range outs {
		v.Insert(out + 1)
	}
	p(d, "filled prioq")

	spoiledLk := make(map[int]int)
	ptr := 0
	prevPtr := 0
	spCnt := 0
	prevSpCnt := 0
	gooderCnt := 0
	for ptr != v.Max() {
		ptr = v.Succ(ptr)
		spCnt += inOutLk[ptr]
		spoiledLk[ptr] = spCnt
		if spCnt == 0 && prevSpCnt > 0 {
			gooderCnt += ptr - prevPtr
		} else if spCnt > 0 && prevSpCnt == 0 {
			prevPtr = ptr
		}
		prevSpCnt = spCnt
	}

	goodCnt := 0
	for _, ing := range ings {
		sc := v.Pred(ing)
		if sp, _ := spoiledLk[sc]; sp > 0 {
			goodCnt++
		}
	}

	p(r, goodCnt)
	p(r, gooderCnt)
}
