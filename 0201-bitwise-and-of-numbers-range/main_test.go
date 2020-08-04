package _201_bitwise_and_of_numbers_range

import (
	"testing"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{m: 5, n: 7}, want: 4},
		{args: args{m: 0, n: 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
