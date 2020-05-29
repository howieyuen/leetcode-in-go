package _068_text_justification

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  "},
		},
		{
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        "},
		},
		{
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20,
			},
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  "},
		},
		{
			args: args{
				words: []string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you",
					"can", "do", "for", "your", "country"},
				maxWidth: 16,
			},
			want: []string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country"},
		},
		{
			args: args{
				words:    []string{"Give", "me", "my", "Romeo;", "and,", "when", "he", "shall", "die,", "Take", "him", "and", "cut", "him", "out", "in", "little", "stars,", "And", "he", "will", "make", "the", "face", "of", "heaven", "so", "fine", "That", "all", "the", "world", "will", "be", "in", "love", "with", "night", "And", "pay", "no", "worship", "to", "the", "garish", "sun."},
				maxWidth: 25,
			},
			want: []string{
				"Give  me  my  Romeo; and,",
				"when  he  shall die, Take",
				"him  and  cut  him out in",
				"little stars, And he will",
				"make  the  face of heaven",
				"so   fine  That  all  the",
				"world  will  be  in  love",
				"with  night  And  pay  no",
				"worship   to  the  garish",
				"sun.                     "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
