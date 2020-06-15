package _171_excel_sheet_column_number

import (
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{s: "Z"}, want: 26},
		{args: args{s: "AB"}, want: 28},
		{args: args{s: "ZY"}, want: 701},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.s); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
