package _842_split_array_into_fibonacci_sequence

import (
	"reflect"
	"testing"
)

func Test_splitIntoFibonacci(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{S: "123456579"}, want: []int{123, 456, 579}},
		{args: args{S: "11235813"}, want: []int{1, 1, 2, 3, 5, 8, 13}},
		{args: args{S: "112358130"}, want: []int{}},
		{args: args{S: "0123"}, want: []int{}},
		{args: args{S: "1101111"}, want: []int{11, 0, 11, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitIntoFibonacci(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci() = %v, want %v", got, tt.want)
			}
			if got := splitIntoFibonacci1(tt.args.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitIntoFibonacci1() = %v, want %v", got, tt.want)
			}
		})
	}
}
