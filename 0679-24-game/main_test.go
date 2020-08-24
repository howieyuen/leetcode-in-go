package _679_24_game

import (
	"testing"
)

func Test_judgePoint24(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{nums: []int{4, 1, 8, 7}}, want: true},
		{args: args{nums: []int{1, 2, 1, 2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24(tt.args.nums); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}
