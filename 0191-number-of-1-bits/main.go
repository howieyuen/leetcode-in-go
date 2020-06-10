package _191_number_of_1_bits

import (
	"strconv"
	"strings"
)

func hammingWeight1(num uint32) int {
	return strings.Count(strconv.FormatUint(uint64(num), 2), "1")
}

func hammingWeight(num uint32) int {
	var res int
	for num > 0 {
		if num&1 != 0 {
			res++
		}
		num = num >> 1
	}
	return res
}