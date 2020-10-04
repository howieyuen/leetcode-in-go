package _002_Add_Two_Numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: buildListNode([]int{2, 4, 3}),
				l2: buildListNode([]int{5, 6, 4}),
			},
			want: buildListNode([]int{7, 0, 8}),
		},
		{
			args: args{
				l1: buildListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: buildListNode([]int{9, 9, 9, 9}),
			},
			want: buildListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
			if got := addTwoNumbers1(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildListNode(input []int) *ListNode {
	var faker = new(ListNode)
	p := faker
	for _, x := range input {
		p.Next = &ListNode{Val: x}
		p = p.Next
	}
	return faker.Next
}
