package _2_02

func kthToLast(head *ListNode, k int) int {
	p, q := head, head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q.Val
}

type ListNode struct {
	Val  int
	Next *ListNode
}
