package _945_minimum_increment_to_make_array_unique

import (
	"testing"
)

func Test_minIncrementForUnique(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{A: []int{1, 2, 2}}, want: 1},
		{args: args{A: []int{3, 2, 1, 2, 1, 7}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minIncrementForUnique(tt.args.A); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
