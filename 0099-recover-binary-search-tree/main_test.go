package _099_recover_binary_search_tree

import (
	"testing"
)

var root1 = &TreeNode{
	Val:  3,
	Left: &TreeNode{Val: 1},
	Right: &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 2},
		Right: nil,
	},
}

var root2 = &TreeNode{
	Val:  2,
	Left: &TreeNode{Val: 1},
	Right: &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 3},
		Right: nil,
	},
}

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{root: root1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
			println(deepEquals(tt.args.root, root2))
		})
	}
}

func deepEquals(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	} else if root1 != nil && root2 != nil {
		if root1.Val != root2.Val {
			return false
		} else {
			return deepEquals(root1.Left, root2.Left) && deepEquals(root1.Right, root2.Right)
		}
	} else {
		return false
	}
}
