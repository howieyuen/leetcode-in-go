package _485_max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			if max < count {
				max = count
			}
			count = 0
		}
	}
	if max < count {
		max = count
	}
	return max
}
