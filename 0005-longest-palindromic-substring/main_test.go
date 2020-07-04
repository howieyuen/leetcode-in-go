package main

import (
	"testing"
)

func Test_longestPalindrome1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "babad"}, want: "aba"},
		{args: args{s: "abcdbbfcba"}, want: "bb"},
		{args: args{s: "adx"}, want: "x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "babad"}, want: "aba"},
		{args: args{s: "abcdbbfcba"}, want: "bb"},
		{args: args{s: "adx"}, want: "x"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPalindrome3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "babad"}, want: "bab"},
		{args: args{s: "abcdbbfcba"}, want: "bb"},
		{args: args{s: "adx"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome3(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome3() = %v, want %v", got, tt.want)
			}
		})
	}
}