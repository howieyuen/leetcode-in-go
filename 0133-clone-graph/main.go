package _133_clone_graph

import (
	"container/list"
)

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := visited[node]; ok {
			return v
		}
		clone := &Node{
			Val:       node.Val,
			Neighbors: make([]*Node, len(node.Neighbors)),
		}
		visited[node] = clone
		for i, v := range node.Neighbors {
			clone.Neighbors[i] = dfs(v)
		}
		return clone
	}
	return dfs(node)
}

func cloneGraph1(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(node)
	visited := make(map[*Node]*Node)
	visited[node] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(*Node)
		for i, v := range node.Neighbors {
			if _, ok := visited[v]; !ok {
				queue.PushBack(v)
				visited[v] = &Node{
					Val:       v.Val,
					Neighbors: make([]*Node, len(v.Neighbors)),
				}
			}
			visited[node].Neighbors[i] = visited[v]
		}
	}
	return visited[node]
}

type Node struct {
	Val       int
	Neighbors []*Node
}
