package _300_longest_increasing_subsequence

// d[i],表示长度为i的最长上升子序列的末尾元素的最小值
func lengthOfLIS(nums []int) int {
	var maxL = 0
	var dp = make([]int, len(nums))
	for i := range nums {
		low, high := 0, maxL
		for low < high {
			mid := (low + high) >> 1
			if dp[mid] < nums[i] {
				low = mid + 1
			} else {
				high = mid
			}
		}
		dp[low] = nums[i]
		if low == maxL {
			maxL++
		}
	}
	return maxL
}

func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp = make([]int, len(nums))
	dp[0] = 1
	var ans = 1
	for i := 1; i < len(nums); i++ {
		var maxVal = 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxVal = max(maxVal, dp[j])
			}
		}
		dp[i] = maxVal + 1
		ans = max(ans, dp[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
