package test4

// 从0~n，遍历样本数
// 动态规划：假设已知dp[i]，求dp[i+1]
// 本题是完全平方数，因此是已知dp[i]，可以推算dp[i+j*j]
// 先拿着的人赢return true，也就是dp[i]是true， dp[i+j*j]就是false，反之亦然
func winnerSquareGame(n int) bool {
	var dp = make([]bool, n+1)
	for i := 0; i <= n; i++ {
		for j := 1; j*j+i <= n; j++ {
			if !dp[i] {
				dp[i+j*j] = true
			}
		}
	}
	return dp[n]
}
