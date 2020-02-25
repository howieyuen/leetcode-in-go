package _102_binary_tree_level_order_traversal

import (
	"container/list"
)

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var row []int
		size := len(queue)
		for i := 0; i < size; i++ {
			root = queue[0]
			queue = queue[1:]
			row = append(row, root.Val)
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		ans = append(ans, row)
	}
	return ans
}

/*
	bfs: 使用go lib的双向队列
*/
func levelOrder2(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	dQueue := list.New()
	dQueue.PushFront(root)
	for dQueue.Len() > 0 {
		row := []int{}
		size := dQueue.Len()
		for i := 0; i < size; i++ {
			root = dQueue.Remove(dQueue.Back()).(*TreeNode)
			row = append(row, root.Val)
			if root.Left != nil {
				dQueue.PushFront(root.Left)
			}
			if root.Right != nil {
				dQueue.PushFront(root.Right)
			}
		}
		ans = append(ans, row)
	}
	return ans
}

// dfs
func levelOrder1(root *TreeNode) [][]int {
	var ans [][]int
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == len(ans) {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], root.Val)
		if root.Left != nil {
			dfs(root.Left, level+1)
		}
		if root.Right != nil {
			dfs(root.Right, level+1)
		}
	}
	dfs(root, 0)
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
