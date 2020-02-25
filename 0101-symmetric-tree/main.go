package _101_symmetric_tree

import (
	"container/list"
)

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
}

func isSymmetric1(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(root)
	for queue.Len() > 0 {
		root1 := queue.Remove(queue.Front()).(*TreeNode)
		root2 := queue.Remove(queue.Front()).(*TreeNode)
		if root1 == nil && root2 == nil {
			continue
		}
		if root1 == nil || root2 == nil {
			return false
		}
		if root1.Val != root2.Val {
			return false
		}
		queue.PushBack(root1.Left)
		queue.PushBack(root2.Right)
		queue.PushBack(root1.Right)
		queue.PushBack(root2.Left)
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
