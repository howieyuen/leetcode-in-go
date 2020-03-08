package _099_recover_binary_search_tree

import (
	"container/list"
)

/*
二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。
*/

var x, y, pred *TreeNode

func recoverTree(root *TreeNode) {
	findTwoSwapped(root)
	x.Val, y.Val = y.Val, x.Val
}

func findTwoSwapped(root *TreeNode) {
	if root == nil {
		return
	}
	findTwoSwapped(root.Left)
	if pred != nil && root.Val < pred.Val {
		y = root
		if x == nil {
			x = pred
		} else {
			return
		}
	}
	pred = root
	findTwoSwapped(root.Right)
}

func recoverTree1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := list.New()
	var x, y, pred *TreeNode
	for stack.Len() > 0 || root != nil {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		root = stack.Remove(stack.Back()).(*TreeNode)
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
