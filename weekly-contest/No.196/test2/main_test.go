package test2

import (
	"testing"
)

func Test_getLastMoment(t *testing.T) {
	type args struct {
		n     int
		left  []int
		right []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:     12,
				left:  []int{7, 1},
				right: []int{12, 11},
			},
			want: 7,
		},
		{
			args: args{
				n:     4,
				left:  []int{4, 3},
				right: []int{0, 1},
			},
			want: 4,
		},
		{
			args: args{
				n:     7,
				left:  []int{},
				right: []int{0, 1, 2, 3, 4, 5, 6, 7},
			},
			want: 7,
		},
		{
			args: args{
				n:     7,
				left:  []int{0, 1, 2, 3, 4, 5, 6, 7},
				right: []int{},
			},
			want: 7,
		},
		{
			args: args{
				n:     9,
				left:  []int{5},
				right: []int{4},
			},
			want: 5,
		},
		{
			args: args{
				n:     6,
				left:  []int{6},
				right: []int{0},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastMoment(tt.args.n, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("getLastMoment() = %v, want %v", got, tt.want)
			}
		})
	}
}
