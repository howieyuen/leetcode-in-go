package _160_intersection_of_two_linked_lists

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	len1 := getLength(headA)
	len2 := getLength(headB)
	if len1 < len2 {
		headA, headB = headB, headA
		len1, len2 = len2, len1
	}
	dis := len1 - len2
	for dis > 0 {
		headA = headA.Next
		dis--
	}
	for headA != nil && headB != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

func getLength(head *ListNode) int {
	n := 0
	for head != nil {
		n++
		head = head.Next
	}
	return n
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

type ListNode struct {
	Val  int
	Next *ListNode
}
