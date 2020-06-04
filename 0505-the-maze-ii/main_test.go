package _505_the_maze_ii

import (
	"testing"
)

func Test_shortestDistance(t *testing.T) {
	type args struct {
		maze        [][]int
		start       []int
		destination []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				maze: [][]int{
					{0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0},
					{1, 1, 0, 1, 1},
					{0, 0, 0, 0, 0},
				},
				start:       []int{0, 4},
				destination: []int{4, 4},
			},
			want: 12,
		},
		{
			args: args{
				maze: [][]int{
					{0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0},
					{0, 0, 0, 1, 0},
					{1, 1, 0, 1, 1},
					{0, 0, 0, 0, 0},
				},
				start:       []int{0, 4},
				destination: []int{3, 2},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistance(tt.args.maze, tt.args.start, tt.args.destination); got != tt.want {
				t.Errorf("shortestDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
