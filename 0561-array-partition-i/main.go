package _561_array_partition_i

import (
	`sort`
)

func arrayPairSum(nums []int) int {
	var ans int
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return ans
}
