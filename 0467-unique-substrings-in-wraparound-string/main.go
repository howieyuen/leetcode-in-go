package _467_unique_substrings_in_wraparound_string

func findSubstringInWraproundString(p string) int {
	var dp = make([]int, 26)
	var cur int
	for i := 0; i < len(p); i++ {
		if i > 0 && (p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25) {
			cur++
		} else {
			cur = 1
		}
		dp[p[i]-'a'] = max(dp[p[i]-'a'], cur)
	}
	var res int
	for _, v := range dp {
		res += v
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
