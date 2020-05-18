package _8

func deleteNode(head *ListNode, val int) *ListNode {
	var faker = &ListNode{Next: head}
	p := faker
	for p.Next != nil {
		if p.Next.Val == val {
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
