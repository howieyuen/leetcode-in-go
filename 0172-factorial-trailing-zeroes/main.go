package _172_factorial_trailing_zeroes

// 计算5的个数即可
func trailingZeroes(n int) int {
	var c = 0
	for n >= 5 {
		n /= 5
		c+=n
	}
	return c
}
