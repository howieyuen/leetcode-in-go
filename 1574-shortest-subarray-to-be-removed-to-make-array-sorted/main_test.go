package _574_shortest_subarray_to_be_removed_to_make_array_sorted

import "testing"

func Test_findLengthOfShortestSubarray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{arr: []int{1, 2, 3, 10, 4, 2, 3, 5}}, want: 3},
		{args: args{arr: []int{5, 4, 3, 2, 1}}, want: 4},
		{args: args{arr: []int{1, 2, 3}}, want: 0},
		{args: args{arr: []int{1, 5, 5, 5, 10, 3, 5, 5, 5, 20}}, want: 2},
		{args: args{arr: []int{2, 1, 2, 3, 4, 5, 3}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfShortestSubarray(tt.args.arr); got != tt.want {
				t.Errorf("findLengthOfShortestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
