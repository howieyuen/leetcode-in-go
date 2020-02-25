package _073_set_matrix_zeroes

/*
用matrix第一行和第一列记录该行该列是否有0,作为标志位

但是对于第一行和第一列要设置一个标志位,为了防止自己这一行(一列)也有0的情况
*/
func setZeroes(matrix [][]int) {
	isCol := false
	rows := len(matrix)
	cols := len(matrix[0])
	for i := 0; i < rows; i++ {
		if matrix[i][0] == 0 {
			isCol = true
		}
		for j := 1; j < cols; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 1; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if isCol {
			matrix[i][0] = 0
		}
	}
}
