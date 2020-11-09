package _998_maximum_binary_tree_ii

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val < val {
		return &TreeNode{Val: val, Left: root}
	}
	root.Right = insertIntoMaxTree(root.Right, val)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
