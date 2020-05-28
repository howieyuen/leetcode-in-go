package _091_decode_ways

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, cur := 1, 1
	for i := 1; i < len(s); i++ {
		tmp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || s[i-1] == '2' && s[i] < '7' {
			cur = cur + pre
		}
		pre = tmp
	}
	return cur
}

func numDecodings1(s string) int {
	if len(s) == 0 {
		return 0
	}
	var dp = make([]int, len(s)+1)
	dp[0] = 1
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			dp[i+1] = 0
		}
		if s[i] != '0' {
			dp[i+1] = dp[i]
		}
		if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6')) {
			dp[i+1] += dp[i-1]
		}
	}
	return dp[len(s)]
}
