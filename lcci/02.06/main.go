package _2_06

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	for slow.Next != nil {
		next := slow.Next
		slow.Next = next.Next
		next.Next = slow
		slow = next
	}
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
