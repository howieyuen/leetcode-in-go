package _466_count_the_repetitions

import (
	"testing"
)

func Test_getMaxRepetitions(t *testing.T) {
	type args struct {
		s1 string
		n1 int
		s2 string
		n2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s1: "abc",
				n1: 4,
				s2: "ab",
				n2: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxRepetitions(tt.args.s1, tt.args.n1, tt.args.s2, tt.args.n2); got != tt.want {
				t.Errorf("getMaxRepetitions() = %v, want %v", got, tt.want)
			}
			if got := getMaxRepetitions1(tt.args.s1, tt.args.n1, tt.args.s2, tt.args.n2); got != tt.want {
				t.Errorf("getMaxRepetitions1() = %v, want %v", got, tt.want)
			}
		})
	}
}
