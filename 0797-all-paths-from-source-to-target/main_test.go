package _797_all_paths_from_source_to_target

import (
	"reflect"
	"testing"
)

func Test_allPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				graph: [][]int{{1, 2}, {3}, {3}, {}},
			},
			want: [][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			args: args{
				graph: [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			},
			want: [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}},
		},
		{
			args: args{
				graph: [][]int{{1}, {}},
			},
			want: [][]int{{0, 1}},
		},
		{
			args: args{
				graph: [][]int{{1, 2, 3}, {2}, {3}, {}},
			},
			want: [][]int{{0, 1, 2, 3}, {0, 2, 3}, {0, 3}},
		},
		{
			args: args{
				graph: [][]int{{1, 3}, {2}, {3}, {}},
			},
			want: [][]int{{0, 1, 2, 3}, {0, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget1() = %v, want %v", got, tt.want)
			}
			if got := allPathsSourceTarget1(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPathsSourceTarget1() = %v, want %v", got, tt.want)
			}
		})
	}
}
