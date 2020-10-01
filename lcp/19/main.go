package _9

import (
	"math"
)

// ryr
// j = LeftRed or MidYellow or RightRed
// dp[i][j] 表示 [0, i] 的叶子且 leaves[i] = j 的操作次数
func minimumOperations(leaves string) int {
	const (
		LR = iota
		MY
		RR
	)
	
	n := len(leaves)
	var dp = make([][3]int, n)
	if leaves[0] == 'y' {
		dp[0][LR] = 1
	}
	dp[0][MY] = math.MaxInt32
	dp[0][RR] = math.MaxInt32
	dp[1][RR] = math.MaxInt32
	for i := 1; i < n; i++ {
		isRed := 0
		if leaves[i] == 'r' {
			isRed = 1
		}
		isYellow := 0
		if leaves[i] == 'y' {
			isYellow = 1
		}
		dp[i][LR] = dp[i-1][LR] + isYellow
		dp[i][MY] = min(dp[i-1][LR], dp[i-1][MY]) + isRed
		if i >= 2 {
			dp[i][RR] = min(dp[i-1][MY], dp[i-1][RR]) + isYellow
		}
	}
	return dp[n-1][RR]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 降维
func minimumOperations1(leaves string) int {
	isRed := func(i int) int {
		if leaves[i] == 'r' {
			return 1
		}
		return 0
	}
	
	isYellow := func(i int) int {
		if leaves[i] == 'y' {
			return 1
		}
		return 0
	}
	
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	
	leftR := isRed(0)
	midY := math.MaxInt32
	rightR := math.MaxInt32
	for i := 1; i < len(leaves); i++ {
		red, yellow := isRed(i), isYellow(i)
		rightR = min(midY, rightR) + red
		midY = min(leftR, midY) + red
		leftR = leftR + yellow
	}
	return rightR
}
