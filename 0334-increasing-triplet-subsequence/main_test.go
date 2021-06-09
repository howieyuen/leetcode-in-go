package _334_increasing_triplet_subsequence

import (
	"testing"
)

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: true,
		},
		{
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: false,
		},
		{
			args: args{nums: []int{2, 1, 5, 0, 4, 6}},
			want: true,
		},
		{
			args: args{nums: []int{2, 4, -2, -3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
			if got := increasingTriplet1(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet1() = %v, want %v", got, tt.want)
			}
		})
	}
}
