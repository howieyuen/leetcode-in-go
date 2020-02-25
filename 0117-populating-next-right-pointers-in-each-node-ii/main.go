package _117_populating_next_right_pointers_in_each_node_ii

func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		dummy := &TreeNode{}
		head := dummy
		for cur != nil {
			if cur.Left != nil {
				head.Next = cur.Left
				head = head.Next
			}
			if cur.Right != nil {
				head.Next = cur.Right
				head = head.Next
			}
			cur = cur.Next
		}
		cur = dummy.Next
		dummy.Next = nil
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}
