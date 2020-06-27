package _2_01

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	faker := &ListNode{}
	prev := faker
	existed := make(map[int]bool)
	for head != nil {
		if existed[head.Val] {
			prev.Next = head.Next
		} else {
			existed[head.Val] = true
			prev.Next = head
			prev = head
		}
		head = head.Next
	}
	return faker.Next
}

func removeDuplicateNodes1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var exist = map[int]bool{head.Val: true}
	var p = head
	for p.Next != nil {
		if exist[p.Next.Val] {
			p.Next = p.Next.Next
		} else {
			exist[p.Next.Val] = true
			p = p.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
