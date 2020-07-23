package _392_is_subsequence

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// dp[i][j]表示字符串t每个位置i的下一个要匹配的字符的位置
func isSubsequence1(s string, t string) bool {
	t = " " + t
	n := len(t)
	var dp = make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	for char := 0; char < 26; char++ {
		pos := -1
		for i := n - 1; i >= 0; i-- {
			dp[i][char] = pos
			if t[i] == byte('a'+char) {
				pos = i
			}
		}
	}
	i := 0
	for j := range s {
		if i = dp[i][s[j]-'a']; i == -1 {
			return false
		}
	}
	return true
}
