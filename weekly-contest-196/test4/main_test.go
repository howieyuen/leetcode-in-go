package test4

import (
	"testing"
)

func Test_minInteger(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{num: "4321", k: 4,}, want: "1342"},
		{args: args{num: "100", k: 1,}, want: "010"},
		{args: args{num: "36789", k: 1000,}, want: "36789"},
		{args: args{num: "22", k: 22,}, want: "22"},
		{args: args{num: "9438957234785635408", k: 23,}, want: "0345989723478563548"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minInteger(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("minInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
