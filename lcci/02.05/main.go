package _2_05

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	
	var pre = new(ListNode)
	p := pre
	carry := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: carry % 10}
		carry = carry / 10
		p = p.Next
	}
	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}
	return pre.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
