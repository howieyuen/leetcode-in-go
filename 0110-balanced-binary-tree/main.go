package _110_balanced_binary_tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left)
	right := depth(root.Right)
	return abs(left-right) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
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
