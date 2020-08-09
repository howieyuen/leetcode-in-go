package _214_shortest_palindrome

import (
	"testing"
)

func Test_shortestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "aacecaaa"}, want: "aaacecaaa"},
		{args: args{s: "abcd"}, want: "dcbabcd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome1() = %v, want %v", got, tt.want)
			}
			if got := shortestPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome2() = %v, want %v", got, tt.want)
			}
			if got := shortestPalindrome3(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome3() = %v, want %v", got, tt.want)
			}
		})
	}
}
