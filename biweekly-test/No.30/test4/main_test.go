package test4

import (
	"testing"
)

func Test_winnerSquareGame(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{n: 1}, want: true},
		{args: args{n: 2}, want: false},
		{args: args{n: 4}, want: true},
		{args: args{n: 7}, want: false},
		{args: args{n: 17}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerSquareGame(tt.args.n); got != tt.want {
				t.Errorf("winnerSquareGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
