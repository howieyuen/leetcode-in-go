package _222_count_complete_tree_nodes

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	if left == right {
		return countNodes(root.Right) + 1<<left
	}
	return countNodes(root.Left) + 1<<right
}

func getDepth(root *TreeNode) int {
	var depth int
	for root != nil {
		root = root.Left
		depth++
	}
	return depth
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
