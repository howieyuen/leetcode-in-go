package _5

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var exits = make(map[*Node]*Node)
	var copyNode func(head *Node) *Node
	copyNode = func(head *Node) *Node {
		if head == nil {
			return nil
		}
		if node, ok := exits[head]; ok {
			return node
		}
		node := &Node{Val: head.Val}
		exits[head] = node
		node.Next = copyNode(head.Next)
		node.Random = copyNode(head.Random)
		return node
	}
	return copyNode(head)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
