package _052_n_queens_ii

import (
	"testing"
)

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 4}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
			if got := totalNQueens1(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens1() = %v, want %v", got, tt.want)
			}
			if got := totalNQueens2(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens2() = %v, want %v", got, tt.want)
			}
		})
	}
}
