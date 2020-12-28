package _188_best_time_to_buy_and_sell_stock_iv

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{prices: []int{2, 4, 1}, k: 2}, want: 2},
		{args: args{prices: []int{3, 2, 6, 5, 0, 3}, k: 2}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
			if got := maxProfit1(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit1() = %v, want %v", got, tt.want)
			}
		})
	}
}
