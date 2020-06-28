package _6_II

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{3, 4, 3, 3}}, want: 4},
		{args: args{nums: []int{9, 1, 7, 9, 7, 9, 7}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
			if got := singleNumber1(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}
