package _357_count_numbers_with_unique_digits

import "testing"

func Test_countNumbersWithUniqueDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 0}, want: 1},
		{args: args{n: 1}, want: 10},
		{args: args{n: 2}, want: 91},
		{args: args{n: 3}, want: 739},
		{args: args{n: 4}, want: 5275},
		{args: args{n: 5}, want: 32491},
		{args: args{n: 6}, want: 168571},
		{args: args{n: 7}, want: 712891},
		{args: args{n: 8}, want: 2345851},
		{args: args{n: 9}, want: 5611771},
		{args: args{n: 10}, want: 8877691},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNumbersWithUniqueDigits(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
			}
			if got := countNumbersWithUniqueDigits1(tt.args.n); got != tt.want {
				t.Errorf("countNumbersWithUniqueDigits1() = %v, want %v", got, tt.want)
			}
		})
	}
}
