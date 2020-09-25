package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				inorder:   []int{15, 9, 10, 3, 20, 5, 7, 8, 4},
				postorder: []int{15, 10, 9, 5, 4, 8, 7, 20, 3},
			},
			want: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   9,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 10},
				},
				Right: &TreeNode{
					Val: 20,
					Right: &TreeNode{
						Val:   7,
						Left:  &TreeNode{Val: 5},
						Right: &TreeNode{Val: 8, Right: &TreeNode{Val: 4}},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree1(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
			if got := buildTree1(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree1() = %v, want %v", got, tt.want)
			}
		})
	}
}
