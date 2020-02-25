package _081_search_in_rotated_sorted_array_ii

func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
			continue
		}
		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[right] >= target && target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}
