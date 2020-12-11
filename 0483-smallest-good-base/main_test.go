package _483_smallest_good_base

import (
	"testing"
)

func Test_smallestGoodBase(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{n: "14919921443713777"}, want: "496"},
		{args: args{n: "13"}, want: "3"},
		{args: args{n: "4681"}, want: "8"},
		{args: args{n: "1000000000000000000"}, want: "999999999999999999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestGoodBase(tt.args.n); got != tt.want {
				t.Errorf("smallestGoodBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
