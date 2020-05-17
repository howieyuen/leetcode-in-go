package _586_lcof_17

import (
	"math"
)

func printNumbers(n int) []int {
	end := int(math.Pow(10, float64(n)))
	ans := make([]int, end-1)
	for i := 0; i < end-1; i++ {
		ans[i] = i + 1
	}
	return ans
}
