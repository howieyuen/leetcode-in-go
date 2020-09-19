package _404_sum_of_left_leaves

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

func sumOfLeftLeaves1(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	var res int
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
			if root != nil && root.Left == nil && root.Right == nil {
				res += root.Val
			}
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = node.Right
		}
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
