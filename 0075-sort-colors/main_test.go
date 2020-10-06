package _075_sort_colors

import (
	"reflect"
	"testing"
)

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
		want []int
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{nums: []int{2, 0, 2, 1, 1, 0}, want: []int{0, 0, 1, 1, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sortColors(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.args.want) {
				t.Errorf("got: %v, want: %v", tt.args.nums, tt.args.want)
			}
			if sortColors1(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.args.want) {
				t.Errorf("got: %v, want: %v", tt.args.nums, tt.args.want)
			}
		})
	}
}
