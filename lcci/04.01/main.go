package _4_01

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	var visited = make([]bool, n)
	var next = make([][]int, n)
	for _, e := range graph {
		next[e[0]] = append(next[e[0]], e[1])
	}
	var q = []int{start}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		visited[cur] = true
		if cur == target {
			return true
		}
		for _, nei := range next[cur] {
			if !visited[nei] {
				q = append(q, nei)
			}
		}
	}
	return false
}

func findWhetherExistsPath1(n int, graph [][]int, start int, target int) bool {
	var visited = make([]bool, n)
	var next = make([][]int, n)
	for _, e := range graph {
		next[e[0]] = append(next[e[0]], e[1])
	}
	var dfs func(cur int) bool
	dfs = func(cur int) bool {
		if cur == target {
			return true
		}
		visited[cur] = true
		for _, nei := range next[cur] {
			if visited[nei] {
				continue
			}
			if dfs(nei) {
				return true
			}
		}
		return false
	}
	return dfs(start)
}
