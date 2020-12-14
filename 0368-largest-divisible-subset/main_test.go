package _368_largest_divisible_subset

import (
	"reflect"
	"testing"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{1}}, want: []int{1}},
		{args: args{nums: []int{1, 2, 3}}, want: []int{1, 2}},
		{args: args{nums: []int{1, 2, 4, 8}}, want: []int{1, 2, 4, 8}},
		{args: args{nums: []int{5, 9, 18, 54, 108, 540, 90, 180, 360, 720}}, want: []int{9, 18, 90, 180, 360, 720}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestDivisibleSubset(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset() = %v, want %v", got, tt.want)
			}
			if got := largestDivisibleSubset1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestDivisibleSubset1() = %v, want %v", got, tt.want)
			}
		})
	}
}
