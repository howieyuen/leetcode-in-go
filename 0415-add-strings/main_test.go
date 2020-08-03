package _415_add_strings

import (
	"testing"
)

func Test_addStrings(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{num1: "123", num2: "321"}, want: "444"},
		{args: args{num1: "999", num2: "1001"}, want: "2000"},
		{args: args{num1: "999", num2: "999"}, want: "1998"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
