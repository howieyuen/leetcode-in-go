package _015_3Sum

import (
	`sort`
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	if nums == nil || len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left+1] == nums[left] {
					left++
				}
				for left < right && nums[right-1] == nums[right] {
					right--
				}
				left++
				right--
			}
		}
	}
	return ans
}
