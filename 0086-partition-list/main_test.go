package _086_partition_list

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: buildList([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			want: buildList([]int{1, 2, 2, 4, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(transform(got), transform(tt.want)) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func transform(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	p := head
	data := []int{}
	for p != nil {
		data = append(data, p.Val)
		p = p.Next
	}
	return data
}

func buildList(data []int) *ListNode {
	faker := &ListNode{}
	p := faker
	for i := range data {
		p.Next = &ListNode{Val: data[i]}
		p = p.Next
	}
	return faker.Next
}
