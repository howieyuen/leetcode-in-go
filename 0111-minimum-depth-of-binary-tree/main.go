package _111_minimum_depth_of_binary_tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d1 := minDepth(root.Left)
	d2 := minDepth(root.Right)
	if root.Left == nil {
		return 1 + d2
	} else if root.Right == nil {
		return 1 + d1
	} else {
		return 1 + min(d1, d2)
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var depth int
	var size int
	for len(queue) > 0 {
		size = len(queue)
		depth++
		for i := 0; i < size; i++ {
			root := queue[i]
			if root.Left == nil && root.Right == nil {
				return depth
			}
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		queue = queue[size:]
	}
	return depth
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
