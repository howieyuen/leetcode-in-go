package _310_minimum_height_trees

import (
	"reflect"
	"testing"
)

func Test_findMinHeightTrees(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				edges: [][]int{{1, 0}, {1, 2}, {1, 3}},
				n:     4,
			},
			want: []int{1},
		},
		{
			args: args{
				edges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
				n:     6,
			},
			want: []int{3, 4},
		},
		{
			args: args{
				edges: [][]int{},
				n:     1,
			},
			want: []int{0},
		},
		{
			args: args{
				edges: [][]int{{1, 0}},
				n:     2,
			},
			want: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinHeightTrees(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findMinHeightTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
