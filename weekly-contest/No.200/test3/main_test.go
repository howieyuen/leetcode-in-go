package test3

import (
	"testing"
)

func Test_minSwaps(t *testing.T) {
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
				{0, 0, 1},
				{1, 1, 0},
				{1, 0, 0},
			}},
			want: 3,
		},
		{
			args: args{grid: [][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 1},
			}},
			want: 0,
		},
		{
			args: args{grid: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
				{0, 1, 1, 0},
			}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.grid); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
