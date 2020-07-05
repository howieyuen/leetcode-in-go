package _277_count_square_submatrices_with_all_ones

// 计算全1正方形
// 以每个点(i,j)作为正方形的右下角、左下角、右上角，计算总和
func countSquares(matrix [][]int) int {
	var ans int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				continue
			} else if i == 0 || j == 0 {
				ans++
			} else {
				matrix[i][j] = min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1])) + 1
				ans += matrix[i][j]
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
