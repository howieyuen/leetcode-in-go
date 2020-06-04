package _115_distinct_subsequences

func numDistinct(s string, t string) int {
	var dp = make([]int, len(t)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		for j := len(t) - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[j+1] += dp[j]
			}
		}
	}
	return dp[len(t)]
}

func numDistinct2(s string, t string) int {
	var dp = make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
		dp[i][0] = 1
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(t); j++ {
			dp[i][j] += dp[i-1][j]
			if s[i-1] == t[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[len(s)][len(t)]
}

func numDistinct1(s string, t string) int {
	var count int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if j == len(t) {
			count++
			return
		}
		if i == len(s) {
			return
		}
		dfs(i+1, j)
		if s[i] == t[j] {
			dfs(i+1, j+1)
		}
	}
	dfs(0, 0)
	return count
}
