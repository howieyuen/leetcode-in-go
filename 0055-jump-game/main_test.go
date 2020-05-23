package _055_jump_game

import (
	"testing"
)

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{nums: []int{2, 3, 1, 1, 4}}, want: true},
		{args: args{nums: []int{3, 2, 1, 0, 4}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
			if got := canJump1(tt.args.nums); got != tt.want {
				t.Errorf("canJump1() = %v, want %v", got, tt.want)
			}
		})
	}
}
