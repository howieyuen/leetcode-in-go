package main

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				str: "-91283472332",
			},
			want: -2147483648,
		},
		{
			name: "test2",
			args: args{
				str: "-32",
			},
			want: -32,
		},
		{
			name: "test3",
			args: args{
				str: "   -32",
			},
			want: -32,
		},
		{
			name: "test4",
			args: args{
				str: "4193 with words",
			},
			want: 4193,
		},
		{
			name: "test5",
			args: args{
				str: "with words 4193",
			},
			want: 0,
		},
		{
			name: "test6",
			args: args{
				str: "+-2",
			},
			want: 0,
		},
		{
			name: "test7",
			args: args{
				str: "-2147483648",
			},
			want: -2147483648,
		},
		{
			name: "test8",
			args: args{
				str: "2147483647",
			},
			want: 2147483647,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
