package _556_next_greater_element_iii

import (
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{n: 6987}, want: 7689},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement1(tt.args.n); got != tt.want {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
