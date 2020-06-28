package _209_minimum_size_subarray_sum

func minSubArrayLen(s int, nums []int) int {
	var start, end int
	minLen := len(nums) + 1
	sum := 0
	for start < len(nums) {
		if sum < s && end < len(nums) {
			sum += nums[end]
			end++
		} else {
			sum -= nums[start]
			start++
		}
		if sum >= s && minLen > (end-start) {
			minLen = end - start
		}
	}
	if minLen > len(nums) {
		return 0
	}
	return minLen
}
