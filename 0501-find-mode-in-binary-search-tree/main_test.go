package _501_find_mode_in_binary_search_tree

import (
	"reflect"
	"testing"
)

func Test_findMode(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 2}},
					Right: nil,
				},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMode(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
