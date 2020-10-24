package _449_form_largest_integer_with_digits_that_add_up_to_target

import (
	"testing"
)

func Test_largestNumber(t *testing.T) {
	type args struct {
		cost   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{cost: []int{4, 3, 2, 5, 6, 7, 2, 5, 5}, target: 9}, want: "7772"},
		{args: args{cost: []int{7, 6, 5, 5, 5, 6, 8, 7, 8}, target: 12}, want: "85"},
		{args: args{cost: []int{2, 4, 6, 2, 4, 6, 4, 4, 4}, target: 5}, want: "0"},
		{args: args{cost: []int{6, 10, 15, 40, 40, 40, 40, 40, 40}, target: 47}, want: "32211"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumberN(tt.args.cost, tt.args.target); got != tt.want {
				t.Errorf("largestNumberN() = %v, want %v", got, tt.want)
			}
			if got := largestNumber2(tt.args.cost, tt.args.target); got != tt.want {
				t.Errorf("largestNumber2() = %v, want %v", got, tt.want)
			}
			if got := largestNumber1(tt.args.cost, tt.args.target); got != tt.want {
				t.Errorf("largestNumber1() = %v, want %v", got, tt.want)
			}
			if got := largestNumber(tt.args.cost, tt.args.target); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
