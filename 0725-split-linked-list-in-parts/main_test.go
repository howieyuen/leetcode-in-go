package _725_split_linked_list_in_parts

import (
	"reflect"
	"testing"
)

func Test_splitListToParts(t *testing.T) {
	type args struct {
		root *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []*ListNode
	}{
		{
			args: args{
				root: nil,
				k:    2,
			},
			want: []*ListNode{nil, nil},
		},
		{
			args: args{
				root: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
				k:    5,
			},
			want: []*ListNode{{Val: 1}, {Val: 2}, {Val: 3}, nil, nil},
		},
		{
			args: args{
				root: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
				k:    2,
			},
			want: []*ListNode{{Val: 1, Next: &ListNode{Val: 2}}, {Val: 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitListToParts(tt.args.root, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitListToParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
