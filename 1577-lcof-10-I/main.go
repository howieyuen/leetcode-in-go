package _577_lcof_10_I

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	var ans int
	for i := 2; i <= n; i++ {
		ans = (a + b) % (1e9 + 7)
		a = b % (1e9 + 7)
		b = ans
	}
	return ans
}
