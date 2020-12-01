package _498_number_of_subsequences_that_satisfy_the_given_sum_condition

import "testing"

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{3, 5, 6, 7}, target: 9}, want: 4},
		{args: args{nums: []int{3, 3, 6, 8}, target: 10}, want: 6},
		{args: args{nums: []int{2, 3, 3, 4, 6, 7}, target: 12}, want: 61},
		{args: args{nums: []int{5, 2, 4, 1, 7, 6, 8}, target: 16}, want: 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubseq(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
