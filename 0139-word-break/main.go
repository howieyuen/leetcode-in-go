package _139_word_break

func wordBreak(s string, wordDict []string) bool {
	n:= len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	
	for i:= 0; i<= n; i++{
		for _, word := range wordDict{
			wordLen := len(word)
			if i >= wordLen && s[i-wordLen:i] == word{
				dp[i] = dp[i] || dp[i-wordLen]
			}
		}
	}
	return dp[n]
}
