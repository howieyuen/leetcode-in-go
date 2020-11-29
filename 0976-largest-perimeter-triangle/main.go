package _976_largest_perimeter_triangle

import (
	"sort"
)

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 2; i-- {
		a, b, c := A[i], A[i-1], A[i-2]
		if a < b+c {
			return a + b + c
		}
	}
	return 0
}
