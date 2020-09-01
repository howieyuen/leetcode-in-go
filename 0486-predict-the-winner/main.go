package _486_predict_the_winner

// 玩家1因为先手优势，在总数为偶数情况下，一定会赢；奇数情况，最后一个取数的一定是玩家1
// dp[i][j] 表示当数组[i, j]，当前玩家与另一个玩家的分数之差的最大值
// 如果区间只有一个元素，即[i,i]，对于玩家1来说，因为先手取数，即：
//     dp[i][i] = nums[i]
// 如果区间元素大于1，即j-i>1，当前取数的玩家的最优解，比较两端取数的结果，取最优，即：
//     max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	if n%2 == 0 {
		return true
	}
	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
