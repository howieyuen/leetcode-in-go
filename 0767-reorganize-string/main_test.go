package _767_reorganize_string

import (
	"testing"
)

func Test_reorganizeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "aab"}, want: "aba"},
		{args: args{s: "aaab"}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorganizeString(tt.args.s); got != tt.want {
				t.Errorf("reorganizeString() = %v, want %v", got, tt.want)
			}
			if got := reorganizeString1(tt.args.s); got != tt.want {
				t.Errorf("reorganizeString1() = %v, want %v", got, tt.want)
			}
		})
	}
}
