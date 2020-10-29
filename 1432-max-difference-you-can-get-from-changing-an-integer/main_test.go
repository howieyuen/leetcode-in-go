package _432_max_difference_you_can_get_from_changing_an_integer

import "testing"

func Test_maxDiff(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{num: 555}, want: 888},
		{args: args{num: 9}, want: 8},
		{args: args{num: 123456}, want: 820000},
		{args: args{num: 10000}, want: 80000},
		{args: args{num: 9288}, want: 8700},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDiff(tt.args.num); got != tt.want {
				t.Errorf("maxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
