package _065_valid_number

import (
	"testing"
)

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "0"}, want: true},
		{args: args{s: " 0.1 "}, want: true},
		{args: args{s: "abc"}, want: false},
		{args: args{s: "1 a"}, want: false},
		{args: args{s: "2e10"}, want: true},
		{args: args{s: " -90e3 "}, want: true},
		{args: args{s: " 1e"}, want: false},
		{args: args{s: "e3"}, want: false},
		{args: args{s: " 6e-1"}, want: true},
		{args: args{s: " 99e2.5"}, want: false},
		{args: args{s: "53.5e93"}, want: true},
		{args: args{s: " --6 "}, want: false},
		{args: args{s: "-+3"}, want: false},
		{args: args{s: "95a54e53"}, want: false},
		{args: args{s: ".1"}, want: true},
		{args: args{s: "3."}, want: true},
		{args: args{s: "1  "}, want: true},
		{args: args{s: "4e+"}, want: false},
		{args: args{s: "+.8"}, want: true},
		{args: args{s: "."}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber(%s) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
