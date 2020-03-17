package _509_fibonacci_number

func fib(N int) int {
	if N == 0 {
		return 0
	}
	dp := make([]int, N+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}

