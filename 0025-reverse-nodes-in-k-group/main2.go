package _025_reverse_nodes_in_k_group

func reverseKGroup1(head *ListNode, k int) *ListNode {
	// faker head
	var faker = &ListNode{}
	faker.Next = head
	// 待反转链表的头和尾
	var pre, post = faker, faker
	for post.Next != nil {
		// 子链表的尾节点
		for i := 0; i < k && post != nil; i++ {
			post = post.Next
		}
		// 反转结束
		if post == nil {
			break
		}
		// 待反转子链表起始节点
		var start = pre.Next
		// 下个待反转子链表起始节点
		var nextStart = post.Next
		
		// 断开待反转子链表和剩余节点
		post.Next = nil
		// 反转后的头是pre的后继
		pre.Next = reverse1(start)
		
		// start反转后是尾节点，续上断开节点
		start.Next = nextStart
		// 更新pre和post
		pre = start
		post = pre
	}
	return faker.Next
}

func reverse1(head *ListNode) *ListNode {
	var p = head
	for head.Next != nil {
		cur := head.Next
		head.Next = cur.Next
		cur.Next = p
		p = cur
	}
	return p
}
