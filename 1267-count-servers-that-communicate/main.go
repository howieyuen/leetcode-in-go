package _267_count_servers_that_communicate

func countServers(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	var countRow, countCol = make([]int, rows), make([]int, cols)
	for i := range grid {
		for j := range grid[i] {
			countRow[i] += grid[i][j]
			countCol[j] += grid[i][j]
		}
	}

	var ans int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 注意避免重复计算
			if grid[i][j] == 1 && (countRow[i] > 1 || countCol[j] > 1) {
				ans++
			}
		}
	}
	return ans
}
