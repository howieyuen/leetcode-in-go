package _209_minimum_size_subarray_sum

import (
	`math`
)

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	left := 0
	ans := math.MaxInt32
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= s {
			if ans > i-left+1 {
				ans = i - left + 1
			}
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}