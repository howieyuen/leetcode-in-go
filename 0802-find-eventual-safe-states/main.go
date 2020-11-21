package _802_find_eventual_safe_states

import (
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	outDegrees := make([]int, n)
	revGraph := make([][]int, n)
	for i := range graph {
		outDegrees[i] = len(graph[i])
		for _, j := range graph[i] {
			revGraph[j] = append(revGraph[j], i)
		}
	}
	var queue []int
	for i := range outDegrees {
		if outDegrees[i] == 0 {
			queue = append(queue, i)
		}
	}
	
	var ans []int
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		ans = append(ans, i)
		for _, j := range revGraph[i] {
			outDegrees[j]--
			if outDegrees[j] == 0 {
				queue = append(queue, j)
			}
		}
	}
	sort.Ints(ans)
	return ans
}
