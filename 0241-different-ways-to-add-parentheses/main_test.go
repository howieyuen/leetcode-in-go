package _241_different_ways_to_add_parentheses

import (
	"reflect"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{input: "2-1-1"}, want: []int{2, 0}},
		{args: args{input: "2*3-4*5"}, want: []int{-34, -10, -14, -10, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diffWaysToCompute(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
