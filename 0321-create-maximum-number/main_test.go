package _321_create_maximum_number

import (
	"reflect"
	"testing"
)

func Test_maxNumber(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums1: []int{6, 7},
				nums2: []int{6, 0, 4},
				k:     5,
			},
			want: []int{6, 7, 6, 0, 4},
		},
		{
			args: args{
				nums1: []int{3, 4, 6, 5},
				nums2: []int{9, 1, 2, 5, 8, 3},
				k:     5,
			},
			want: []int{9, 8, 6, 5, 3},
		},
		{
			args: args{
				nums1: []int{3, 9},
				nums2: []int{8, 9},
				k:     3,
			},
			want: []int{9, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumber(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
