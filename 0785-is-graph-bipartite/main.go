package _785_is_graph_bipartite

// dfs
func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	var ans = true
	var dfs func(node int, color int)
	dfs = func(node int, color int) {
		colors[node] = color
		cNei := 1
		if color == cNei {
			cNei = 2
		}
		for _, j := range graph[node] {
			if colors[j] == 0 {
				dfs(j, cNei)
				if !ans {
					return
				}
			} else if colors[j] != cNei {
				ans = false
				return
			}
		}
	}
	for i := 0; i < len(graph); i++ {
		if colors[i] == 0 {
			dfs(i, 1)
		}
	}
	return ans
}

// bfs
func isBipartite1(graph [][]int) bool {
	n := len(graph)
	colors := make([]int, n)
	for i := 0; i < n; i++ {
		if colors[i] == 0 {
			var queue []int
			queue = append(queue, i)
			colors[i] = 1
			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]
				cNei := 1
				if cNei == colors[node] {
					cNei = 2
				}
				for _, neighbor := range graph[node] {
					if colors[neighbor] == 0 {
						queue = append(queue, neighbor)
						colors[neighbor] = cNei
					} else if colors[neighbor] != cNei {
						return false
					}
				}
			}
		}
	}
	return true
}
