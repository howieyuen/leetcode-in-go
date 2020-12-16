package _474_ones_and_zeroes

func findMaxForm(strs []string, m int, n int) int {
	var dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, v := range strs {
		ones := 0
		zeros := 0
		for j := range v {
			if v[j] == '1' {
				ones++
			} else {
				zeros++
			}
		}
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
