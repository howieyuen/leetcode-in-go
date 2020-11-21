package _886_possible_bipartition

func possibleBipartition(N int, dislikes [][]int) bool {
	var graph = make([][]int, N+1)
	var colors = make([]int, N+1)
	for _, d := range dislikes {
		graph[d[0]] = append(graph[d[0]], d[1])
		graph[d[1]] = append(graph[d[1]], d[0])
	}
	
	var dfs func(node int, color int) bool
	dfs = func(node int, color int) bool {
		if colors[node] != 0 {
			return colors[node] == color
		}
		colors[node] = color
		for _, nei := range graph[node] {
			if !dfs(nei, -color) {
				return false
			}
		}
		return true
	}
	
	for node := 1; node <= N; node++ {
		if colors[node] == 0 && !dfs(node, 1) {
			return false
		}
	}
	return true
}

func possibleBipartition1(N int, dislikes [][]int) bool {
	var graph = make([][]int, N+1)
	var colors = make([]int, N+1)
	for _, d := range dislikes {
		graph[d[0]] = append(graph[d[0]], d[1])
		graph[d[1]] = append(graph[d[1]], d[0])
	}
	
	for node := 1; node <= N; node++ {
		if colors[node] != 0 {
			continue
		}
		colors[node] = 1
		var q = []int{node}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			for _, nei := range graph[cur] {
				if colors[nei] == colors[cur] {
					return false
				} else if colors[nei] == 0 {
					q = append(q, nei)
					colors[nei] = -colors[cur]
				}
			}
		}
	}
	return true
}

func possibleBipartition2(N int, dislikes [][]int) bool {
	var graph = make([][]int, N+1)
	for _, d := range dislikes {
		graph[d[0]] = append(graph[d[0]], d[1])
		graph[d[1]] = append(graph[d[1]], d[0])
	}
	
	uf := newUnionFind(N + 1)
	for node := 1; node <= N; node++ {
		neighbours := graph[node]
		for _, nei := range neighbours {
			if uf.connected(node, nei) {
				return false
			}
			uf.union(neighbours[0], nei)
		}
	}
	return true
}

type UnionFind struct {
	parent []int
}

func newUnionFind(n int) *UnionFind {
	var parent = make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &UnionFind{parent: parent}
}

func (uf *UnionFind) find(x int) int {
	if uf.parent[x] == x {
		return x
	}
	return uf.find(uf.parent[x])
}

func (uf *UnionFind) union(x, y int) {
	px, py := uf.find(x), uf.find(y)
	if px != py {
		uf.parent[px] = py
	}
}

func (uf *UnionFind) connected(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

var pre []int

func find(x int) int {
	if pre[x] != x {
		pre[x] = find(pre[x])
	}
	return pre[x]
}

// 并查集，主要用来确定集合关系，如果两个顶点在同一个集合，我们就可以用并查集建立它们的关系
// 反之两个顶点不在同一个集合，无法使用并查集。
// 问题的关键是根据两个顶点的关系，确定属于同一个集合的顶点，使用并查集。
// （已知条件，必然和推导条件，有关系。本题中就是，i与n+i肯定不在同一个集合）
// 最后找到与所建立的并查集相矛盾的地方。
func possibleBipartition3(N int, dislikes [][]int) bool {
	pre = make([]int, 2*N+1)
	for i := range pre {
		pre[i] = i
	}
	
	for _, dislike := range dislikes {
		u, v := dislike[0], dislike[1]
		pu, pv := find(u), find(v)
		if pu == pv {
			// 等价于u与n+u在同一组，或
			// v与n+v在同一组
			return false
		}
		// u与n+v可以在同一组
		pre[pu] = find(v + N)
		// v与u+n可以在同一组
		pre[pv] = find(u + N)
	}
	return true
}
