package _037_sudoku_solver

/*
解数独
按照行列顺序，依次尝试填写数字，如果能正常进行，则顺利到结束；如果无法进行，就回退到上次，重新填写数字
1. row=9,col=9,完成数独
2. board[i][j]='.'，需要填写数字，回溯法依次尝试填写数字

*/
func solveSudoku(board [][]byte) {
	rows := [9][10]bool{}
	cols := [9][10]bool{}
	boxes := [3][3][10]bool{}
	for i := 0; i < 9; i++ {
		for j, num := range board[i] {
			if num == '.' {
				continue
			}
			num -= '0'
			rows[i][num] = true
			cols[j][num] = true
			boxes[i/3][j/3][num] = true
		}
	}
	recursiveSolveSudoku(board, rows, cols, boxes, 0, 0)
}

func recursiveSolveSudoku(board [][]byte, rows, cols [9][10]bool, boxes [3][3][10]bool, row, col int) bool {
	if col == len(board[0]) {
		row++
		if row == len(board) {
			return true
		}
		col = 0
	}
	if board[row][col] == '.' {
		for num := 1; num < 10; num++ {
			used := rows[row][num] || cols[col][num] || boxes[row/3][col/3][num]
			if !used {
				rows[row][num] = true
				cols[col][num] = true
				boxes[row/3][col/3][num] = true
				board[row][col] = byte(num + '0')
				if recursiveSolveSudoku(board, rows, cols, boxes, row, col+1) {
					return true
				}
				rows[row][num] = false
				cols[col][num] = false
				boxes[row/3][col/3][num] = false
				board[row][col] = '.'
			}
		}
	} else {
		return recursiveSolveSudoku(board, rows, cols, boxes, row, col+1)
	}
	return false
}
