package _845_longest_mountain_in_array

import (
	"testing"
)

func Test_longestMountain(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{A: []int{2, 1, 4, 7, 3, 2, 5}}, want: 5},
		{args: args{A: []int{2, 2, 2}}, want: 0},
		{args: args{A: []int{875, 884, 239, 731, 723, 685}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.args.A); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
