package _954_array_of_doubled_pairs

import (
	"testing"
)

func Test_canReorderDoubled(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{A: []int{4, -2, 2, -4}}, want: true},
		{args: args{A: []int{3, 1, 3, 6}}, want: false},
		{args: args{A: []int{-5, -2}}, want: false},
		{args: args{A: []int{2, 1, 2, 6}}, want: false},
		{args: args{A: []int{1, 2, 4, 16, 8, 4}}, want: false},
		{args: args{A: []int{0, 0, 0, 0, 0, 0}}, want: true},
		{args: args{A: []int{-62, 86, 96, -18, 43, -24, -76, 13, -31, -26, -88, -13, 96, -44, 9, -20, -42, 100, -44, -24, -36, 28, -32, 58, -72, 20, 48, -36, -45, 14, 24, -64, -90, -44, -16, 86, -33, 48, 26, 29, -84, 10, 46, 50, -66, -8, -38, 92, -19, 43, 48, -38, -22, 18, -32, -48, -64, -10, -22, -48}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReorderDoubled(tt.args.A); got != tt.want {
				t.Errorf("canReorderDoubled() = %v, want %v", got, tt.want)
			}
			if got := canReorderDoubled1(tt.args.A); got != tt.want {
				t.Errorf("canReorderDoubled1() = %v, want %v", got, tt.want)
			}
		})
	}
}
