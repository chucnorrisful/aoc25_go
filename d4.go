package main

func d4() {
	input := readFileAsLines(4, 1)

	yBound := len(input)
	xBound := len(input[0])
	grid1 := make([][]int, 0)
	grid2 := make([][]int, 0)
	for range yBound {
		grid1 = append(grid1, make([]int, xBound))
		grid2 = append(grid2, make([]int, xBound))
	}

	for y, line := range input {
		for x, ele := range line {
			if ele == '@' {
				grid1[y][x] = -1
				grid2[y][x] = -1
			}
		}
	}

	pt1Cnt := 0
	for y, line := range grid1 {
		for x, ele := range line {
			if ele != -1 {
				continue
			}

			cnt := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					checkY, checkX := y+i, x+j

					if checkY < 0 || checkY >= yBound || checkX < 0 || checkX >= xBound || (checkX == x && checkY == y) {
						continue
					}

					if grid1[checkY][checkX] == -1 {
						cnt++
					}
				}
			}

			if cnt < 4 {
				pt1Cnt++
			}
		}
	}

	iter := 1
	cntSt3NbAll := 0
	for {
		cntSt3Nb := 0
		for y, line := range grid2 {
			for x, ele := range line {
				if ele != -1 {
					continue
				}

				cnt := 0
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						checkY, checkX := y+i, x+j

						if checkY < 0 || checkY >= yBound || checkX < 0 || checkX >= xBound || (checkX == x && checkY == y) {
							continue
						}

						if grid2[checkY][checkX] == -1 {
							cnt++
						}
					}
				}

				if cnt < 4 {
					cntSt3Nb++
					grid2[y][x] = iter
				}
			}
		}
		if cntSt3Nb == 0 {
			break
		}
		iter++
		cntSt3NbAll += cntSt3Nb
	}

	p(d, iter)
	p(r, pt1Cnt)
	p(r, cntSt3NbAll)
}
