package _2_05

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				l1: []int{7, 1, 6},
				l2: []int{5, 9, 2},
			},
			want: []int{2, 1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(buildList(tt.args.l1), buildList(tt.args.l2)); !reflect.DeepEqual(got, buildList(tt.want)) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildList(nums []int) *ListNode {
	var pre = new(ListNode)
	p := pre
	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}
	return pre.Next
}
