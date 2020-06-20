package _162_as_far_from_land_as_possible

import (
	"testing"
)

func Test_maxDistance(t *testing.T) {
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
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			}},
			want: 2,
		},
		{
			args: args{grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistance(tt.args.grid); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
