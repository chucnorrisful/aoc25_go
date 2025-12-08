package main

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

type point struct{ x, y, z int }

func (pt *point) distSq(pt2 *point) float64 {
	return math.Pow(float64(pt.x-pt2.x), 2) +
		math.Pow(float64(pt.y-pt2.y), 2) +
		math.Pow(float64(pt.z-pt2.z), 2)
}
func (pt *point) dist(pt2 *point) float64 {
	return math.Sqrt(math.Pow(float64(pt.x-pt2.x), 2) +
		math.Pow(float64(pt.y-pt2.y), 2) +
		math.Pow(float64(pt.z-pt2.z), 2))
}

func d8() {
	ver := 1
	input := readFileAsLines(8, ver)
	setStepsPt1 := 10 //
	if ver == 1 {
		setStepsPt1 = 1000
	}

	points := make([]*point, 0, len(input))
	for _, line := range input {
		raw := strings.Split(line, ",")
		x, _ := strconv.Atoi(raw[0])
		y, _ := strconv.Atoi(raw[1])
		z, _ := strconv.Atoi(raw[2])
		points = append(points, &point{x, y, z})
	}

	type distance struct {
		id1, id2 int
		dist     float64
	}
	//dists := make([][]distance, len(input))
	//for i, _ := range dists {
	//	dists[i] = make([]distance, len(input))
	//}

	distsList := make([]distance, 0)
	for i, pt := range points {
		for j, pt2 := range points {
			if i >= j {
				continue
			}

			dist := pt.dist(pt2)
			//dists[i][j] = distance{dist: dist, id1: i, id2: j}
			distsList = append(distsList, distance{dist: dist, id1: i, id2: j})
		}
	}

	slices.SortFunc(distsList, func(a, b distance) int {
		return cmp.Compare(a.dist, b.dist)
	})

	//circuits := make([][]*point, 0)
	circuits := make([][]int, 0)
	for pid, _ := range points {
		circuits = append(circuits, []int{pid})
	}

	k := 0
	aggPt1 := 1
	pt1C, pt2C := -1, -1
	pt1ID, pt2ID := -1, -1
	for len(circuits[0]) < len(points) {
		minDist := distsList[k]

		pt1C, pt2C = -1, -1
		for cID, circuit := range circuits {
			for _, cPt := range circuit {
				if cPt == minDist.id1 {
					pt1C = cID
					pt1ID = cPt
				}
				if cPt == minDist.id2 {
					pt2C = cID
					pt2ID = cPt
				}
			}
		}

		if pt1C == -1 && pt2C == -1 {
			// not found: new circuit with the two points
			circuits = append(circuits, []int{minDist.id1, minDist.id2})
		} else if pt1C == -1 {
			circuits[pt2C] = append(circuits[pt2C], minDist.id1)
		} else if pt2C == -1 {
			circuits[pt1C] = append(circuits[pt1C], minDist.id2)
		} else if pt1C != pt2C {
			circuits[pt1C] = append(circuits[pt1C], circuits[pt2C]...)
			circuits = append(circuits[:pt2C], circuits[pt2C+1:]...)
		}

		//pt1 result
		if k == setStepsPt1-1 {
			//find the biggest circuits and add

			slices.SortFunc(circuits, func(a, b []int) int {
				return len(b) - len(a)
			})
			for i := 0; i < 3; i++ {
				aggPt1 *= len(circuits[i])
			}
		}
		k++
	}

	p(r, aggPt1)
	p(d, points[pt1ID], points[pt2ID])
	p(r, points[pt1ID].x*points[pt2ID].x)
}
