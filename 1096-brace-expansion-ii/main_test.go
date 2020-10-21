package _096_brace_expansion_ii

import (
	"reflect"
	"testing"
)

func Test_braceExpansionII(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{args: args{expression: "{a,b}{c,{d,e}}"}, want: []string{"ac", "ad", "ae", "bc", "bd", "be"}},
		{args: args{expression: "{{a,z},a{b,c},{ab,z}}"}, want: []string{"a", "ab", "ac", "z"}},
		{args: args{expression: "abc{x,y{m,n}},{p,q}"}, want: []string{"abcx", "abcym", "abcyn", "p", "q"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := braceExpansionII(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceExpansionII() = %v, want %v", got, tt.want)
			}
			if got := braceExpansionII1(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceExpansionII1() = %v, want %v", got, tt.want)
			}
			if got := braceExpansionII2(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("braceExpansionII2() = %v, want %v", got, tt.want)
			}
		})
	}
}
