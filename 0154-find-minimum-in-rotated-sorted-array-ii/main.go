package _154_find_minimum_in_rotated_sorted_array_ii

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

请找出其中最小的元素。

注意数组中可能存在重复的元素。
*/
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] { // right not sorted
			left = mid + 1
		} else if nums[mid] < nums[left] { // left not sorted
			right = mid
		} else {
			right--
		}
	}
	return nums[left]
}

func findMin1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[left]
	}
	for left < right {
		mid := left + (right-left)/2
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[left] == nums[mid] {
			left++
			continue
		}
		if nums[0] < nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}