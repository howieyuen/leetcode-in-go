package _094_binary_tree_inorder_traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	var ans []int
	inOrder(root, &ans)
	return ans
}

func inOrder(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inOrder(root.Left, ans)
	}
	*ans = append(*ans, root.Val)
	if root.Right != nil {
		inOrder(root.Right, ans)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
