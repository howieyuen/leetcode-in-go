package _973_k_closest_points_to_origin

import (
	"math/rand"
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	var euclid = func(point []int) int {
		return point[0]*point[0] + point[1]*point[1]
	}

	var quickSort func(left, right int)
	quickSort = func(left, right int) {
		if left == right {
			return
		}
		lessCount := left
		for i := left; i < right; i++ {
			if euclid(points[i]) < euclid(points[right]) {
				points[i], points[lessCount] = points[lessCount], points[i]
				lessCount++
			}
		}
		points[lessCount], points[right] = points[right], points[lessCount]

		if lessCount+1 == K {
			return
		} else if lessCount+1 < K {
			quickSort(lessCount+1, right)
		} else {
			quickSort(left, lessCount-1)
		}
	}

	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})
	quickSort(0, len(points)-1)
	return points[:K]
}

func kClosest1(points [][]int, K int) [][]int {
	var euclid = func(point []int) int {
		return point[0]*point[0] + point[1]*point[1]
	}
	for i := 0; i < K; i++ {
		exchanged := i
		min := euclid(points[i])
		for j := i + 1; j < len(points); j++ {
			if cur := euclid(points[j]); cur < min {
				min = cur
				exchanged = j
			}
		}
		points[i], points[exchanged] = points[exchanged], points[i]
	}
	return points[:K]
}

func kClosest2(points [][]int, K int) [][]int {
	var euclid = func(point []int) int {
		return point[0]*point[0] + point[1]*point[1]
	}
	sort.Slice(points, func(i, j int) bool {
		return euclid(points[i]) < euclid(points[j])
	})
	return points[:K]
}
