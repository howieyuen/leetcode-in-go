package _329_longest_increasing_path_in_a_matrix

func longestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	if cols == 0 {
		return 0
	}
	var visited = make([][]int, rows)
	for i := range visited {
		visited[i] = make([]int, cols)
	}
	var ans int
	var directions = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if visited[i][j] != 0 {
			return visited[i][j]
		}
		visited[i][j]++
		for _, dir := range directions {
			row, col := i+dir[0], j+dir[1]
			if row < 0 || row >= rows || col < 0 || col >= cols {
				continue
			}
			if matrix[row][col] <= matrix[i][j] {
				continue
			}
			visited[i][j] = max(visited[i][j], 1+dfs(row, col))
		}
		return visited[i][j]
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
