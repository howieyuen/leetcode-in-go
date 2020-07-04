package _647_palindromic_substrings

import (
	"testing"
)

func Test_countSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "abc"}, want: 3},
		{args: args{s: "aaa"}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubstrings(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings() = %v, want %v", got, tt.want)
			}
			if got := countSubstrings1(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings1() = %v, want %v", got, tt.want)
			}
			if got := countSubstrings2(tt.args.s); got != tt.want {
				t.Errorf("countSubstrings2() = %v, want %v", got, tt.want)
			}
		})
	}
}
