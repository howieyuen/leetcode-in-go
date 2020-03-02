package _095_unique_binary_search_trees_ii

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateNode(1, n)
}

func generateNode(start, end int) []*TreeNode {
	trees := []*TreeNode{}
	if start > end {
		trees = append(trees, nil)
		return trees
	}
	for i := start; i <= end; i++ {
		leftSubTrees := generateNode(start, i-1)
		rightSubTrees := generateNode(i+1, end)
		for _, left := range leftSubTrees {
			for _, right := range rightSubTrees {
				root := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				trees = append(trees, root)
			}
		}
	}
	return trees
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
