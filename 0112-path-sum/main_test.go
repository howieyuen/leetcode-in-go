package _112_path_sum

import (
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
		sum  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val:   11,
							Left:  &TreeNode{Val: 7},
							Right: &TreeNode{Val: 2},
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:  8,
						Left: &TreeNode{Val: 13},
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: &TreeNode{Val: 1},
						},
					},
				},
				sum: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
			if got := hasPathSum1(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("hasPathSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}
