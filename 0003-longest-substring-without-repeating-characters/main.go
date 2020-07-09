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

func lengthOfLongestSubstring1(s string) int {
	var left, right = 0, 0
	var ans int
	var windows = make(map[byte]int)
	for right < len(s) {
		c := s[right]
		// 无脑窗口扩大
		right++
		// 窗口数据更新
		windows[c]++
		// 窗口过大，尝试左侧收缩
		for windows[c] > 1 {
			d := s[left]
			// 左侧无脑收缩
			left++
			// 窗口数据更新，可能跳出循环
			windows[d]--
		}
		ans = maxInt(ans, right-left)
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
