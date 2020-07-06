package _980_unique_paths_iii

func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 待走的空格总数，用于终止条件判断
	var steps int
	// 起点坐标
	var row, col int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				row, col = i, j
			} else if grid[i][j] == 0 {
				steps++
			}
		}
	}
	// 回溯(当前坐标，已经走过的步数)
	var bfs func(row, col int, curSteps int)
	// 合法路径数
	var ans int
	var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bfs = func(row, col int, curSteps int) {
		// 越界判断
		if row < 0 || row >= m || col < 0 || col >= n {
			return
		}
		// 重复走过的坐标
		if grid[row][col] == -1 || grid[row][col] == 1 {
			return
		}
		// 走到终点，并且已经走过的步数=待走的步数，路径数+1
		if grid[row][col] == 2 {
			if curSteps == steps {
				ans++
			}
			return
		}
		// 添加回溯标记
		grid[row][col] = -1
		for i := range dirs {
			// 上下左右深搜
			bfs(row+dirs[i][0], col+dirs[i][1], curSteps+1)
		}
		// 去除回溯标记
		grid[row][col] = 0
	}
	for i := range dirs {
		// 从起点的上下左右四个方向开始搜索
		bfs(row+dirs[i][0], col+dirs[i][1], 0)
	}
	return ans
}
