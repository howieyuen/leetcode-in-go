package _234_palindrome_linked_list

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{[]int{1, 2}}, want: false},
		{args: args{[]int{1, 2, 2, 1}}, want: true},
		{args: args{[]int{1, 2, 1}}, want: true},
		{args: args{[]int{1, 2, 2, 3}}, want: false},
		{args: args{[]int{1, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(buildListNode(tt.args.head)); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
			if got := isPalindrome1(buildListNode(tt.args.head)); got != tt.want {
				t.Errorf("isPalindrome1() = %v, want %v", got, tt.want)
			}
			if got := isPalindrome2(buildListNode(tt.args.head)); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	var head = &ListNode{Val: arr[0]}
	p := head
	for _, x := range arr[1:] {
		p.Next = &ListNode{Val: x}
		p = p.Next
	}
	return head
}
