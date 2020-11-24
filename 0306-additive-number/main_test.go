package _306_additive_number

import "testing"

func Test_isAdditiveNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{num: "101"}, want: true},
		{args: args{num: "199100199"}, want: true},
		{args: args{num: "111"}, want: false},
		{args: args{num: "112358"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdditiveNumber(tt.args.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
