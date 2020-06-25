package _139_word_break

import (
	"strings"
)

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	
	for i := 0; i <= n; i++ {
		for _, word := range wordDict {
			wordLen := len(word)
			if i >= wordLen && s[i-wordLen:i] == word {
				dp[i] = dp[i] || dp[i-wordLen]
			}
		}
	}
	return dp[n]
}

// 超时
func wordBreak1(s string, wordDict []string) bool {
	var dfs func(s string) bool
	dfs = func(s string) bool {
		if s == "" {
			return true
		}
		for _, word := range wordDict {
			if strings.Index(s, word) == 0 {
				if dfs(s[len(word):]) {
					return true
				}
			}
		}
		return false
	}
	return dfs(s)
}
