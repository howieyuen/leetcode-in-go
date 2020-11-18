package _765_couples_holding_hands

import (
	"testing"
)

func Test_minSwapsCouples(t *testing.T) {
	type args struct {
		row []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{row: []int{0, 2, 1, 3}}, want: 1},
		{args: args{row: []int{3, 2, 0, 1}}, want: 0},
		{args: args{row: []int{0, 2, 4, 6, 7, 1, 3, 5}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwapsCouples(tt.args.row); got != tt.want {
				t.Errorf("minSwapsCouples() = %v, want %v", got, tt.want)
			}
		})
	}
}
