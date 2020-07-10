package _122_best_time_to_buy_and_sell_stock_ii

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 7},
		{args: args{prices: []int{1, 2, 3, 4, 5}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
			if got := maxProfit1(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit1() = %v, want %v", got, tt.want)
			}
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
