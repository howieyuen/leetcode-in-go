package _514_freedom_trail

import (
	"testing"
)

func Test_findRotateSteps(t *testing.T) {
	type args struct {
		ring string
		key  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{ring: "zvyii", key: "iivyz"}, want: 11},
		{args: args{ring: "edcba", key: "abcde"}, want: 10},
		{args: args{ring: "abcde", key: "ade"}, want: 6},
		{args: args{ring: "godding", key: "gd"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRotateSteps1(tt.args.ring, tt.args.key); got != tt.want {
				t.Errorf("findRotateSteps1() = %v, want %v", got, tt.want)
			}
			if got := findRotateSteps(tt.args.ring, tt.args.key); got != tt.want {
				t.Errorf("findRotateSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
