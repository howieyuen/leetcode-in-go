package _572_subtree_of_another_tree

import (
	"testing"
)

func Test_isSubtree(t *testing.T) {
	type args struct {
		s *TreeNode
		t *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: &TreeNode{
					Val:  1,
					Left: &TreeNode{Val: 1},
				},
				t: &TreeNode{Val: 1},
			},
			want: true,
		},
		{
			args: args{
				s: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:  1,
							Left: &TreeNode{Val: 0},
						},
						Right: &TreeNode{Val: 2},
					},
					Right: &TreeNode{Val: 5},
				},
				t: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubtree(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
