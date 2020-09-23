package _208_get_equal_substrings_within_budget

import (
	"testing"
)

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "pxezla", t: "loewbi", maxCost: 25}, want: 4},
		{args: args{s: "krrgw", t: "zjxss", maxCost: 19}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
