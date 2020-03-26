package _237_delete_node_in_a_linked_list

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}