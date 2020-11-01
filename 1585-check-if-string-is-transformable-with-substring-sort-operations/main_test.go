package _585_check_if_string_is_transformable_with_substring_sort_operations

import (
	"testing"
)

func Test_isTransformable(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{s: "84532", t: "34852"}, want: true},
		{args: args{s: "34521", t: "23415"}, want: true},
		{args: args{s: "12345", t: "12435"}, want: false},
		{args: args{s: "1", t: "2"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTransformable(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isTransformable() = %v, want %v", got, tt.want)
			}
		})
	}
}
