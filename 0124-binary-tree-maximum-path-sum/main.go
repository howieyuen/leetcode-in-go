package _124_binary_tree_maximum_path_sum

import (
	"math"
)

func maxPathSum(root *TreeNode) int {
	var maxPath = math.MinInt64
	var bfs func(root *TreeNode) int
	// bfs return root.Val for no children's TreeNode
	bfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// Val may be negative
		left := max(bfs(root.Left), 0)
		right := max(bfs(root.Right), 0)
		curSum := left + right + root.Val
		maxPath = max(maxPath, curSum)
		return root.Val + max(left, right)
	}
	bfs(root)
	return maxPath
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
