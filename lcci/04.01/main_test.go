package _4_01

import (
	"testing"
)

func Test_findWhetherExistsPath(t *testing.T) {
	type args struct {
		n      int
		graph  [][]int
		start  int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				n:      3,
				graph:  [][]int{{0, 1}, {0, 2}, {1, 2}, {1, 2}},
				start:  0,
				target: 2,
			},
			want: true,
		},
		{
			args: args{
				n:      5,
				graph:  [][]int{{0, 1}, {0, 2}, {0, 4}, {0, 4}, {0, 1}, {1, 3}, {1, 4}, {1, 3}, {2, 3}, {3, 4}},
				start:  0,
				target: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWhetherExistsPath(tt.args.n, tt.args.graph, tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("findWhetherExistsPath() = %v, want %v", got, tt.want)
			}
			if got := findWhetherExistsPath1(tt.args.n, tt.args.graph, tt.args.start, tt.args.target); got != tt.want {
				t.Errorf("findWhetherExistsPath1() = %v, want %v", got, tt.want)
			}
		})
	}
}
