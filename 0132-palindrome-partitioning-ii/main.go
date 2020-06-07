package _132_palindrome_partitioning_ii

/*
当切割次数最少使得切割后的所有字符串都是回文时，也正是这些回文子串最长的时候，那么如果说能找到以每个字符为中心的最长回文串
中心扩展法
dp[i]表示，s[0:i]的最小切割次数
*/
func minCut(s string) int {
	var dp = make([]int, len(s))
	for i := range dp {
		dp[i] = i
	}
	for i := range s {
		extend(s, i, i, dp)
		extend(s, i, i+1, dp)
	}
	return dp[len(s)-1]
}

func extend(s string, left, right int, dp []int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		if left == 0 {
			dp[right] = 0
		} else {
			dp[right] = min(dp[right], dp[left-1]+1)
		}
		left--
		right++
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCut1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var dp = make([]int, len(s)+1)
	dp[0] = -1
	for right := 1; right <= len(s); right++ {
		dp[right] = right - 1
		for left := 0; left < right; left++ {
			if isPalindrome(s[left:right]) {
				dp[right] = min(dp[right], dp[left]+1)
			}
		}
	}
	return dp[len(s)]
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
