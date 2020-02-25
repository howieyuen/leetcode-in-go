package _035_search_insert_position

/*
	【寻找target插入排序数组的位置】
	正常二分查找，返回left的位置
	如果target<nums[0], return 0
	如果target> nums[len[nums]-1], return len(nums)
*/
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	if target < nums[left] {
		return 0
	}
	if target > nums[right] {
		return right + 1
	}
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
