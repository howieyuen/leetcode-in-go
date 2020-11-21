package _283_move_zeroes

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{0, 1, 0, 3, 12}}, want: []int{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if moveZeroes(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, got %v", tt.args.nums, tt.want)
			}
			if moveZeroes1(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes1() = %v, got %v", tt.args.nums,tt.want)
			}
		})
	}
}
