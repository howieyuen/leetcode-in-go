package _201_bitwise_and_of_numbers_range

func rangeBitwiseAnd(m int, n int) int {
	var shift int
	for m < n {
		m >>= 1
		n >>= 1
		shift++
	}
	return n << shift
}

// 返回按位与，实际上就是返回公共前缀
func rangeBitwiseAnd1(m int, n int) int {
	for m < n {
		n = n & (n - 1)
	}
	return m & n
}
