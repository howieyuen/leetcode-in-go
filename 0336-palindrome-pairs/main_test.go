package _336_palindrome_pairs

import (
	"reflect"
	"testing"
)

func Test_palindromePairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{words: []string{"abcd", "dcba", "lls", "s", "sssll"}},
			want: [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}},
		},
		{
			args: args{words: []string{"bat", "tab", "cat"}},
			want: [][]int{{0, 1}, {1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePairs(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
