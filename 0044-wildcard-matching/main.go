package _044_wildcard_matching

/*
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。

dp[i][j]表示，s[:i+1]和p[:j+1]的匹配关系
显然，
情况1：s[i]==p[j]，则dp[i][j] = dp[i-1][j-1]
情况2：p[j]=='?'，和情况1一样
情况3：p[j]=='*'，'*'可以匹配0个，也可以匹配多个；匹配0个，即dp[i][j] = dp[i][j-1]；匹配多个，dp[i][j] = dp[i-1][j]
*/
func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	var dp = make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] && p[i-1] == '*'
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i] == p[j] || p[j] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j] == '*' {
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}
