package _445_add_two_numbers_ii

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := generateStack(l1)
	stack2 := generateStack(l2)
	i, j := len(stack1)-1, len(stack2)-1
	var head = &ListNode{}
	p := head
	carry := 0
	for i >= 0 && j >= 0 {
		sum := stack1[i] + stack2[j] + carry
		p.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		p = p.Next
		carry = sum / 10
		i--
		j--
	}
	for i >= 0 {
		sum := stack1[i] + carry
		p.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		p = p.Next
		carry = sum / 10
		i--
	}
	for j >= 0 {
		sum := stack2[j] + carry
		p.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		p = p.Next
		carry = sum / 10
		j--
	}
	if carry == 1 {
		p.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	p = reverse(head.Next)
	return p
}

func reverse(head *ListNode) *ListNode {
	p := head
	for head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = p
		p = next
	}
	return p
}

func generateStack(head *ListNode) []int {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	return stack
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := reverse(l1)
	h2 := reverse(l2)
	var temp *ListNode
	var newHead *ListNode
	up := 0
	for {
		if h1 == nil && h2 == nil && up == 0 {
			break
		}
		newNode, newUp := add(h1, h2, up)
		up = newUp
		if temp != nil {
			temp.Next = newNode
		} else {
			newHead = newNode
		}
		temp = newNode
		if h1 != nil {
			h1 = h1.Next
		}
		if h2 != nil {
			h2 = h2.Next
		}
	}
	return reverse(newHead)
}

func add(l1 *ListNode, l2 *ListNode, up int) (*ListNode, int) {
	var res int
	var newUp int
	if l1 == nil && l2 == nil {
		res = up
		newUp = 0
	} else if l1 == nil {
		res = (l2.Val + up) % 10
		newUp = (l2.Val + up) / 10
	} else if l2 == nil {
		res = (l1.Val + up) % 10
		newUp = (l1.Val + up) / 10
	} else {
		res = (l1.Val + l2.Val + up) % 10
		newUp = (l1.Val + l2.Val + up) / 10
	}
	return &ListNode{res, nil}, newUp
}

func reverse(head *ListNode) *ListNode {
	var front *ListNode
	var next *ListNode
	for {
		if head == nil {
			break
		}
		next = head.Next
		head.Next = front
		front = head
		head = next
	}
	return front
}
*/