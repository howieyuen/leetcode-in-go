package _844_backspace_string_compare

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{S: "ab#c", T: "ad#c"}, want: true},
		{args: args{S: "ab##", T: "c#d#"}, want: true},
		{args: args{S: "a##c", T: "#a#c"}, want: true},
		{args: args{S: "a#c", T: "b"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
			if got := backspaceCompare1(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare1() = %v, want %v", got, tt.want)
			}
		})
	}
}
