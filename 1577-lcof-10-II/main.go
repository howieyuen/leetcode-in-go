package _577_lcof_10_II

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	var ans int
	for i := 2; i <= n; i++ {
		ans = (a + b) % (1e9 + 7)
		a = b % (1e9 + 7)
		b = ans
	}
	return ans
}
