package _034_find_first_and_last_position_of_element_in_sorted_array

import "sort"

/*
	在排序数组中搜索元素出现的第一个位置和最后位置，要求时间复杂度为O(log n)
	1. 二分查找元素的位置index，
	2. 以此位置向左left=index-1和向右right=index+1分别二分，直到左侧和右侧均不存在此元素，返回left和right
*/
func searchRange(nums []int, target int) []int {
	var start = 0
	var end = len(nums) - 1
	if index := binarySearch(nums, start, end, target); index != -1 {
		var left = index
		var right = index
		var lastLeft = left
		var lastRight = right
		for left != -1 {
			lastLeft = left
			left = binarySearch(nums, start, left-1, target)
		}
		for right != -1 {
			lastRight = right
			right = binarySearch(nums, right+1, end, target)
		}
		return []int{lastLeft, lastRight}
	}
	return []int{-1, -1}
}

func binarySearch(nums []int, start, end, target int) int {
	for start <= end {
		var mid = (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func searchRange1(nums []int, target int) []int {
	leftMost := sort.SearchInts(nums, target)
	if leftMost == len(nums) || nums[leftMost] != target {
		return []int{-1, -1}
	}
	rightMost := sort.SearchInts(nums, target+1) - 1
	return []int{leftMost, rightMost}
}
