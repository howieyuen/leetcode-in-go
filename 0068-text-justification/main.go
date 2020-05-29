package _068_text_justification

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var ans []string
	var curWidth int
	i := 0
	for i < len(words) {
		curWidth = len(words[i])
		j := i + 1
		for ; j < len(words); j++ {
			if curWidth+len(words[j])+j-i <= maxWidth {
				curWidth += len(words[j])
				continue
			}
			
			interval := j - i - 1
			total := maxWidth - curWidth
			if interval == 0 {
				ans = append(ans, words[i]+strings.Repeat(" ", total))
				i = j
				break
			}
			
			quotient := total / interval
			remainder := total % interval
			sep := strings.Repeat(" ", quotient)
			if remainder == 0 {
				ans = append(ans, strings.Join(words[i:j], sep))
				i = j
				break
			}
			
			front := strings.Join(words[i:i+remainder+1], strings.Repeat(" ", quotient+1))
			var end string
			if (j-i)-(remainder+1) > 1 {
				end = strings.Join(words[i+remainder+1:j], sep)
			} else {
				end = words[i+remainder+1]
			}
			ans = append(ans, front+sep+end)
			i = j
			break
		}
		if j == len(words) {
			break
		}
	}
	if i < len(words) {
		ans = append(ans, strings.Join(words[i:], " ")+strings.Repeat(" ", maxWidth-curWidth-(len(words)-1-i)))
	}
	return ans
}
