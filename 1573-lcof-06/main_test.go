package _573_lcof_06

import (
	"reflect"
	"testing"
)

func Test_reversePrint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}, want: []int{3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrint(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
			if got := reversePrint1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint1() = %v, want %v", got, tt.want)
			}
		})
	}
}
