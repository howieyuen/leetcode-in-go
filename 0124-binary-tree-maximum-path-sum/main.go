package _124_binary_tree_maximum_path_sum

import (
	"math"
)

/*
	从一个节点到另一个节点的路径最大值
*/
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	maxGain(root, &maxSum)
	return maxSum
}

/*
	递归中更新最长路径
	maxGain的返回值，是以当前节点为根出发，左子树和右子树的最大路径和
*/
func maxGain(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}
	left := max(maxGain(root.Left, maxSum), 0)
	right := max(maxGain(root.Right, maxSum), 0)
	newPath := left + right + root.Val
	*maxSum = max(*maxSum, newPath)
	return root.Val + max(left, right)
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
