package _042_trapping_rain_water

import (
	"testing"
)

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, want: 6},
		{args: args{height: []int{5, 2, 1, 2, 1, 5}}, want: 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap1(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
			if got := trap2(tt.args.height); got != tt.want {
				t.Errorf("trap2() = %v, want %v", got, tt.want)
			}
			if got := trap3(tt.args.height); got != tt.want {
				t.Errorf("trap3() = %v, want %v", got, tt.want)
			}
			if got := trap4(tt.args.height); got != tt.want {
				t.Errorf("trap4() = %v, want %v", got, tt.want)
			}
		})
	}
}
