package _392_longest_happy_prefix

import (
	"testing"
)

func Test_longestPrefix(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "level"}, want: "l"},
		{args: args{s: "ababab"}, want: "abab"},
		{args: args{s: "leetcodeleet"}, want: "leet"},
		{args: args{s: "a"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPrefix(tt.args.s); got != tt.want {
				t.Errorf("longestPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
