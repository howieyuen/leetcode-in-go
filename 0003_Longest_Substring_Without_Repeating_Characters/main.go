package _003_Longest_Substring_Without_Repeating_Characters

func lengthOfLongestSubstring(s string) int {
	dp := [128]int{}
	for i := range dp {
		dp[i] = -1
	}
	front := 0
	maxLen := 0
	for i, char := range s {
		if dp[char] >= front {
			front = dp[char] + 1
		}
		dp[char] = i
		maxLen = maxInt(maxLen, i-front+1)
	}
	return maxLen
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
