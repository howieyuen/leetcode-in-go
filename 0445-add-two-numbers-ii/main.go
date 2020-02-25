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
