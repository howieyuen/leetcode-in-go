package _547_friend_circles

import (
	"testing"
)

func Test_findCircleNum(t *testing.T) {
	type args struct {
		M [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{M: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			}},
			want: 1,
		},
		{
			args: args{M: [][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			}},
			want: 2,
		},
		{
			args: args{M: [][]int{
				{1, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNum(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum() = %v, want %v", got, tt.want)
			}
			if got := findCircleNum1(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum1() = %v, want %v", got, tt.want)
			}
			if got := findCircleNum2(tt.args.M); got != tt.want {
				t.Errorf("findCircleNum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
