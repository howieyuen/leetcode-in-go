package _027_longest_arithmetic_subsequence

func longestArithSeqLength(A []int) int {
	n := len(A)
	var dp = make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	var res int
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			d := A[i] - A[j]
			if dp[j][d] == 0 {
				dp[i][d] = 2
			} else {
				dp[i][d] = max(dp[i][d], dp[j][d]+1)
			}
			res = max(res, dp[i][d])
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
