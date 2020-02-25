package _025_reverse_nodes_in_k_group

func reverseKGroup(head *ListNode, k int) *ListNode {
	var tmp = &ListNode{}
	tmp.Next = head
	var pre = tmp  // 待反转链表首节点的前驱
	var post = tmp // 待反转链表尾节点
	for post.Next != nil {
		for i := 0; i < k && post != nil; i++ { // 移动尾节点
			post = post.Next
		}
		if post == nil {
			break
		}
		var start = pre.Next      // 记录开始反转节点
		var end = post.Next       // 记录未反转链表首节点
		post.Next = nil           // 断开待反转和未反转
		pre.Next = reverse(start) // 反转后的首节点是pre的后继
		start.Next = end          // start节点经过反转成为尾节点，与未反转链表首节点连接
		pre = start               // 记录下次反转的首节点
		post = pre                // 记录下次反转的尾节点的初始位置，向后移动K步
	}
	return tmp.Next
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode = nil
	var cur = head
	for cur != nil {
		var next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
