package _513_number_of_substrings_with_only_1s

import (
	"testing"
)

func Test_numSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "0110111"}, want: 9},
		{args: args{s: "101"}, want: 2},
		{args: args{s: "111111"}, want: 21},
		{args: args{s: "0000"}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSub(tt.args.s); got != tt.want {
				t.Errorf("numSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
