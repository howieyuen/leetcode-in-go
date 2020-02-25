package _116_populating_next_right_pointers_in_each_node

func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head = root
	for head.Left != nil {
		cur := head
		for cur != nil {
			cur.Left.Next = head.Right
			if cur.Next != nil {
				cur.Right.Next = head.Next.Left
			}
			cur = cur.Next
		}
		head = head.Left
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}
