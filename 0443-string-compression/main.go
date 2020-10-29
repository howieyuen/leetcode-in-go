package _443_string_compression

import (
	"strconv"
)

func compress(chars []byte) int {
	var write, anchor int
	for read, c := range chars {
		if read+1 == len(chars) || chars[read+1] != c {
			chars[write] = chars[read]
			write++
			if read > anchor {
				for _, ch := range strconv.Itoa(read - anchor + 1) {
					chars[write] = byte(ch)
					write++
				}
			}
			anchor = read + 1
		}
	}
	return write
}
