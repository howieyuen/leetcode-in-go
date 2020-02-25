package _209_minimum_size_subarray_sum

import (
	`testing`
)

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		s    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s:    7,
				nums: []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
