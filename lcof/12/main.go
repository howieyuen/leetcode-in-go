package _2

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[i] {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i, j int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[index] {
		return false
	}
	tmp := board[i][j]
	board[i][j] = ' '
	flag := dfs(board, i-1, j, word, index+1) ||
		dfs(board, i+1, j, word, index+1) ||
		dfs(board, i, j-1, word, index+1) ||
		dfs(board, i, j+1, word, index+1)
	board[i][j] = tmp
	return flag
}
