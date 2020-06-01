package _103_binary_tree_zigzag_level_order_traversal

import (
	"container/list"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	leftToRight := true
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if leftToRight {
				level[i] = node.Val
			} else {
				level[size-i-1] = node.Val
			}
		}
		res = append(res, level)
		leftToRight = !leftToRight
		queue = queue[size:]
	}
	return res
}

func zigzagLevelOrder1(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var dequeue = list.New()
	dequeue.PushBack(root)
	leftToRight := true
	for dequeue.Len() > 0 {
		size := dequeue.Len()
		var level = make([]int, size)
		for i := 0; i < size; i++ {
			if leftToRight {
				root = dequeue.Remove(dequeue.Front()).(*TreeNode)
			} else {
				root = dequeue.Remove(dequeue.Back()).(*TreeNode)
			}
			level[i] = root.Val
			if !leftToRight {
				if root.Right != nil {
					dequeue.PushFront(root.Right)
				}
				if root.Left != nil {
					dequeue.PushFront(root.Left)
				}
			} else {
				if root.Left != nil {
					dequeue.PushBack(root.Left)
				}
				if root.Right != nil {
					dequeue.PushBack(root.Right)
				}
			}
		}
		ans = append(ans, level)
		leftToRight = !leftToRight
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
