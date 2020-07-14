package _300_longest_increasing_subsequence

/*
dp[i]: 所有长度为i+1的递增子序列中, 最小的那个序列尾数.
由定义知dp数组必然是一个递增数组, 可以用 maxL 来表示最长递增子序列的长度.
对数组进行迭代, 依次判断每个数num将其插入dp数组相应的位置:
1. num[i] > dp[maxL], 表示num[i]比所有已知递增序列的尾数都大,
   将num[i]添加入dp数组尾部, 并将最长递增序列长度maxL++
2. dp[i-1] < num[i] <= dp[i], 只更新相应的dp[i]
*/
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
