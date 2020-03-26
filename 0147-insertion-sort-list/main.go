package _147_insertion_sort_list

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	faker := &ListNode{}
	p := head.Next
	faker.Next = head
	head.Next = nil
	for p != nil {
		prev := faker
		next := p.Next
		
		for prev.Next != nil && prev.Next.Val <= p.Val {
			prev = prev.Next
		}
		
		p.Next = prev.Next
		prev.Next = p
		p = next
	}
	return faker.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}