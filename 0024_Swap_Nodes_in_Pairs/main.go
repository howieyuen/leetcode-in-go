package _024_Swap_Nodes_in_Pairs

func swapPairs(head *ListNode) *ListNode {
	var pre = &ListNode{}
	pre.Next = head
	var cur = pre
	for cur.Next != nil && cur.Next.Next != nil {
		var s = cur.Next
		var e = cur.Next.Next
		cur.Next = e
		s.Next = e.Next
		e.Next = s
		cur = s
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
