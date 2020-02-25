package _136_single_number

func singleNumber(nums []int) int {
	single := nums[0]
	for i := 1; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}
