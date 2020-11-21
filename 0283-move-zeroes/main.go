package _283_move_zeroes

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

func moveZeroes1(nums []int) {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
}
