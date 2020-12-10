package _441_arranging_coins

import "testing"

func Test_arrangeCoins(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 0}, want: 0},
		{args: args{n: 1}, want: 1},
		{args: args{n: 2}, want: 1},
		{args: args{n: 3}, want: 2},
		{args: args{n: 4}, want: 2},
		{args: args{n: 5}, want: 2},
		{args: args{n: 6}, want: 3},
		{args: args{n: 7}, want: 3},
		{args: args{n: 8}, want: 3},
		{args: args{n: 9}, want: 3},
		{args: args{n: 10}, want: 4},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoins(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
			}
			if got := arrangeCoins1(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins1() = %v, want %v", got, tt.want)
			}
			if got := arrangeCoins2(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoins2() = %v, want %v", got, tt.want)
			}
		})
	}
}
