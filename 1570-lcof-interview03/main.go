package _570_lcof_interview03

import (
	"sort"
)

func findRepeatNumber(nums []int) int {
	for i := range nums {
		for nums[i] != i {
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}

func findRepeatNumber1(nums []int) int {
	tmp := make([]bool, len(nums))
	for i := range nums {
		if tmp[nums[i]] {
			return nums[i]
		}
		tmp[nums[i]] = true
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return nums[i-1]
		}
	}
	return -1
}
