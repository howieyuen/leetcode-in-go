package _673_number_of_longest_increasing_subsequence

func findNumberOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var length = make([]int, len(nums))
	var count = make([]int, len(nums))
	var maxLen = 1
	for i := 0; i < len(nums); i++ {
		length[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if length[j]+1 > length[i] {
					length[i] = length[j] + 1
					count[i] = count[j]
				} else if length[j]+1 == length[i] {
					count[i] += count[j]
				}
			}
		}
		maxLen = max(maxLen, length[i])
	}
	var ans int
	for i := range length {
		if length[i] == maxLen {
			ans += count[i]
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
