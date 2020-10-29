package _553_optimal_division

import "testing"

func Test_optimalDivision(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{nums: []int{1000, 100, 10, 2}}, want: "1000/(100/10/2)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimalDivision(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision() = %v, want %v", got, tt.want)
			}
			if got := optimalDivision1(tt.args.nums); got != tt.want {
				t.Errorf("optimalDivision1() = %v, want %v", got, tt.want)
			}
		})
	}
}
