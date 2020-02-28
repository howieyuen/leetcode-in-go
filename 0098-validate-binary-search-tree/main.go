package _098_validate_binary_search_tree

import (
	"math"
)

/*
	验证二叉搜索树
*/
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	min := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Val <= min {
				return false
			} else {
				min = root.Val
			}
			root = root.Right
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
