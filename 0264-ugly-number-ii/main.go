package _264_ugly_number_ii

func nthUglyNumber(n int) int {
	var q2, q3, q5 = 0, 0, 0
	var dp = make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(min(dp[q2]*2, dp[q3]*3), dp[q5]*5)
		if dp[i] == dp[q2]*2 {
			q2++
		}
		if dp[i] == dp[q3]*3 {
			q3++
		}
		if dp[i] == dp[q5]*5 {
			q5++
		}
	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
