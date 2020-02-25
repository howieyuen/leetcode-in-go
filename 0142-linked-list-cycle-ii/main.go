package _142_linked_list_cycle_ii

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			for fast != slow {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
	}
	return nil
}

func detectCycle1(head *ListNode) *ListNode {
	nodeExist := make(map[*ListNode]int)
	cur := head
	for cur != nil {
		if _, exist := nodeExist[cur]; exist {
			return cur
		}
		nodeExist[cur] = 1
		cur = cur.Next
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}
