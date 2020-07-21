package _095_unique_binary_search_trees_ii

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateNode(1, n)
}

func generateNode(start, end int) []*TreeNode {
	var ans []*TreeNode
	if start > end {
		ans = append(ans, nil)
		return ans
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
				ans = append(ans, root)
			}
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
