package _166_fraction_to_recurring_decimal

import (
	"testing"
)

func Test_fractionToDecimal(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{numerator: 1, denominator: 2,}, want: "0.5"},
		{args: args{numerator: 2, denominator: 1,}, want: "2"},
		{args: args{numerator: 1, denominator: 3,}, want: "0.(3)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionToDecimal(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("fractionToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
