package _221_maximal_square

/*
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

动态规划：
dp[i][j]代表，从(0,0)位置到(i,j)位置的区域的最大正方形的边长
从matrix[i][j]=1开始，正方形能否扩展与左边，上边，和左上角三个位置最大正方形边长有关
因此dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
以下第一种情况正方形可扩展，后三种虽然扩展位置为1，但另外三个位置不满足扩张要求，因此取三者最小值。
1 1 | 1 0 | 0 1 | 1 1
1 1 | 1 1 | 1 1 | 0 1
*/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxEdge := 0
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = 1 + min(dp[i][j], min(dp[i+1][j], dp[i][j+1]))
				maxEdge = max(maxEdge, dp[i+1][j+1])
			}
		}
	}
	return maxEdge * maxEdge
}

func maximalSquare1(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxEdge := 0
	rows, cols := len(matrix), len(matrix[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				maxEdge = max(maxEdge, 1)
				maxSide := min(rows-i, cols-j)
				for k := 1; k < maxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for m := 0; m < k; m++ {
						if matrix[i+m][j+k] == '0' || matrix[i+k][j+m] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxEdge = max(maxEdge, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxEdge * maxEdge
}

/*
采用位运算的方式，将每行看做是一个二进制数，依次相与
与运算结果不为0，二进制序列中至少有一个1，即高度至少为1
宽度即为云运算结果中连续的1个数
取高度和宽度的最小值，即为正方形的边长
*/
func maximalSquare2(matrix [][]byte) int {
	dp := make([]uint, len(matrix))
	// 每行转换成二进制
	for i := range matrix {
		for j := range matrix[i] {
			dp[i] = dp[i]<<1 + uint(matrix[i][j]-'0')
		}
	}
	maxEdge := 0
	// 从第i行开始，依次与第i+1行到第n行做与运算
	for i := 0; i < len(dp)-1; i++ {
		andValue := dp[i]
		for j := i + 1; j < len(dp); j++ {
			// 待查找矩形区域的最大高度
			maxHeight := j - i + 1
			andValue &= dp[j]
			// 计算最大宽度
			maxWidth := 0
			for tmp := andValue; tmp > 0; tmp &= tmp << 1 {
				maxWidth++
			}
			maxEdge = max(maxEdge, min(maxWidth, maxHeight))
		}
	}
	return maxEdge*maxEdge
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
