package _030_matrix_cells_in_distance_order

import (
	`reflect`
	`testing`
)

func Test_allCellsDistOrder(t *testing.T) {
	type args struct {
		rows int
		cols int
		r0   int
		c0   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				rows: 2,
				cols: 3,
				r0:   1,
				c0:   2,
			},
			want: [][]int{{1, 2}, {0, 2}, {1, 1}, {0, 1}, {1, 0}, {0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder(tt.args.rows, tt.args.cols, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
