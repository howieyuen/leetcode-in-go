package _148_sort_list

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	fast = slow.Next
	slow.Next = nil
	slow = head
	
	left := sortList(slow)
	right := sortList(fast)
	faker := &ListNode{}
	p := faker
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}
	if left != nil {
		p.Next = left
	}
	if right != nil {
		p.Next = right
	}
	return faker.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
