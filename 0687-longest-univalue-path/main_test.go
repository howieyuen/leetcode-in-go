package _687_longest_univalue_path

import (
	"testing"
)

func Test_longestUnivaluePath(t *testing.T) {
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
					Val:   4,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 5},
				},
			}},
			want: 2,
		},
		{
			args: args{root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 1},
				},
				Right: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 5},
				},
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestUnivaluePath(tt.args.root); got != tt.want {
				t.Errorf("longestUnivaluePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
