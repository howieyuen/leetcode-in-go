package _526_beautiful_arrangement

import (
	"testing"
)

func Test_countArrangement(t *testing.T) {
	type args struct {
		N []int
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{N: []int{1, 2, 3, 8, 10, 36, 41, 132, 250, 700, 750, 4010, 4237, 10680, 24679}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range tt.args.N {
				if got := countArrangement(i + 1); got != tt.args.N[i] {
					t.Errorf("countArrangement() = %v, want %v", got, tt.args.N[i])
				}
			}
		})
	}
}
