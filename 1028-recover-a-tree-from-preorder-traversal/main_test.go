package _028_recover_a_tree_from_preorder_traversal

import (
	"testing"
)

func Test_recoverFromPreorder(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{S: "1-2--3--4-5--6--7"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
		},
		{
			args: args{S: "1-2--3---4-5--6---7"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  3,
						Left: &TreeNode{Val: 4},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 7},
					},
				},
			},
		},
		{
			args: args{S: "1-401--349---90--88"},
			want: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 401,
					Left: &TreeNode{
						Val:  349,
						Left: &TreeNode{Val: 90},
					},
					Right: &TreeNode{Val: 88},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recoverFromPreorder(tt.args.S); !deepEqual(got, tt.want) {
				t.Errorf("recoverFromPreorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deepEqual(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Val == root2.Val && deepEqual(root1.Left, root2.Left) && deepEqual(root1.Right, root2.Right)
}
