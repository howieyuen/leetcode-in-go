package _260_single_number_iii

import (
	"reflect"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{1, 2, 1, 3, 2, 5}}, want: []int{3, 5}},
		{args: args{nums: []int{2, 1, 2, 3, 4, 1}}, want: []int{3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
