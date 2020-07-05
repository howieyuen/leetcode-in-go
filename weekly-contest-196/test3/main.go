package test3

import (
	"math"
)

// 参考 1277.统计全为 1 的正方形子矩阵
// dp[i][j]表示矩阵i行，到(i,j)这个位置，最多有多少个连续的1
// 对于每个点(i,j)，我们固定子矩形的右下角为(i,j)，利用dp从该行i向上寻找子矩阵左上角为第k行的矩阵个数
// 每次将子矩阵个数加到答案中即可。
func numSubmat(mat [][]int) int {
	rows, cols := len(mat), len(mat[0])
	var dp = make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, cols)
	}
	for i := 0; i < rows; i++ {
		serialOnes := 0
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				serialOnes++
			} else {
				serialOnes = 0
			}
			dp[i][j] = serialOnes
		}
	}
	var ans int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			minX := math.MaxInt64
			for k := i; k >= 0; k-- {
				minX = min(minX, dp[k][j])
				ans += minX
			}
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
