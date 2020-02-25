package _062_unique_paths

/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。

机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。

问总共有多少条不同的路径？
*/

func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

func uniquePaths2(m int, n int) int {
	pre := make([]int, n)
	cur := make([]int, n)
	for i := range pre {
		pre[i] = 1
		cur[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 这里pre[j] = cur[j] 再优化
			cur[j] = cur[j-1] + pre[j]
		}
		copy(pre, cur)
	}
	return cur[n-1]
}

func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 根据状态转移方程可知，只需要记录当前行和上一行的数据即可
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
