package _546_remove_boxes

import (
	"testing"
)

func Test_removeBoxes(t *testing.T) {
	type args struct {
		boxes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{boxes: []int{1, 3, 2, 2, 2, 3, 4, 3, 1}}, want: 23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBoxes(tt.args.boxes); got != tt.want {
				t.Errorf("removeBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
