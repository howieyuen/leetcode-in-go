package _3

import (
	"testing"
)

func Test_minimalSteps(t *testing.T) {
	type args struct {
		maze []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{maze: []string{"S#O", "M..", "M.T"}}, want: 16},
		{args: args{maze: []string{"S#O", "M.#", "M.T"}}, want: -1},
		{args: args{maze: []string{"S#O", "M.T", "M.."}}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalSteps(tt.args.maze); got != tt.want {
				t.Errorf("minimalSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
