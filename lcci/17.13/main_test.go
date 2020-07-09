package _7_13

import (
	"testing"
)

func Test_respace(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				dictionary: []string{"looked", "just", "like", "her", "brother"},
				sentence:   "jesslookedjustliketimherbrother",
			},
			want: 7,
		},
		{
			args: args{
				dictionary: []string{"potimzz"},
				sentence:   "potimzzpotimzz",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := respace(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("respace() = %v, want %v", got, tt.want)
			}
		})
	}
}
