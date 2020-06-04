package _837_new_21_game

import (
	"testing"
)

func Test_new21Game(t *testing.T) {
	type args struct {
		n int
		k int
		w int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{args: args{n: 6, k: 1, w: 10}, want: 0.60000},
		{args: args{n: 21, k: 17, w: 10}, want: 0.73278},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new21Game(tt.args.n, tt.args.k, tt.args.w); got != tt.want {
				t.Errorf("new21Game() = %v, want %v", got, tt.want)
			}
		})
	}
}
