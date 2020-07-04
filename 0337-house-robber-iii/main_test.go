package _337_house_robber_iii

import (
	"testing"
)

func Test_rob(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
					Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}},
				},
			},
			want: 7,
		},
		{
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 1},
					},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
			if got := rob1(tt.args.root); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
			}
			if got := rob2(tt.args.root); got != tt.want {
				t.Errorf("rob2() = %v, want %v", got, tt.want)
			}
		})
	}
}
