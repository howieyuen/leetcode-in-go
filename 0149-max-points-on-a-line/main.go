package _149_max_points_on_a_line

import (
	"math"
)

type coordinate struct {
	x int
	y int
}

func maxPoints(points [][]int) int {
	coordinateMap := map[coordinate]int{}
	for _, p := range points {
		coordinateMap[coordinate{p[0], p[1]}]++
	}
	n := len(coordinateMap)
	if n <= 2 {
		return len(points)
	}
	coordinateSlice := make([]coordinate, n)
	i := 0
	for k := range coordinateMap {
		coordinateSlice[i] = k
		i++
	}
	ans, inf := 0, math.Inf(1)
	for i := range coordinateSlice {
		x, y := coordinateSlice[i].x, coordinateSlice[i].y
		slopes := map[float64]int{}
		tmp := 0
		for j := 0; j < i; j++ {
			if x == coordinateSlice[j].x {
				slopes[inf] += coordinateMap[coordinateSlice[j]]
				tmp = max(tmp, slopes[inf])
			} else {
				slope := float64(y-coordinateSlice[j].y) / float64(x-coordinateSlice[j].x)
				slopes[slope] += coordinateMap[coordinateSlice[j]]
				tmp = max(tmp, slopes[slope])
			}
		}
		ans = max(ans, tmp+coordinateMap[coordinateSlice[i]])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
