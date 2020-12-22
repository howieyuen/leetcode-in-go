package _516_longest_palindromic_subsequence

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	// dp[i][j]表示s[i:j]的回文子序列的最大长度
	// s[i]==s[j], dp[i][j] = dp[i+1][j-1]+2
	// s[i]!=s[j], dp[i][j] = max(dp[i+1][j],dp[i][j-1])
	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
