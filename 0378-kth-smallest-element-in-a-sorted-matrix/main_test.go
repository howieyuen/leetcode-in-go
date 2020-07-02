package _378_kth_smallest_element_in_a_sorted_matrix

import (
	"testing"
)

func Test_kthSmallest(t *testing.T) {
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
					{1, 5, 9},
					{0, 11, 13},
					{2, 13, 15},
				},
				k: 8,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest1(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest1() = %v, want %v", got, tt.want)
			}
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
