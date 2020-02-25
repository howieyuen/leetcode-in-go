package main

func isMatch(s string, p string) bool {
	n := len(s) + 1
	m := len(p) + 1
	var dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, m)
	}
	dp[0][0] = true
	for j := 2; j < m; j++ {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' { // 当前比较字符相等或者pattern为'.'
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' { // 当前模式为'*'，首先判断是否可以借助前一位匹配
				if p[j-2] == s[i-1] || p[j-2] == '.' { // 可以借助前一位
					dp[i][j] = dp[i-1][j] || // * 匹配多个前一位字符
						dp[i][j-1] || // * 匹配1个前一位字符
						dp[i][j-2] // * 匹配0个前一位字符
				} else { // 不可以借助，前一位连同'*'丢弃
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[n-1][m-1]
}
