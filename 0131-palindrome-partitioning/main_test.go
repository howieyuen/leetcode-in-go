package _131_palindrome_partitioning

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			args: args{
				s: "aab",
			},
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		
		{
			args: args{
				s: "abbab",
			},
			want: [][]string{
				{"a", "b", "b", "a", "b"},
				{"a", "b", "bab"},
				{"a", "bb", "a", "b"},
				{"abba", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
