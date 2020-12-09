package _363_max_sum_of_rectangle_no_larger_than_k

import (
	"testing"
)

func Test_maxSumSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				matrix: [][]int{
					{2, 2, -1},
				},
				k: 0,
			},
			want: -1,
		},
		{
			args: args{
				matrix: [][]int{
					{1, 0, 1},
					{0, -2, 3},
				},
				k: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumSubmatrix(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("maxSumSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
