package _229_majority_element_ii

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{3, 2, 3}}, want: []int{3}},
		{args: args{nums: []int{1, 1, 1, 3, 3, 2, 2, 2}}, want: []int{1, 2}},
		{args: args{nums: []int{2, 2, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9}}, want: []int{3, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
