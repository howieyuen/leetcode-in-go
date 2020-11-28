package _691_stickers_to_spell_word

import (
	"testing"
)

func Test_minStickers(t *testing.T) {
	type args struct {
		stickers []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{stickers: []string{"with", "example", "science"}, target: "thehat"}, want: 3},
		{args: args{stickers: []string{"notice", "possible"}, target: "basicbasic"}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minStickers(tt.args.stickers, tt.args.target); got != tt.want {
				t.Errorf("minStickers() = %v, want %v", got, tt.want)
			}
			if got := minStickers1(tt.args.stickers, tt.args.target); got != tt.want {
				t.Errorf("minStickers1() = %v, want %v", got, tt.want)
			}
		})
	}
}
