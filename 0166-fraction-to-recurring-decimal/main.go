package _166_fraction_to_recurring_decimal

import (
	"bytes"
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var fraction bytes.Buffer
	if numerator > 0 && denominator < 0 || numerator < 0 && denominator > 0 {
		fraction.WriteByte('-')
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	fraction.WriteString(strconv.Itoa(numerator / denominator))
	numerator %= denominator
	if numerator == 0 {
		return fraction.String()
	}
	fraction.WriteByte('.')
	var m = make(map[int]int)
	var start = -1
	for numerator != 0 {
		numerator *= 10
		if pos, ok := m[numerator]; ok {
			start = pos
			break
		}
		m[numerator] = fraction.Len()
		fraction.WriteString(strconv.Itoa(numerator / denominator))
		numerator %= denominator
	}
	if start == -1 {
		return fraction.String()
	}
	str := fraction.String()
	return fmt.Sprintf("%s(%s)", str[:start], str[start:])
}
