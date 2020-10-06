package _075_sort_colors

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for i := 0; i < right; i++ {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			i--
			right--
		}
	}
}

func sortColors1(nums []int) {
	left, right := 0, len(nums)-1
	i := 0
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
	}
}
