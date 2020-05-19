package _054_spiral_matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	rows, cols := len(matrix), len(matrix[0])
	ans := make([]int, rows*cols)
	left, right := 0, cols-1
	up, down := 0, rows-1
	index := 0
	for {
		// left to right
		for j := left; j <= right; j++ {
			ans[index] = matrix[up][j]
			index++
		}
		if up++; up > down {
			break
		}
		
		// up to down
		for i := up; i <= down; i++ {
			ans[index] = matrix[i][right]
			index++
		}
		if right--; left > right {
			break
		}
		
		// right to left
		for j := right; j >= left; j-- {
			ans[index] = matrix[down][j]
			index++
		}
		if down--; up > down {
			break
		}
		
		// down to up
		for i := down; i >= up; i-- {
			ans[index] = matrix[i][left]
			index++
		}
		if left++; left > right {
			break
		}
	}
	return ans
}
