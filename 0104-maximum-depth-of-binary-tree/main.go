package _104_maximum_depth_of_binary_tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
