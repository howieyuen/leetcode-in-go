package _074_search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}
	cols := len(matrix[0])
	left, right := 0, rows*cols-1
	for left <= right {
		mid := (left + right) >> 1
		r, c := mid/cols, mid%cols
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
