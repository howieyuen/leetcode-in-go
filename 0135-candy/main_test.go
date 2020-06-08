package _135_candy

import (
	"testing"
)

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{ratings: []int{1, 0, 2}}, want: 5},
		{args: args{ratings: []int{1, 2, 2}}, want: 4},
		{args: args{ratings: []int{1, 3, 4, 5, 2}}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candy(tt.args.ratings); got != tt.want {
				t.Errorf("candy() = %v, want %v", got, tt.want)
			}
		})
	}
}
