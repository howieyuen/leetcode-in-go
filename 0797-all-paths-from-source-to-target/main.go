package _797_all_paths_from_source_to_target

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	var ans [][]int
	var queue = [][]int{{0}}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			next := cur[len(cur)-1]
			for _, v := range graph[next] {
				tmp := make([]int, len(cur))
				copy(tmp, cur)
				tmp = append(tmp, v)
				if v == n-1 {
					ans = append(ans, tmp)
				} else {
					queue = append(queue, tmp)
				}
			}
		}
		queue = queue[size:]
	}
	return ans
}

func allPathsSourceTarget1(graph [][]int) [][]int {
	var ans [][]int
	var dfs func(index int, cur []int)
	dfs = func(index int, cur []int) {
		if index == len(graph)-1 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}
		for _, v := range graph[index] {
			cur = append(cur, v)
			dfs(v, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, []int{0})
	return ans
}
