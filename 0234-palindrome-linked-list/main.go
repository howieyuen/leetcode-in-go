package _234_palindrome_linked_list

func isPalindrome(head *ListNode) bool {
	arr := make([]int, 0)
	cur := head
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}

	left := 0
	right := len(arr) - 1
	for left <= right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var firstHalfEnd = getEndOfFirstHalf(head)
	var secondHalfStart = reverse(firstHalfEnd.Next)

	var p1 = head
	var p2 = secondHalfStart
	res := true
	for res && p2 != nil {
		if p1.Val != p2.Val {
			res = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	firstHalfEnd.Next = reverse(secondHalfStart)
	return res
}

func getEndOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	var curr = head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	
	head2 := func(head *ListNode) *ListNode {
		// get mid
		var p, q = head, head
		for q.Next != nil && q.Next.Next != nil {
			p = p.Next
			q = q.Next.Next
		}
		// reverse
		var pre *ListNode
		var cur = p
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}(head)
	
	head1 := head
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
