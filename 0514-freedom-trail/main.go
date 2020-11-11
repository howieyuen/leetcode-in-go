package _514_freedom_trail

import (
	"math"
)

func findRotateSteps(ring string, key string) int {
	const inf = math.MaxInt32
	n, m := len(ring), len(key)
	pos := [26][]int{}
	// 保存字符对应的下标数组
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	// dp[i][j]表示key[i]，ring[j]转到12:00方向的最小步数
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	// 计算初始位置的最小步数
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = min(p, n-p) + 1
	}
	for i := 1; i < m; i++ { // 从key[1]开始
		for _, j := range pos[key[i]-'a'] { // key[i]对应的下标数组
			for _, k := range pos[key[i-1]-'a'] { // key[i-1]对应的下标数组
				dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
			}
		}
	}
	return min(dp[m-1]...)
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
