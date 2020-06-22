package _6_18

import (
	"testing"
)

func Test_patternMatching(t *testing.T) {
	type args struct {
		pattern string
		value   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{args: args{pattern: "abba", value: "dogcatcatdog",}, want: true},
		{args: args{pattern: "abba", value: "dogcatcatfish",}, want: false},
		{args: args{pattern: "aaaa", value: "dogcatcatdog",}, want: false},
		{args: args{pattern: "abba", value: "dogdogdogdog",}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := patternMatching(tt.args.pattern, tt.args.value); got != tt.want {
				t.Errorf("patternMatching() = %v, want %v", got, tt.want)
			}
		})
	}
}
