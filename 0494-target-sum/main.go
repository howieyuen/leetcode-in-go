package _494_target_sum

/*
	数组非空，且长度不会超过20。
	初始的数组的和不会超过1000。 [-1000,1000] => [0,2000]
	保证返回的最终结果能被32位整数存下。
*/

func findTargetSumWays(nums []int, S int) int {
	count := 0
	calculate(nums, 0, 0, S, &count)
	return count
}

func calculate(nums []int, index int, sum, target int, count *int) {
	if index == len(nums) {
		if sum == target {
			*count++
		}
		return
	}
	calculate(nums, index+1, sum+nums[index], target, count)
	calculate(nums, index+1, sum-nums[index], target, count)
}

func findTargetSumWays1(nums []int, S int) int {
	// dp[i][j] 表示用数组中的前 i 个元素，组成和为 j 的方案数
	var dp = make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2001)
	}
	// 由于数组中所有数的和不超过 1000，那么 j 的最小值可以达到 -1000
	// 给 dp[i][j] 的第二维预先增加 1000，满足下标索引
	dp[0][nums[0]+1000] = 1
	dp[0][-nums[0]+1000] += 1
	for i := 1; i < len(nums); i++ {
		for sum := -1000; sum <= 1000; sum++ {
			if dp[i-1][sum+1000] > 0 {
				dp[i][sum+nums[i]+1000] += dp[i-1][sum+1000]
				dp[i][sum-nums[i]+1000] += dp[i-1][sum+1000]

			}
		}
	}
	var ans = 0
	if S <= 1000 {
		ans = dp[len(nums)-1][S+1000]
	}
	return ans
}

// 动态规划的状态转移方程中，dp[i][...] 只和 dp[i - 1][...] 有关
func findTargetSumWays2(nums []int, S int) int {
	var dp = make([]int, 2001)
	dp[nums[0]+1000] = 1
	dp[-nums[0]+1000] += 1
	for i := 1; i < len(nums); i++ {
		var next = make([]int, 2001)
		for sum := -1000; sum <= 1000; sum++ {
			if dp[sum+1000] > 0 {
				next[sum+nums[i]+1000] += dp[sum+1000]
				next[sum-nums[i]+1000] += dp[sum+1000]
			}
		}
		dp = next
	}
	var ans = 0
	if S <= 1000 {
		ans = dp[S+1000]
	}
	return ans
}
