package main

func d4() {
	input := readFileAsLines(4, 1)

	yBound := len(input)
	xBound := len(input[0])
	grid := make([][]int, 0)
	grid3 := make([][]int, 0)
	for range yBound {
		grid = append(grid, make([]int, xBound))
		grid3 = append(grid3, make([]int, xBound))
	}

	for y, line := range input {
		for x, ele := range line {
			if ele == '@' {
				grid[y][x] = -1
				grid3[y][x] = -1
			}
		}
	}

	cntSt3Nb3 := 0
	for y, line := range grid3 {
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

					if grid3[checkY][checkX] == -1 {
						cnt++
					}
				}
			}

			if cnt < 4 {
				cntSt3Nb3++
			}
		}
	}

	iter := 1
	cntSt3NbAll := 0
	for {
		cntSt3Nb := 0
		for y, line := range grid {
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

						if grid[checkY][checkX] == -1 {
							cnt++
						}
					}
				}

				if cnt < 4 {
					cntSt3Nb++
					grid[y][x] = iter
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
	p(r, cntSt3Nb3)
	p(r, cntSt3NbAll)
}
