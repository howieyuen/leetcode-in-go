package _463_island_perimeter

func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	var ans int
	var dfs func(row, col int)
	dfs = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == 0 {
			ans++
			return
		}
		if grid[row][col] == -1 {
			return
		}
		grid[row][col] = -1
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				dfs(i, j)
				break
			}
		}
	}
	return ans
}
