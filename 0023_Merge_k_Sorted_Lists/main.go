package _023_Merge_k_Sorted_Lists

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return binaryMerge(lists, 0, len(lists)-1)

}

func binaryMerge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	} else {
		mid := (left + right) / 2
		p1 := binaryMerge(lists, left, mid)
		p2 := binaryMerge(lists, mid+1, right)
		return mergeTwoLists(p1, p2)
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var t = &ListNode{}
	var p = t
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return t.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
