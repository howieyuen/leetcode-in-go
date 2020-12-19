package _446_arithmetic_slices_ii_subsequence

import (
	"math"
)

// 函数要返回数组 A 中所有等差子序列的个数。
func numberOfArithmeticSlices(A []int) int {
	n := len(A)
	if n <= 2 {
		return 0
	}
	// dp[i][d] ，代表以 A[i] 结束且公差为 d 的弱等差数列个数
	// 弱等差数列是至少有两个元素的子序列，任意两个连续元素的差相等
	// 任意2个元素都可以组成弱等差数列
	var dp = make([]map[int]int, n)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	var res int
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			d := A[i] - A[j]
			if d < math.MinInt32 || d > math.MaxInt32 {
				continue
			}
			// 将A[i]加入公差为d、以A[j]结尾的弱等差序列
			// 原有的弱等差数列个数dp[j][d]和{A[j],A[i]}新组建的弱等差数列
			dp[i][d] += dp[j][d] + 1
			// dp[j][d]是原有弱等差数列的个数
			// 现在加入A[i]，则弱等差数列均成为等差数列
			res += dp[j][d]
		}
	}
	return res
}
