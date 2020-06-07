package _130_surrounded_regions

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
		want  [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
				want: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if solve(tt.args.board); !reflect.DeepEqual(tt.args.board, tt.args.want) {
				t.Errorf("longestConsecutive() = %v, want %v", tt.args.board, tt.args.want)
			}
		})
	}
}
