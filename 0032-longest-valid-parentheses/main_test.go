package _032_longest_valid_parentheses

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "(()"}, want: 2},
		{args: args{s: ")()())"}, want: 4},
		{args: args{s: "()(()"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
			if got := longestValidParentheses1(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses1() = %v, want %v", got, tt.want)
			}
			if got := longestValidParentheses2(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses2() = %v, want %v", got, tt.want)
			}
		})
	}
}
