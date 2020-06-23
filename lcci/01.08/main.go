package _1_08

func setZeroes(matrix [][]int) {
	rowFlag, colFlag := false, false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			colFlag = true
			break
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			rowFlag = true
			break
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
	
	if colFlag {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	
	if rowFlag {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
}
