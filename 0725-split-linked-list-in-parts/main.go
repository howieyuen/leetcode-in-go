package _725_split_linked_list_in_parts

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(root *ListNode, k int) []*ListNode {
	var ans []*ListNode
	p := root
	n := 0
	for p != nil {
		p = p.Next
		n++
	}
	average := n / k
	additional := n % k
	pre := new(ListNode)
	pre.Next = root
	p = root
	for ; k > 0; k-- {
		head := p
		for i := 0; i < average; i++ {
			pre = p
			p = p.Next
		}

		if additional > 0 {
			pre = p
			p = p.Next
			additional--
		}
		pre.Next = nil
		ans = append(ans, head)
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}
