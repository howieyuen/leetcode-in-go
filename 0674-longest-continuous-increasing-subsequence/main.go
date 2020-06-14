package _674_longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var ans = 1
	var cur = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			ans = max(ans, cur)
			cur = 1
		}
	}
	ans = max(ans, cur)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
