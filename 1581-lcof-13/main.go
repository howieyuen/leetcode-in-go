package _581_lcof_13

func movingCount(m int, n int, k int) int {
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	return bfs(0, 0, visited, k)
}

func bfs(i, j int, visited [][]bool, k int) int {
	if i < 0 || i >= len(visited) || j < 0 || j >= len(visited[0]) {
		return 0
	}
	if visited[i][j] {
		return 0
	}
	visited[i][j] = true
	if numerical(i)+numerical(j) > k {
		return 0
	}
	return 1 + bfs(i-1, j, visited, k) + bfs(i+1, j, visited, k) + bfs(i, j-1, visited, k) + bfs(i, j+1, visited, k)
}

func numerical(n int) int {
	if n < 10 {
		return n
	}
	c := 0
	for n > 0 {
		c += n % 10
		n /= 10
	}
	return c
}
