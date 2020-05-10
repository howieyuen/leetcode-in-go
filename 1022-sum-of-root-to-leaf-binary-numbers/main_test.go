package _022_sum_of_root_to_leaf_binary_numbers

import (
	"testing"
)

func Test_sumRootToLeaf(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 1},
				},
			}},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumRootToLeaf(tt.args.root); got != tt.want {
				t.Errorf("sumRootToLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
