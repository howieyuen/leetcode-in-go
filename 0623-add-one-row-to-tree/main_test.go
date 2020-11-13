package _623_add_one_row_to_tree

import (
	"reflect"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root *TreeNode
		v    int
		d    int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 1},
					},
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 5},
						Right: nil,
					},
				},
				v: 1,
				d: 2,
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  &TreeNode{Val: 3},
						Right: &TreeNode{Val: 1},
					},
					Right: nil,
				},
				Right: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   6,
						Left:  &TreeNode{Val: 5},
						Right: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addOneRow(tt.args.root, tt.args.v, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
