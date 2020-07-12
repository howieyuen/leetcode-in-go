package test3

import (
	"sort"
)

// 操作三次，消除3个最大、消除3个最小、消除2个最大1个最小、消除1个最大2个最小
// 最后4种情况选最优
func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	return min(nums[n-4]-nums[0], min(nums[n-3]-nums[1], min(nums[n-2]-nums[2], nums[n-1]-nums[3])))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
