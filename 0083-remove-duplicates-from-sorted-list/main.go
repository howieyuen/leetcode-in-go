package _083_remove_duplicates_from_sorted_list

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var faker = &ListNode{}
	faker.Next = head
	var p = head
	for p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return faker.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
