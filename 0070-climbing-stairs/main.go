package _070_climbing_stairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	pre, next := 1, 2
	for i := 3; i <= n; i++ {
		pre, next = next, pre+next
	}
	return next
}
