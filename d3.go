package main

import "math"

func d3() {

	input := readFileAsLines(3, 1)

	agg := 0
	agg2 := 0
	for _, bankRaw := range input {
		bank := []byte(bankRaw)
		for i := 0; i < len(bank); i++ {
			bank[i] -= 48 //convert ascii
		}

		agg += naiveHighest2Joltage(bank)
		agg2 += highestNJoltage(bank, 12)
	}

	p(r, agg)
	p(r, agg2)
}

func highestNJoltage(bank []byte, n int) int {
	highestInds := make([]int, n)

	leadingInd := -1
	for i := n - 1; i >= 0; i-- {
		highest := 0
		shortBank := bank[leadingInd+1 : len(bank)-i]
		for ind, jolt := range shortBank {
			if int(jolt) > highest {
				highestInds[i] = ind + leadingInd + 1
				highest = int(jolt)
			}
		}

		leadingInd = highestInds[i]
	}

	agg := 0
	for i := n - 1; i >= 0; i-- {
		agg += int(bank[highestInds[i]]) * int(math.Pow10(i))
	}
	return agg
}

func naiveHighest2Joltage(bank []byte) int {

	highest := 0
	highestTInd := -1
	highestEInd := -1

	for i := 0; i < len(bank)-1; i++ { //skip rightmost battery as we cant use it as 10^1 place
		if int(bank[i]) > highest {
			highest = int(bank[i])
			highestTInd = i
		}
	}
	highest = 0
	for i := 0; i < len(bank); i++ {
		if i <= highestTInd {
			continue
		}
		if int(bank[i]) > highest {
			highest = int(bank[i])
			highestEInd = i
		}
	}

	//fmt.Printf("%v\n%v %v\n%v %v\n", bank, highestTInd, highestEInd, int(bank[highestTInd]), int(bank[highestEInd]))

	return int(bank[highestTInd])*10 + int(bank[highestEInd])
}
