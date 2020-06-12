package _334_increasing_triplet_subsequence

import (
	"math"
)

// 递增三元子序列
// 维护a,b两个数，a < b，找到大于b的数即可
func increasingTriplet(nums []int) bool {
	var a, b = math.MaxInt64, math.MaxInt64
	for _, num := range nums {
		if a >= num {
			a = num
		} else if b >= num {
			b = num
		} else {
			return true
		}
	}
	return false
}
