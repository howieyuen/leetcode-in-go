package _5

/*
see https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var t = &ListNode{}
	var p = t
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return t.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
