package _041_first_missing_positive

import (
	`math`
)

func firstMissingPositive(nums []int) int {
	var n = len(nums)

	// 判断是否不存在1
	var contains = 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			contains++
			break
		}
	}
	if contains == 0 {
		return 1
	}

	// nums只有1个元素
	if n == 1 {
		return 2
	}

	// 负数，0，大于n的数字都置为1，数组中均>0
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = 1
		}
	}
	// 使用下标判断1~n的正整数是否存在
	// 使用绝对值，是为了保证操作重复元素时，不会负负得正，导致误判
	for i := 0; i < n; i++ {
		var a = int(math.Abs(float64(nums[i])))
		if a == n {
			nums[0] = - int(math.Abs(float64(nums[0])))
		} else {
			nums[a] = - int(math.Abs(float64(nums[a])))
		}
	}

	// 找到第一个正数即为解
	for i := 1; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}
	if nums[0] > 0 {
		return n
	}

	// 数组刚好是1~n的集合
	return n + 1
}
