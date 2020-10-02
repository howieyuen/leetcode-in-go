package _7_21

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.args.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
			if got := trap1(tt.args.height); got != tt.want {
				t.Errorf("trap1() = %v, want %v", got, tt.want)
			}
			if got := trap2(tt.args.height); got != tt.want {
				t.Errorf("trap2() = %v, want %v", got, tt.want)
			}
			if got := trap3(tt.args.height); got != tt.want {
				t.Errorf("trap3() = %v, want %v", got, tt.want)
			}
		})
	}
}
