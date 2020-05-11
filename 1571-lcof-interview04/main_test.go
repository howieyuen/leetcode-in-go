package _571_lcof_interview04

import (
	"testing"
)

func Test_findNumberIn2DArray(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	var matrix = [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{matrix: matrix, target: 5,}, want: true},
		{args: args{matrix: matrix, target: 20,}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArray(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArray() = %v, want %v", got, tt.want)
			}
			if got := findNumberIn2DArray1(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
