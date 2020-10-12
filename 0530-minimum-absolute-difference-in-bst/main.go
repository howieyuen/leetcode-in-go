package _530_minimum_absolute_difference_in_bst

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	var ans, pre = math.MaxInt32, -1
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		if pre != -1 && root.Val-pre < ans {
			ans = root.Val - pre
		}
		pre = root.Val
		inOrder(root.Right)
	}
	inOrder(root)
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
