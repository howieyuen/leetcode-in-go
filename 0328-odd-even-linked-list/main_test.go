package _328_odd_even_linked_list

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{head: []int{1, 2, 3, 4, 5}}, want: []int{1, 3, 5, 2, 4}},
		{args: args{head: []int{2, 1, 3, 5, 6, 4, 7}}, want: []int{2, 3, 6, 7, 1, 5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(buildList(tt.args.head)); !reflect.DeepEqual(readList(got), tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readList(node *ListNode) []int {
	var vals []int
	for node != nil {
		vals = append(vals, node.Val)
		node = node.Next
	}
	return vals
}

func buildList(vals []int) *ListNode {
	var pre = new(ListNode)
	var p = pre
	for _, val := range vals {
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return pre.Next
}
