package _487_making_file_names_unique

import (
	"reflect"
	"testing"
)

func Test_getFolderNames(t *testing.T) {
	type args struct {
		names []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{args: args{names: []string{"pes", "fifa", "gta", "pes(2019)"}},
			want: []string{"pes", "fifa", "gta", "pes(2019)"}},
		{args: args{names: []string{"kaido", "kaido(1)", "kaido", "kaido(1)"}},
			want: []string{"kaido", "kaido(1)", "kaido(2)", "kaido(1)(1)"}},
		{args: args{names: []string{"wano", "wano", "wano", "wano"}},
			want: []string{"wano", "wano(1)", "wano(2)", "wano(3)"}},
		{args: args{names: []string{"gta", "gta(1)", "gta", "avalon"}},
			want: []string{"gta", "gta(1)", "gta(2)", "avalon"}},
		{args: args{names: []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}},
			want: []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece(4)"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFolderNames(tt.args.names); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFolderNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
