package _162_find_peak_element

/*
	找到任何一个峰值所在的下标
	1、正常就是顺序遍历，nums[i-1] < nums[i] > nums[i+1]，同时满足此条件时，返回i
	2、如果nums本身是降序排序的，那么nums[i]>nums[i+1]，返回下标0，在左侧
	3、如果nums本身是升序排序的，那么nums[i]<nums[i+1]，返回下标len(nums)-1，在右侧
	所以，二分查找到mid，nums[mid]>nums[mid+1], 递归查左边，反之右侧
*/

func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func findPeakElement1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}
