package test3

import (
	"testing"
)

func Test_numSubmat(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{mat: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			}},
			want: 13,
		},
		{
			args: args{mat: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 1},
				{1, 1, 1, 0},
			}},
			want: 24,
		},
		{
			args: args{mat: [][]int{
				{1, 1, 1, 1, 1, 1},
			}},
			want: 21,
		},
		{
			args: args{mat: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
			}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubmat(tt.args.mat); got != tt.want {
				t.Errorf("numSubmat() = %v, want %v", got, tt.want)
			}
		})
	}
}
