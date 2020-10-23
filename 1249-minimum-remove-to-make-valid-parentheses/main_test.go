package _249_minimum_remove_to_make_valid_parentheses

import (
	"testing"
)

func Test_minRemoveToMakeValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "lee(t(c)o)de)"}, want: "lee(t(c)o)de"},
		{args: args{s: "a)b(c)d"}, want: "ab(c)d"},
		{args: args{s: "))(("}, want: ""},
		{args: args{s: "(a(b(c)d)"}, want: "a(b(c)d)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.args.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
