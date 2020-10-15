package _116_populating_next_right_pointers_in_each_node

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var levelHead = root
	for levelHead.Left != nil {
		p := levelHead
		for p != nil {
			p.Left.Next = p.Right
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
			p = p.Next
		}
		levelHead = levelHead.Left
	}
	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	// 判断是否是叶子节点
	if root.Left != nil {
		root.Left.Next = root.Right
		// 判断节点是否为同一层的最后一个，如果是，那么该节点的右子节点也是最后一个
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		connect(root.Left)
		connect(root.Right)
	}
	return root
}

func connect1(root *Node) *Node {
	var dfs func(cur, next *Node)
	dfs = func(cur, next *Node) {
		if cur == nil {
			return
		}
		cur.Next = next
		dfs(cur.Left, cur.Right)
		if cur.Next == nil {
			dfs(cur.Right, nil)
		} else {
			dfs(cur.Right, cur.Next.Left)
		}
	}
	dfs(root, nil)
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
