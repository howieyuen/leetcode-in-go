package _171_remove_zero_sum_consecutive_nodes_from_linked_list

func removeZeroSumSublists(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var faker = new(ListNode)
	faker.Next = head
	for p := faker; p != nil; p = p.Next {
		sum := 0
		pre := p
		for q := p.Next; q != nil; q = q.Next {
			sum += q.Val
			if sum == 0 {
				pre.Next = q.Next
			}
		}
	}
	return faker.Next
}

// 节点i到节点j之间的和为0，那么0~i和0~j之间的的和是相等
// 所以利用哈希存储前缀和，重复的被后面的节点代替，也是删除了和为0的子串
func removeZeroSumSublists1(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	faker := &ListNode{Next: head}
	prefixSum := make(map[int]*ListNode)
	var sum int
	p := faker
	for p != nil {
		sum += p.Val
		prefixSum[sum] = p
		p = p.Next
	}
	sum = 0
	p = faker
	for p != nil {
		sum += p.Val
		p.Next = prefixSum[sum].Next
		p = p.Next
	}
	return faker.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
