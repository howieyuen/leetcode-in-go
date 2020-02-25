package _328_odd_even_linked_list

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var odd = head
	var even = head.Next
	var q = even
	for q != nil && q.Next != nil {
		odd.Next = q.Next
		odd = odd.Next
		q.Next = odd.Next
		q = q.Next
	}
	odd.Next = even
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
