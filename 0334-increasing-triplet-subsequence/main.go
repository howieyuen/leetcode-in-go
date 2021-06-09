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

func increasingTriplet1(nums []int) bool {
	var max = math.MaxInt32
	var min = math.MaxInt32
	for _, v := range nums{
		if v > max { // 比最大值还大，满足
			return true
		} else if v > min { // 介于最大值和最小值之间，更新最大值，
			max = v
		} else { // 比最小值还小，更新最小值
			min = v
		}
	}
	return false
}