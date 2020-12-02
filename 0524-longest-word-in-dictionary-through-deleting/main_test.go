package _524_longest_word_in_dictionary_through_deleting

import (
	"testing"
)

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s string
		d []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{s: "abpcplea", d: []string{"ale", "apple", "monkey", "plea"}}, want: "apple"},
		{args: args{s: "abpcplea", d: []string{"a", "b", "c"}}, want: "a"},
		{args: args{s: "aewfafwafjlwajflwajflwafj", d: []string{"apple", "ewaf", "awefawfwaf", "awef", "awefe", "ewafeffewafewf"}}, want: "ewaf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
