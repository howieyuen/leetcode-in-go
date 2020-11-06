package _290_convert_binary_number_in_a_linked_list_to_integer

func getDecimalValue(head *ListNode) int {
	var ans int
	for head != nil {
		ans = ans<<1 + head.Val
		head = head.Next
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}
