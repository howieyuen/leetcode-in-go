package _738_monotone_increasing_digits

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{N: 1234}, want: 1234},
		{args: args{N: 10}, want: 9},
		{args: args{N: 4321}, want: 3999},
		{args: args{N: 332}, want: 299},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.N); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
