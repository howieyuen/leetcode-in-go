package _413_arithmetic_slices

// 函数要返回数组 A 中所有为等差数组的子数组个数。
func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n <= 2 {
		return 0
	}
	var res, add int
	for i := 2; i < n; i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			add++
			res += add
		} else {
			add = 0
		}
	}
	return res
}
