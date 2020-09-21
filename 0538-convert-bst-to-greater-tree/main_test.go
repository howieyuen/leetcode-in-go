package _538_convert_bst_to_greater_tree

import (
	"reflect"
	"testing"
)

func Test_convertBST(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{input: []int{4, 1, 6, 0, 2, 5, 7, -1, -1, -1, 3, -1, -1, -1, 8}},
			want: []int{30, 36, 21, 36, 35, 26, 15, -1, -1, -1, 33, -1, -1, -1, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := convertBST(buildTree(tt.args.input))
			order := levelOrder(root)[:len(tt.args.input)]
			if !reflect.DeepEqual(order, tt.want) {
				t.Errorf("convertBST() = %v, want %v", order, tt.want)
			}
		})
	}
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	var ans []int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i] == nil {
				ans = append(ans, -1)
			} else {
				ans = append(ans, queue[i].Val)
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
	}
	return ans
}

func buildTree(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}
	root := &TreeNode{Val: input[0]}
	var queue []*TreeNode
	queue = append(queue, root)
	for i := 1; i < len(input); {
		cur := queue[0]
		if input[i] != -1 {
			cur.Left = &TreeNode{Val: input[i]}
			queue = append(queue, cur.Left)
		}
		i++
		if i >= len(input) {
			break
		}
		if input[i] != -1 {
			cur.Right = &TreeNode{Val: input[i]}
			queue = append(queue, cur.Right)
		}
		i++
		queue = queue[1:]
	}
	return root
}
