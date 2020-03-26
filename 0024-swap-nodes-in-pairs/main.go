package _024_Swap_Nodes_in_Pairs

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first := head
	second := head.Next
	first.Next = swapPairs(second.Next)
	second.Next = first
	return second
}

func swapPairs1(head *ListNode) *ListNode {
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
