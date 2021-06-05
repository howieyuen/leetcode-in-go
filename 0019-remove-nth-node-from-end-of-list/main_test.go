package _019_Remove_Nth_Node_From_End_of_List

import (
	"testing"
)

func Test_removeNthFromEnd1(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: buildListNode([]int{1, 2, 3, 5}),
		},
		{
			args: args{
				head: []int{1},
				n:    1,
			},
			want: buildListNode(nil),
		},
		{
			args: args{
				head: []int{1, 2},
				n:    2,
			},
			want: buildListNode([]int{2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(buildListNode(tt.args.head), tt.args.n); !shallowEquals(got, tt.want) {
				t.Errorf("removeNthFromEnd1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var faker = new(ListNode)
	p := faker
	for _, val := range nums {
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return faker.Next
}

func shallowEquals(head1, head2 *ListNode) bool {
	p, q := head1, head2
	for p != nil && q != nil && p.Val == q.Val {
		p = p.Next
		q = q.Next
	}
	return p == nil && q == nil
}
