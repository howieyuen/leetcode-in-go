package _202_smallest_string_with_swaps

import "testing"

func Test_smallestStringWithSwaps(t *testing.T) {
	type args struct {
		s     string
		pairs [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "dcab", pairs: [][]int{{0,3}, {1,2}}}, want: "bacd"},
		{args: args{s: "dcab", pairs: [][]int{{0,3}, {1,2}, {0, 2}}}, want: "abcd"},
		{args: args{s: "cba", pairs: [][]int{{0,1}, {1,2}}}, want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStringWithSwaps(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("smallestStringWithSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
