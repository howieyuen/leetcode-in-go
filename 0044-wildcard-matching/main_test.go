package _044_wildcard_matching

import (
	"testing"
)

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "aa", p: "a"}, want: false},
		{args: args{s: "aa", p: "*"}, want: true},
		{args: args{s: "cb", p: "?a"}, want: false},
		{args: args{s: "adceb", p: "*a*b"}, want: true},
		{args: args{s: "acdcb", p: "a*c?b"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
