package _8_03

func findMagicIndex(nums []int) int {
	for i, num := range nums {
		if i == num {
			return i
		} else if i < nums[i] {
			i = nums[i] - 1
		}
	}
	return -1
}
