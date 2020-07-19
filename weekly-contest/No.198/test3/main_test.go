package test3

import (
	"reflect"
	"testing"
)

func Test_maxNumOfSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{args: args{s: "abcaba"}, want: []string{"c"}},
		{args: args{s: "adefaddaccc"}, want: []string{"f", "e", "ccc"}},
		{args: args{s: "abbaccd"}, want: []string{"d", "bb", "cc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumOfSubstrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxNumOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
