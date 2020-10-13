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

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var fakerHead = new(ListNode)
	fakerHead.Next = head
	pre, cur := fakerHead, head
	for cur != nil && cur.Next != nil {
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = cur
		pre = cur
		cur = cur.Next
	}
	return fakerHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
