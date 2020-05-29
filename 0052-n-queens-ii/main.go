package _052_n_queens_ii

func totalNQueens(n int) int {
	var ans int
	// row: 当前放置皇后的行号
	// hills: 主对角线占据情况 [1 = 被占据，0 = 未被占据]
	// nextRow: 下一行被占据的情况 [1 = 被占据，0 = 未被占据]
	// dales: 次对角线占据情况 [1 = 被占据，0 = 未被占据]
	var backtrace func(row, hill, nextRow, dale int)
	backtrace = func(row, hill, nextRow, dale int) {
		// 棋盘所有的列都可放置， 即按位表示为 n 个 '1'
		// bin(cols) = 0b1111 (n = 4), bin(cols) = 0b111 (n = 3)
		// [1 = 可放置]
		columns := (1 << uint(n)) - 1
		if row == n { // 如果已经放置了 n 个皇后
			ans++ // 累加可行解
		} else {
			// 当前行可用的列
			// ! 表示 0 和 1 的含义对于变量 hills, nextRow and dales的含义是相反的
			// [1 = 未被占据，0 = 被占据]
			freeColumns := columns & ^(hill | nextRow | dale)
			// 找到可以放置下一个皇后的列
			for freeColumns != 0 {
				// freeColumns 的第一个为 '1' 的位,在该列我们放置当前皇后
				curColumn := -freeColumns & freeColumns
				
				// 放置皇后,并且排除对应的列
				freeColumns ^= curColumn
				backtrace(row+1, (hill|curColumn)<<1, nextRow|curColumn, (dale|curColumn)>>1)
			}
		}
	}
	backtrace(0, 0, 0, 0)
	return ans
}

func totalNQueens1(n int) int {
	if n == 0 {
		return 0
	}
	cols := make([]bool, n)
	dale := make([]bool, 2*n) // 记录 '\' 方向的对角线的占用情况 col - row = constant
	hill := make([]bool, 2*n) // 记录 '/' 方向的对角线的占用情况 col + row = constant
	ans := 0
	dfs(0, cols, dale, hill, &ans)
	return ans
}

func dfs(row int, cols, dale, hill []bool, ans *int) {
	n := len(cols)
	if row == n {
		*ans++
		return
	}
	for col := 0; col < n; col++ {
		index1 := row - col + n
		index2 := 2*n - row - col - 1
		if !cols[col] && !dale[index1] && !hill[index2] {
			cols[col], dale[index1], hill[index2] = true, true, true
			dfs(row+1, cols, dale, hill, ans)
			cols[col], dale[index1], hill[index2] = false, false, false
		}
	}
}

func totalNQueens2(n int) int {
	var ans int
	var board = make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	backtrace(board, 0, &ans)
	return ans
}

func backtrace(board [][]byte, row int, ans *int) {
	if row == len(board) {
		*ans += 1
		return
	}
	for col := 0; col < len(board); col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = 'Q'
		backtrace(board, row+1, ans)
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
