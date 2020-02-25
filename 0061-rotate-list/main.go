package _061_rotate_list

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	k %= n
	if k == 0 {
		return head
	}
	k = n - k - 1
	p = head
	for k > 0 {
		p = p.Next
		k--
	}
	newHead := p.Next
	p.Next = nil
	p = newHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = head
	return newHead
}

type ListNode struct {
	Val  int
	Next *ListNode
}
