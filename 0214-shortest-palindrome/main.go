package _214_shortest_palindrome

// 暴力破解
func shortestPalindrome1(s string) string {
	n := len(s)
	rev := reverse(s)
	for i := 0; i < n; i++ {
		if s[:n-i] == rev[i:] {
			return rev[:i] + s
		}
	}
	return ""
}

func reverse(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

// 双指针+递归，减少每次检查字符串的长度
func shortestPalindrome2(s string) string {
	n := len(s)
	i := 0
	for j := n - 1; j >= 0; j-- {
		if s[i] == s[j] {
			i++
		}
	}
	// 本身是回文串
	if i == n {
		return s
	}
	rev := reverse(s[i:])
	return rev + shortestPalindrome2(s[:i]) + s[i:]
}

// KMP算法
func shortestPalindrome3(s string) string {
	rev := reverse(s)
	n := len(s)
	newS := s + "#" + rev
	next := make([]int, len(newS))
	for i := 1; i < len(newS); i++ {
		j := next[i-1]
		for j > 0 && newS[i] != newS[j] {
			j = next[j-1]
		}
		if newS[i] == newS[j] {
			j++
		}
		next[i] = j
	}
	return rev[:n-next[len(newS)-1]] + s
}
