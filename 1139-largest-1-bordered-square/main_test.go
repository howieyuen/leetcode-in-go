package _139_largest_1_bordered_square

import (
	"testing"
)

func Test_largest1BorderedSquare(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			}},
			want: 9,
		},
		{
			args: args{grid: [][]int{
				{1, 1},
				{1, 0},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largest1BorderedSquare(tt.args.grid); got != tt.want {
				t.Errorf("largest1BorderedSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
