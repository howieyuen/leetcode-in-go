package _6

import (
	"testing"
)

func Test_isSubStructure(t *testing.T) {
	type args struct {
		a *TreeNode
		b *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				a: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 2},
					},
					Right: &TreeNode{Val: 5},
				},
				b: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 1},
				},
			},
			want: true,
		},
		{
			args: args{
				a: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   4,
							Left:  &TreeNode{Val: 8},
							Right: &TreeNode{Val: 9},
						},
						Right: &TreeNode{Val: 5},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  &TreeNode{Val: 6},
						Right: &TreeNode{Val: 7},
					},
				},
				b: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 9},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
