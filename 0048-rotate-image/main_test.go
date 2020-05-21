package _048_rotate_image

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
		want   [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
				want: [][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.args.want) {
				t.Errorf("got %v, want: %v", tt.args.matrix, tt.args.want)
			}
		})
	}
}
