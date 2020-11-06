package _2_03

func deleteNode(node *ListNode) {
	var pre = node
	for node.Next != nil {
		node.Val = node.Next.Val
		pre = node
		node = node.Next
	}
	pre.Next = nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
