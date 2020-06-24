package _289_game_of_life

// 1表示存活，0表示死亡
// 用int二进制位表示当前状态和下次状态，可节省空间
// 低位0/1表示当前存活状态，高位0/1表示下次存活状态
func gameOfLife(board [][]int) {
	var dx = []int{0, 0, 1, -1, 1, 1, -1, -1}
	var dy = []int{1, -1, 0, 0, 1, -1, 1, -1}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			live := 0
			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				live += board[x][y] & 1
			}
			if board[i][j]&1 == 1 {
				if live >= 2 && live <= 3 {
					board[i][j] = 0b11
				}
			} else if live == 3 {
				board[i][j] = 0b10
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1
		}
	}
}
