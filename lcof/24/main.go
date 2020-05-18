package _4

/*
see https://leetcode-cn.com/problems/reverse-linked-list/
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var p = head
	for head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = p
		p = next
	}
	return p
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

type ListNode struct {
	Val  int
	Next *ListNode
}
