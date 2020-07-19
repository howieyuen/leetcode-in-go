package test2

import (
	"reflect"
	"testing"
)

func Test_countSubTrees(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		labels string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				n:      4,
				edges:  [][]int{{0, 2}, {0, 3}, {1, 2}},
				labels: "aeed",
			},
			want: []int{1, 1, 2, 1},
		},
		{
			args: args{
				n:      7,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}},
				labels: "abaedcd",
			},
			want: []int{2, 1, 1, 1, 1, 1, 1},
		},
		{
			args: args{
				n:      4,
				edges:  [][]int{{0, 1}, {1, 2}, {0, 3}},
				labels: "bbbb",
			},
			want: []int{4, 2, 1, 1},
		},
		{
			args: args{
				n:      5,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}},
				labels: "aabab",
			},
			want: []int{3, 2, 1, 1, 1},
		},
		{
			args: args{
				n:      6,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {3, 4}, {4, 5}},
				labels: "cbabaa",
			},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			args: args{
				n:      7,
				edges:  [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}},
				labels: "aaabaaa",
			},
			want: []int{6, 5, 4, 1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubTrees(tt.args.n, tt.args.edges, tt.args.labels); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSubTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
