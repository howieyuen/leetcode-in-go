package _647_palindromic_substrings

// 一维动态规划
func countSubstrings2(s string) int {
	var ans int
	n := len(s)
	var dp = make([]bool, n)
	for j := 0; j < n; j++ {
		dp[j] = true
		ans++
		// 对照二维动态规划，dp[i][j]=true, 要求s[i]==s[j]&&dp[i+1][j-1]
		// dp[i]表示，s[i,j]是否为回文
		// 下一轮遍历j++，因此dp[i+1]取的是上一轮的结果，即dp[i+1][j-1]，
		for i := 0; i < j; i++ {
			if s[i] == s[j] && dp[i+1] {
				dp[i] = true
				ans++
			} else {
				dp[i] = false
			}
		}
	}
	return ans
}

// 二维动态规划
func countSubstrings1(s string) int {
	var ans int
	n := len(s)
	var dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				ans++
			}
		}
	}
	return ans
}

// 中心扩展
func countSubstrings(s string) int {
	var ans int
	var count func(left, right int)
	count = func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			ans++
			left--
			right++
		}
	}
	for i := range s {
		count(i, i)
		count(i, i+1)
	}
	return ans
}
