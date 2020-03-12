package _140_word_break_ii

import "strings"

func wordBreak(s string, wordDict []string) []string {
	dicts := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		dicts[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dicts[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	ans := []string{}
	if !dp[len(s)] {
		return ans
	}
	dfs(s, []string{}, dicts, &ans)
	return ans
}

func dfs(s string, splits []string, dicts map[string]bool, ans *[]string) {
	if len(s) == 0 {
		*ans = append(*ans, strings.Join(splits, " "))
		return
	}
	for i := 1; i <= len(s); i++ {
		if dicts[s[:i]] {
			dfs(s[i:], append(splits, s[:i]), dicts, ans)
		}
	}
}
