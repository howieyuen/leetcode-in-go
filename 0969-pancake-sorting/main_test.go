package _969_pancake_sorting

import (
	"reflect"
	"testing"
)

func Test_pancakeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{arr: []int{3, 2, 4, 1}}, want: []int{3, 4, 2, 3, 2}},
		{args: args{arr: []int{3, 1, 2}}, want: []int{3, 2}},
		{args: args{arr: []int{2, 1, 3}}, want: []int{2}},
		{args: args{arr: []int{1, 2, 3}}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pancakeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pancakeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
