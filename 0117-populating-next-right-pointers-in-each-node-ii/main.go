package _117_populating_next_right_pointers_in_each_node_ii

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		dummy := &Node{}
		head := dummy
		for cur != nil {
			if cur.Left != nil {
				head.Next = cur.Left
				head = head.Next
			}
			if cur.Right != nil {
				head.Next = cur.Right
				head = head.Next
			}
			cur = cur.Next
		}
		cur = dummy.Next
		dummy.Next = nil
	}
	return root
}

func connect1(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, pre *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
