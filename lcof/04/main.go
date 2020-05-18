package _4

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}

func findNumberIn2DArray1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var binarySearch func(up, down int, left, right int) bool
	binarySearch = func(up, down int, left, right int) bool {
		if up > down || left > right {
			return false
		}
		if matrix[up][left] > target || matrix[down][right] < target {
			return false
		}
		row := up
		col := left + (right-left)/2
		for row <= down && matrix[row][col] <= target {
			if matrix[row][col] == target {
				return true
			}
			row++
		}
		// 右上角和左下角
		return binarySearch(up, row-1, col+1, right) || binarySearch(row, down, left, col-1)
	}
	return binarySearch(0, len(matrix)-1, 0, len(matrix[0])-1)
}
