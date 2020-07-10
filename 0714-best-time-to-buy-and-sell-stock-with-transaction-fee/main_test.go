package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

import (
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{prices: []int{1, 3, 2, 8, 4, 9}, fee: 2}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
