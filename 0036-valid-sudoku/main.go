package _036_valid_sudoku

/*
	判断9*9的数独是否有效
	1. 判断行是否有重复
	2. 判断列是否有重复
	3. 判断3*3小格是否有重复
	每轮判断后，将行，列，小格中对应位置标记已经出现的数字
*/
func isValidSudoku(board [][]byte) bool {
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	grid := [9][9]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			key := board[i][j] - '1'
			if rows[i][key] || cols[j][key] || grid[i/3*3+j/3][key] {
				return false
			}
			rows[i][key] = true
			cols[j][key] = true
			grid[i/3*3+j/3][key] = true
		}
	}
	return true
}
