package _138_copy_list_with_random_pointer

func copyRandomList(head *Node) *Node {
	visited := make(map[*Node]*Node)
	return recursive(head, visited)
}

func recursive(head *Node, visited map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}
	if node, ok := visited[head]; ok {
		return node
	}
	node := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	visited[head] = node
	node.Next = recursive(head.Next, visited)
	node.Random = recursive(head.Random, visited)
	return node
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
