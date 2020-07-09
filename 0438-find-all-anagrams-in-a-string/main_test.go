package _438_find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{s: "cbaebabacd", p: "abc"}, want: []int{0, 6}},
		{args: args{s: "abab", p: "ab"}, want: []int{0, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
			if got := findAnagrams1(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams1() = %v, want %v", got, tt.want)
			}
		})
	}
}
