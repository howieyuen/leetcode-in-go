package _356_sort_integers_by_the_number_of_1_bits

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}}, want: []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortByBits(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
