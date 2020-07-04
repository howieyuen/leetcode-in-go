package main

/*
中心扩展算法
*/
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	start, end := 1, 0
	for i := 0; i < len(s); i++ {
		len1 := extend(s, i, i)
		len2 := extend(s, i, i+1)
		len3 := max(len1, len2)
		if len3 > end-start {
			start = i - (len3-1)/2
			end = i + len3/2
		}
	}
	return s[start : end+1]
}

func extend(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

/*
动态规划
*/
func longestPalindrome2(s string) string {
	res := ""
	n := len(s)
	f := make([][]bool, n)
	for i := 0; i < n; i++ {
		f[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 1 || f[i+1][j-1]) {
				f[i][j] = true
				if j-i >= len(res) {
					res = s[i : j+1]
				}
			}
		}
	}
	return res
}

/*
Manacher 算法
*/
func longestPalindrome3(s string) string {
	t := rebuildString(s)
	n := len(t)
	var p = make([]int, n)
	c, r := 0, 0
	for i := 1; i < n-1; i++ {
		mirrorI := 2*c - i
		if r > i {
			p[i] = min(r-i, p[mirrorI])
		} else {
			p[i] = 0
		}
		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i]++
		}
		if i+p[i] > r {
			c = i
			r = c + p[i]
		}
	}
	maxLen := 0
	centerIndex := 0
	for i := 1; i < n-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}
	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

func rebuildString(s string) string {
	if len(s) == 0 {
		return "^$"
	}
	ret := "^"
	for i := 0; i < len(s); i++ {
		ret += "#" + string(s[i])
	}
	ret += "#$"
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
