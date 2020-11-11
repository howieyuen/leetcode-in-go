package _508_most_frequent_subtree_sum

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findFrequentTreeSum(t *testing.T) {
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
					Val:   5,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: -3},
				},
			},
			want: []int{2, -3, 4},
		},
		{
			args: args{
				root: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 2},
					Right: &TreeNode{Val: -5},
				},
			},
			want: []int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFrequentTreeSum(tt.args.root)
			sort.Ints(got)
			sort.Ints(tt.want)
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findFrequentTreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
