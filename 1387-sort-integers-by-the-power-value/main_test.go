package _387_sort_integers_by_the_power_value

import "testing"

func Test_getKth(t *testing.T) {
	type args struct {
		lo int
		hi int
		k  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{lo: 12, hi: 15, k: 2}, want: 13},
		{args: args{lo: 1, hi: 1, k: 1}, want: 1},
		{args: args{lo: 10, hi: 20, k: 5}, want: 13},
		{args: args{lo: 1, hi: 1000, k: 777}, want: 570},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKth(tt.args.lo, tt.args.hi, tt.args.k); got != tt.want {
				t.Errorf("getKth() = %v, want %v", got, tt.want)
			}
			if got := getKth1(tt.args.lo, tt.args.hi, tt.args.k); got != tt.want {
				t.Errorf("getKth1() = %v, want %v", got, tt.want)
			}
			if got := getKth2(tt.args.lo, tt.args.hi, tt.args.k); got != tt.want {
				t.Errorf("getKth2() = %v, want %v", got, tt.want)
			}
		})
	}
}
