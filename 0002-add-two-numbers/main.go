package _002_Add_Two_Numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		sum := n1 + n2 + carry
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}
	return head.Next
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var faker = &ListNode{Val: -1}
	p := faker
	var carry int
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		carry = val / 10
		val %= 10
		p.Next = &ListNode{Val: val}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + carry
		carry = val / 10
		val %= 10
		p.Next = &ListNode{Val: val}
		p = p.Next
		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + carry
		carry = val / 10
		val %= 10
		p.Next = &ListNode{Val: val}
		p = p.Next
		l2 = l2.Next
	}
	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}
	return faker.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
