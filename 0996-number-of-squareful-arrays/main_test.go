package _996_number_of_squareful_arrays

import (
	"testing"
)

func Test_numSquarefulPerms(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{A: []int{65, 44, 5, 11}}, want: 0},
		{args: args{A: []int{1, 17, 8}}, want: 2},
		{args: args{A: []int{2, 2, 2}}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquarefulPerms(tt.args.A); got != tt.want {
				t.Errorf("numSquarefulPerms() = %v, want %v", got, tt.want)
			}
		})
	}
}
