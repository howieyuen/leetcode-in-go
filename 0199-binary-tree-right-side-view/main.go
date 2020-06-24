package _199_binary_tree_right_side_view

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			root = queue[i]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		res = append(res, queue[size-1].Val)
		queue = queue[size:]
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
