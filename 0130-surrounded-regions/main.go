package _130_surrounded_regions

func solve(board [][]byte) {
	var dfs func(i, j int)
	var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(i, j int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
			return
		}
		if board[i][j] == 'X' || board[i][j] == '#' {
			return
		}
		if board[i][j] == 'O' {
			board[i][j] = '#'
		}
		for _, dir := range dirs {
			dfs(i+dir[0], j+dir[1])
		}
	}
	
	for i := range board {
		for j := range board[i] {
			if (i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1) && board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
	
	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case '#':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			default:
			
			}
		}
	}
}
