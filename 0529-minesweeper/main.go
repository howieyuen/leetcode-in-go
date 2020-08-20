package _529_minesweeper

func updateBoard(board [][]byte, click []int) [][]byte {
	rows, cols := len(board), len(board[0])
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rows || y < 0 || y >= cols {
			return
		}
		switch board[x][y] {
		case 'X': // 挖出的地雷
			return
		case 'B': // 没有地雷相邻的空方块
			return
		case 'E': // 空方块
			mineNum := 0
			for _, dir := range directions {
				a, b := x+dir[0], y+dir[1]
				if a < 0 || a >= rows || b < 0 || b >= cols {
					continue
				}
				if board[a][b] == 'M' {
					mineNum++
				}
			}
			if mineNum != 0 {
				board[x][y] = byte('0' + mineNum)
				return
			}
			board[x][y] = 'B'
			for _, dir := range directions {
				dfs(x+dir[0], y+dir[1])
			}
		case 'M': // 未挖的地雷
			board[x][y] = 'X'
		}
	}
	dfs(click[0], click[1])
	return board
}
