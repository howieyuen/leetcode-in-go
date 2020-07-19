package test2

func countSubTrees(n int, edges [][]int, labels string) []int {
	relations := generateRelations(n, edges)
	var ans = make([]int, n)
	var dfs func(root int, father int) []int
	dfs = func(root int, father int) []int {
		rootCounts := make([]int, 26)
		rootCounts[labels[root]-'a']++
		
		for _, node := range relations[root] {
			if node == father {
				continue
			}
			nodeCounts := dfs(node, root)
			for i := range nodeCounts {
				rootCounts[i] += nodeCounts[i]
			}
		}
		
		ans[root] = rootCounts[labels[root]-'a']
		return rootCounts
	}
	dfs(0, -1)
	return ans
}

func generateRelations(n int, edges [][]int) [][]int {
	var relations = make([][]int, n)
	for _, edge := range edges {
		relations[edge[0]] = append(relations[edge[0]], edge[1])
		relations[edge[1]] = append(relations[edge[1]], edge[0])
	}
	return relations
}
