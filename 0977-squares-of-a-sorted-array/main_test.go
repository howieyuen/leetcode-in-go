package _977_squares_of_a_sorted_array

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{A: []int{-1, 2, 2}}, want: []int{1, 4, 4}},
		{args: args{A: []int{-4, -1, 0, 3, 10}}, want: []int{0, 1, 9, 16, 100}},
		{args: args{A: []int{-7, -3, 2, 3, 11}}, want: []int{4, 9, 9, 49, 121}},
		{args: args{A: []int{-4, -3, -2, -1}}, want: []int{1, 4, 9, 16}},
		{args: args{A: []int{-4, -3, -2, -1, 0}}, want: []int{0, 1, 4, 9, 16}},
		{args: args{A: []int{0, 1, 2, 3, 4}}, want: []int{0, 1, 4, 9, 16}},
		{args: args{A: []int{1, 2, 3, 4}}, want: []int{1, 4, 9, 16}},
		{args: args{A: []int{1}}, want: []int{1}},
		{args: args{A: []int{0}}, want: []int{0}},
		{args: args{A: []int{-1}}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
			if got := sortedSquares1(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares1() = %v, want %v", got, tt.want)
			}
		})
	}
}
