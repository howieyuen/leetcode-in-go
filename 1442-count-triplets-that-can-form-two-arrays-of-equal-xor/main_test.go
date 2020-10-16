package _442_count_triplets_that_can_form_two_arrays_of_equal_xor

import "testing"

func Test_countTriplets(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{arr: []int{2, 3, 1, 6, 7}}, want: 4},
		{args: args{arr: []int{1, 1, 1, 1, 1}}, want: 10},
		{args: args{arr: []int{1, 3, 5, 7, 9}}, want: 3},
		{args: args{arr: []int{2, 3}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countTriplets(tt.args.arr); got != tt.want {
				t.Errorf("countTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
