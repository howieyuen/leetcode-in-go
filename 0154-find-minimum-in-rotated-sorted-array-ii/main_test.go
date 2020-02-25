package _154_find_minimum_in_rotated_sorted_array_ii

import (
	"testing"
)

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{[]int{2, 2, 2, 0, 1}},
			want: 0,
		},
		{
			args: args{[]int{1, 1}},
			want: 1,
		},
		{
			args: args{[]int{4, 5, 6, 7, 0, 1, 2}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
