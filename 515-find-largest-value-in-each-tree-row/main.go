package _15_find_largest_value_in_each_tree_row

import (
	"math"
)

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var level []int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		max := math.MinInt32
		for i := 0; i < size; i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		level = append(level, max)
		queue = queue[size:]
	}
	return level
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
