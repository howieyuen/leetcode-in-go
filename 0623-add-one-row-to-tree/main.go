package _623_add_one_row_to_tree

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	var depth = 1
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		if depth+1 == d {
			for i := 0; i < size; i++ {
				left := queue[i].Left
				right := queue[i].Right
				queue[i].Left = &TreeNode{Val: v, Left: left}
				queue[i].Right = &TreeNode{Val: v, Right: right}
			}
			break
		} else {
			for i := 0; i < size; i++ {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			depth++
			queue = queue[size:]
		}
	}
	if depth+1 == d {
		return root
	}
	return &TreeNode{Val: v, Left: root}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
