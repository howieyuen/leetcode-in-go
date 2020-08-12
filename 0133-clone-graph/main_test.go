package _133_clone_graph

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		adjust [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{args: args{adjust: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}}},
		{args: args{adjust: [][]int{{}}}},
		{args: args{adjust: [][]int{}}},
		{args: args{adjust: [][]int{{2}, {1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := buildNode(tt.args.adjust)
			tt.want = node
			if got := cloneGraph(node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildNode(adjust [][]int) *Node {
	if len(adjust) == 0 {
		return nil
	}
	var nodes = make([]*Node, len(adjust))
	for i := 0; i < len(adjust); i++ {
		nodes[i] = &Node{
			Val:       i + 1,
			Neighbors: make([]*Node, len(adjust[i])),
		}
	}
	for i := range adjust {
		for j := range adjust[i] {
			nodes[i].Neighbors[j] = nodes[adjust[i][j]-1]
		}
	}
	return nodes[0]
}
