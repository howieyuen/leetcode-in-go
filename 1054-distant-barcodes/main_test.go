package _054_distant_barcodes

import (
	"testing"
)

func Test_rearrangeBarcodes(t *testing.T) {
	type args struct {
		barcodes []int
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{barcodes: []int{1, 1, 1, 1, 2, 2, 2}}},
		{args: args{barcodes: []int{1, 1, 1, 1, 2, 2, 3, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rearrangeBarcodes(tt.args.barcodes)
			for i := 1; i < len(got); i++ {
				if got[i] == got[i-1] {
					t.Errorf("same neighbours")
				}
			}
		})
	}
}
