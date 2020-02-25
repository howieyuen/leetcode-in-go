package _297_serialize_and_deserialize_binary_tree

import (
	"reflect"
	"testing"
)

var root = &TreeNode{
	Val:  1,
	Left: &TreeNode{Val: 2},
	Right: &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 5},
	},
}

var want = "[1,2,3,null,null,4,5]"

func Test_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{root: root},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{data: want},
			want: root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
