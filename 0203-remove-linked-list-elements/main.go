package _203_remove_linked_list_elements

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	var faker = &ListNode{}
	p := faker
	for head != nil {
		if head.Val != val {
			p.Next = head
			p = p.Next
		}
		head = head.Next
	}
	p.Next = nil
	return faker.Next
}

func removeElements1(head *ListNode, val int) *ListNode {
	prev := &ListNode{
		Next: head,
	}
	for p := prev; p.Next != nil; {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return prev.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
