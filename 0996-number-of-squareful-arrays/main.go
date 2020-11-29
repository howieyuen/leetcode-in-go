package _996_number_of_squareful_arrays

import (
	"math"
	"sort"
)

func numSquarefulPerms(A []int) int {
	var visited = make([]bool, len(A))
	var ans int
	var backtrace func(count int, pre int)
	backtrace = func(count int, pre int) {
		if count == len(A) {
			ans++
			return
		}
		for i := range A {
			if visited[i] {
				continue
			}
			if i > 0 && A[i] == A[i-1] && !visited[i-1] {
				continue
			}
			sum := float64(pre + A[i])
			x := math.Sqrt(sum)
			if count > 0 && x!=float64(int(x)) {
				continue
			}
			visited[i] = true
			backtrace(count+1, A[i])
			visited[i] = false
		}
	}
	sort.Ints(A)
	backtrace(0, 0)
	return ans
}
