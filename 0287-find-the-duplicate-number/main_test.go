package _287_find_the_duplicate_number

import (
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{[]int{5, 2, 3, 1, 1, 4}}, want: 1},
		{args: args{[]int{1, 3, 4, 2, 2}}, want: 2},
		{args: args{[]int{3, 1, 3, 4, 2}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
