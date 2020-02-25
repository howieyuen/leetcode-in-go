package _206_reverse_linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head
	for head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = newHead
		newHead = next
	}
	return newHead
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

// 1->2->3->4->5->NULL
// 2  1  3  4  5

type ListNode struct {
	Val  int
	Next *ListNode
}
