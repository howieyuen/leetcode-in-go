package _069_sqrtx

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{x: 9},
			want: 3,
		},
		{
			args: args{x: 8},
			want: 2,
		},
		{
			args: args{x: 7},
			want: 2,
		},
		{
			args: args{x: 6},
			want: 2,
		},
		{
			args: args{x: 5},
			want: 2,
		},
		{
			args: args{x: 4},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
