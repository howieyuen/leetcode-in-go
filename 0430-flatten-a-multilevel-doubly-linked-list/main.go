package _430_flatten_a_multilevel_doubly_linked_list

func flatten(head *Node) *Node {
	if head == nil {
		return head
	}
	faker := &Node{
		Val:   0,
		Pre:   nil,
		Child: nil,
		Next:  head,
	}
	dfs(faker, head)
	faker.Next.Pre = nil
	return faker.Next
}

func dfs(pre, cur *Node) *Node {
	if cur == nil {
		return cur
	}
	cur.Pre = pre
	pre.Next = cur
	tmp := cur.Next
	tail := dfs(cur, cur.Child)
	cur.Child = nil
	return dfs(tail, tmp)
}

func flatten1(head *Node) *Node {
	if head == nil {
		return nil
	}
	faker := &Node{
		Val:   0,
		Pre:   nil,
		Child: nil,
		Next:  head,
	}
	var pre, cur *Node = faker, nil
	stack := make([]*Node, 0)
	index := 0
	stack = append(stack, head)
	index++
	for index > 0 {
		cur = stack[index-1]
		index--
		pre.Next = cur
		cur.Pre = pre
		pre = cur
		if cur.Next != nil {
			stack = append(stack, cur.Next)
			index++
		}
		if cur.Child != nil {
			stack = append(stack, cur.Child)
			index++
			cur.Child = nil
		}
		pre = cur
	}
	faker.Next.Pre = nil
	return faker.Next
}

type Node struct {
	Val   int
	Pre   *Node
	Child *Node
	Next  *Node
}
