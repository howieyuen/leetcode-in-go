package _719_find_k_th_smallest_pair_distance

import (
	"sort"
)

/*
给定一个整数数组，返回所有数对之间的第 k 个最小距离。一对 (A, B) 的距离被定义为 A 和 B 之间的绝对差值。
*/
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	max := nums[len(nums)-1] - nums[0]
	min := 0
	for min < max {
		mid := (max + min) / 2
		left := 0
		count := 0
		// 更新左右指针，直到nums[right]-nums[left]<=mid时，[left,right]范围内的距离，都小于mid
		for right := 0; right < len(nums); right++ {
			for nums[right]-nums[left] > mid {
				left++
			}
			// 计数，从[0, len(nums)-1]之间存在多少对数的距离小于等于mid
			count += right - left
		}
		if count >= k { // mid取大了，范围向左移动
			max = mid
		} else {
			min = mid + 1
		}
	}
	return min
}
