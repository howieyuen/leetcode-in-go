package _540_single_element_in_a_sorted_array

func singleNonDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}

func singleNonDuplicate1(nums []int) int {
	var ans int
	for _, v := range nums {
		ans ^= v
	}
	return ans
}
