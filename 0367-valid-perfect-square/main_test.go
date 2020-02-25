package _367_valid_perfect_square

import (
	"testing"
)

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{num: 0},
			want: true,
		},
		{
			args: args{num: 1},
			want: true,
		},
		{
			args: args{num: 2},
			want: false,
		},
		{
			args: args{num: 3},
			want: false,
		},
		{
			args: args{num: 4},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
