package _073_set_matrix_zeroes

import (
	"fmt"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{matrix: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			}},
		},
		{
			args: args{matrix: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			fmt.Println(tt.args.matrix)
		})
	}
}
