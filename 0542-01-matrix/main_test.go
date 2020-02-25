package _542_01_matrix

import (
	`reflect`
	`testing`
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{matrix: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			args: args{matrix: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			}},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
			if got := updateMatrix1(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix1() = %v, want %v", got, tt.want)
			}
			if got := updateMatrix2(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix2() = %v, want %v", got, tt.want)
			}
		})
	}
}
