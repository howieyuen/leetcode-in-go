package _107_binary_tree_level_order_traversal_ii

import (
	"container/list"
)

func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var queue = list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var cur []int
		size := queue.Len()
		for i := 0; i < size; i++ {
			root = queue.Remove(queue.Front()).(*TreeNode)
			cur = append(cur, root.Val)
			if root.Left != nil {
				queue.PushBack(root.Left)
			}
			if root.Right != nil {
				queue.PushBack(root.Right)
			}
		}
		ans = append(ans, cur)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i],ans[j] = ans[j], ans[i]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
