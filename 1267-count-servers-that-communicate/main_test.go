package _267_count_servers_that_communicate

import "testing"

func Test_countServers(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: 0,
		},
		{
			args: args{
				grid: [][]int{
					{1, 0},
					{1, 1},
				},
			},
			want: 3,
		},
		{
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countServers(tt.args.grid); got != tt.want {
				t.Errorf("countServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
