package _0

import (
	"testing"
)

func Test_isNumber1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// "-123"、"3.1416"、"-1E-16"、"0123"
		{args: args{s: "+100"}, want: true},
		{args: args{s: "5e2"}, want: true},
		{args: args{s: "-123"}, want: true},
		{args: args{s: "3.1416"}, want: true},
		{args: args{s: "-1E-16"}, want: true},
		{args: args{s: "0123"}, want: true},
		// "12e"、"1a3.14"、"1.2.3"、"+-5"、"12e+5.4"
		{args: args{s: "12e"}, want: false},
		{args: args{s: "1a3.14"}, want: false},
		{args: args{s: "1.2.3"}, want: false},
		{args: args{s: "+-5"}, want: false},
		{args: args{s: "12e+5.4"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber1(tt.args.s); got != tt.want {
				t.Errorf("isNumber1() = %v, want %v", got, tt.want)
			}
		})
	}
}
