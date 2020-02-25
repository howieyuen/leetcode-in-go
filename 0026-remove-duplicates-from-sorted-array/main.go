package _026_remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	var i = 0
	for j := 1; j < len(nums); j++ {
		for j < len(nums)-1 && nums[i] == nums[j] {
			j++
		}
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
