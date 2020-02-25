package _033_search_in_rotated_sorted_array

/*
	搜索旋转排序数组
	1. 二分查找的思想不变，重点在于判断nums[left:mid]有序或nums[mid:right]有序
	2. 再判断target在左侧or右侧
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		if nums[mid] <= nums[right] {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
