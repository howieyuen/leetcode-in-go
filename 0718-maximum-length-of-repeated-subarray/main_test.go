package _718_maximum_length_of_repeated_subarray

import (
	"testing"
)

func Test_findLength(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: []int{1, 2, 3, 2, 1, 7},
				b: []int{3, 2, 1, 4, 7, 4},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
