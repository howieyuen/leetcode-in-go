package _153_find_minimum_in_rotated_sorted_array

/*
升序数组旋转后，找最小元素
1、正常情况是nums[mid-1] < nums[mid] < nums[mid+1]
2、最小元素必然是转折点，即出现nums[mid-1]>nums[mid] 或者 nums[mid]>nums[mid+1]，最小值找到
3、否则，根据当前升降序，缩小查找范围
*/
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[left]
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[0] < nums[mid] { // left is in order
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
