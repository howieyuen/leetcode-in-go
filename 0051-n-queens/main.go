package _051_n_queens

func solveNQueens(n int) [][]string {
	var ans [][]string
	var board = make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	dfs(board, 0, &ans)
	return ans
}

func dfs(board [][]byte, row int, ans *[][]string) {
	if row == len(board) {
		var tmp = make([]string, row)
		for i := range board {
			tmp[i] = string(board[i])
		}
		*ans = append(*ans, tmp)
		return
	}
	for col := 0; col < len(board); col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = 'Q'
		dfs(board, row+1, ans)
		board[row][col] = '.'
	}
}

func isValid(board [][]byte, row, col int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func solveNQueens1(n int) [][]string {
	var board = make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	
	var admit = func(row, col int) bool {
		for i := row; i >= 0; i-- {
			if board[i][col] == 'Q' {
				return false
			}
		}
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}
		return true
	}
	
	var ans [][]string
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			tmp := make([]string, n)
			for i := range board {
				tmp[i] = string(board[i])
			}
			ans = append(ans, tmp)
			return
		}
		for j := 0; j < n; j++ {
			if !admit(row, j) {
				continue
			}
			board[row][j] = 'Q'
			dfs(row + 1)
			board[row][j] = '.'
		}
	}
	
	dfs(0)
	return ans
}
