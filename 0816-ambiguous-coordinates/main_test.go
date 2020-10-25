package _816_ambiguous_coordinates

import (
	"reflect"
	"sort"
	"testing"
)

func Test_ambiguousCoordinates(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{args: args{S: "(123)"}, want: []string{"(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"}},
		{args: args{S: "(00011)"}, want: []string{"(0.001, 1)", "(0, 0.011)"}},
		{args: args{S: "(0123)"}, want: []string{"(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"}},
		{args: args{S: "(100)"}, want: []string{"(10, 0)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ambiguousCoordinates(tt.args.S)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ambiguousCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
