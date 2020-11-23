package _452_minimum_number_of_arrows_to_burst_balloons

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	arrows := 1
	y := points[0][1]
	for i := 1; i < len(points); i++ {
		if y < points[i][0] {
			arrows++
			y = points[i][1]
		}
	}
	return arrows
}
