package _2_04

func partition(head *ListNode, x int) *ListNode {
	p, q := head, head
	for q != nil {
		if q.Val < x {
			p.Val, q.Val = q.Val, p.Val
			p = p.Next
		}
		q = q.Next
	}
	return head
}

func partition1(head *ListNode, x int) *ListNode {
	less := new(ListNode)
	more := new(ListNode)
	p, q := less, more
	for head != nil {
		next := head.Next
		if head.Val < x {
			p.Next = head
			p = p.Next
		} else {
			q.Next = head
			q = q.Next
		}
		head.Next = nil
		head = next
	}
	p.Next = more.Next
	return less.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
