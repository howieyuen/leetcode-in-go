package _021_remove_outermost_parentheses

import (
	`testing`
)

func Test_removeOuterParentheses(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{S: "(()())(())"},
			want: "()()()",
		},
		{
			args: args{S: "(()())(())(()(()))"},
			want: "()()()()(())",
		},
		{
			args: args{S: "()()"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.S); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
