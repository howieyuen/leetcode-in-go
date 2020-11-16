package _310_minimum_height_trees

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	
	// 度，无向图不分出度和入度
	var degrees = make([]int, n)
	var graph = make(map[int][]int)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
		degrees[e[0]]++
		degrees[e[1]]++
	}
	// 叶子节点
	var queue []int
	for i := range degrees {
		if degrees[i] == 1 {
			queue = append(queue, i)
		}
	}
	
	var ans []int
	for len(queue) > 0 {
		ans = []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			ans = append(ans, node)
			neighbours := graph[node]
			for _, neighbour := range neighbours {
				degrees[neighbour]--
				if degrees[neighbour] == 1 {
					queue = append(queue, neighbour)
				}
			}
		}
		queue = queue[size:]
	}
	return ans
}

func findMinHeightTrees1(n int, edges [][]int) []int {
	var graph = make([][]bool, n)
	for i := range graph {
		graph[i] = make([]bool, n)
	}
	
	var degrees = make([]int, n)
	
	for _, edge := range edges {
		graph[edge[0]][edge[1]] = true
		graph[edge[1]][edge[0]] = true
		degrees[edge[0]]++
		degrees[edge[1]]++
	}
	
	var visited = make([]bool, n)
	for n > 2 {
		var queue []int
		for i := range degrees {
			if degrees[i] == 1 {
				queue = append(queue, i)
			}
		}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			degrees[node]--
			n--
			visited[node] = true
			for i := 0; i < len(graph[node]); i++ {
				if graph[node][i] {
					graph[node][i] = false
					graph[i][node] = false
					degrees[i]--
				}
			}
		}
	}
	
	var ans []int
	for i := range visited {
		if !visited[i] {
			ans = append(ans, i)
		}
	}
	return ans
}
