package _404_sum_of_left_leaves

import (
	"strconv"
	"strings"
	"testing"
)

func Test_sumOfLeftLeaves(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{root: buildTree("3,9,20,null,null,15,7")},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfLeftLeaves(tt.args.root); got != tt.want {
				t.Errorf("sumOfLeftLeaves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTree(val string) *TreeNode {
	if len(val) == 0 {
		return nil
	}
	values := strings.Split(val, ",")
	rootVal, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: rootVal}
	var queue []*TreeNode
	queue = append(queue, root)
	
	var buildLeaf func(cur *TreeNode, val string, left bool)
	buildLeaf = func(cur *TreeNode, val string, left bool) {
		if val == "null" {
			if left {
				cur.Left = nil
			} else {
				cur.Right = nil
			}
		} else {
			v, _ := strconv.Atoi(val)
			node := &TreeNode{Val: v}
			if left {
				cur.Left = node
			} else {
				cur.Right = node
			}
			queue = append(queue, node)
		}
	}
	
	for i := 1; i < len(values);{
		cur := queue[0]
		buildLeaf(cur, values[i], i%2 == 1)
		i++
		buildLeaf(cur, values[i], i%2 == 1)
		i++
		queue = queue[1:]
	}
	return root
}
