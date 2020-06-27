package _111_maximum_nesting_depth_of_two_valid_parentheses_strings

import (
	"reflect"
	"testing"
)

func Test_maxDepthAfterSplit(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{seq: "(()())"}, want: []int{0, 1, 1, 1, 1, 0}},
		{args: args{seq: "()(())()"}, want: []int{0, 0, 0, 1, 1, 0, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepthAfterSplit(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthAfterSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
