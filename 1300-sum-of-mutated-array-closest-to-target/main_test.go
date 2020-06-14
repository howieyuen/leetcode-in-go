package _300_sum_of_mutated_array_closest_to_target

import (
	"testing"
)

func Test_findBestValue(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr:    []int{4, 9, 3},
				target: 10,
			},
			want: 3,
		},
		{
			args: args{
				arr:    []int{2, 3, 5},
				target: 10,
			},
			want: 5,
		},
		{
			args: args{
				arr:    []int{60864, 25176, 27249, 21296, 20204},
				target: 56803,
			},
			want: 11361,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBestValue(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("findBestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
