package _948_bag_of_tokens

import (
	"testing"
)

func Test_bagOfTokensScore(t *testing.T) {
	type args struct {
		tokens []int
		P      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{tokens: []int{100}, P: 50}, want: 0},
		{args: args{tokens: []int{100, 200}, P: 150}, want: 1},
		{args: args{tokens: []int{100, 200, 300, 400}, P: 200}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagOfTokensScore(tt.args.tokens, tt.args.P); got != tt.want {
				t.Errorf("bagOfTokensScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
