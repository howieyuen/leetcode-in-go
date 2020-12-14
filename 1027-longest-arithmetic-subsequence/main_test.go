package _027_longest_arithmetic_subsequence

import "testing"

func Test_longestArithSeqLength(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{A: []int{3, 6, 9, 12}}, want: 4},
		{args: args{A: []int{9, 4, 7, 2, 10}}, want: 3},
		{args: args{A: []int{20, 1, 15, 3, 10, 5, 8}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestArithSeqLength(tt.args.A); got != tt.want {
				t.Errorf("longestArithSeqLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
