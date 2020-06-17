package _014_best_sightseeing_pair

import (
	"math"
)

// a[i]+a[j]+i-j => (a[i]+i)+(a[j]-j)
// 统计以j为最后一个景点的最大组合时，a[j]-j是固定的，a[i]+i在变
// 遍历时，更新a[i]+i的最大值maxI，结果即为maxI+a[j]-j
func maxScoreSightseeingPair(A []int) int {
	var res = math.MinInt64
	var left = A[0]
	for i := 1; i < len(A); i++ {
		res = max(res, left+A[i]-i)
		left = max(left, A[i]+i)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
