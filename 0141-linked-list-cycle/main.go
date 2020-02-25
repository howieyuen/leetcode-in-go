package _141_linked_list_cycle

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle1(head *ListNode) bool {
	hash := make(map[*ListNode]int)
	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}
		hash[head] = head.Val
		head = head.Next
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
