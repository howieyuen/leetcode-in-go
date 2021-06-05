package _019_Remove_Nth_Node_From_End_of_List

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var t = &ListNode{}
	t.Next = head
	var p = t
	var q = t
	for i := 0; i <= n; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return t.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
