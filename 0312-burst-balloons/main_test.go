package _312_burst_balloons

import (
	"testing"
)

func Test_maxCoins(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{nums: []int{3, 1, 5, 8}}, want: 167},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoins1(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins1() = %v, want %v", got, tt.want)
			}
			if got := maxCoins2(tt.args.nums); got != tt.want {
				t.Errorf("maxCoins2() = %v, want %v", got, tt.want)
			}
		})
	}
}
