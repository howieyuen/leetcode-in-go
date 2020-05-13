package _573_lcof_interview06

func reversePrint(head *ListNode) []int {
	return reverseArray(traverseList(head))
}

func traverseList(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	return ans
}

func reverseArray(ans []int) []int {
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func reversePrint1(head *ListNode) []int {
	return traverseList(reverseList(head))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	for p.Next != nil {
		q := p.Next
		p.Next = q.Next
		q.Next = head
		head = q
	}
	return head
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverseList1(head.Next)
	head.Next.Next = next
	return next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
