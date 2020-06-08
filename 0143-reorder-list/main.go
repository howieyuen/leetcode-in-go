package _143_reorder_list

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := getMid(head)
	right := mid.Next
	mid.Next = nil
	right = reverse(right)
	left := head
	merge(left, right)
}

func merge(left, right *ListNode) {
	var p *ListNode
	var q *ListNode
	for left.Next != nil && right != nil {
		p = left.Next
		q = right.Next
		left.Next = right
		right.Next = p
		left = p
		right = q
	}
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func getMid(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var p = head
	var q = head
	for q != nil && q.Next != nil {
		p = p.Next
		q = q.Next.Next
	}
	return p
}

type ListNode struct {
	Val  int
	Next *ListNode
}
