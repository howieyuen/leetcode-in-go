package _968_binary_tree_cameras

import (
	"testing"
)

func Test_minCameraCover(t *testing.T) {
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
				Val: 0,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
			}},
			want: 1,
		},
		{
			args:
			args{root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 0,
						Left: &TreeNode{
							Val:   0,
							Right: &TreeNode{Val: 0},
						},
					},
				},
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCameraCover(tt.args.root); got != tt.want {
				t.Errorf("minCameraCover() = %v, want %v", got, tt.want)
			}
		})
	}
}
