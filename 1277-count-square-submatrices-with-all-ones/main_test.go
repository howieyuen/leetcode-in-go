package _277_count_square_submatrices_with_all_ones

import (
	"testing"
)

func Test_countSquares(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			}},
			want: 15,
		},
		{
			args: args{matrix: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			}},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSquares(tt.args.matrix); got != tt.want {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
