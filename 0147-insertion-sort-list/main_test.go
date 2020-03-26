package _147_insertion_sort_list

import (
	"testing"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}},
		},
		{
			args: args{
				head: &ListNode{
					Val: -1,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{Val: 0}}}}},
			},
			want: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{Val: 5}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !equal(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(head1, head2 *ListNode) bool {
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return head1 == nil && head2 == nil
}
