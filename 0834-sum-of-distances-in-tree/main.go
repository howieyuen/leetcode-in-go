package _834_sum_of_distances_in_tree

func sumOfDistancesInTree(N int, edges [][]int) []int {
	// graph表示无向图的连通关系
	var graph = make([][]int, N)
	for i := range edges {
		node1, node2 := edges[i][0], edges[i][1]
		graph[node1] = append(graph[node1], node2)
		graph[node2] = append(graph[node2], node1)
	}
	
	// disSum[i]表示节点i到其他所有节点之和
	var disSum = make([]int, N)
	// nodeNum[i]表示与节点i作为根的子树的节点个数，包括节点i本身
	var nodeNum = make([]int, N)
	
	// 自底向上递归处理每个节点的子树
	var postOrder func(cur, parent int)
	postOrder = func(cur, parent int) {
		// 初始化nodeSum[cur] = 1
		nodeNum[cur] = 1
		neighbors := graph[cur]
		for _, neighbor := range neighbors {
			if neighbor == parent {
				continue
			}
			postOrder(neighbor, cur)
			nodeNum[cur] += nodeNum[neighbor]
			disSum[cur] += nodeNum[neighbor] + disSum[neighbor]
		}
	}
	
	// 自上向下，更新每个节点作为根时的disSum
	var preOrder func(cur, parent int)
	preOrder = func(cur, parent int) {
		neighbors := graph[cur]
		for _, neighbor := range neighbors {
			if neighbor == parent {
				continue
			}
			// disSum[neighbor]是以0为根，现在要更新成cur为根
			// 子树内：原先neighbor子树贡献的距离nodeNum[neighbor]，现在不需要
			// 子树外：不在子树内的节点，都需要额外再走一步
			disSum[neighbor] = disSum[cur] - nodeNum[neighbor] + (N - nodeNum[neighbor])
			preOrder(neighbor, cur)
		}
	}
	
	postOrder(0, -1)
	preOrder(0, -1)
	return disSum
}
