package _092_reverse_linked_list_ii

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	p := pre
	n -= m
	for m > 1 {
		p = head
		head = head.Next
		m--
	}
	newHead := head
	for n > 0 {
		tmp := head.Next
		head.Next = tmp.Next
		tmp.Next = newHead
		newHead = tmp
		p.Next = tmp
		n--
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
