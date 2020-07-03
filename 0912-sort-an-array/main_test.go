package _912_sort_an_array

import (
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{args: args{nums: []int{5, 2, 3, 1}}, want: []int{1, 2, 3, 5}},
		{args: args{nums: []int{5, 1, 1, 2, 0, 0}}, want: []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray1() = %v, want %v", got, tt.want)
			}
			if got := sortArray2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray2() = %v, want %v", got, tt.want)
			}
			if got := sortArray3(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray3() = %v, want %v", got, tt.want)
			}
			if got := sortArray4(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray4() = %v, want %v", got, tt.want)
			}
			if got := sortArray5(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray5() = %v, want %v", got, tt.want)
			}
			if got := sortArray6(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray6() = %v, want %v", got, tt.want)
			}
			if got := sortArray7(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray7() = %v, want %v", got, tt.want)
			}
		})
	}
}
