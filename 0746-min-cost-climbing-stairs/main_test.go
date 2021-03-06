package _746_min_cost_climbing_stairs

import (
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{cost: []int{10, 15, 20}},
			want: 15,
		},
		{
			args: args{cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
			if got := minCostClimbingStairs1(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs1() = %v, want %v", got, tt.want)
			}
		})
	}
}
