package _809_expressive_words

import "testing"

func Test_expressiveWords(t *testing.T) {
	type args struct {
		S     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: args{S: "heeellooo", words: []string{"helo", "hello", "hi"}}, want: 1},
		{args: args{S: "aaabbbaaa", words: []string{"a"}}, want: 0},
		{args: args{S: "zzzzzyyyyy", words: []string{"zzyy", "zy", "zyy"}}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressiveWords(tt.args.S, tt.args.words); got != tt.want {
				t.Errorf("expressiveWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
