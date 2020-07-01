package _945_minimum_increment_to_make_array_unique

import (
	"sort"
)

func minIncrementForUnique(A []int) int {
	if len(A) == 0 {
		return 0
	}
	sort.Ints(A)
	ans := 0
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			ans += A[i-1] + 1 - A[i]
			A[i] = A[i-1] + 1
		}
	}
	return ans
}
