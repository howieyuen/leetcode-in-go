package test1

import (
	"testing"
)

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{numBottles: 9, numExchange: 3}, want: 13},
		{args: args{numBottles: 15, numExchange: 4}, want: 19},
		{args: args{numBottles: 5, numExchange: 5}, want: 6},
		{args: args{numBottles: 2, numExchange: 3}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
