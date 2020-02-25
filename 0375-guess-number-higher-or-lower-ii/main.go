package _375_guess_number_higher_or_lower_ii

func getMoneyAmount(n int) int {
	var dp = make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for length := 2; length <= n; length++ {
		for start := 1; start <= n-(length-1); start++ {
			var minRes = n * n
			for pivot := start + (length-1)/2; pivot < start+(length-1); pivot++ {
				res := pivot + max(dp[start][pivot-1], dp[pivot+1][start+length-1])
				minRes = min(res, minRes)
			}
			dp[start][start+length-1] = minRes
		}
	}
	return dp[1][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
