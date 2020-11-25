package _370_increasing_decreasing_string

import (
	"testing"
)

func Test_sortString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "aaaabbbbcccc"}, want: "abccbaabccba"},
		{args: args{s: "rat"}, want: "art"},
		{args: args{s: "leetcode"}, want: "cdelotee"},
		{args: args{s: "ggggggg"}, want: "ggggggg"},
		{args: args{s: "spo"}, want: "ops"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.args.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
